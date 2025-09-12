package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/keeper"
	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
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

func (q *queryServer) QueryAssets(c context.Context, req *v1.QueryAssetsRequest) (*v1.QueryAssetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return q.HandleQueryAssets(ctx, req)
}

func (q *queryServer) QueryAsset(c context.Context, req *v1.QueryAssetRequest) (*v1.QueryAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return q.HandleQueryAsset(ctx, req)
}

func (q *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return q.HandleQueryParams(ctx, req)
}
