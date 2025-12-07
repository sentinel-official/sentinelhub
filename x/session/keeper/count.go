package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/sentinelhub/v13/x/session/types"
)

// SetSessionCount sets the count value in the KVStore.
func (k *Keeper) SetSessionCount(ctx sdk.Context, count uint64) {
	store := k.Store(ctx)
	key := types.CountKey
	value := k.cdc.MustMarshal(&protobuf.UInt64Value{Value: count})

	store.Set(key, value)
}

// GetSessionCount retrieves the count value from the KVStore.
// If the count value does not exist, it returns 0 as the default.
func (k *Keeper) GetSessionCount(ctx sdk.Context) uint64 {
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
