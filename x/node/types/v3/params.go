package v3

import (
	"errors"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/sentinelhub/v13/types/v1"
)

var (
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

	// DefaultStatusTimeout defines the default duration before a status is considered outdated.
	DefaultStatusTimeout = 30 * time.Second
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
	if err := validateDeposit(m.Deposit); err != nil {
		return fmt.Errorf("invalid deposit: %w", err)
	}

	if err := validateMinGigabytePrices(m.MinGigabytePrices); err != nil {
		return fmt.Errorf("invalid min_gigabyte_prices: %w", err)
	}

	if err := validateMinHourlyPrices(m.MinHourlyPrices); err != nil {
		return fmt.Errorf("invalid min_hourly_prices: %w", err)
	}

	if err := validateStatusTimeout(m.StatusTimeout); err != nil {
		return fmt.Errorf("invalid status_timeout: %w", err)
	}

	return nil
}

// NewParams creates a new Params instance with the provided values.
func NewParams(
	deposit sdk.Coin, minGigabytePrices, minHourlyPrices v1base.Prices, statusTimeout time.Duration,
) Params {
	return Params{
		Deposit:           deposit,
		MinGigabytePrices: minGigabytePrices,
		MinHourlyPrices:   minHourlyPrices,
		StatusTimeout:     statusTimeout,
	}
}

// DefaultParams returns the default parameters for the module.
func DefaultParams() Params {
	return NewParams(
		DefaultDeposit,
		DefaultMinGigabytePrices,
		DefaultMinHourlyPrices,
		DefaultStatusTimeout,
	)
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

// validateStatusTimeout checks that the status duration is greater than zero.
func validateStatusTimeout(v time.Duration) error {
	if v == 0 {
		return errors.New("value cannot be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
}
