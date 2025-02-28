package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types"
)

type Keeper struct {
	authority        string
	feeCollectorName string
	cdc              codec.BinaryCodec
	key              storetypes.StoreKey
	router           *baseapp.MsgServiceRouter

	bank    BankKeeper
	node    NodeKeeper
	oracle  OracleKeeper
	plan    PlanKeeper
	session SessionKeeper
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

func (k *Keeper) WithBankKeeper(keeper BankKeeper) {
	k.bank = keeper
}

func (k *Keeper) WithNodeKeeper(keeper NodeKeeper) {
	k.node = keeper
}

func (k *Keeper) WithOracleKeeper(keeper OracleKeeper) {
	k.oracle = keeper
}

func (k *Keeper) WithPlanKeeper(keeper PlanKeeper) {
	k.plan = keeper
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
