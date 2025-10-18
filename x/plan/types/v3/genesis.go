package v3

import (
	"fmt"
)

type (
	GenesisPlans []GenesisPlan
)

func NewGenesisState(plans GenesisPlans) *GenesisState {
	return &GenesisState{
		Plans: plans,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil)
}

func ValidateGenesis(state *GenesisState) error {
	m := make(map[uint64]bool)
	for _, item := range state.Plans {
		if m[item.Plan.ID] {
			return fmt.Errorf("found a duplicate plan for id %d", item.Plan.ID)
		}

		m[item.Plan.ID] = true
	}

	for _, item := range state.Plans {
		m := make(map[string]bool)
		for _, addr := range item.Nodes {
			if m[addr] {
				return fmt.Errorf("found a duplicate node %s for the plan %d", addr, item.Plan.ID)
			}

			m[addr] = true
		}
	}

	for _, item := range state.Plans {
		if err := item.Plan.Validate(); err != nil {
			return err
		}
	}

	return nil
}
