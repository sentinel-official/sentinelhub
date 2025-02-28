package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) HandleMsgCancelSubscription(ctx sdk.Context, msg *v3.MsgCancelSubscriptionRequest) (*v3.MsgCancelSubscriptionResponse, error) {
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

	if err := k.SubscriptionInactivePendingPreHook(ctx, subscription.ID); err != nil {
		return nil, err
	}

	k.DeleteSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	subscription.RenewalPricePolicy = v1base.RenewalPricePolicyUnspecified
	subscription.InactiveAt = k.GetInactiveAt(ctx)
	subscription.Status = v1base.StatusInactivePending
	subscription.StatusAt = ctx.BlockTime()

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)

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

func (k *Keeper) HandleMsgRenewSubscription(ctx sdk.Context, msg *v3.MsgRenewSubscriptionRequest) (*v3.MsgRenewSubscriptionResponse, error) {
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

	plan, found := k.GetPlan(ctx, subscription.PlanID)
	if !found {
		return nil, types.NewErrorPlanNotFound(subscription.PlanID)
	}
	if !plan.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	price, found := plan.Price(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	price, err := price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	if err := subscription.ValidateRenewalPolicies(price); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidRenewalPolicy, err.Error())
	}

	k.DeleteSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	share := k.StakingShare(ctx)
	totalPayment := price.QuotePrice()

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	reward := baseutils.GetProportionOfCoin(totalPayment, share)
	if err := k.SendCoinFromAccountToModule(ctx, accAddr, k.feeCollectorName, reward); err != nil {
		return nil, err
	}

	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	payment := totalPayment.Sub(reward)
	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), payment); err != nil {
		return nil, err
	}

	subscription = v3.Subscription{
		ID:                 subscription.ID,
		AccAddress:         subscription.AccAddress,
		PlanID:             subscription.PlanID,
		Price:              price,
		RenewalPricePolicy: subscription.RenewalPricePolicy,
		Status:             v1base.StatusActive,
		InactiveAt:         ctx.BlockTime().Add(plan.GetHours()),
		StartAt:            ctx.BlockTime(),
		StatusAt:           ctx.BlockTime(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

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

func (k *Keeper) HandleMsgShareSubscription(ctx sdk.Context, msg *v3.MsgShareSubscriptionRequest) (*v3.MsgShareSubscriptionResponse, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

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

	grantedBytes := fromAlloc.GrantedBytes.Add(toAlloc.GrantedBytes)
	utilisedBytes := fromAlloc.UtilisedBytes.Add(toAlloc.UtilisedBytes)
	availableBytes := grantedBytes.Sub(utilisedBytes)

	if msg.Bytes.GT(availableBytes) {
		return nil, types.NewErrorInsufficientBytes(subscription.ID, msg.Bytes)
	}

	fromAlloc.GrantedBytes = availableBytes.Sub(msg.Bytes)
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

func (k *Keeper) HandleMsgStartSubscription(ctx sdk.Context, msg *v3.MsgStartSubscriptionRequest) (*v3.MsgStartSubscriptionResponse, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if !plan.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	if plan.IsPrivate() && !accAddr.Equals(provAddr) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	price, found := plan.Price(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	price, err = price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	count := k.GetSubscriptionCount(ctx)
	subscription := v3.Subscription{
		ID:                 count + 1,
		AccAddress:         accAddr.String(),
		PlanID:             plan.ID,
		Price:              price,
		RenewalPricePolicy: msg.RenewalPricePolicy,
		Status:             v1base.StatusActive,
		InactiveAt:         ctx.BlockTime().Add(plan.GetHours()),
		StartAt:            ctx.BlockTime(),
		StatusAt:           ctx.BlockTime(),
	}

	share := k.StakingShare(ctx)
	totalPayment := price.QuotePrice()

	reward := baseutils.GetProportionOfCoin(totalPayment, share)
	if err := k.SendCoinFromAccountToModule(ctx, accAddr, k.feeCollectorName, reward); err != nil {
		return nil, err
	}

	payment := totalPayment.Sub(reward)
	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), payment); err != nil {
		return nil, err
	}

	k.SetSubscriptionCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.ID)
	k.SetSubscriptionForPlan(ctx, subscription.PlanID, subscription.ID)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

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

	alloc := v2.Allocation{
		ID:            subscription.ID,
		Address:       subscription.AccAddress,
		GrantedBytes:  plan.GetGigabytes(),
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

func (k *Keeper) HandleMsgUpdateSubscription(ctx sdk.Context, msg *v3.MsgUpdateSubscriptionRequest) (*v3.MsgUpdateSubscriptionResponse, error) {
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

	k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

	subscription.RenewalPricePolicy = msg.RenewalPricePolicy

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

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

func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if !subscription.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.ID, subscription.Status)
	}

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

	if !k.HasNodeForPlan(ctx, subscription.PlanID, nodeAddr) {
		return nil, types.NewErrorNodeForPlanNotFound(subscription.PlanID, nodeAddr)
	}

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

	k.SetSessionCount(ctx, count+1)
	k.SetSession(ctx, session)
	k.SetSessionForAccount(ctx, accAddr, session.ID)
	k.SetSessionForNode(ctx, nodeAddr, session.ID)
	k.SetSessionForPlanByNode(ctx, subscription.PlanID, nodeAddr, session.ID)
	k.SetSessionForSubscription(ctx, subscription.ID, session.ID)
	k.SetSessionForAllocation(ctx, subscription.ID, accAddr, session.ID)
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

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

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.SetParams(ctx, msg.Params)
	return &v3.MsgUpdateParamsResponse{}, nil
}
