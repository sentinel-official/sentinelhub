package v1

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
)

var (
	DefaultMaxHours     int64 = 10
	DefaultMinHours     int64 = 1
	DefaultStakingShare       = sdkmath.LegacyNewDecWithPrec(1, 1)
)

func (m *Params) Validate() error {
	if err := validateMaxHours(m.MaxHours); err != nil {
		return err
	}
	if err := validateMinHours(m.MinHours); err != nil {
		return err
	}
	if err := validateStakingShare(m.StakingShare); err != nil {
		return err
	}

	return nil
}

func NewParams(maxHours, minHours int64, stakingShare sdkmath.LegacyDec) Params {
	return Params{
		MaxHours:     maxHours,
		MinHours:     minHours,
		StakingShare: stakingShare,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultMaxHours,
		DefaultMinHours,
		DefaultStakingShare,
	)
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

func validateStakingShare(v interface{}) error {
	value, ok := v.(sdkmath.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("staking_share cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("staking_share cannot be negative")
	}
	if value.GT(sdkmath.LegacyOneDec()) {
		return fmt.Errorf("staking_share cannot be greater than 1")
	}

	return nil
}
