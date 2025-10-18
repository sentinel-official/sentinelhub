package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types/v3"
)

// LeaseInactivePreHook handles the necessary operations when a lease becomes inactive.
func (k *Keeper) LeaseInactivePreHook(ctx sdk.Context, id uint64) error {
	k.Logger(ctx).Info("Running lease inactive pre-hook", "id", id)

	// Retrieve the lease by ID and check if it exists, return an error if not found.
	lease, found := k.GetLease(ctx, id)
	if !found {
		return types.NewErrorLeaseNotFound(id)
	}

	// Convert the node address and provider address from Bech32 format.
	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return err
	}

	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return err
	}

	// Iterate through plans linked to the node by the provider, and send unlink requests for the node.
	return k.IteratePlansForNodeByProvider(ctx, nodeAddr, provAddr, func(_ int, item v3.Plan) (bool, error) {
		msg := &v3.MsgUnlinkNodeRequest{
			From:        item.ProvAddress,
			ID:          item.ID,
			NodeAddress: lease.NodeAddress,
		}

		// Retrieve the handler for the unlink message and check if it exists, return an error if nil.
		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		// Execute the handler to process the unlink request.
		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		// Emit any events generated during the unlink process.
		ctx.EventManager().EmitEvents(resp.GetEvents())

		return false, nil
	})
}

// ProviderInactivePreHook handles the necessary operations when a provider becomes inactive.
func (k *Keeper) ProviderInactivePreHook(ctx sdk.Context, addr base.ProvAddress) error {
	k.Logger(ctx).Info("Running provider inactive pre-hook", "address", addr.String())

	// Iterate through all plans associated with the given provider address.
	return k.IteratePlansForProvider(ctx, addr, func(_ int, item v3.Plan) (bool, error) {
		// Check if the plan status is active; if not, skip to the next plan.
		if !item.Status.Equal(v1base.StatusActive) {
			return false, nil
		}

		// Create a message to update the plan status to inactive.
		msg := &v3.MsgUpdatePlanStatusRequest{
			From:   item.ProvAddress,
			ID:     item.ID,
			Status: v1base.StatusInactive,
		}

		// Retrieve the handler for the update status message and check if it exists, return an error if nil.
		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		// Execute the handler to process the status update request.
		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		// Emit any events generated during the status update process.
		ctx.EventManager().EmitEvents(resp.GetEvents())

		return false, nil
	})
}
