package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountKeeper interface {
	GetAccount(ctx context.Context, address sdk.AccAddress) sdk.AccountI
}

type BankKeeper interface {
	SendCoins(ctx context.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, from string, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromModuleToModule(ctx context.Context, from, to string, coins sdk.Coins) error
	SpendableCoins(ctx context.Context, address sdk.AccAddress) sdk.Coins
}

type DistributionKeeper interface {
	FundCommunityPool(ctx context.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type OracleKeeper interface {
	GetQuotePrice(ctx context.Context, price sdk.DecCoin) (sdk.Coin, error)
}
