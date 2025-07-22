package v3

import (
	"errors"
	"fmt"
	"time"

	"cosmossdk.io/math"
)

// Default parameter values for the Params struct
var (
	DefaultMaxGigabytes             int64 = 10                                  // Default maximum allowed gigabytes
	DefaultMinGigabytes             int64 = 1                                   // Default minimum allowed gigabytes
	DefaultMaxHours                 int64 = 10                                  // Default maximum allowed hours
	DefaultMinHours                 int64 = 1                                   // Default minimum allowed hours
	DefaultProofVerificationEnabled       = false                               // Default proof verification flag
	DefaultStakingShare                   = math.LegacyMustNewDecFromStr("0.1") // Default staking share: 0.1
	DefaultStatusChangeDelay              = 1 * time.Minute                     // Default delay before validator status change
)

// Validate checks whether the Params fields are valid according to defined rules.
func (m *Params) Validate() error {
	if err := validateMaxGigabytes(m.MaxGigabytes); err != nil {
		return err
	}
	if err := validateMinGigabytes(m.MinGigabytes); err != nil {
		return err
	}
	if err := validateMaxHours(m.MaxHours); err != nil {
		return err
	}
	if err := validateMinHours(m.MinHours); err != nil {
		return err
	}
	if err := validateProofVerificationEnabled(m.ProofVerificationEnabled); err != nil {
		return err
	}
	if err := validateStakingShare(m.StakingShare); err != nil {
		return err
	}
	if err := validateStatusChangeDelay(m.StatusChangeDelay); err != nil {
		return err
	}

	return nil
}

// NewParams creates a new Params instance with custom values.
func NewParams(
	maxGigabytes, minGigabytes, maxHours, minHours int64, proofVerificationEnabled bool, stakingShare math.LegacyDec,
	statusChangeDelay time.Duration,
) Params {
	return Params{
		MaxGigabytes:             maxGigabytes,
		MinGigabytes:             minGigabytes,
		MaxHours:                 maxHours,
		MinHours:                 minHours,
		ProofVerificationEnabled: proofVerificationEnabled,
		StakingShare:             stakingShare,
		StatusChangeDelay:        statusChangeDelay,
	}
}

// DefaultParams returns a Params struct initialized with default values.
func DefaultParams() Params {
	return NewParams(
		DefaultMaxGigabytes,
		DefaultMinGigabytes,
		DefaultMaxHours,
		DefaultMinHours,
		DefaultProofVerificationEnabled,
		DefaultStakingShare,
		DefaultStatusChangeDelay,
	)
}

// validateMaxGigabytes ensures maxGigabytes is a positive non-zero integer.
func validateMaxGigabytes(v int64) error {
	if v < 0 {
		return errors.New("max_gigabytes cannot be negative")
	}
	if v == 0 {
		return errors.New("max_gigabytes cannot be zero")
	}

	return nil
}

// validateMinGigabytes ensures minGigabytes is a positive non-zero integer.
func validateMinGigabytes(v int64) error {
	if v < 0 {
		return errors.New("min_gigabytes cannot be negative")
	}
	if v == 0 {
		return errors.New("min_gigabytes cannot be zero")
	}

	return nil
}

// validateMaxHours ensures maxHours is a positive non-zero integer.
func validateMaxHours(v int64) error {
	if v < 0 {
		return errors.New("max_hours cannot be negative")
	}
	if v == 0 {
		return errors.New("max_hours cannot be zero")
	}

	return nil
}

// validateMinHours ensures minHours is a positive non-zero integer.
func validateMinHours(v int64) error {
	if v < 0 {
		return errors.New("min_hours cannot be negative")
	}
	if v == 0 {
		return errors.New("min_hours cannot be zero")
	}

	return nil
}

// validateProofVerificationEnabled always returns nil as the type is bool.
func validateProofVerificationEnabled(v bool) error {
	// Bool type needs no validation in this context
	return nil
}

// validateStakingShare ensures stakingShare is not nil, not negative, and ≤ 1.
func validateStakingShare(v math.LegacyDec) error {
	if v.IsNil() {
		return fmt.Errorf("staking_share cannot be nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("staking_share cannot be negative")
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("staking_share cannot be greater than 1")
	}

	return nil
}

// validateStatusChangeDelay ensures the delay is positive and non-zero.
func validateStatusChangeDelay(v time.Duration) error {
	if v < 0 {
		return errors.New("status_change_delay cannot be negative")
	}
	if v == 0 {
		return errors.New("status_change_delay cannot be zero")
	}

	return nil
}
