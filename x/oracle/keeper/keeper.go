package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
)

type Keeper struct {
	authority string
	cdc       codec.Codec
	key       storetypes.StoreKey

	capability ibcexported.ScopedKeeper
	ics4       ibcporttypes.ICS4Wrapper
}

func NewKeeper(
	cdc codec.Codec, key storetypes.StoreKey, capability ibcexported.ScopedKeeper, ics4 ibcporttypes.ICS4Wrapper,
	authority string,
) Keeper {
	return Keeper{
		authority:  authority,
		cdc:        cdc,
		key:        key,
		capability: capability,
		ics4:       ics4,
	}
}

func (k *Keeper) GetAuthority() string {
	return k.authority
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}
