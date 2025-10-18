package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/deposit/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
	for _, item := range state.Deposits {
		k.SetDeposit(ctx, item)
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return v1.NewGenesisState(
		k.GetDeposits(ctx),
	)
}
