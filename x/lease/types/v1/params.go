package v1

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
)

// Default parameter values for the Params struct.
var (
	DefaultMaxHours     int64 = 10                                     // Default value for maximum hours
	DefaultMinHours     int64 = 1                                      // Default value for minimum hours
	DefaultStakingShare       = sdkmath.LegacyMustNewDecFromStr("0.1") // Default staking share: 0.1
)

// Validate checks whether the Params fields are valid according to defined rules.
func (m *Params) Validate() error {
	if err := validateMaxHours(m.MaxHours); err != nil {
		return fmt.Errorf("invalid max_hours: %w", err)
	}

	if err := validateMinHours(m.MinHours); err != nil {
		return fmt.Errorf("invalid min_hours: %w", err)
	}

	if err := validateStakingShare(m.StakingShare); err != nil {
		return fmt.Errorf("invalid staking_share: %w", err)
	}

	if m.MinHours > m.MaxHours {
		return errors.New("min_hours cannot be greater than max_hours")
	}

	return nil
}

// NewParams creates a new Params instance with custom values.
func NewParams(maxHours, minHours int64, stakingShare sdkmath.LegacyDec) Params {
	return Params{
		MaxHours:     maxHours,
		MinHours:     minHours,
		StakingShare: stakingShare,
	}
}

// DefaultParams returns a Params struct initialized with default values.
func DefaultParams() Params {
	return NewParams(
		DefaultMaxHours,
		DefaultMinHours,
		DefaultStakingShare,
	)
}

// validateMaxHours checks that maxHours is a positive integer.
func validateMaxHours(v int64) error {
	if v == 0 {
		return errors.New("value cannot be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
}

// validateMinHours checks that minHours is a positive integer.
func validateMinHours(v int64) error {
	if v == 0 {
		return errors.New("value cannot be zero")
	}

	if v < 0 {
		return errors.New("value cannot be negative")
	}

	return nil
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
