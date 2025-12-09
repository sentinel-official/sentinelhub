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

	"github.com/sentinel-official/sentinelhub/v13/x/plan/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	router       *baseapp.MsgServiceRouter
	storeService store.KVStoreService

	lease        LeaseKeeper
	node         NodeKeeper
	provider     ProviderKeeper
	session      SessionKeeper
	subscription SubscriptionKeeper
}

func NewKeeper(cdc codec.BinaryCodec, storeService store.KVStoreService, router *baseapp.MsgServiceRouter) Keeper {
	return Keeper{
		cdc:          cdc,
		router:       router,
		storeService: storeService,
	}
}

func (k *Keeper) WithLeaseKeeper(keeper LeaseKeeper)               { k.lease = keeper }
func (k *Keeper) WithNodeKeeper(keeper NodeKeeper)                 { k.node = keeper }
func (k *Keeper) WithProviderKeeper(keeper ProviderKeeper)         { k.provider = keeper }
func (k *Keeper) WithSessionKeeper(keeper SessionKeeper)           { k.session = keeper }
func (k *Keeper) WithSubscriptionKeeper(keeper SubscriptionKeeper) { k.subscription = keeper }

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) storetypes.KVStore {
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	child := types.ModuleName + "/"

	return prefix.NewStore(kvStore, []byte(child))
}
