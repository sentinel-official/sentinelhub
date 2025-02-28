package v2

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *Allocation) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.GrantedBytes.IsNil() {
		return fmt.Errorf("granted_bytes cannot be nil")
	}
	if m.GrantedBytes.IsNegative() {
		return fmt.Errorf("granted_bytes cannot be negative")
	}
	if m.UtilisedBytes.IsNil() {
		return fmt.Errorf("utilised_bytes cannot be nil")
	}
	if m.UtilisedBytes.IsNegative() {
		return fmt.Errorf("utilised_bytes cannot be negative")
	}
	if m.UtilisedBytes.GT(m.GrantedBytes) {
		return fmt.Errorf("utilised_bytes cannot be greater than granted_bytes")
	}

	return nil
}

type (
	Allocations []Allocation
)
