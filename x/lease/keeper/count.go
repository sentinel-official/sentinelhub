package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/sentinelhub/v12/x/lease/types"
)

// SetLeaseCount stores the lease count in the module's KVStore.
func (k *Keeper) SetLeaseCount(ctx sdk.Context, count uint64) {
	store := k.Store(ctx)
	key := types.CountKey
	value := k.cdc.MustMarshal(&protobuf.UInt64Value{Value: count})

	store.Set(key, value)
}

// GetLeaseCount retrieves the lease count from the module's KVStore.
// If the count is not found, it returns 0.
func (k *Keeper) GetLeaseCount(ctx sdk.Context) uint64 {
	store := k.Store(ctx)
	key := types.CountKey
	value := store.Get(key)

	if value == nil {
		return 0
	}

	var count protobuf.UInt64Value
	k.cdc.MustUnmarshal(value, &count)

	return count.GetValue()
}
