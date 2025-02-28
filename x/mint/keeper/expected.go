package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

type MintKeeper interface {
	GetMinter(ctx sdk.Context) minttypes.Minter
	SetMinter(ctx sdk.Context, minter minttypes.Minter)
	GetParams(ctx sdk.Context) minttypes.Params
	SetParams(ctx sdk.Context, params minttypes.Params) error
}
