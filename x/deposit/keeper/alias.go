package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return k.bank.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

func (k *Keeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.bank.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (k *Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	return k.bank.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
}
