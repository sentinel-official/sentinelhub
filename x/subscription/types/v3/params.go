package v3

import (
	"errors"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
)

// Default parameter values for the Params struct.
var (
	DefaultStakingShare      = sdkmath.LegacyMustNewDecFromStr("0.1") // Default staking share: 0.1
	DefaultStatusChangeDelay = 2 * time.Minute                        // Default delay before status change
)

// Validate checks whether the Params fields are valid according to defined rules.
func (m *Params) Validate() error {
	if err := validateStakingShare(m.StakingShare); err != nil {
		return fmt.Errorf("invalid staking_share: %w", err)
	}

	if err := validateStatusChangeDelay(m.StatusChangeDelay); err != nil {
		return fmt.Errorf("invalid status_change_delay: %w", err)
	}

	return nil
}

// NewParams creates a new Params instance with custom values.
func NewParams(stakingShare sdkmath.LegacyDec, statusChangeDelay time.Duration) Params {
	return Params{
		StakingShare:      stakingShare,
		StatusChangeDelay: statusChangeDelay,
	}
}

// DefaultParams returns a Params struct initialized with default values.
func DefaultParams() Params {
	return NewParams(
		DefaultStakingShare,
		DefaultStatusChangeDelay,
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

// validateStatusChangeDelay checks that statusChangeDelay is a positive duration.
func validateStatusChangeDelay(v time.Duration) error {
	if v == 0 {
		return errors.New("value cannot be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
}
