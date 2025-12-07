package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"

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

func (k *Keeper) BindPort(ctx sdk.Context, portID string) error {
	capability := k.port.BindPort(ctx, portID)

	return k.ClaimCapability(ctx, capability, ibchost.PortPath(portID))
}

func (k *Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.capability.GetCapability(ctx, ibchost.PortPath(portID))

	return ok
}
