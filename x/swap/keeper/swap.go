package keeper

import (
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/swap/types"
	"github.com/sentinel-official/sentinelhub/v13/x/swap/types/v1"
)

func (k *Keeper) SetSwap(ctx sdk.Context, swap v1.Swap) {
	key := types.SwapKey(swap.GetTxHash())
	value := k.cdc.MustMarshal(&swap)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetSwap(ctx sdk.Context, txHash types.EthereumHash) (swap v1.Swap, found bool) {
	store := k.Store(ctx)

	key := types.SwapKey(txHash)

	value := store.Get(key)
	if value == nil {
		return swap, false
	}

	k.cdc.MustUnmarshal(value, &swap)

	return swap, true
}

func (k *Keeper) HasSwap(ctx sdk.Context, txHash types.EthereumHash) bool {
	key := types.SwapKey(txHash)

	store := k.Store(ctx)

	return store.Has(key)
}

func (k *Keeper) GetSwaps(ctx sdk.Context) (items v1.Swaps) {
	var (
		store    = k.Store(ctx)
		iterator = storetypes.KVStorePrefixIterator(store, types.SwapKeyPrefix)
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v1.Swap
		k.cdc.MustUnmarshal(iterator.Value(), &item)
		items = append(items, item)
	}

	return items
}
