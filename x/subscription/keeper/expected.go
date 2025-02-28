package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v3"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v3"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	AddDeposit(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	SubtractDeposit(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error
}

type NodeKeeper interface {
	HasNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) bool
	GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool)
}

type OracleKeeper interface {
	GetQuotePrice(ctx sdk.Context, price sdk.DecCoin) (sdk.Coin, error)
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool)
}

type SessionKeeper interface {
	DeleteSession(ctx sdk.Context, id uint64)
	DeleteSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64)
	DeleteSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64)
	DeleteSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64)
	DeleteSessionForPlanByNode(ctx sdk.Context, planID uint64, addr base.NodeAddress, sessionID uint64)
	DeleteSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64)
	GetInactiveAt(ctx sdk.Context) time.Time
	GetSessionCount(ctx sdk.Context) uint64
	GetSession(ctx sdk.Context, id uint64) (sessiontypes.Session, bool)
	SetSessionCount(ctx sdk.Context, count uint64)
	SetSession(ctx sdk.Context, session sessiontypes.Session)
	SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64)
	SetSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64)
	SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64)
	SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64)
	SetSessionForPlanByNode(ctx sdk.Context, planID uint64, addr base.NodeAddress, sessionID uint64)
	SetSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64)
	SubscriptionInactivePendingPreHook(ctx sdk.Context, id uint64) error
}
