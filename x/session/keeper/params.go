package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

// SetParams stores the parameters for the module in the KVStore.
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

// MaxGigabytes retrieves the maximum gigabytes parameter from the module's parameters.
func (k *Keeper) MaxGigabytes(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MaxGigabytes
}

// MinGigabytes retrieves the minimum gigabytes parameter from the module's parameters.
func (k *Keeper) MinGigabytes(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MinGigabytes
}

// MaxHours retrieves the maximum hours parameter from the module's parameters.
func (k *Keeper) MaxHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MaxHours
}

// MinHours retrieves the minimum hours parameter from the module's parameters.
func (k *Keeper) MinHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MinHours
}

// ProofVerificationEnabled returns whether proof verification is enabled from the module's parameters.
func (k *Keeper) ProofVerificationEnabled(ctx sdk.Context) bool {
	return k.GetParams(ctx).ProofVerificationEnabled
}

// StakingShare retrieves the staking share parameter from the module's parameters.
func (k *Keeper) StakingShare(ctx sdk.Context) sdkmath.LegacyDec {
	return k.GetParams(ctx).StakingShare
}

// StatusChangeDelay returns the delay for status changes from the module's parameters.
func (k *Keeper) StatusChangeDelay(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).StatusChangeDelay
}

// IsValidGigabytes checks if the provided gigabytes are within the valid range defined by the module's parameters.
func (k *Keeper) IsValidGigabytes(ctx sdk.Context, gigabytes int64) bool {
	if gigabytes > k.MaxGigabytes(ctx) {
		return false
	}
	if gigabytes < k.MinGigabytes(ctx) {
		return false
	}

	return true
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

// GetInactiveAt returns the inactive time by adding StatusChangeDelay to the current block time.
func (k *Keeper) GetInactiveAt(ctx sdk.Context) time.Time {
	d := k.StatusChangeDelay(ctx)
	return ctx.BlockTime().Add(d)
}
