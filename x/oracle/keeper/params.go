package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
)

// SetParams stores the oracle module parameters in the module's KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the oracle module parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)

	return v
}

// GetBlockInterval retrieves the BlockInterval parameter from the module's parameters.
func (k *Keeper) GetBlockInterval(ctx sdk.Context) int64 {
	return k.GetParams(ctx).BlockInterval
}

// GetChannelID retrieves the ChannelID parameter from the module's parameters.
func (k *Keeper) GetChannelID(ctx sdk.Context) string {
	return k.GetParams(ctx).ChannelID
}

// GetTimeout retrieves the Timeout parameter from the module's parameters.
func (k *Keeper) GetTimeout(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).Timeout
}

// GetQueryTimeout returns the current block time adjusted by the module's timeout parameter in Unix nanoseconds.
func (k *Keeper) GetQueryTimeout(ctx sdk.Context) int64 {
	t := k.GetTimeout(ctx)

	return ctx.BlockTime().Add(t).UnixNano()
}
