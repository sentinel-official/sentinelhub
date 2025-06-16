package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
)

// SetAsset stores an asset in the module's KVStore.
func (k *Keeper) SetAsset(ctx sdk.Context, asset v1.Asset) {
	store := k.Store(ctx)
	key := types.AssetKey(asset.Denom)
	value := k.cdc.MustMarshal(&asset)

	store.Set(key, value)
}

// GetAsset retrieves an asset from the module's KVStore based on the asset denomination.
// If the asset exists, it returns the asset and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetAsset(ctx sdk.Context, denom string) (v v1.Asset, found bool) {
	store := k.Store(ctx)
	key := types.AssetKey(denom)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// HasAsset checks if an asset with the given denomination exists in the module's KVStore.
func (k *Keeper) HasAsset(ctx sdk.Context, denom string) bool {
	store := k.Store(ctx)
	key := types.AssetKey(denom)

	return store.Has(key)
}

// DeleteAsset removes an asset from the module's KVStore based on the asset denomination.
func (k *Keeper) DeleteAsset(ctx sdk.Context, denom string) {
	store := k.Store(ctx)
	key := types.AssetKey(denom)

	store.Delete(key)
}

// IterateAssets iterates over all assets stored in the module's KVStore and calls the provided function for each asset.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateAssets(ctx sdk.Context, fn func(int, v1.Asset) bool) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.AssetKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v1.Asset
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetDenomForPacket stores a denomination for a packet in the module's KVStore.
func (k *Keeper) SetDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64, denom string) {
	store := k.Store(ctx)
	key := ibchost.PacketCommitmentKey(portID, channelID, sequence)
	value := k.cdc.MustMarshal(&protobuf.StringValue{Value: denom})

	store.Set(key, value)
}

// GetDenomForPacket retrieves a denomination for a packet from the module's KVStore based on the port ID, channel ID, and sequence.
// If the denomination exists, it returns the denomination and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) (v string, found bool) {
	store := k.Store(ctx)
	key := ibchost.PacketCommitmentKey(portID, channelID, sequence)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	var denom protobuf.StringValue
	k.cdc.MustUnmarshal(value, &denom)

	return denom.GetValue(), true
}

// DeleteDenomForPacket removes a denomination for a packet from the module's KVStore based on the port ID, channel ID, and sequence.
func (k *Keeper) DeleteDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) {
	store := k.Store(ctx)
	key := ibchost.PacketCommitmentKey(portID, channelID, sequence)

	store.Delete(key)
}

// GetAssetForPacket retrieves an asset for a packet from the module's KVStore based on the port ID, channel ID, and sequence.
// If the asset exists, it returns the asset; otherwise, it returns an error.
func (k *Keeper) GetAssetForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) (v v1.Asset, err error) {
	denom, found := k.GetDenomForPacket(ctx, portID, channelID, sequence)
	if !found {
		return v, types.NewErrorDenomForPacketNotFound(portID, channelID, sequence)
	}

	v, found = k.GetAsset(ctx, denom)
	if !found {
		return v, types.NewErrorAssetNotFound(denom)
	}

	return v, nil
}

// GetQuotePrice retrieves a quote for a given DecCoin, applying asset pricing if available.
// If the asset for the given denomination is not found, it returns an error.
func (k *Keeper) GetQuotePrice(ctx sdk.Context, basePrice sdk.DecCoin) (sdk.Coin, error) {
	// Retrieve the asset associated with the coin denomination.
	asset, found := k.GetAsset(ctx, basePrice.Denom)
	if !found {
		return sdk.Coin{Amount: sdk.ZeroInt()}, types.NewErrorAssetNotFound(basePrice.Denom)
	}

	// Calculate the quote price by multiplying the base price with the asset's price.
	amount := basePrice.Amount.MulInt(asset.Price).TruncateInt()

	// Return the resulting coin with the same denomination as the base price.
	return sdk.Coin{Denom: basePrice.Denom, Amount: amount}, nil
}
