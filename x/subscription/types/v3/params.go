package v3

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
)

var (
	DefaultStakingShare      = math.LegacyMustNewDecFromStr("0.1")
	DefaultStatusChangeDelay = 2 * time.Minute
)

func (m *Params) Validate() error {
	if err := validateStakingShare(m.StakingShare); err != nil {
		return err
	}
	if err := validateStatusChangeDelay(m.StatusChangeDelay); err != nil {
		return err
	}

	return nil
}

func NewParams(stakingShare math.LegacyDec, statusChangeDelay time.Duration) Params {
	return Params{
		StakingShare:      stakingShare,
		StatusChangeDelay: statusChangeDelay,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultStakingShare,
		DefaultStatusChangeDelay,
	)
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
