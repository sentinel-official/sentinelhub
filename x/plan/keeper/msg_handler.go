package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types/v3"
	subscriptiontypes "github.com/sentinel-official/sentinelhub/v12/x/subscription/types/v3"
)

// HandleMsgCreatePlan handles a request to create a new plan.
// It validates the provider, constructs the plan, stores it, and emits a creation event.
func (k *Keeper) HandleMsgCreatePlan(ctx sdk.Context, msg *v3.MsgCreatePlanRequest) (*v3.MsgCreatePlanResponse, error) {
	// Parse and validate the provider address
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Verify the provider exists and is in an active state
	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	if !provider.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
	}

	// Construct the new plan with an incremented ID and inactive status
	count := k.GetPlanCount(ctx)
	plan := v3.Plan{
		ID:          count + 1,
		ProvAddress: provAddr.String(),
		Bytes:       msg.Bytes,
		Duration:    msg.Duration,
		Prices:      msg.Prices,
		Private:     msg.Private,
		Status:      v1base.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	// Persist the plan and update indexes
	k.SetPlanCount(ctx, count+1)
	k.SetPlan(ctx, plan)
	k.SetPlanForProvider(ctx, provAddr, plan.ID)

	// Emit event to indicate plan creation
	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreate{
			PlanID:      plan.ID,
			ProvAddress: plan.ProvAddress,
			Bytes:       plan.Bytes.String(),
			Duration:    plan.Duration.String(),
			Prices:      plan.GetPrices().String(),
		},
	)

	return &v3.MsgCreatePlanResponse{
		ID: plan.ID,
	}, nil
}

// HandleMsgLinkNode handles a request to associate a node with a plan.
// It ensures authorization, validates node and lease state, and updates indexes accordingly.
func (k *Keeper) HandleMsgLinkNode(ctx sdk.Context, msg *v3.MsgLinkNodeRequest) (*v3.MsgLinkNodeResponse, error) {
	// Fetch and authorize the plan
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	if msg.From != plan.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Parse and validate the node address
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Prevent duplicate node linkage for the same plan
	if k.HasNodeForPlan(ctx, plan.ID, nodeAddr) {
		return nil, types.NewErrorDuplicateNodeForPlan(plan.ID, nodeAddr)
	}

	// Convert and validate the provider address from the plan
	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	// Ensure an active lease exists between the node and provider
	if !k.HasAnyLeaseForNodeByProvider(ctx, nodeAddr, provAddr) {
		return nil, types.NewErrorLeaseForNodeByProviderNotFound(nodeAddr, provAddr)
	}

	// Store the node-plan association and update indexes
	k.SetNodeForPlan(ctx, plan.ID, nodeAddr)
	k.SetPlanForNodeByProvider(ctx, nodeAddr, provAddr, plan.ID)

	// Emit event to signal node linkage
	ctx.EventManager().EmitTypedEvent(
		&v3.EventLinkNode{
			PlanID:      plan.ID,
			ProvAddress: plan.ProvAddress,
			NodeAddress: nodeAddr.String(),
		},
	)

	return &v3.MsgLinkNodeResponse{}, nil
}

// HandleMsgUnlinkNode handles a request to dissociate a node from a plan.
// It performs checks, calls pre-unlink hook, and removes the mapping.
func (k *Keeper) HandleMsgUnlinkNode(ctx sdk.Context, msg *v3.MsgUnlinkNodeRequest) (*v3.MsgUnlinkNodeResponse, error) {
	// Fetch and authorize the plan
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	if msg.From != plan.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Parse and validate the node address
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Ensure the node is currently linked to the plan
	if !k.HasNodeForPlan(ctx, plan.ID, nodeAddr) {
		return nil, types.NewErrorNodeForPlanNotFound(plan.ID, nodeAddr)
	}

	// Call unlink pre-hook to allow any custom logic or validation
	if err := k.PlanUnlinkNodePreHook(ctx, plan.ID, nodeAddr); err != nil {
		return nil, err
	}

	// Convert and validate the provider address from the plan
	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	// Remove the node-plan association and update indexes
	k.DeleteNodeForPlan(ctx, plan.ID, nodeAddr)
	k.DeletePlanForNodeByProvider(ctx, nodeAddr, provAddr, plan.ID)

	// Emit event to indicate node unlink
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUnlinkNode{
			PlanID:      plan.ID,
			ProvAddress: plan.ProvAddress,
			NodeAddress: nodeAddr.String(),
		},
	)

	return &v3.MsgUnlinkNodeResponse{}, nil
}

// HandleMsgUpdatePlanDetails processes a request to update the details of a plan.
// It updates the plan's privacy setting and ensures that the plan's details are persisted.
func (k *Keeper) HandleMsgUpdatePlanDetails(ctx sdk.Context, msg *v3.MsgUpdatePlanDetailsRequest) (*v3.MsgUpdatePlanDetailsResponse, error) {
	// Fetch and authorize the plan
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	if msg.From != plan.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Apply the plan's privacy setting and persist the changes
	plan.Private = msg.Private

	k.SetPlan(ctx, plan)

	// Emit event to indicate details update
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateDetails{
			PlanID:      plan.ID,
			ProvAddress: plan.ProvAddress,
			Private:     plan.Private,
		},
	)

	return &v3.MsgUpdatePlanDetailsResponse{}, nil
}

// HandleMsgUpdatePlanStatus handles a request to update the activation status of a plan.
// It updates the plan's status and cleans up the old status index.
func (k *Keeper) HandleMsgUpdatePlanStatus(ctx sdk.Context, msg *v3.MsgUpdatePlanStatusRequest) (*v3.MsgUpdatePlanStatusResponse, error) {
	// Fetch and authorize the plan
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	if msg.From != plan.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if msg.Status.Equal(v1base.StatusActive) {
		provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
		if err != nil {
			return nil, err
		}

		// Ensure provider exists and is currently active
		provider, found := k.GetProvider(ctx, provAddr)
		if !found {
			return nil, types.NewErrorProviderNotFound(provAddr)
		}

		if !provider.Status.Equal(v1base.StatusActive) {
			return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
		}
	}

	// Remove old status index if transitioning state
	if msg.Status.Equal(v1base.StatusActive) {
		if plan.Status.Equal(v1base.StatusInactive) {
			k.DeleteInactivePlan(ctx, plan.ID)
		}
	}

	if msg.Status.Equal(v1base.StatusInactive) {
		if plan.Status.Equal(v1base.StatusActive) {
			k.DeleteActivePlan(ctx, plan.ID)
		}
	}

	// Apply new status and persist the plan
	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)

	// Emit event to indicate status update
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateStatus{
			PlanID:      plan.ID,
			ProvAddress: plan.ProvAddress,
			Status:      plan.Status.String(),
		},
	)

	return &v3.MsgUpdatePlanStatusResponse{}, nil
}

// HandleMsgStartSession handles a request to start a new session under a plan.
// It triggers subscription and session workflows via delegation to the subscription module.
func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	// Forward subscription creation request to subscription module
	subscriptionReq := &subscriptiontypes.MsgStartSubscriptionRequest{
		From:               msg.From,
		ID:                 msg.ID,
		Denom:              msg.Denom,
		RenewalPricePolicy: msg.RenewalPricePolicy,
	}

	subscriptionResp, err := k.subscription.HandleMsgStartSubscription(ctx, subscriptionReq)
	if err != nil {
		return nil, err
	}

	// Forward session start request based on returned subscription ID
	sessionReq := &subscriptiontypes.MsgStartSessionRequest{
		From:        msg.From,
		ID:          subscriptionResp.ID,
		NodeAddress: msg.NodeAddress,
	}

	sessionResp, err := k.subscription.HandleMsgStartSession(ctx, sessionReq)
	if err != nil {
		return nil, err
	}

	return &v3.MsgStartSessionResponse{
		ID: sessionResp.ID,
	}, nil
}
