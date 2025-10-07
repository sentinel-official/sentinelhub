package v1

import (
	"errors"
	"fmt"
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

	if a.ProtoRevPoolRequest.BaseDenom == "" {
		return errors.New("base_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(a.ProtoRevPoolRequest.BaseDenom); err != nil {
		return fmt.Errorf("invalid base_denom: %w", err)
	}

	if a.ProtoRevPoolRequest.OtherDenom == "" {
		return errors.New("other_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(a.ProtoRevPoolRequest.OtherDenom); err != nil {
		return fmt.Errorf("invalid other_denom: %w", err)
	}

	if a.ProtoRevPoolRequest.BaseDenom == a.ProtoRevPoolRequest.OtherDenom {
		return errors.New("base_denom and other_denom cannot be the same")
	}

	if a.SpotPriceRequest.BaseAssetDenom == "" {
		return errors.New("base_asset_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(a.SpotPriceRequest.BaseAssetDenom); err != nil {
		return fmt.Errorf("invalid base_asset_denom: %w", err)
	}

	if a.SpotPriceRequest.QuoteAssetDenom == "" {
		return errors.New("quote_asset_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(a.SpotPriceRequest.QuoteAssetDenom); err != nil {
		return fmt.Errorf("invalid quote_asset_denom: %w", err)
	}

	if a.SpotPriceRequest.BaseAssetDenom == a.SpotPriceRequest.QuoteAssetDenom {
		return errors.New("base_asset_denom and quote_asset_denom cannot be the same")
	}

	if a.Height < 0 {
		return errors.New("height cannot be negative")
	}

	if a.SpotPrice.IsNil() {
		return errors.New("spot_price cannot be nil")
	}

	if a.SpotPrice.IsNegative() {
		return errors.New("spot_price cannot be negative")
	}

	return nil
}
