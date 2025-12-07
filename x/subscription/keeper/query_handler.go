package keeper

import (
	"fmt"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types"
	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types/v2"
	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types/v3"
)

// HandleQueryAllocation handles a query to fetch a specific allocation by subscription ID and address.
// Validates the address format and returns a NotFound error if the allocation does not exist.
func (k *Keeper) HandleQueryAllocation(ctx sdk.Context, req *v2.QueryAllocationRequest) (*v2.QueryAllocationResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	// Retrieve allocation for subscription ID and address
	item, found := k.GetAllocation(ctx, req.Id, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "allocation %d/%s does not exist", req.Id, req.Address)
	}

	return &v2.QueryAllocationResponse{Allocation: item}, nil
}

// HandleQueryAllocations handles a paginated query to fetch all allocations under a specific subscription.
// Uses a prefixed store based on subscription ID to retrieve allocations.
func (k *Keeper) HandleQueryAllocations(ctx sdk.Context, req *v2.QueryAllocationsRequest) (*v2.QueryAllocationsResponse, error) {
	var (
		items v2.Allocations                                                                       // Collected allocations
		store = prefix.NewStore(k.Store(ctx), types.GetAllocationForSubscriptionKeyPrefix(req.Id)) // Scoped store for allocations
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Allocation
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryAllocationsResponse{Allocations: items, Pagination: pagination}, nil
}

// HandleQueryParams handles a request to retrieve the current subscription module parameters.
// Returns the stored Params object as-is.
func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	params := k.GetParams(ctx)

	return &v3.QueryParamsResponse{Params: params}, nil
}

// HandleQuerySubscription handles a query to retrieve a single subscription by its ID.
// Returns a NotFound error if the subscription does not exist.
func (k *Keeper) HandleQuerySubscription(ctx sdk.Context, req *v3.QuerySubscriptionRequest) (*v3.QuerySubscriptionResponse, error) {
	item, found := k.GetSubscription(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "subscription %d does not exist", req.Id)
	}

	return &v3.QuerySubscriptionResponse{Subscription: item}, nil
}

// HandleQuerySubscriptions handles a paginated query to list all subscriptions in the store.
// Uses SubscriptionKeyPrefix to scope the iteration.
func (k *Keeper) HandleQuerySubscriptions(ctx sdk.Context, req *v3.QuerySubscriptionsRequest) (*v3.QuerySubscriptionsResponse, error) {
	var (
		items []v3.Subscription                                            // Collected subscriptions
		store = prefix.NewStore(k.Store(ctx), types.SubscriptionKeyPrefix) // Store of all subscriptions
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v3.Subscription
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySubscriptionsResponse{Subscriptions: items, Pagination: pagination}, nil
}

// HandleQuerySubscriptionsForAccount handles a paginated query for all subscriptions tied to a specific account address.
// Validates the address format and uses account-based prefix for store access.
func (k *Keeper) HandleQuerySubscriptionsForAccount(ctx sdk.Context, req *v3.QuerySubscriptionsForAccountRequest) (*v3.QuerySubscriptionsForAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []v3.Subscription                                                               // Collected subscriptions
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForAccountKeyPrefix(addr)) // Scoped store by account
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySubscriptionsForAccountResponse{Subscriptions: items, Pagination: pagination}, nil
}

// HandleQuerySubscriptionsForPlan handles a paginated query for all subscriptions under a specific plan ID.
// Uses the plan ID prefix to locate all related subscriptions in the store.
func (k *Keeper) HandleQuerySubscriptionsForPlan(ctx sdk.Context, req *v3.QuerySubscriptionsForPlanRequest) (*v3.QuerySubscriptionsForPlanResponse, error) {
	var (
		items []v3.Subscription                                                              // Collected subscriptions under the plan
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForPlanKeyPrefix(req.Id)) // Store for plan subscriptions
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		// Retrieve full subscription using key-derived ID
		item, found := k.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySubscriptionsForPlanResponse{Subscriptions: items, Pagination: pagination}, nil
}
