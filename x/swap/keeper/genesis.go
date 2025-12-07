package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/swap/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Swaps {
		k.SetSwap(ctx, item)
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return v1.NewGenesisState(
		k.GetSwaps(ctx),
		k.GetParams(ctx),
	)
}
