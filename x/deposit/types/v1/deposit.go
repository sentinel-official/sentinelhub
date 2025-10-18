package v1

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *Deposit) Validate() error {
	if m.Address == "" {
		return errors.New("address cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("invalid address %s: %w", m.Address, err)
	}

	if m.Coins == nil {
		return errors.New("coins cannot be empty")
	}

	if m.Coins.Len() == 0 {
		return errors.New("coins length cannot be zero")
	}

	if m.Coins.IsAnyNil() {
		return errors.New("coins cannot be nil")
	}

	if !m.Coins.IsValid() {
		return errors.New("coins must be valid")
	}

	return nil
}

type (
	Deposits []Deposit
)
