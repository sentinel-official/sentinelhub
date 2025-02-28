package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

var (
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v2.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QueryAllocation(c context.Context, req *v2.QueryAllocationRequest) (*v2.QueryAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryAllocation(ctx, req)
}

func (q *queryServer) QueryAllocations(c context.Context, req *v2.QueryAllocationsRequest) (*v2.QueryAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryAllocations(ctx, req)
}

func (q *queryServer) QuerySubscriptions(_ context.Context, _ *v2.QuerySubscriptionsRequest) (*v2.QuerySubscriptionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySubscriptionsForAccount(_ context.Context, _ *v2.QuerySubscriptionsForAccountRequest) (*v2.QuerySubscriptionsForAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySubscriptionsForNode(_ context.Context, _ *v2.QuerySubscriptionsForNodeRequest) (*v2.QuerySubscriptionsForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySubscriptionsForPlan(_ context.Context, _ *v2.QuerySubscriptionsForPlanRequest) (*v2.QuerySubscriptionsForPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySubscription(_ context.Context, _ *v2.QuerySubscriptionRequest) (*v2.QuerySubscriptionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPayouts(_ context.Context, _ *v2.QueryPayoutsRequest) (*v2.QueryPayoutsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPayoutsForAccount(_ context.Context, _ *v2.QueryPayoutsForAccountRequest) (*v2.QueryPayoutsForAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPayoutsForNode(_ context.Context, _ *v2.QueryPayoutsForNodeRequest) (*v2.QueryPayoutsForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPayout(_ context.Context, _ *v2.QueryPayoutRequest) (*v2.QueryPayoutResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryParams(c context.Context, req *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
