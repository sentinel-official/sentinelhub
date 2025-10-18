package v1

import (
	"fmt"
	"time"
)

func NewGenesisState(inflations []Inflation) *GenesisState {
	return &GenesisState{
		Inflations: inflations,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil)
}

func (gs *GenesisState) Validate() error {
	m := make(map[time.Time]bool)
	for _, item := range gs.Inflations {
		if m[item.Timestamp] {
			return fmt.Errorf("found a duplicate inflation for timestamp %s", item.Timestamp)
		}

		m[item.Timestamp] = true
	}

	for _, item := range gs.Inflations {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
