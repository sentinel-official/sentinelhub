package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.Lease.BeginBlock(ctx)
	k.Node.BeginBlock(ctx)
	k.Session.BeginBlock(ctx)
	k.Subscription.BeginBlock(ctx)
}
