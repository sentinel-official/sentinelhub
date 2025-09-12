package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

// NodeInactivePreHook handles the necessary operations when a node becomes inactive.
func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	k.Logger(ctx).Info("Running node inactive pre-hook", "address", addr.String())

	// Iterate through all active sessions associated with the given node address.
	return k.IterateSessionsForNode(ctx, addr, func(_ int, item v3.Session) (bool, error) {
		// Skip the session if it is not active.
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false, nil
		}

		// Create a message to cancel the active session.
		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		// Retrieve the handler for the cancel session message, and return an error if it is nil.
		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		// Execute the handler to process the session cancellation.
		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		// Emit any events generated during the session cancellation process.
		ctx.EventManager().EmitEvents(resp.GetEvents())

		return false, nil
	})
}

// SubscriptionInactivePendingPreHook handles the necessary operations when a subscription becomes inactive pending.
func (k *Keeper) SubscriptionInactivePendingPreHook(ctx sdk.Context, id uint64) error {
	k.Logger(ctx).Info("Running subscription inactive pending pre-hook", "id", id)

	// Iterate through all active sessions associated with the given subscription ID.
	return k.IterateSessionsForSubscription(ctx, id, func(_ int, item v3.Session) (bool, error) {
		// Skip the session if it is not active.
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false, nil
		}

		// Create a message to cancel the active session.
		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		// Retrieve the handler for the cancel session message, and return an error if it is nil.
		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		// Execute the handler to process the session cancellation.
		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		// Emit any events generated during the session cancellation process.
		ctx.EventManager().EmitEvents(resp.GetEvents())

		return false, nil
	})
}

// PlanUnlinkNodePreHook handles the necessary operations when unlinking a node from a plan.
func (k *Keeper) PlanUnlinkNodePreHook(ctx sdk.Context, id uint64, addr base.NodeAddress) error {
	k.Logger(ctx).Info("Running plan unlink node pre-hook", "id", id, "address", addr.String())

	// Iterate through all active sessions associated with the given plan ID and node address.
	return k.IterateSessionsForPlanByNode(ctx, id, addr, func(_ int, item v3.Session) (bool, error) {
		// Skip the session if it is not active.
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false, nil
		}

		// Create a message to cancel the active session.
		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		// Retrieve the handler for the cancel session message, and return an error if it is nil.
		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		// Execute the handler to process the session cancellation.
		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		// Emit any events generated during the session cancellation process.
		ctx.EventManager().EmitEvents(resp.GetEvents())

		return false, nil
	})
}
