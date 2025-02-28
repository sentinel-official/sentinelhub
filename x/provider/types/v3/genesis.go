package v3

import (
	"fmt"

	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

func NewGenesisState(providers v2.Providers, params Params) *GenesisState {
	return &GenesisState{
		Providers: providers,
		Params:    params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	m := make(map[string]bool)
	for _, item := range state.Providers {
		if m[item.Address] {
			return fmt.Errorf("found a duplicate provider %s", item.Address)
		}

		m[item.Address] = true
	}

	for _, item := range state.Providers {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
