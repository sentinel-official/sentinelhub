package v3

import (
	"errors"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

var (
	// DefaultActiveDuration defines the default duration a node remains active after registration.
	DefaultActiveDuration = 30 * time.Second

	// DefaultDeposit defines the default amount of deposit required for a node.
	DefaultDeposit = sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(10))

	// DefaultMinGigabytePrices defines the default minimum price per gigabyte a node can charge.
	DefaultMinGigabytePrices = v1base.Prices{
		{
			Denom:      sdk.DefaultBondDenom,
			BaseValue:  sdkmath.LegacyZeroDec(),
			QuoteValue: sdkmath.NewInt(10),
		},
	}

	// DefaultMinHourlyPrices defines the default minimum price per hour a node can charge.
	DefaultMinHourlyPrices = v1base.Prices{
		{
			Denom:      sdk.DefaultBondDenom,
			BaseValue:  sdkmath.LegacyZeroDec(),
			QuoteValue: sdkmath.NewInt(10),
		},
	}
)

// GetMinGigabytePrices returns the minimum gigabyte prices configured in the parameters.
func (m *Params) GetMinGigabytePrices() v1base.Prices {
	return m.MinGigabytePrices
}

// GetMinHourlyPrices returns the minimum hourly prices configured in the parameters.
func (m *Params) GetMinHourlyPrices() v1base.Prices {
	return m.MinHourlyPrices
}

// Validate validates all parameters to ensure they conform to expected rules.
func (m *Params) Validate() error {
	if err := validateActiveDuration(m.ActiveDuration); err != nil {
		return fmt.Errorf("invalid active_duration: %w", err)
	}

	if err := validateDeposit(m.Deposit); err != nil {
		return fmt.Errorf("invalid deposit: %w", err)
	}

	if err := validateMinGigabytePrices(m.MinGigabytePrices); err != nil {
		return fmt.Errorf("invalid min_gigabyte_prices: %w", err)
	}

	if err := validateMinHourlyPrices(m.MinHourlyPrices); err != nil {
		return fmt.Errorf("invalid min_hourly_prices: %w", err)
	}

	return nil
}

// NewParams creates a new Params instance with the provided values.
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

// DefaultParams returns the default parameters for the module.
func DefaultParams() Params {
	return NewParams(
		DefaultActiveDuration,
		DefaultDeposit,
		DefaultMinGigabytePrices,
		DefaultMinHourlyPrices,
	)
}

// validateActiveDuration checks that the active duration is greater than zero.
func validateActiveDuration(v time.Duration) error {
	if v == 0 {
		return errors.New("value cannot be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
}

// validateDeposit checks that the deposit is not nil, not negative, and valid.
func validateDeposit(v sdk.Coin) error {
	if v.IsNil() {
		return errors.New("value cannot be nil")
	}

	if v.IsNegative() {
		return errors.New("value cannot be negative")
	}

	if !v.IsValid() {
		return errors.New("invalid value")
	}

	return nil
}

// validateMinGigabytePrices validates the list of minimum gigabyte prices.
func validateMinGigabytePrices(v []v1base.Price) error {
	if v == nil {
		return nil
	}

	if err := v1base.Prices(v).Validate(); err != nil {
		return fmt.Errorf("invalid value: %w", err)
	}

	return nil
}

// validateMinHourlyPrices validates the list of minimum hourly prices.
func validateMinHourlyPrices(v []v1base.Price) error {
	if v == nil {
		return nil
	}

	if err := v1base.Prices(v).Validate(); err != nil {
		return fmt.Errorf("invalid value: %w", err)
	}

	return nil
}
