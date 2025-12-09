package keeper

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	protorevtypes "github.com/sentinel-official/sentinelhub/v13/third_party/osmosis/x/protorev/types"
	"github.com/sentinel-official/sentinelhub/v13/x/oracle/types/v1"
)

// BeginBlock is called at the beginning of each block to trigger IBC query packets for relevant assets.
func (k *Keeper) BeginBlock(c context.Context) error {
	ctx := sdk.UnwrapSDKContext(c)

	interval := k.GetBlockInterval(ctx)
	if ctx.BlockHeight()%interval != 0 {
		return nil
	}

	portID := k.GetPortID(ctx)
	if portID == "" {
		k.Logger(ctx).Info("PortID is empty, skipping BeginBlock execution")

		return nil
	}

	channelID := k.GetChannelID(ctx)
	if channelID == "" {
		k.Logger(ctx).Info("ChannelID is empty, skipping BeginBlock execution")

		return nil
	}

	timeoutTimestamp := k.GetTimeoutTimestamp(ctx)

	// Iterate over each asset and send a ProtoRevPool query for each.
	k.IterateAssets(ctx, func(_ int, item v1.Asset) bool {
		// Create a request for the GetProtoRevPool query using asset details.
		req := abcitypes.RequestQuery{
			Data: k.cdc.MustMarshal(
				&protorevtypes.QueryGetProtoRevPoolRequest{
					BaseDenom:  item.ProtoRevPoolRequest.BaseDenom,
					OtherDenom: item.ProtoRevPoolRequest.OtherDenom,
				},
			),
			Path: "/osmosis.protorev.v1beta1.Query/GetProtoRevPool",
		}

		// Send the GetProtoRevPool query packet over IBC.
		sequence, err := k.SendQueryPacket(ctx, portID, channelID, timeoutTimestamp, req)
		if err != nil {
			k.Logger(ctx).Error("Failed to send query packet", "asset", item.Denom, "msg", err)

			return false
		}

		// Map the sequence number to the asset denom for tracking.
		k.SetDenomForPacket(ctx, portID, channelID, sequence, item.Denom)

		return false
	})

	return nil
}
