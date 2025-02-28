package v3

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
)

var (
	DefaultMaxGigabytes             int64 = 10
	DefaultMinGigabytes             int64 = 1
	DefaultMaxHours                 int64 = 10
	DefaultMinHours                 int64 = 1
	DefaultProofVerificationEnabled       = false
	DefaultStakingShare                   = math.LegacyNewDecWithPrec(1, 1)
	DefaultStatusChangeDelay              = 1 * time.Minute
)

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

func validateMaxGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("max_gigabytes cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("max_gigabytes cannot be zero")
	}

	return nil
}

func validateMinGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("min_gigabytes cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("min_gigabytes cannot be zero")
	}

	return nil
}

func validateMaxHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("max_hours cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("max_hours cannot be zero")
	}

	return nil
}

func validateMinHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("min_hours cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("min_hours cannot be zero")
	}

	return nil
}

func validateProofVerificationEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	return nil
}

func validateStakingShare(v interface{}) error {
	value, ok := v.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("staking_share cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("staking_share cannot be negative")
	}
	if value.GT(math.LegacyOneDec()) {
		return fmt.Errorf("staking_share cannot be greater than 1")
	}

	return nil
}

func validateStatusChangeDelay(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("status_change_delay cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("status_change_delay cannot be zero")
	}

	return nil
}
