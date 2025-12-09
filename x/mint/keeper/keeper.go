package keeper

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/mint/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService

	mint MintKeeper
}

func NewKeeper(cdc codec.BinaryCodec, storeService store.KVStoreService, mint MintKeeper) Keeper {
	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		mint:         mint,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) storetypes.KVStore {
	return runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
}
