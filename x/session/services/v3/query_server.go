package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

var (
	_ v3.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v3.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QuerySession(c context.Context, req *v3.QuerySessionRequest) (*v3.QuerySessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySession(ctx, req)
}

func (q *queryServer) QuerySessions(c context.Context, req *v3.QuerySessionsRequest) (*v3.QuerySessionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySessions(ctx, req)
}

func (q *queryServer) QuerySessionsForAccount(c context.Context, req *v3.QuerySessionsForAccountRequest) (*v3.QuerySessionsForAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySessionsForAccount(ctx, req)
}

func (q *queryServer) QuerySessionsForNode(c context.Context, req *v3.QuerySessionsForNodeRequest) (*v3.QuerySessionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySessionsForNode(ctx, req)
}

func (q *queryServer) QuerySessionsForSubscription(c context.Context, req *v3.QuerySessionsForSubscriptionRequest) (*v3.QuerySessionsForSubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySessionsForSubscription(ctx, req)
}

func (q *queryServer) QuerySessionsForAllocation(c context.Context, req *v3.QuerySessionsForAllocationRequest) (*v3.QuerySessionsForAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySessionsForAllocation(ctx, req)
}

func (q *queryServer) QueryParams(c context.Context, req *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryParams(ctx, req)
}
