package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) HandleQueryAllocation(ctx sdk.Context, req *v2.QueryAllocationRequest) (*v2.QueryAllocationResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	item, found := k.GetAllocation(ctx, req.Id, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "allocation %d/%s does not exist", req.Id, req.Address)
	}

	return &v2.QueryAllocationResponse{Allocation: item}, nil
}

func (k *Keeper) HandleQueryAllocations(ctx sdk.Context, req *v2.QueryAllocationsRequest) (*v2.QueryAllocationsResponse, error) {
	var (
		items v2.Allocations
		store = prefix.NewStore(k.Store(ctx), types.GetAllocationForSubscriptionKeyPrefix(req.Id))
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

func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	params := k.GetParams(ctx)
	return &v3.QueryParamsResponse{Params: params}, nil
}

func (k *Keeper) HandleQuerySubscription(ctx sdk.Context, req *v3.QuerySubscriptionRequest) (*v3.QuerySubscriptionResponse, error) {
	item, found := k.GetSubscription(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "subscription %d does not exist", req.Id)
	}

	return &v3.QuerySubscriptionResponse{Subscription: item}, nil
}

func (k *Keeper) HandleQuerySubscriptions(ctx sdk.Context, req *v3.QuerySubscriptionsRequest) (*v3.QuerySubscriptionsResponse, error) {
	var (
		items []v3.Subscription
		store = prefix.NewStore(k.Store(ctx), types.SubscriptionKeyPrefix)
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

func (k *Keeper) HandleQuerySubscriptionsForAccount(ctx sdk.Context, req *v3.QuerySubscriptionsForAccountRequest) (*v3.QuerySubscriptionsForAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []v3.Subscription
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForAccountKeyPrefix(addr))
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

func (k *Keeper) HandleQuerySubscriptionsForPlan(ctx sdk.Context, req *v3.QuerySubscriptionsForPlanRequest) (*v3.QuerySubscriptionsForPlanResponse, error) {
	var (
		items []v3.Subscription
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForPlanKeyPrefix(req.Id))
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

	return &v3.QuerySubscriptionsForPlanResponse{Subscriptions: items, Pagination: pagination}, nil
}
