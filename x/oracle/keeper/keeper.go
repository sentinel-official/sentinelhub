package keeper

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcporttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"

	"github.com/sentinel-official/sentinelhub/v13/x/oracle/types"
)

type Keeper struct {
	authority    string
	cdc          codec.Codec
	storeService store.KVStoreService

	ics4 ibcporttypes.ICS4Wrapper
}

func NewKeeper(
	cdc codec.Codec, storeService store.KVStoreService, ics4 ibcporttypes.ICS4Wrapper, authority string,
) Keeper {
	return Keeper{
		authority:    authority,
		cdc:          cdc,
		storeService: storeService,
		ics4:         ics4,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) storetypes.KVStore {
	return runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
}
