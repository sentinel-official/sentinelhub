package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v3.GenesisState) {
	k.SetParams(ctx, state.Params)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v3.GenesisState {
	return v3.NewGenesisState(
		k.GetSessions(ctx),
		k.GetParams(ctx),
	)
}
