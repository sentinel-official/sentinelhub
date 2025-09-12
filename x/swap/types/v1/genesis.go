package v1

import (
	"fmt"
)

func NewGenesisState(swaps Swaps, params Params) *GenesisState {
	return &GenesisState{
		Swaps:  swaps,
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func (m *GenesisState) Validate() error {
	if err := m.Params.Validate(); err != nil {
		return err
	}

	swaps := make(map[string]bool)

	for _, item := range m.Swaps {
		txHash := item.GetTxHash().String()
		if swaps[txHash] {
			return fmt.Errorf("found duplicate swap for tx_hash %s", txHash)
		}

		swaps[txHash] = true
	}

	for _, item := range m.Swaps {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
