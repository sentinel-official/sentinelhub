package v1

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *Deposit) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Coins == nil {
		return fmt.Errorf("coins cannot be empty")
	}
	if m.Coins.Len() == 0 {
		return fmt.Errorf("coins length cannot be zero")
	}
	if m.Coins.IsAnyNil() {
		return fmt.Errorf("coins cannot be nil")
	}
	if !m.Coins.IsValid() {
		return fmt.Errorf("coins must be valid")
	}

	return nil
}

type (
	Deposits []Deposit
)
