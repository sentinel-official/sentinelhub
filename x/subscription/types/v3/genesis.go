package v3

import (
	"fmt"

	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func NewGenesisState(allocations []v2.Allocation, subscriptions []Subscription, params Params) *GenesisState {
	return &GenesisState{
		Allocations:   allocations,
		Subscriptions: subscriptions,
		Params:        params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	m := make(map[uint64]bool)
	for _, item := range state.Subscriptions {
		if m[item.ID] {
			return fmt.Errorf("found a duplicate subscription for id %d", item.ID)
		}

		m[item.ID] = true
		if err := item.Validate(); err != nil {
			return err
		}
	}

	mm := make(map[uint64]map[string]bool)
	for _, item := range state.Allocations {
		if _, ok := mm[item.ID]; !ok {
			mm[item.ID] = make(map[string]bool)
		}
	}

	for _, item := range state.Allocations {
		if mm[item.ID][item.Address] {
			return fmt.Errorf("found a duplicate allocation %d/%s", item.ID, item.Address)
		}

		mm[item.ID][item.Address] = true
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
