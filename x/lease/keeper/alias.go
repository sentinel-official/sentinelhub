package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	nodetypes "github.com/sentinel-official/sentinelhub/v13/x/node/types/v3"
	providertypes "github.com/sentinel-official/sentinelhub/v13/x/provider/types/v2"
)

func (k *Keeper) AddDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.AddDeposit(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) SubtractDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SubtractDeposit(ctx, addr, sdk.NewCoins(coin))
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

func (k *Keeper) GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, addr)
}

func (k *Keeper) GetQuotePrice(ctx context.Context, price sdk.DecCoin) (sdk.Coin, error) {
	return k.oracle.GetQuotePrice(ctx, price)
}

func (k *Keeper) GetProvider(ctx sdk.Context, addr base.ProvAddress) (providertypes.Provider, bool) {
	return k.provider.GetProvider(ctx, addr)
}

func (k *Keeper) LeaseInactivePreHook(ctx sdk.Context, id uint64) error {
	if err := k.plan.LeaseInactivePreHook(ctx, id); err != nil {
		return err
	}

	return nil
}
