package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/provider/types"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

// SetActiveProvider stores an active provider in the module's KVStore.
func (k *Keeper) SetActiveProvider(ctx sdk.Context, v v2.Provider) {
	addr, err := base.ProvAddressFromBech32(v.Address)
	if err != nil {
		panic(err)
	}

	store := k.Store(ctx)
	key := types.ActiveProviderKey(addr)
	value := k.cdc.MustMarshal(&v)

	store.Set(key, value)
}

// HasActiveProvider checks if an active provider exists in the module's KVStore based on the provider address.
func (k *Keeper) HasActiveProvider(ctx sdk.Context, addr base.ProvAddress) bool {
	store := k.Store(ctx)
	key := types.ActiveProviderKey(addr)

	return store.Has(key)
}

// GetActiveProvider retrieves an active provider from the module's KVStore based on the provider address.
// If the provider exists, it returns the provider and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetActiveProvider(ctx sdk.Context, addr base.ProvAddress) (v v2.Provider, found bool) {
	store := k.Store(ctx)
	key := types.ActiveProviderKey(addr)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// DeleteActiveProvider removes an active provider from the module's KVStore based on the provider address.
func (k *Keeper) DeleteActiveProvider(ctx sdk.Context, addr base.ProvAddress) {
	store := k.Store(ctx)
	key := types.ActiveProviderKey(addr)

	store.Delete(key)
}

// SetInactiveProvider stores an inactive provider in the module's KVStore.
func (k *Keeper) SetInactiveProvider(ctx sdk.Context, v v2.Provider) {
	addr, err := base.ProvAddressFromBech32(v.Address)
	if err != nil {
		panic(err)
	}

	store := k.Store(ctx)
	key := types.InactiveProviderKey(addr)
	value := k.cdc.MustMarshal(&v)

	store.Set(key, value)
}

// HasInactiveProvider checks if an inactive provider exists in the module's KVStore based on the provider address.
func (k *Keeper) HasInactiveProvider(ctx sdk.Context, addr base.ProvAddress) bool {
	store := k.Store(ctx)
	key := types.InactiveProviderKey(addr)

	return store.Has(key)
}

// GetInactiveProvider retrieves an inactive provider from the module's KVStore based on the provider address.
// If the provider exists, it returns the provider and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetInactiveProvider(ctx sdk.Context, addr base.ProvAddress) (v v2.Provider, found bool) {
	store := k.Store(ctx)
	key := types.InactiveProviderKey(addr)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// DeleteInactiveProvider removes an inactive provider from the module's KVStore based on the provider address.
func (k *Keeper) DeleteInactiveProvider(ctx sdk.Context, addr base.ProvAddress) {
	store := k.Store(ctx)
	key := types.InactiveProviderKey(addr)

	store.Delete(key)
}

// SetProvider stores a provider in the module's KVStore based on its status.
func (k *Keeper) SetProvider(ctx sdk.Context, provider v2.Provider) {
	switch provider.Status {
	case v1base.StatusActive:
		k.SetActiveProvider(ctx, provider)
	case v1base.StatusInactive:
		k.SetInactiveProvider(ctx, provider)
	default:
		panic(fmt.Errorf("failed to set the provider %v", provider))
	}
}

// HasProvider checks if a provider exists in the module's KVStore based on the provider address.
func (k *Keeper) HasProvider(ctx sdk.Context, addr base.ProvAddress) bool {
	return k.HasActiveProvider(ctx, addr) || k.HasInactiveProvider(ctx, addr)
}

// GetProvider retrieves a provider from the module's KVStore based on the provider address.
// If the provider exists, it returns the provider and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetProvider(ctx sdk.Context, addr base.ProvAddress) (provider v2.Provider, found bool) {
	provider, found = k.GetActiveProvider(ctx, addr)
	if found {
		return provider, true
	}

	provider, found = k.GetInactiveProvider(ctx, addr)
	if found {
		return provider, true
	}

	return provider, false
}

// GetProviders retrieves all providers from the module's KVStore.
func (k *Keeper) GetProviders(ctx sdk.Context) (items v2.Providers) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v2.Provider
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

// IterateProviders iterates over all providers in the module's KVStore and performs the specified action.
func (k *Keeper) IterateProviders(ctx sdk.Context, fn func(index int, item v2.Provider) (stop bool)) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v2.Provider
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}
