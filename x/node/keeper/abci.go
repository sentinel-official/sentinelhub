package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
)

// handleInactiveNodes processes nodes that have become inactive at the current block time.
func (k *Keeper) handleInactiveNodes(ctx sdk.Context) {
	// Iterate through nodes that have become inactive at the current block time
	k.IterateNodesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Node) bool {
		k.Logger(ctx).Info("Handling inactive node", "address", item.Address)

		// Create a message to update the status of the node to inactive
		msg := &v3.MsgUpdateNodeStatusRequest{
			From:   item.Address,
			Status: v1base.StatusInactive,
		}

		// Get the appropriate handler for processing the status update message
		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		// Execute the handler to process the node status update request
		resp, err := handler(ctx, msg)
		if err != nil {
			panic(err)
		}

		// Emit any events generated during the node status update
		ctx.EventManager().EmitEvents(resp.GetEvents())

		return false
	})
}

// BeginBlock is called at the beginning of each block to handle node-related operations.
func (k *Keeper) BeginBlock(ctx sdk.Context) {
	// Handle nodes that have become inactive at the beginning of each block
	k.handleInactiveNodes(ctx)
}
