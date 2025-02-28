package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	protorevtypes "github.com/sentinel-official/hub/v12/third_party/osmosis/x/protorev/types"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

// BeginBlock is called at the beginning of each block to trigger IBC query packets for relevant assets.
func (k *Keeper) BeginBlock(ctx sdk.Context) {
	interval := k.GetBlockInterval(ctx)
	if ctx.BlockHeight()%interval != 0 {
		return
	}

	portID := k.GetPortID(ctx)
	if portID == "" {
		k.Logger(ctx).Info("PortID is empty, skipping BeginBlock execution")
		return
	}

	channelID := k.GetChannelID(ctx)
	if channelID == "" {
		k.Logger(ctx).Info("ChannelID is empty, skipping BeginBlock execution")
		return
	}

	timeout := k.GetQueryTimeout(ctx)

	// Get the channel capability to ensure we have the authority to send packets.
	channelCap, found := k.capability.GetCapability(ctx, ibchost.ChannelCapabilityPath(portID, channelID))
	if !found {
		k.Logger(ctx).Info("Channel capability not found, skipping BeginBlock execution")
		return
	}

	// Iterate over each asset and send a ProtoRevPool query for each.
	k.IterateAssets(ctx, func(_ int, item v1.Asset) bool {
		// Create a request for the GetProtoRevPool query using asset details.
		req := abcitypes.RequestQuery{
			Data: k.cdc.MustMarshal(
				&protorevtypes.QueryGetProtoRevPoolRequest{
					BaseDenom:  item.BaseAssetDenom,
					OtherDenom: item.QuoteAssetDenom,
				},
			),
			Path: "/osmosis.protorev.v1beta1.Query/GetProtoRevPool",
		}

		// Send the GetProtoRevPool query packet over IBC.
		sequence, err := k.SendQueryPacket(ctx, channelCap, portID, channelID, uint64(timeout), req)
		if err != nil {
			k.Logger(ctx).Error("Failed to send query packet", "asset", item.Denom, "msg", err)
			return false
		}

		// Map the sequence number to the asset denom for tracking.
		k.SetDenomForPacket(ctx, portID, channelID, sequence, item.Denom)
		return false
	})
}
