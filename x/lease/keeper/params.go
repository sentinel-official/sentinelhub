package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

// SetParams stores the lease module parameters in the module's KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the lease module parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

// MaxHours retrieves the maximum hours parameter from the module's parameters.
func (k *Keeper) MaxHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MaxHours
}

// MinHours retrieves the minimum hours parameter from the module's parameters.
func (k *Keeper) MinHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MinHours
}

// StakingShare retrieves the staking share parameter from the module's parameters.
func (k *Keeper) StakingShare(ctx sdk.Context) sdkmath.LegacyDec {
	return k.GetParams(ctx).StakingShare
}

// IsValidHours checks if the provided hours are within the valid range defined by the module's parameters.
func (k *Keeper) IsValidHours(ctx sdk.Context, hours int64) bool {
	if hours > k.MaxHours(ctx) {
		return false
	}
	if hours < k.MinHours(ctx) {
		return false
	}

	return true
}
