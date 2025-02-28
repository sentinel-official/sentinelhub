package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, from string, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, from, to string, coins sdk.Coins) error
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type OracleKeeper interface {
	GetQuotePrice(ctx sdk.Context, price sdk.DecCoin) (sdk.Coin, error)
}
