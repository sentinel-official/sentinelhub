package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/deposit/types"
)

type Keeper struct {
	key storetypes.StoreKey
	cdc codec.BinaryCodec

	bank BankKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey) Keeper {
	return Keeper{
		key: key,
		cdc: cdc,
	}
}

func (k *Keeper) WithBankKeeper(keeper BankKeeper) {
	k.bank = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
