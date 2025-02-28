package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) SendCoinFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToModule(ctx sdk.Context, from sdk.AccAddress, to string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToModule(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) SessionInactivePreHook(ctx sdk.Context, id uint64) error {
	if err := k.node.SessionInactivePreHook(ctx, id); err != nil {
		return err
	}
	if err := k.subscription.SessionInactivePreHook(ctx, id); err != nil {
		return err
	}

	return nil
}

func (k *Keeper) SessionUpdatePreHook(ctx sdk.Context, id uint64, currBytes sdkmath.Int) error {
	if err := k.subscription.SessionUpdatePreHook(ctx, id, currBytes); err != nil {
		return err
	}

	return nil
}

func (k *Keeper) UpdateMaxValues(ctx sdk.Context, session v3.Session) error {
	if err := k.node.UpdateSessionMaxValues(ctx, session); err != nil {
		return err
	}
	if err := k.subscription.UpdateSessionMaxValues(ctx, session); err != nil {
		return err
	}

	return nil
}
