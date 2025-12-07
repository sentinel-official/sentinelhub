package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PortKeeper defines the expected IBC port keeper.
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}
