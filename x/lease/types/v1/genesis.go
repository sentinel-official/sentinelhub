package v1

import (
	"fmt"
)

func NewGenesisState(leases []Lease, params Params) *GenesisState {
	return &GenesisState{
		Leases: leases,
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	m := make(map[uint64]bool)
	for _, lease := range state.Leases {
		if m[lease.ID] {
			return fmt.Errorf("found a duplicate lease with ID %d", lease.ID)
		}

		m[lease.ID] = true

		if err := lease.Validate(); err != nil {
			return err
		}
	}

	return nil
}
