package keeper

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/provider/types"
)

type Keeper struct {
	authority string
	cdc       codec.BinaryCodec
	key       storetypes.StoreKey
	router    *baseapp.MsgServiceRouter

	distribution DistributionKeeper
	plan         PlanKeeper
	lease        LeaseKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, router *baseapp.MsgServiceRouter, authority string,
) Keeper {
	return Keeper{
		authority: authority,
		cdc:       cdc,
		key:       key,
		router:    router,
	}
}

func (k *Keeper) WithDistributionKeeper(keeper DistributionKeeper) { k.distribution = keeper }
func (k *Keeper) WithPlanKeeper(keeper PlanKeeper)                 { k.plan = keeper }
func (k *Keeper) WithLeaseKeeper(keeper LeaseKeeper)               { k.lease = keeper }

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := types.ModuleName + "/"

	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
