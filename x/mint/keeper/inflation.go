package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/mint/types"
	"github.com/sentinel-official/sentinelhub/v13/x/mint/types/v1"
)

// SetInflation stores an inflation record in the module's KVStore.
func (k *Keeper) SetInflation(ctx sdk.Context, inflation v1.Inflation) {
	store := k.Store(ctx)
	key := types.InflationKey(inflation.Timestamp)
	value := k.cdc.MustMarshal(&inflation)

	store.Set(key, value)
}

// GetInflation retrieves an inflation record from the module's KVStore based on the timestamp.
// If the inflation record exists, it returns the record and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetInflation(ctx sdk.Context, t time.Time) (inflation v1.Inflation, found bool) {
	store := k.Store(ctx)
	key := types.InflationKey(t)
	value := store.Get(key)

	if value == nil {
		return inflation, false
	}

	k.cdc.MustUnmarshal(value, &inflation)

	return inflation, true
}

// DeleteInflation removes an inflation record from the module's KVStore based on the timestamp.
func (k *Keeper) DeleteInflation(ctx sdk.Context, t time.Time) {
	store := k.Store(ctx)
	key := types.InflationKey(t)

	store.Delete(key)
}

// GetInflations retrieves all inflation records stored in the module's KVStore.
func (k *Keeper) GetInflations(ctx sdk.Context) (items []v1.Inflation) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.InflationKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v1.Inflation
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

// IterateInflations iterates over all inflation records stored in the module's KVStore and calls the provided function for each record.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateInflations(ctx sdk.Context, fn func(index int, item v1.Inflation) (stop bool)) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.InflationKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v1.Inflation
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}

		i++
	}
}
