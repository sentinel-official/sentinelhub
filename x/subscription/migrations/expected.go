package migrations

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
	v3plan "github.com/sentinel-official/hub/v12/x/plan/types/v3"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

type DepositKeeper interface {
	SubtractDeposit(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error
}

type LeaseKeeper interface {
	GetLeaseCount(ctx sdk.Context) uint64
	SetLeaseCount(ctx sdk.Context, count uint64)
	SetLease(ctx sdk.Context, lease v1.Lease)
	SetLeaseForInactiveAt(ctx sdk.Context, time time.Time, id uint64)
	SetLeaseForNodeByProvider(ctx sdk.Context, nodeAddr types.NodeAddress, provAddr types.ProvAddress, id uint64)
	SetLeaseForPayoutAt(ctx sdk.Context, time time.Time, id uint64)
	SetLeaseForProvider(ctx sdk.Context, addr types.ProvAddress, id uint64)
	SetLeaseForRenewalAt(ctx sdk.Context, time time.Time, id uint64)
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plan v3plan.Plan, found bool)
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr types.ProvAddress) bool
}

type SubscriptionKeeper interface {
	DeleteAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress)
	GetAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) (v2.Allocation, bool)
	IterateAllocationsForSubscription(ctx sdk.Context, id uint64, fn func(index int, item v2.Allocation) (stop bool))
	IterateSubscriptions(ctx sdk.Context, fn func(index int, item v3.Subscription) (stop bool))
	SetParams(ctx sdk.Context, params v3.Params)
	SetSubscription(ctx sdk.Context, subscription v3.Subscription)
	SetSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64)
	SetSubscriptionForInactiveAt(ctx sdk.Context, time time.Time, id uint64)
	SetSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64)
	SetSubscriptionForRenewalAt(ctx sdk.Context, time time.Time, id uint64)
	Store(ctx sdk.Context) sdk.KVStore
}
