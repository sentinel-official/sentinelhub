package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/plan/types"
)

type Keeper struct {
	cdc    codec.BinaryCodec
	key    storetypes.StoreKey
	router *baseapp.MsgServiceRouter

	lease        LeaseKeeper
	node         NodeKeeper
	provider     ProviderKeeper
	session      SessionKeeper
	subscription SubscriptionKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, router *baseapp.MsgServiceRouter) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		router: router,
	}
}

func (k *Keeper) WithLeaseKeeper(keeper LeaseKeeper) {
	k.lease = keeper
}

func (k *Keeper) WithNodeKeeper(keeper NodeKeeper) {
	k.node = keeper
}

func (k *Keeper) WithProviderKeeper(keeper ProviderKeeper) {
	k.provider = keeper
}

func (k *Keeper) WithSessionKeeper(keeper SessionKeeper) {
	k.session = keeper
}

func (k *Keeper) WithSubscriptionKeeper(keeper SubscriptionKeeper) {
	k.subscription = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
