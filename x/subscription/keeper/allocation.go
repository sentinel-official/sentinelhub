package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

// SetAllocation sets the allocation for a given ID and address in the KVStore.
func (k *Keeper) SetAllocation(ctx sdk.Context, alloc v2.Allocation) {
	addr, err := sdk.AccAddressFromBech32(alloc.Address)
	if err != nil {
		panic(err)
	}

	store := k.Store(ctx)
	key := types.AllocationKey(alloc.ID, addr)
	value := k.cdc.MustMarshal(&alloc)

	store.Set(key, value)
}

// GetAllocation retrieves the allocation for a given ID and address from the KVStore.
// It returns the allocation and a boolean indicating if the allocation was found.
func (k *Keeper) GetAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) (alloc v2.Allocation, found bool) {
	store := k.Store(ctx)
	key := types.AllocationKey(id, addr)
	value := store.Get(key)

	if value == nil {
		return alloc, false
	}

	k.cdc.MustUnmarshal(value, &alloc)
	return alloc, true
}

// HasAllocation checks if an allocation for a given ID and address exists in the KVStore.
func (k *Keeper) HasAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) bool {
	store := k.Store(ctx)
	key := types.AllocationKey(id, addr)

	return store.Has(key)
}

// DeleteAllocation removes the allocation for a given ID and address from the KVStore.
func (k *Keeper) DeleteAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) {
	store := k.Store(ctx)
	key := types.AllocationKey(id, addr)

	store.Delete(key)
}

// GetAllocationsForSubscription retrieves all allocations for a given subscription ID.
func (k *Keeper) GetAllocationsForSubscription(ctx sdk.Context, id uint64) (items v2.Allocations) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetAllocationForSubscriptionKeyPrefix(id))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v2.Allocation
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

// IterateAllocationsForSubscription iterates over all allocations for a given subscription ID,
// calling the provided function for each allocation.
func (k *Keeper) IterateAllocationsForSubscription(ctx sdk.Context, id uint64, fn func(index int, item v2.Allocation) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStorePrefixIterator(store, types.GetAllocationForSubscriptionKeyPrefix(id))
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v2.Allocation
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}
