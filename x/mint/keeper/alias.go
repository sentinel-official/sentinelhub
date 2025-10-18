package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func (k *Keeper) GetMinter(ctx sdk.Context) minttypes.Minter {
	return k.mint.GetMinter(ctx)
}

func (k *Keeper) SetMinter(ctx sdk.Context, minter minttypes.Minter) {
	k.mint.SetMinter(ctx, minter)
}

func (k *Keeper) GetParams(ctx sdk.Context) minttypes.Params {
	return k.mint.GetParams(ctx)
}

func (k *Keeper) SetParams(ctx sdk.Context, params minttypes.Params) error {
	return k.mint.SetParams(ctx, params)
}
