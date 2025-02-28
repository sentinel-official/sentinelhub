package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

// SetParams stores the node module parameters in the module's KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v3.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the node module parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v3.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

// ActiveDuration retrieves the active duration parameter from the module's parameters.
func (k *Keeper) ActiveDuration(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).ActiveDuration
}

// Deposit retrieves the deposit parameter from the module's parameters.
func (k *Keeper) Deposit(ctx sdk.Context) sdk.Coin {
	return k.GetParams(ctx).Deposit
}

// MinGigabytePrices retrieves the minimum gigabyte prices parameter from the module's parameters.
func (k *Keeper) MinGigabytePrices(ctx sdk.Context) v1base.Prices {
	return k.GetParams(ctx).MinGigabytePrices
}

// MinHourlyPrices retrieves the minimum hourly prices parameter from the module's parameters.
func (k *Keeper) MinHourlyPrices(ctx sdk.Context) v1base.Prices {
	return k.GetParams(ctx).MinHourlyPrices
}

// IsValidGigabytePrices checks if the provided gigabyte prices are valid based on the minimum prices defined in the module's parameters.
func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices v1base.Prices) bool {
	if prices.Len() == 0 {
		return true
	}

	minPrices := k.MinGigabytePrices(ctx)
	for _, price := range minPrices {
		baseValue, quoteValue := prices.AmountOf(price.Denom)
		if !baseValue.IsZero() && baseValue.LT(price.BaseValue) {
			return false
		}
		if !quoteValue.IsZero() && quoteValue.LT(price.QuoteValue) {
			return false
		}
	}

	return true
}

// IsValidHourlyPrices checks if the provided hourly prices are valid based on the minimum prices defined in the module's parameters.
func (k *Keeper) IsValidHourlyPrices(ctx sdk.Context, prices v1base.Prices) bool {
	if prices.Len() == 0 {
		return true
	}

	minPrices := k.MinHourlyPrices(ctx)
	for _, price := range minPrices {
		baseValue, quoteValue := prices.AmountOf(price.Denom)
		if !baseValue.IsZero() && baseValue.LT(price.BaseValue) {
			return false
		}
		if !quoteValue.IsZero() && quoteValue.LT(price.QuoteValue) {
			return false
		}
	}

	return true
}

// GetInactiveAt returns the inactive time by adding ActiveDuration to the current block time.
func (k *Keeper) GetInactiveAt(ctx sdk.Context) time.Time {
	d := k.ActiveDuration(ctx)
	return ctx.BlockTime().Add(d)
}
