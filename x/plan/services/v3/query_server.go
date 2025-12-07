package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/sentinelhub/v13/x/plan/keeper"
	"github.com/sentinel-official/sentinelhub/v13/x/plan/types/v3"
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

func (q *queryServer) QueryPlan(c context.Context, req *v3.QueryPlanRequest) (*v3.QueryPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return q.HandleQueryPlan(ctx, req)
}

func (q *queryServer) QueryPlans(c context.Context, req *v3.QueryPlansRequest) (res *v3.QueryPlansResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return q.HandleQueryPlans(ctx, req)
}

func (q *queryServer) QueryPlansForProvider(c context.Context, req *v3.QueryPlansForProviderRequest) (res *v3.QueryPlansForProviderResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return q.HandleQueryPlansForProvider(ctx, req)
}
