package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
}

type NodeKeeper interface {
	SessionInactivePreHook(ctx sdk.Context, id uint64) error
	UpdateSessionMaxValues(ctx sdk.Context, session v3.Session) error
}

type SubscriptionKeeper interface {
	SessionInactivePreHook(ctx sdk.Context, id uint64) error
	SessionUpdatePreHook(ctx sdk.Context, id uint64, currBytes sdkmath.Int) error
	UpdateSessionMaxValues(ctx sdk.Context, session v3.Session) error
}
