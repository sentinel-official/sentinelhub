package v1

import (
	"errors"
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Multiplier returns 10^Decimals as an integer multiplier for the asset.
func (a *Asset) Multiplier() sdkmath.Int {
	i := new(big.Int).Exp(big.NewInt(10), big.NewInt(a.Decimals), nil)
	return sdkmath.NewIntFromBigInt(i)
}

// Validate checks if the Asset fields are valid.
func (a *Asset) Validate() error {
	if a.Denom == "" {
		return errors.New("denom cannot be empty")
	}
	if err := sdk.ValidateDenom(a.Denom); err != nil {
		return err
	}
	if a.Decimals < 0 {
		return errors.New("decimals cannot be negative")
	}
	if a.BaseAssetDenom == "" {
		return errors.New("base_asset_denom cannot be empty")
	}
	if a.QuoteAssetDenom == "" {
		return errors.New("quote_asset_denom cannot be empty")
	}
	if a.BaseAssetDenom == a.QuoteAssetDenom {
		return errors.New("base_asset_denom and quote_asset_denom cannot be same")
	}
	if a.Price.IsNil() {
		return errors.New("price cannot be nil")
	}
	if a.Price.IsNegative() {
		return errors.New("price cannot be negative")
	}
	if a.Price.IsZero() {
		return errors.New("price cannot be zero")
	}
	if a.Height < 0 {
		return errors.New("height cannot be negative")
	}

	return nil
}
