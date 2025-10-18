package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/subscription/types"
	"github.com/sentinel-official/sentinelhub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/sentinelhub/v12/x/subscription/types/v3"
)

// SetParams stores the given parameters in the module's KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v3.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v3.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)

	return v
}

// MaxAllocations retrieves the max allocations parameter from the module's parameters.
func (k *Keeper) MaxAllocations(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MaxAllocations
}

// StakingShare retrieves the staking share parameter from the module's parameters.
func (k *Keeper) StakingShare(ctx sdk.Context) sdkmath.LegacyDec {
	return k.GetParams(ctx).StakingShare
}

// StatusTimeout returns the status timeout parameter from the module's parameters.
func (k *Keeper) StatusTimeout(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).StatusTimeout
}

// GetInactiveAt returns the inactive time by adding StatusTimeout to the current block time.
func (k *Keeper) GetInactiveAt(ctx sdk.Context) time.Time {
	d := k.StatusTimeout(ctx)

	return ctx.BlockTime().Add(d)
}

// IsMaxAllocationsReached checks if the maximum number of allocations for a given subscription has been reached.
func (k *Keeper) IsMaxAllocationsReached(ctx sdk.Context, id uint64) bool {
	maxAllocations := k.MaxAllocations(ctx)

	var count int64

	k.IterateAllocationsForSubscription(ctx, id, func(_ int, _ v2.Allocation) bool {
		count++
		if count >= maxAllocations {
			return true // Stop iterating when maxAllocations is reached
		}

		return false
	})

	return count >= maxAllocations
}
