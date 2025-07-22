package v3

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultDeposit = sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1000)) // Default required deposit: 1000 stake
)

// Validate checks that all parameters in Params are valid.
func (m *Params) Validate() error {
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}

	return nil
}

// NewParams creates a new Params instance with the given deposit.
func NewParams(deposit sdk.Coin) Params {
	return Params{
		Deposit: deposit,
	}
}

// DefaultParams returns a Params instance initialized with default values.
func DefaultParams() Params {
	return NewParams(DefaultDeposit)
}

// validateDeposit checks that the deposit is a valid, non-negative, non-nil coin.
func validateDeposit(v sdk.Coin) error {
	if v.IsNil() {
		return errors.New("deposit cannot be nil")
	}
	if v.IsNegative() {
		return errors.New("deposit cannot be negative")
	}
	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid deposit: %w", err)
	}

	return nil
}
