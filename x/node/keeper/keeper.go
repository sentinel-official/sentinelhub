package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/node/types"
)

type Keeper struct {
	authority        string
	feeCollectorName string
	cdc              codec.BinaryCodec
	key              storetypes.StoreKey
	router           *baseapp.MsgServiceRouter

	deposit      DepositKeeper
	distribution DistributionKeeper
	lease        LeaseKeeper
	oracle       OracleKeeper
	session      SessionKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, router *baseapp.MsgServiceRouter, authority, feeCollectorName string) Keeper {
	return Keeper{
		authority:        authority,
		feeCollectorName: feeCollectorName,
		cdc:              cdc,
		key:              key,
		router:           router,
	}
}

func (k *Keeper) WithDepositKeeper(keeper DepositKeeper) {
	k.deposit = keeper
}

func (k *Keeper) WithDistributionKeeper(keeper DistributionKeeper) {
	k.distribution = keeper
}

func (k *Keeper) WithLeaseKeeper(keeper LeaseKeeper) {
	k.lease = keeper
}

func (k *Keeper) WithOracleKeeper(keeper OracleKeeper) {
	k.oracle = keeper
}

func (k *Keeper) WithSessionKeeper(keeper SessionKeeper) {
	k.session = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
