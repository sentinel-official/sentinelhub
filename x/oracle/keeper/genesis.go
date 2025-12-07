package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/oracle/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
	k.SetPortID(ctx, state.PortID)

	if !k.IsBound(ctx, state.PortID) {
		if err := k.BindPort(ctx, state.PortID); err != nil {
			panic(fmt.Errorf("claiming capability for port %q: %w", state.PortID, err))
		}
	}

	k.SetParams(ctx, state.Params)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return &v1.GenesisState{}
}
