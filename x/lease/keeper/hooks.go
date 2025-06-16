package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types/v1"
)

// NodeInactivePreHook is triggered when a node becomes inactive. It ends all leases associated with the specified node address.
func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	k.Logger(ctx).Info("Running node inactive pre-hook", "address", addr.String())

	// Iterate through all leases associated with the given node address
	return k.IterateLeasesForNode(ctx, addr, func(_ int, item v1.Lease) (bool, error) {
		// Create a message to end the lease
		msg := &v1.MsgEndLeaseRequest{
			From: item.ProvAddress,
			ID:   item.ID,
		}

		// Get the appropriate handler for processing the end lease message
		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		// Execute the handler to process the lease termination request
		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		// Emit any events generated during the lease termination process
		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false, nil
	})
}

// ProviderInactivePreHook is triggered when a provider becomes inactive. It ends all leases associated with the specified provider address.
func (k *Keeper) ProviderInactivePreHook(ctx sdk.Context, addr base.ProvAddress) error {
	k.Logger(ctx).Info("Running provider inactive pre-hook", "address", addr.String())

	// Iterate through all leases associated with the given provider address
	return k.IterateLeasesForProvider(ctx, addr, func(_ int, item v1.Lease) (bool, error) {
		// Create a message to end the lease
		msg := &v1.MsgEndLeaseRequest{
			From: item.ProvAddress,
			ID:   item.ID,
		}

		// Get the appropriate handler for processing the end lease message
		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		// Execute the handler to process the lease termination request
		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		// Emit any events generated during the lease termination process
		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false, nil
	})
}
