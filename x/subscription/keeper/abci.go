package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

// handleInactivePendingSubscriptions processes pending subscriptions that have become inactive.
func (k *Keeper) handleInactivePendingSubscriptions(ctx sdk.Context) {
	// Iterate through subscriptions that have become inactive at the current block time
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		// Skip the subscription if its status is not active
		if !item.Status.Equal(v1base.StatusActive) {
			return false
		}

		k.Logger(ctx).Info("Handling inactive pending subscription", "id", item.ID)

		// Create a message to cancel the inactive pending subscription
		msg := &v3.MsgCancelSubscriptionRequest{
			From: item.AccAddress,
			ID:   item.ID,
		}

		// Get the appropriate handler for processing the cancel subscription message
		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		// Execute the handler to process the cancel request
		resp, err := handler(ctx, msg)
		if err != nil {
			panic(err)
		}

		// Emit any events generated during the cancel process
		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

// handleInactiveSubscriptions processes subscriptions that are in the inactive pending state.
func (k *Keeper) handleInactiveSubscriptions(ctx sdk.Context) {
	// Iterate through subscriptions that are inactive pending at the current block time
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		// Skip the subscription if its status is not inactive pending
		if !item.Status.Equal(v1base.StatusInactivePending) {
			return false
		}

		k.Logger(ctx).Info("Handling inactive subscription", "id", item.ID)

		// Delete the subscription from the state
		k.DeleteSubscription(ctx, item.ID)
		k.DeleteSubscriptionForPlan(ctx, item.PlanID, item.ID)

		// Iterate through all allocations for the subscription and delete them
		k.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) bool {
			accAddr, err := sdk.AccAddressFromBech32(item.Address)
			if err != nil {
				panic(err)
			}

			// Delete allocation for the given subscription and account
			k.DeleteAllocation(ctx, item.ID, accAddr)
			k.DeleteSubscriptionForAccount(ctx, accAddr, item.ID)
			return false
		})

		// Delete the subscription from the inactive queue
		k.DeleteSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

		// Emit an event indicating the update of the subscription status to inactive
		ctx.EventManager().EmitTypedEvent(
			&v3.EventUpdate{
				ID:         item.ID,
				PlanID:     item.PlanID,
				AccAddress: item.AccAddress,
				Status:     v1base.StatusInactive,
				StatusAt:   ctx.BlockTime().String(),
			},
		)

		return false
	})
}

// handleSubscriptionRenewals processes subscription renewals that are due at the current block time.
func (k *Keeper) handleSubscriptionRenewals(ctx sdk.Context) {
	// Iterate through subscriptions that are due for renewal at the current block time
	k.IterateSubscriptionsForRenewalAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		k.Logger(ctx).Info("Handling subscription renewal", "id", item.ID)

		// Create a message to renew the subscription
		msg := &v3.MsgRenewSubscriptionRequest{
			From:  item.AccAddress,
			ID:    item.ID,
			Denom: item.Price.Denom,
		}

		// Get the appropriate handler for processing the renewal message
		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		// Create a cache context to prevent state changes until renewal is successful
		cc, write := ctx.CacheContext()

		// Execute the handler to process the subscription renewal
		resp, err := handler(cc, msg)
		if err != nil {
			k.Logger(cc).Error("Failed to handle subscription renewal", "id", item.ID, "msg", err)
			return false
		}

		// Commit the changes to the main context if renewal is successful
		defer write()

		// Emit any events generated during the renewal process
		cc.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

// BeginBlock is called at the beginning of each block to handle subscription-related operations.
func (k *Keeper) BeginBlock(ctx sdk.Context) {
	// Handle subscription renewals at the beginning of each block
	k.handleSubscriptionRenewals(ctx)

	// Handle inactive pending subscriptions at the beginning of each block
	k.handleInactivePendingSubscriptions(ctx)

	// Handle subscriptions that are inactive at the beginning of each block
	k.handleInactiveSubscriptions(ctx)
}
