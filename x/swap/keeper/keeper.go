package keeper

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sentinel-official/sentinelhub/v13/x/swap/types"
	"github.com/sentinel-official/sentinelhub/v13/x/swap/types/v1"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	params       paramstypes.Subspace
	storeService store.KVStoreService

	account AccountKeeper
	bank    BankKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec, storeService store.KVStoreService, params paramstypes.Subspace, account AccountKeeper,
	bank BankKeeper,
) Keeper {
	return Keeper{
		cdc:          cdc,
		params:       params.WithKeyTable(v1.ParamsKeyTable()),
		storeService: storeService,
		account:      account,
		bank:         bank,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) storetypes.KVStore {
	return runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
}
