package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/vpn/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
	k.Deposit.InitGenesis(ctx, state.Deposit)
	k.Lease.InitGenesis(ctx, state.Lease)
	k.Node.InitGenesis(ctx, state.Node)
	k.Plan.InitGenesis(ctx, state.Plan)
	k.Provider.InitGenesis(ctx, state.Provider)
	k.Session.InitGenesis(ctx, state.Session)
	k.Subscription.InitGenesis(ctx, state.Subscription)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return &v1.GenesisState{
		Deposit:      k.Deposit.ExportGenesis(ctx),
		Lease:        k.Lease.ExportGenesis(ctx),
		Node:         k.Node.ExportGenesis(ctx),
		Plan:         k.Plan.ExportGenesis(ctx),
		Provider:     k.Provider.ExportGenesis(ctx),
		Session:      k.Session.ExportGenesis(ctx),
		Subscription: k.Subscription.ExportGenesis(ctx),
	}
}
