package keeper

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) FundCommunityPool(ctx sdk.Context, fromAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), fromAddr)
}

func (k *Keeper) AddDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.AddDeposit(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, fromAddr, toAddr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToModule(ctx, fromAddr, toModule, sdk.NewCoins(coin))
}

func (k *Keeper) SubtractDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SubtractDeposit(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) QuotePriceFunc(ctx sdk.Context, price sdk.DecCoin) (sdk.Coin, error) {
	return k.oracle.GetQuotePrice(ctx, price)
}

func (k *Keeper) DeleteSession(ctx sdk.Context, id uint64) {
	k.session.DeleteSession(ctx, id)
}

func (k *Keeper) DeleteSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	k.session.DeleteSessionForAccount(ctx, addr, id)
}

func (k *Keeper) DeleteSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	k.session.DeleteSessionForNode(ctx, addr, id)
}

func (k *Keeper) IsValidSessionGigabytes(ctx sdk.Context, gigabytes int64) bool {
	return k.session.IsValidGigabytes(ctx, gigabytes)
}

func (k *Keeper) IsValidSessionHours(ctx sdk.Context, hours int64) bool {
	return k.session.IsValidHours(ctx, hours)
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

func (k *Keeper) SessionStakingShare(ctx sdk.Context) math.LegacyDec {
	return k.session.StakingShare(ctx)
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

func (k *Keeper) SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	k.session.SetSessionForInactiveAt(ctx, at, id)
}

func (k *Keeper) SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	k.session.SetSessionForNode(ctx, addr, id)
}

func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	if err := k.lease.NodeInactivePreHook(ctx, addr); err != nil {
		return err
	}
	if err := k.session.NodeInactivePreHook(ctx, addr); err != nil {
		return err
	}

	return nil
}
