package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/oracle/types"
)

// GetPortID returns the portID for the oracle module.
func (k *Keeper) GetPortID(ctx sdk.Context) string {
	store := k.Store(ctx)

	return string(store.Get(types.PortIDKey))
}

// SetPortID sets the portID for the oracle module.
func (k *Keeper) SetPortID(ctx sdk.Context, portID string) {
	store := k.Store(ctx)
	store.Set(types.PortIDKey, []byte(portID))
}
