package migrations

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

type NodeKeeper interface {
	SetNode(ctx sdk.Context, node v3.Node)
	SetNodeForInactiveAt(ctx sdk.Context, time time.Time, addr types.NodeAddress)
	SetNodeForPlan(ctx sdk.Context, id uint64, addr types.NodeAddress)
	SetParams(ctx sdk.Context, params v3.Params)
	Store(ctx sdk.Context) sdk.KVStore
}
