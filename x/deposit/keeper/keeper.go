package keeper

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/deposit/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService

	bank BankKeeper
}

func NewKeeper(cdc codec.BinaryCodec, storeService store.KVStoreService) Keeper {
	return Keeper{
		cdc:          cdc,
		storeService: storeService,
	}
}

func (k *Keeper) WithBankKeeper(keeper BankKeeper) { k.bank = keeper }

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) storetypes.KVStore {
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	child := types.ModuleName + "/"

	return prefix.NewStore(kvStore, []byte(child))
}
