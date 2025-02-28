package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v3"
)

type NodeKeeper interface {
	Store(ctx sdk.Context) sdk.KVStore
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (v3.Plan, bool)
	SetPlan(ctx sdk.Context, plan v3.Plan)
	SetPlanForNodeByProvider(ctx sdk.Context, nodeAddr types.NodeAddress, provAddr types.ProvAddress, id uint64)
	SetPlanForProvider(ctx sdk.Context, addr types.ProvAddress, id uint64)
	Store(ctx sdk.Context) sdk.KVStore
}
