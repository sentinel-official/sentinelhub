package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/mint/types"
)

type Keeper struct {
	cdc codec.BinaryCodec
	key storetypes.StoreKey

	mint MintKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, mint MintKeeper) Keeper {
	return Keeper{
		cdc:  cdc,
		key:  key,
		mint: mint,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}
