package keeper

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/session/types"
)

type Keeper struct {
	authority        string
	cdc              codec.BinaryCodec
	feeCollectorName string
	router           *baseapp.MsgServiceRouter
	storeService     store.KVStoreService

	account      AccountKeeper
	deposit      DepositKeeper
	node         NodeKeeper
	subscription SubscriptionKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec, storeService store.KVStoreService, router *baseapp.MsgServiceRouter,
	authority, feeCollectorName string,
) Keeper {
	return Keeper{
		authority:        authority,
		cdc:              cdc,
		feeCollectorName: feeCollectorName,
		router:           router,
		storeService:     storeService,
	}
}

func (k *Keeper) WithAccountKeeper(keeper AccountKeeper)           { k.account = keeper }
func (k *Keeper) WithDepositKeeper(keeper DepositKeeper)           { k.deposit = keeper }
func (k *Keeper) WithNodeKeeper(keeper NodeKeeper)                 { k.node = keeper }
func (k *Keeper) WithSubscriptionKeeper(keeper SubscriptionKeeper) { k.subscription = keeper }

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) storetypes.KVStore {
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	child := types.ModuleName + "/"

	return prefix.NewStore(kvStore, []byte(child))
}
