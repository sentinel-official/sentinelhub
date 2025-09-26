package v3

import (
	"errors"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
)

// Default parameter values for the Params struct.
var (
	DefaultStakingShare  = sdkmath.LegacyMustNewDecFromStr("0.1") // Default staking share: 0.1
	DefaultStatusTimeout = 2 * time.Minute                        // Default timeout for status change
)

// Validate checks whether the Params fields are valid according to defined rules.
func (m *Params) Validate() error {
	if err := validateStakingShare(m.StakingShare); err != nil {
		return fmt.Errorf("invalid staking_share: %w", err)
	}

	if err := validateStatusTimeout(m.StatusTimeout); err != nil {
		return fmt.Errorf("invalid status_timeout: %w", err)
	}

	return nil
}

// NewParams creates a new Params instance with custom values.
func NewParams(stakingShare sdkmath.LegacyDec, statusTimeout time.Duration) Params {
	return Params{
		StakingShare:  stakingShare,
		StatusTimeout: statusTimeout,
	}
}

// DefaultParams returns a Params struct initialized with default values.
func DefaultParams() Params {
	return NewParams(
		DefaultStakingShare,
		DefaultStatusTimeout,
	)
}

// validateStakingShare ensures that the staking share is:
// - Not nil
// - Not negative
// - Not greater than 1 (100%).
func validateStakingShare(v sdkmath.LegacyDec) error {
	if v.IsNil() {
		return errors.New("value cannot be nil")
	}

	if v.IsNegative() {
		return errors.New("value cannot be negative")
	}

	if v.GT(sdkmath.LegacyOneDec()) {
		return errors.New("value cannot be greater than 1")
	}

	return nil
}

// validateStatusTimeout checks that statusTimeout is a positive duration.
func validateStatusTimeout(v time.Duration) error {
	if v == 0 {
		return errors.New("value cannot be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
}
