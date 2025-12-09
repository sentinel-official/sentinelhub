package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountKeeper interface {
	GetAccount(ctx context.Context, address sdk.AccAddress) sdk.AccountI
}

type BankKeeper interface {
	MintCoins(ctx context.Context, name string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, name string, address sdk.AccAddress, coins sdk.Coins) error
}
