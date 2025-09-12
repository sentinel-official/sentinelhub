package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibcicqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v7/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	"github.com/sentinel-official/sentinelhub/v12/third_party/osmosis/x/poolmanager/client/queryproto"
	protorevtypes "github.com/sentinel-official/sentinelhub/v12/third_party/osmosis/x/protorev/types"
	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
)

// SendQueryPacket serializes query requests and sends them as an IBC packet to a destination chain.
func (k *Keeper) SendQueryPacket(
	ctx sdk.Context, channelCap *capabilitytypes.Capability, portID, channelID string, timeout uint64,
	reqs ...abcitypes.RequestQuery,
) (uint64, error) {
	// Serialize the Cosmos query requests into binary format.
	data, err := ibcicqtypes.SerializeCosmosQuery(reqs)
	if err != nil {
		return 0, err
	}

	// Create packet data with the serialized queries and validate it.
	packetData := ibcicqtypes.InterchainQueryPacketData{Data: data}
	if err := packetData.ValidateBasic(); err != nil {
		return 0, err
	}

	// Use the ICS-04 interface to send the packet over IBC.
	return k.ics4.SendPacket(
		ctx, channelCap, portID, channelID, ibcclienttypes.ZeroHeight(), timeout, packetData.GetBytes(),
	)
}

// OnAcknowledgementPacket processes the acknowledgement packet received after sending an IBC query packet.
func (k *Keeper) OnAcknowledgementPacket(
	ctx sdk.Context, packet ibcchanneltypes.Packet, ack ibcchanneltypes.Acknowledgement,
) error {
	// Retrieve the source port, channel, and sequence number from the packet.
	portID := packet.GetSourcePort()
	channelID := packet.GetSourceChannel()
	sequence := packet.GetSequence()

	// Ensure the denom mapping for the packet is deleted after processing the acknowledgement.
	defer k.DeleteDenomForPacket(ctx, portID, channelID, sequence)

	// If the acknowledgement indicates failure, there's no further processing required.
	if !ack.Success() {
		return nil
	}

	// Unmarshal the packet data to get the interchain query packet details.
	var packetData ibcicqtypes.InterchainQueryPacketData
	if err := k.cdc.UnmarshalJSON(packet.GetData(), &packetData); err != nil {
		return err
	}

	// Deserialize the Cosmos queries from the packet data.
	reqs, err := ibcicqtypes.DeserializeCosmosQuery(packetData.Data)
	if err != nil {
		return err
	}

	// Unmarshal the acknowledgement result to obtain the query responses.
	var packetAck ibcicqtypes.InterchainQueryPacketAck
	if err := k.cdc.UnmarshalJSON(ack.GetResult(), &packetAck); err != nil {
		return err
	}

	// Deserialize the Cosmos responses from the acknowledgement data.
	resps, err := ibcicqtypes.DeserializeCosmosResponse(packetAck.Data)
	if err != nil {
		return err
	}

	// Verify that the number of responses matches the number of requests.
	if len(reqs) != len(resps) {
		return fmt.Errorf("invalid response count %d; expected %d", len(resps), len(reqs))
	}

	// Retrieve the asset associated with the packet using port, channel, and sequence information.
	asset, err := k.GetAssetForPacket(ctx, portID, channelID, sequence)
	if err != nil {
		return err
	}

	// Iterate through each request-response pair and update the asset accordingly.
	for i := range reqs {
		// Handle specific query paths to extract the required data and update the asset.
		switch reqs[i].Path {
		case "/osmosis.poolmanager.v1beta1.Query/SpotPrice":
			if err := k.handleSpotPriceQueryResponse(ctx, asset, &resps[i]); err != nil {
				return err
			}
		case "/osmosis.protorev.v1beta1.Query/GetProtoRevPool":
			if err := k.handleProtoRevPoolQueryResponse(ctx, asset, &resps[i]); err != nil {
				return err
			}
		}
	}

	return nil
}

// handleSpotPriceQueryResponse handles the response for the SpotPrice query.
func (k *Keeper) handleSpotPriceQueryResponse(ctx sdk.Context, asset v1.Asset, resp *abcitypes.ResponseQuery) error {
	// Skip updates if the response height is older than the current asset height.
	if resp.GetHeight() < asset.Height {
		return nil
	}

	// Unmarshal the response data to extract the spot price details.
	var res queryproto.SpotPriceResponse
	if err := k.cdc.Unmarshal(resp.GetValue(), &res); err != nil {
		return err
	}

	// Convert the spot price to a decimal value.
	spotPrice, err := sdkmath.LegacyNewDecFromStr(res.GetSpotPrice())
	if err != nil {
		return err
	}

	// Update the asset price using the spot price and its multiplier.
	asset.Price = spotPrice.MulInt(asset.Multiplier()).TruncateInt()
	asset.Height = resp.GetHeight()

	// Persist the updated asset information in the store.
	k.SetAsset(ctx, asset)

	return nil
}

// handleProtoRevPoolQueryResponse handles the response for the GetProtoRevPool query.
func (k *Keeper) handleProtoRevPoolQueryResponse(ctx sdk.Context, asset v1.Asset, resp *abcitypes.ResponseQuery) error {
	// Unmarshal the response to extract the pool ID and other relevant details.
	var res protorevtypes.QueryGetProtoRevPoolResponse
	if err := k.cdc.Unmarshal(resp.GetValue(), &res); err != nil {
		return err
	}

	// Retrieve necessary information for sending a follow-up query.
	portID := k.GetPortID(ctx)
	channelID := k.GetChannelID(ctx)
	timeout := k.GetQueryTimeout(ctx)

	// Get the channel capability to ensure we have the authority to send packets.
	channelCap, found := k.capability.GetCapability(ctx, ibchost.ChannelCapabilityPath(portID, channelID))
	if !found {
		return nil
	}

	// Create a new request for the SpotPrice query using the pool ID and asset details.
	req := abcitypes.RequestQuery{
		Data: k.cdc.MustMarshal(
			&queryproto.SpotPriceRequest{
				PoolId:          res.GetPoolId(),
				BaseAssetDenom:  asset.BaseAssetDenom,
				QuoteAssetDenom: asset.QuoteAssetDenom,
			},
		),
		Path: "/osmosis.poolmanager.v1beta1.Query/SpotPrice",
	}

	// Send the SpotPrice query packet over IBC.
	sequence, err := k.SendQueryPacket(ctx, channelCap, portID, channelID, uint64(timeout), req)
	if err != nil {
		return err
	}

	// Map the sequence number to the asset denom for tracking.
	k.SetDenomForPacket(ctx, portID, channelID, sequence, asset.Denom)

	return nil
}

// OnTimeoutPacket handles the case when a packet times out before receiving an acknowledgement.
func (k *Keeper) OnTimeoutPacket(ctx sdk.Context, packet ibcchanneltypes.Packet) error {
	// Retrieve the source port, channel, and sequence number from the packet.
	portID := packet.GetSourcePort()
	channelID := packet.GetSourceChannel()
	sequence := packet.GetSequence()

	// Delete the denom mapping associated with the timed-out packet.
	k.DeleteDenomForPacket(ctx, portID, channelID, sequence)

	return nil
}
