package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

// SessionInactivePreHook performs cleanup operations when a session transitions to an inactive state.
func (k *Keeper) SessionInactivePreHook(ctx sdk.Context, id uint64) error {
	k.Logger(ctx).Info("Running session inactive pre-hook", "id", id)

	// Retrieve the session by ID; return an error if it doesn't exist.
	item, found := k.GetSession(ctx, id)
	if !found {
		return types.NewErrorSessionNotFound(id)
	}

	// Ensure the session is of type v3.Session; do nothing if it's not.
	session, ok := item.(*v3.Session)
	if !ok {
		return nil
	}

	// Verify that the session's status is "InactivePending"; otherwise, return an error.
	if !session.Status.Equal(v1base.StatusInactivePending) {
		return types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Fetch the subscription associated with the session; return an error if it doesn't exist.
	subscription, found := k.GetSubscription(ctx, session.SubscriptionID)
	if !found {
		return types.NewErrorSubscriptionNotFound(session.SubscriptionID)
	}

	// Decode the session's account address from Bech32 format.
	accAddr, err := sdk.AccAddressFromBech32(session.AccAddress)
	if err != nil {
		return err
	}

	// Decode the session's node address from Bech32 format.
	nodeAddr, err := base.NodeAddressFromBech32(session.NodeAddress)
	if err != nil {
		return err
	}

	// Remove session references for allocation, node, plan, and subscription.
	k.DeleteSessionForAllocation(ctx, subscription.ID, accAddr, session.ID)
	k.DeleteSessionForPlanByNode(ctx, subscription.PlanID, nodeAddr, session.ID)
	k.DeleteSessionForSubscription(ctx, subscription.ID, session.ID)

	return nil
}

// SessionUpdatePreHook updates session and allocation details during a session update.
func (k *Keeper) SessionUpdatePreHook(ctx sdk.Context, id uint64, currBytes sdkmath.Int) error {
	k.Logger(ctx).Info("Running session update pre-hook", "id", id)

	// Retrieve the session by ID; return an error if it doesn't exist.
	item, found := k.GetSession(ctx, id)
	if !found {
		return types.NewErrorSessionNotFound(id)
	}

	// Ensure the session is of type v3.Session; do nothing if it's not.
	session, ok := item.(*v3.Session)
	if !ok {
		return nil
	}

	// Ensure the session is not in the "Inactive" state; return an error if it is.
	if session.Status.Equal(v1base.StatusInactive) {
		return types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Decode the session's account address from Bech32 format.
	accAddr, err := sdk.AccAddressFromBech32(session.AccAddress)
	if err != nil {
		return err
	}

	// Fetch the allocation for the subscription and account; return an error if it doesn't exist.
	alloc, found := k.GetAllocation(ctx, session.SubscriptionID, accAddr)
	if !found {
		return types.NewErrorAllocationNotFound(session.SubscriptionID, accAddr)
	}

	// Update allocation's utilised bytes based on the difference between current and previous session bytes.
	diffBytes := currBytes.Sub(session.Bytes())
	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(diffBytes)

	// Store the updated allocation in the keeper.
	k.SetAllocation(ctx, alloc)

	// Emit an event logging the updated allocation details.
	ctx.EventManager().EmitTypedEvent(
		&v3.EventAllocate{
			ID:            alloc.ID,
			AccAddress:    alloc.Address,
			GrantedBytes:  alloc.GrantedBytes.String(),
			UtilisedBytes: alloc.UtilisedBytes.String(),
		},
	)

	return nil
}
