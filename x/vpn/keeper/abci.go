package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) BeginBlock(c context.Context) error {
	ctx := sdk.UnwrapSDKContext(c)

	k.Lease.BeginBlock(ctx)
	k.Node.BeginBlock(ctx)
	k.Session.BeginBlock(ctx)
	k.Subscription.BeginBlock(ctx)

	return nil
}
