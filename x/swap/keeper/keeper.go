package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sentinel-official/hub/v12/x/swap/types"
	"github.com/sentinel-official/hub/v12/x/swap/types/v1"
)

type Keeper struct {
	cdc    codec.BinaryCodec
	key    storetypes.StoreKey
	params paramstypes.Subspace

	account AccountKeeper
	bank    BankKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, params paramstypes.Subspace, account AccountKeeper, bank BankKeeper) Keeper {
	return Keeper{
		cdc:     cdc,
		key:     key,
		params:  params.WithKeyTable(v1.ParamsKeyTable()),
		account: account,
		bank:    bank,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}
