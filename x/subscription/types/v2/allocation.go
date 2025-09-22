package v2

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *Allocation) Validate() error {
	if m.Address == "" {
		return errors.New("address cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("invalid address %s: %w", m.Address, err)
	}

	if m.GrantedBytes.IsNil() {
		return errors.New("granted_bytes cannot be nil")
	}

	if m.GrantedBytes.IsNegative() {
		return errors.New("granted_bytes cannot be negative")
	}

	if m.UtilisedBytes.IsNil() {
		return errors.New("utilised_bytes cannot be nil")
	}

	if m.UtilisedBytes.IsNegative() {
		return errors.New("utilised_bytes cannot be negative")
	}

	if m.UtilisedBytes.GT(m.GrantedBytes) {
		return errors.New("utilised_bytes cannot be greater than granted_bytes")
	}

	return nil
}

type (
	Allocations []Allocation
)
