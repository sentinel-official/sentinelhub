package v1

import (
	"github.com/sentinel-official/sentinelhub/v13/x/oracle/types"
)

func NewGenesisState(assets []Asset, params Params, portID string) *GenesisState {
	return &GenesisState{
		Assets: assets,
		Params: params,
		PortID: portID,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		nil,
		DefaultParams(),
		types.PortID,
	)
}

func (m *GenesisState) Validate() error {
	return nil
}
