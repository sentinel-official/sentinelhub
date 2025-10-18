package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	var state = DefaultGenesisState()
	require.Equal(t, &GenesisState{Deposits: nil}, state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		deposits Deposits
		state    *GenesisState
	)

	state = NewGenesisState(nil)
	require.Equal(t, &GenesisState{Deposits: nil}, state)
	require.Empty(t, state.Deposits)

	state = NewGenesisState(deposits)
	require.Equal(t, &GenesisState{Deposits: deposits}, state)
	require.Empty(t, state.Deposits)

	deposits = append(deposits,
		Deposit{},
	)
	state = NewGenesisState(deposits)
	require.Equal(t, &GenesisState{Deposits: deposits}, state)
	require.Len(t, state.Deposits, 1)

	deposits = append(deposits,
		Deposit{},
	)
	state = NewGenesisState(deposits)
	require.Equal(t, &GenesisState{Deposits: deposits}, state)
	require.Len(t, state.Deposits, 2)
}
