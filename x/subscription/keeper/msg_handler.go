package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	baseutils "github.com/sentinel-official/sentinelhub/v12/utils"
	sessiontypes "github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
	"github.com/sentinel-official/sentinelhub/v12/x/subscription/types"
	"github.com/sentinel-official/sentinelhub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/sentinelhub/v12/x/subscription/types/v3"
)

// HandleMsgCancelSubscription handles a request to cancel an active subscription.
// It verifies authorization, marks the subscription as inactive pending, updates state and emits a cancellation event.
func (k *Keeper) HandleMsgCancelSubscription(ctx sdk.Context, msg *v3.MsgCancelSubscriptionRequest) (*v3.MsgCancelSubscriptionResponse, error) {
	// Fetch the subscription, check if active, and verify ownership
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	if !subscription.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.ID, subscription.Status)
	}

	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Run the pre-hook to prepare the subscription for deactivation
	if err := k.SubscriptionInactivePendingPreHook(ctx, subscription.ID); err != nil {
		return nil, err
	}

	// Remove the subscription from existing indexes tied to future timestamps
	k.DeleteSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	// Clear renewal policy and mark subscription as inactive pending
	subscription.RenewalPricePolicy = v1base.RenewalPricePolicyUnspecified
	subscription.InactiveAt = k.GetInactiveAt(ctx)
	subscription.Status = v1base.StatusInactivePending
	subscription.StatusAt = ctx.BlockTime()

	// Persist the updated subscription and re-index it for inactive tracking
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)

	// Emit event to signal subscription status update
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdate{
			ID:                 subscription.ID,
			PlanID:             subscription.PlanID,
			AccAddress:         subscription.AccAddress,
			RenewalPricePolicy: subscription.RenewalPricePolicy.String(),
			Status:             subscription.Status,
			InactiveAt:         subscription.InactiveAt.String(),
			StatusAt:           subscription.StatusAt.String(),
		},
	)

	return &v3.MsgCancelSubscriptionResponse{}, nil
}

// HandleMsgRenewSubscription handles a request to renew an active subscription.
// It verifies authorization, validates pricing, processes payment and rewards, updates the subscription, and emits events.
func (k *Keeper) HandleMsgRenewSubscription(ctx sdk.Context, msg *v3.MsgRenewSubscriptionRequest) (*v3.MsgRenewSubscriptionResponse, error) {
	// Validate the subscription and requester ownership
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if !subscription.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.ID, subscription.Status)
	}

	// Fetch and validate the associated plan
	plan, found := k.GetPlan(ctx, subscription.PlanID)
	if !found {
		return nil, types.NewErrorPlanNotFound(subscription.PlanID)
	}

	if !plan.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	// Retrieve and quote the plan's price for the given denomination
	price, found := plan.Price(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	price, err := price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	// Validate that the renewal is allowed under the subscription's price policy
	if err := subscription.ValidateRenewalPolicies(price); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidRenewalPolicy, err.Error())
	}

	// Remove old index entries before updating subscription
	k.DeleteSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	// Calculate payment and reward amounts
	share := k.StakingShare(ctx)
	total := price.QuotePrice()

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	reward := baseutils.GetProportionOfCoin(total, share)
	if err := k.SendCoinFromAccountToModule(ctx, accAddr, k.feeCollectorName, reward); err != nil {
		return nil, err
	}

	// Transfer remaining payment to the provider
	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	payment := total.Sub(reward)
	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), payment); err != nil {
		return nil, err
	}

	// Construct the renewed subscription with updated times and pricing
	inactiveAt := ctx.BlockTime().Add(plan.GetDuration())
	subscription = v3.Subscription{
		ID:                 subscription.ID,
		AccAddress:         subscription.AccAddress,
		PlanID:             subscription.PlanID,
		Price:              price,
		RenewalPricePolicy: subscription.RenewalPricePolicy,
		Status:             v1base.StatusActive,
		InactiveAt:         inactiveAt,
		StartAt:            ctx.BlockTime(),
		StatusAt:           ctx.BlockTime(),
	}

	// Save updated subscription and re-index
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	// Emit renewal and payment events
	ctx.EventManager().EmitTypedEvents(
		&v3.EventRenew{
			ID:          subscription.ID,
			PlanID:      subscription.PlanID,
			AccAddress:  subscription.AccAddress,
			ProvAddress: provAddr.String(),
			Price:       subscription.Price.String(),
		},
		&v3.EventPay{
			ID:            subscription.ID,
			PlanID:        subscription.PlanID,
			AccAddress:    subscription.AccAddress,
			ProvAddress:   provAddr.String(),
			Payment:       payment.String(),
			StakingReward: reward.String(),
		},
	)

	// Reset utilisation on all allocations under this subscription
	k.IterateAllocationsForSubscription(ctx, subscription.ID, func(_ int, item v2.Allocation) bool {
		item.UtilisedBytes = sdkmath.ZeroInt()

		k.SetAllocation(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&v3.EventAllocate{
				ID:            item.ID,
				AccAddress:    item.Address,
				GrantedBytes:  item.GrantedBytes.String(),
				UtilisedBytes: item.UtilisedBytes.String(),
			},
		)

		return false
	})

	return &v3.MsgRenewSubscriptionResponse{}, nil
}

// HandleMsgShareSubscription handles a request to share granted subscription bytes with another account.
// It validates ownership and availability, updates allocations, and emits events for both sender and recipient.
func (k *Keeper) HandleMsgShareSubscription(ctx sdk.Context, msg *v3.MsgShareSubscriptionRequest) (*v3.MsgShareSubscriptionResponse, error) {
	// Fetch the subscription and verify the sender is the owner
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Load allocation from sender and parse both account addresses
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	fromAlloc, found := k.GetAllocation(ctx, subscription.ID, fromAddr)
	if !found {
		return nil, types.NewErrorAllocationNotFound(subscription.ID, fromAddr)
	}

	toAddr, err := sdk.AccAddressFromBech32(msg.AccAddress)
	if err != nil {
		return nil, err
	}

	// Load or initialize recipient allocation
	toAlloc, found := k.GetAllocation(ctx, subscription.ID, toAddr)
	if !found {
		toAlloc = v2.Allocation{
			ID:            subscription.ID,
			Address:       toAddr.String(),
			GrantedBytes:  sdkmath.ZeroInt(),
			UtilisedBytes: sdkmath.ZeroInt(),
		}

		k.SetSubscriptionForAccount(ctx, toAddr, subscription.ID)
	}

	// Compute shared and available bytes; validate sufficiency
	grantedBytes := fromAlloc.GrantedBytes.Add(toAlloc.GrantedBytes)
	utilisedBytes := fromAlloc.UtilisedBytes.Add(toAlloc.UtilisedBytes)
	availableBytes := grantedBytes.Sub(utilisedBytes)

	if msg.Bytes.GT(availableBytes) {
		return nil, types.NewErrorInsufficientBytes(subscription.ID, msg.Bytes)
	}

	// Update sender's allocation after subtracting granted bytes
	fromAlloc.GrantedBytes = grantedBytes.Sub(msg.Bytes)
	if fromAlloc.GrantedBytes.LT(fromAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.ID, fromAddr)
	}

	k.SetAllocation(ctx, fromAlloc)
	ctx.EventManager().EmitTypedEvent(
		&v3.EventAllocate{
			ID:            fromAlloc.ID,
			AccAddress:    fromAlloc.Address,
			GrantedBytes:  fromAlloc.GrantedBytes.String(),
			UtilisedBytes: fromAlloc.GrantedBytes.String(),
		},
	)

	// Update recipient allocation with granted bytes
	toAlloc.GrantedBytes = msg.Bytes
	if toAlloc.GrantedBytes.LT(toAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.ID, toAddr)
	}

	k.SetAllocation(ctx, toAlloc)
	ctx.EventManager().EmitTypedEvent(
		&v3.EventAllocate{
			ID:            toAlloc.ID,
			AccAddress:    toAlloc.Address,
			GrantedBytes:  toAlloc.GrantedBytes.String(),
			UtilisedBytes: toAlloc.GrantedBytes.String(),
		},
	)

	return &v3.MsgShareSubscriptionResponse{}, nil
}

// HandleMsgStartSubscription handles a request to start a new subscription.
// It validates the plan and pricing, processes payment and rewards, persists the subscription and allocation, and emits events.
func (k *Keeper) HandleMsgStartSubscription(ctx sdk.Context, msg *v3.MsgStartSubscriptionRequest) (*v3.MsgStartSubscriptionResponse, error) {
	// Fetch the plan and verify it's active
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	if !plan.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	// Parse subscriber and provider addresses
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	// Restrict private plans to the provider only
	if plan.IsPrivate() && !accAddr.Equals(provAddr) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Fetch and quote the plan price for the requested denomination
	price, found := plan.Price(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	price, err = price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	// Build the subscription object with a new ID
	count := k.GetSubscriptionCount(ctx)
	inactiveAt := ctx.BlockTime().Add(plan.GetDuration())
	subscription := v3.Subscription{
		ID:                 count + 1,
		AccAddress:         accAddr.String(),
		PlanID:             plan.ID,
		Price:              price,
		RenewalPricePolicy: msg.RenewalPricePolicy,
		Status:             v1base.StatusActive,
		InactiveAt:         inactiveAt,
		StartAt:            ctx.BlockTime(),
		StatusAt:           ctx.BlockTime(),
	}

	// Process payment and reward
	share := k.StakingShare(ctx)
	total := price.QuotePrice()

	reward := baseutils.GetProportionOfCoin(total, share)
	if err := k.SendCoinFromAccountToModule(ctx, accAddr, k.feeCollectorName, reward); err != nil {
		return nil, err
	}

	payment := total.Sub(reward)
	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), payment); err != nil {
		return nil, err
	}

	// Save subscription and update all related indexes
	k.SetSubscriptionCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.ID)
	k.SetSubscriptionForPlan(ctx, subscription.PlanID, subscription.ID)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	// Emit subscription creation and payment events
	ctx.EventManager().EmitTypedEvents(
		&v3.EventCreate{
			ID:          subscription.ID,
			PlanID:      subscription.PlanID,
			AccAddress:  subscription.AccAddress,
			ProvAddress: provAddr.String(),
			Price:       subscription.Price.String(),
		},
		&v3.EventPay{
			ID:            subscription.ID,
			PlanID:        subscription.PlanID,
			AccAddress:    subscription.AccAddress,
			ProvAddress:   provAddr.String(),
			Payment:       payment.String(),
			StakingReward: reward.String(),
		},
	)

	// Initialize allocation for the subscriber
	alloc := v2.Allocation{
		ID:            subscription.ID,
		Address:       subscription.AccAddress,
		GrantedBytes:  plan.GetBytes(),
		UtilisedBytes: sdkmath.ZeroInt(),
	}

	k.SetAllocation(ctx, alloc)
	ctx.EventManager().EmitTypedEvent(
		&v3.EventAllocate{
			ID:            alloc.ID,
			AccAddress:    alloc.Address,
			GrantedBytes:  alloc.GrantedBytes.String(),
			UtilisedBytes: alloc.UtilisedBytes.String(),
		},
	)

	return &v3.MsgStartSubscriptionResponse{
		ID: subscription.ID,
	}, nil
}

// HandleMsgUpdateSubscription handles a request to update the renewal price policy of an active subscription.
// It verifies ownership and status, updates the policy, refreshes indexes, and emits an update event.
func (k *Keeper) HandleMsgUpdateSubscription(ctx sdk.Context, msg *v3.MsgUpdateSubscriptionRequest) (*v3.MsgUpdateSubscriptionResponse, error) {
	// Fetch the subscription, ensure it's active, and verify sender ownership
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	if !subscription.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.ID, subscription.Status)
	}

	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Remove old renewal index before applying update
	k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	// Apply the new renewal policy
	subscription.RenewalPricePolicy = msg.RenewalPricePolicy

	// Persist the updated subscription and re-index
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	// Emit subscription update event
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdate{
			ID:                 subscription.ID,
			PlanID:             subscription.PlanID,
			AccAddress:         subscription.AccAddress,
			RenewalPricePolicy: subscription.RenewalPricePolicy.String(),
		},
	)

	return &v3.MsgUpdateSubscriptionResponse{}, nil
}

// HandleMsgStartSession handles a request to start a new session under an active subscription.
// It validates the subscription, node, and allocation, creates the session, indexes it, and emits a session creation event.
func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	// Validate the subscription and ensure it's active
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	if !subscription.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.ID, subscription.Status)
	}

	// Parse and validate the node address and ensure node is active
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Ensure the node is authorized to serve this plan
	if !k.HasNodeForPlan(ctx, subscription.PlanID, nodeAddr) {
		return nil, types.NewErrorNodeForPlanNotFound(subscription.PlanID, nodeAddr)
	}

	// Parse account address and validate allocation availability
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	alloc, found := k.GetAllocation(ctx, subscription.ID, accAddr)
	if !found {
		return nil, types.NewErrorAllocationNotFound(subscription.ID, accAddr)
	}

	if alloc.UtilisedBytes.GTE(alloc.GrantedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.ID, accAddr)
	}

	// Build a new session with default usage and time fields
	count := k.GetSessionCount(ctx)
	inactiveAt := k.GetSessionInactiveAt(ctx)
	session := &v3.Session{
		BaseSession: &sessiontypes.BaseSession{
			ID:            count + 1,
			AccAddress:    accAddr.String(),
			NodeAddress:   nodeAddr.String(),
			DownloadBytes: sdkmath.ZeroInt(),
			UploadBytes:   sdkmath.ZeroInt(),
			MaxBytes:      sdkmath.ZeroInt(),
			Duration:      0,
			MaxDuration:   0,
			Status:        v1base.StatusActive,
			InactiveAt:    inactiveAt,
			StartAt:       ctx.BlockTime(),
			StatusAt:      ctx.BlockTime(),
		},
		SubscriptionID: subscription.ID,
	}

	// Save the session and register all related indexes
	k.SetSessionCount(ctx, count+1)
	k.SetSession(ctx, session)
	k.SetSessionForAccount(ctx, accAddr, session.ID)
	k.SetSessionForNode(ctx, nodeAddr, session.ID)
	k.SetSessionForPlanByNode(ctx, subscription.PlanID, nodeAddr, session.ID)
	k.SetSessionForSubscription(ctx, subscription.ID, session.ID)
	k.SetSessionForAllocation(ctx, subscription.ID, accAddr, session.ID)
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Emit session creation event
	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreateSession{
			ID:             session.ID,
			AccAddress:     session.AccAddress,
			NodeAddress:    session.NodeAddress,
			SubscriptionID: session.SubscriptionID,
		},
	)

	return &v3.MsgStartSessionResponse{
		ID: session.ID,
	}, nil
}

// HandleMsgUpdateParams handles a request to update module-wide parameters.
// It verifies the authority and stores the new parameter configuration.
func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	// Ensure the sender is authorized to update parameters
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Apply and persist the new parameter set
	k.SetParams(ctx, msg.Params)

	return &v3.MsgUpdateParamsResponse{}, nil
}
