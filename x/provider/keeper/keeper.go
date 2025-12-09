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

	"github.com/sentinel-official/sentinelhub/v13/x/provider/types"
)

type Keeper struct {
	authority    string
	cdc          codec.BinaryCodec
	router       *baseapp.MsgServiceRouter
	storeService store.KVStoreService

	distribution DistributionKeeper
	plan         PlanKeeper
	lease        LeaseKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec, storeService store.KVStoreService, router *baseapp.MsgServiceRouter, authority string,
) Keeper {
	return Keeper{
		authority:    authority,
		cdc:          cdc,
		router:       router,
		storeService: storeService,
	}
}

func (k *Keeper) WithDistributionKeeper(keeper DistributionKeeper) { k.distribution = keeper }
func (k *Keeper) WithPlanKeeper(keeper PlanKeeper)                 { k.plan = keeper }
func (k *Keeper) WithLeaseKeeper(keeper LeaseKeeper)               { k.lease = keeper }

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) storetypes.KVStore {
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	child := types.ModuleName + "/"

	return prefix.NewStore(kvStore, []byte(child))
}
