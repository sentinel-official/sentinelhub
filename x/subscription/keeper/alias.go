package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v3"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v3"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) SendCoin(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromAccountToModule(ctx sdk.Context, from sdk.AccAddress, to string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.SendCoinsFromAccountToModule(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, addr)
}

func (k *Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) bool {
	return k.node.HasNodeForPlan(ctx, id, addr)
}

func (k *Keeper) QuotePriceFunc(ctx sdk.Context, price sdk.DecCoin) (sdk.Coin, error) {
	return k.oracle.GetQuotePrice(ctx, price)
}

func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool) {
	return k.plan.GetPlan(ctx, id)
}

func (k *Keeper) DeleteSession(ctx sdk.Context, id uint64) {
	k.session.DeleteSession(ctx, id)
}

func (k *Keeper) DeleteSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	k.session.DeleteSessionForAccount(ctx, addr, id)
}

func (k *Keeper) DeleteSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	k.session.DeleteSessionForAllocation(ctx, subscriptionID, addr, sessionID)
}

func (k *Keeper) DeleteSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	k.session.DeleteSessionForNode(ctx, addr, id)
}

func (k *Keeper) DeleteSessionForPlanByNode(ctx sdk.Context, planID uint64, addr base.NodeAddress, sessionID uint64) {
	k.session.DeleteSessionForPlanByNode(ctx, planID, addr, sessionID)
}

func (k *Keeper) DeleteSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	k.session.DeleteSessionForSubscription(ctx, subscriptionID, sessionID)
}

func (k *Keeper) GetSessionCount(ctx sdk.Context) uint64 {
	return k.session.GetSessionCount(ctx)
}

func (k *Keeper) GetSessionInactiveAt(ctx sdk.Context) time.Time {
	return k.session.GetInactiveAt(ctx)
}

func (k *Keeper) GetSession(ctx sdk.Context, id uint64) (sessiontypes.Session, bool) {
	return k.session.GetSession(ctx, id)
}

func (k *Keeper) SetSessionCount(ctx sdk.Context, count uint64) {
	k.session.SetSessionCount(ctx, count)
}

func (k *Keeper) SetSession(ctx sdk.Context, session sessiontypes.Session) {
	k.session.SetSession(ctx, session)
}

func (k *Keeper) SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	k.session.SetSessionForAccount(ctx, addr, id)
}

func (k *Keeper) SetSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	k.session.SetSessionForAllocation(ctx, subscriptionID, addr, sessionID)
}

func (k *Keeper) SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	k.session.SetSessionForInactiveAt(ctx, at, id)
}

func (k *Keeper) SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	k.session.SetSessionForNode(ctx, addr, id)
}

func (k *Keeper) SetSessionForPlanByNode(ctx sdk.Context, planID uint64, addr base.NodeAddress, sessionID uint64) {
	k.session.SetSessionForPlanByNode(ctx, planID, addr, sessionID)
}

func (k *Keeper) SetSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	k.session.SetSessionForSubscription(ctx, subscriptionID, sessionID)
}

func (k *Keeper) SubscriptionInactivePendingPreHook(ctx sdk.Context, id uint64) error {
	if err := k.session.SubscriptionInactivePendingPreHook(ctx, id); err != nil {
		return err
	}

	return nil
}
