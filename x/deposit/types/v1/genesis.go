package v1

import (
	"fmt"
)

func NewGenesisState(deposits Deposits) *GenesisState {
	return &GenesisState{
		Deposits: deposits,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil)
}

func ValidateGenesisState(state *GenesisState) error {
	m := make(map[string]bool)
	for _, item := range state.Deposits {
		if m[item.Address] {
			return fmt.Errorf("found a duplicate deposit for address %s", item.Address)
		}

		m[item.Address] = true
	}

	for _, item := range state.Deposits {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
