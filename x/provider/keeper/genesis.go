package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types/v3"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v3.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Providers {
		k.SetProvider(ctx, item)
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v3.GenesisState {
	return v3.NewGenesisState(
		k.GetProviders(ctx),
		k.GetParams(ctx),
	)
}
