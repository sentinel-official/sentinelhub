package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
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

// StatusTimeout retrieves the status timeout parameter from the module's parameters.
func (k *Keeper) StatusTimeout(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).StatusTimeout
}

// ValidateGigabytePrices validates prices against the minimum gigabyte prices.
func (k *Keeper) ValidateGigabytePrices(ctx sdk.Context, prices v1base.Prices) error {
	minPrices := k.MinGigabytePrices(ctx)

	return validatePrices(prices, minPrices)
}

// ValidateHourlyPrices validates prices against the minimum hourly prices.
func (k *Keeper) ValidateHourlyPrices(ctx sdk.Context, prices v1base.Prices) error {
	minPrices := k.MinHourlyPrices(ctx)

	return validatePrices(prices, minPrices)
}

// GetInactiveAt returns the inactive time by adding StatusTimeout to the current block time.
func (k *Keeper) GetInactiveAt(ctx sdk.Context) time.Time {
	d := k.StatusTimeout(ctx)

	return ctx.BlockTime().Add(d)
}

// validatePrices checks that all provided prices use allowed denoms and
// meet or exceed the given minimums (zero values and empty lists are valid).
func validatePrices(prices v1base.Prices, minPrices v1base.Prices) error {
	// Empty set of prices is valid: treated as zero prices for all denoms.
	if prices.Len() == 0 {
		return nil
	}

	// Build a lookup map for minPrices
	m := minPrices.Map()

	for _, price := range prices {
		minPrice, ok := m[price.Denom]
		if !ok {
			return fmt.Errorf("denom %s is not allowed", price.Denom)
		}

		// Only check non-zero values
		if !price.BaseValue.IsZero() && price.BaseValue.LT(minPrice.BaseValue) {
			return fmt.Errorf("base value %s for denom %s is lesser than %s",
				price.BaseValue, price.Denom, minPrice.BaseValue)
		}

		if !price.QuoteValue.IsZero() && price.QuoteValue.LT(minPrice.QuoteValue) {
			return fmt.Errorf("quote value %s for denom %s is lesser than %s",
				price.QuoteValue, price.Denom, minPrice.QuoteValue)
		}
	}

	return nil
}
