package v1

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultParams(),
	)
}

func (m *GenesisState) Validate() error {
	return nil
}
