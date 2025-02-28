package v3

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultDeposit = sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1000))
)

func (m *Params) Validate() error {
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}

	return nil
}

func NewParams(deposit sdk.Coin) Params {
	return Params{
		Deposit: deposit,
	}
}

func DefaultParams() Params {
	return NewParams(DefaultDeposit)
}

func validateDeposit(v interface{}) error {
	value, ok := v.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("deposit cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !value.IsValid() {
		return fmt.Errorf("invalid deposit %s", value)
	}

	return nil
}
