package v3

func NewGenesisState(_ []Session, params Params) *GenesisState {
	return &GenesisState{
		Sessions: nil,
		Params:   params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	return nil
}
