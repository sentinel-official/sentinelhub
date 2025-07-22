package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/sentinelhub/v12/x/lease/keeper"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types/v1"
)

var (
	_ v1.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v1.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QueryLease(c context.Context, req *v1.QueryLeaseRequest) (*v1.QueryLeaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryLease(ctx, req)
}

func (q *queryServer) QueryLeases(c context.Context, req *v1.QueryLeasesRequest) (*v1.QueryLeasesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryLeases(ctx, req)
}

func (q *queryServer) QueryLeasesForNode(c context.Context, req *v1.QueryLeasesForNodeRequest) (*v1.QueryLeasesForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryLeasesForNode(ctx, req)
}

func (q *queryServer) QueryLeasesForProvider(c context.Context, req *v1.QueryLeasesForProviderRequest) (*v1.QueryLeasesForProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryLeasesForProvider(ctx, req)
}

func (q *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryParams(ctx, req)
}
