package v3

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	DefaultActiveDuration    = 30 * time.Second
	DefaultDeposit           = sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(10))
	DefaultMinGigabytePrices = v1base.Prices{{Denom: sdk.DefaultBondDenom, BaseValue: sdkmath.LegacyZeroDec(), QuoteValue: sdkmath.NewInt(10)}}
	DefaultMinHourlyPrices   = v1base.Prices{{Denom: sdk.DefaultBondDenom, BaseValue: sdkmath.LegacyZeroDec(), QuoteValue: sdkmath.NewInt(10)}}
)

func (m *Params) GetMinGigabytePrices() v1base.Prices {
	return m.MinGigabytePrices
}

func (m *Params) GetMinHourlyPrices() v1base.Prices {
	return m.MinHourlyPrices
}

func (m *Params) Validate() error {
	if err := validateActiveDuration(m.ActiveDuration); err != nil {
		return err
	}
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}
	if err := validateMinGigabytePrices(m.MinGigabytePrices); err != nil {
		return err
	}
	if err := validateMinHourlyPrices(m.MinHourlyPrices); err != nil {
		return err
	}

	return nil
}

func NewParams(
	activeDuration time.Duration, deposit sdk.Coin, minGigabytePrices, minHourlyPrices v1base.Prices,
) Params {
	return Params{
		ActiveDuration:    activeDuration,
		Deposit:           deposit,
		MinGigabytePrices: minGigabytePrices,
		MinHourlyPrices:   minHourlyPrices,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultActiveDuration,
		DefaultDeposit,
		DefaultMinGigabytePrices,
		DefaultMinHourlyPrices,
	)
}

func validateActiveDuration(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("active_duration cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("active_duration cannot be zero")
	}

	return nil
}

func validateDeposit(v interface{}) error {
	value, ok := v.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("deposit cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !value.IsValid() {
		return fmt.Errorf("invalid deposit %s", value)
	}

	return nil
}

func validateMinGigabytePrices(v interface{}) error {
	value, ok := v.([]v1base.Price)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if !v1base.Prices(value).IsValid() {
		return fmt.Errorf("min_gigabyte_prices must be valid")
	}

	return nil
}

func validateMinHourlyPrices(v interface{}) error {
	value, ok := v.([]v1base.Price)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if !v1base.Prices(value).IsValid() {
		return fmt.Errorf("min_hourly_prices must be valid")
	}

	return nil
}
