package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types/v3"
)

func (k *Keeper) HandleQueryPlan(ctx sdk.Context, req *v3.QueryPlanRequest) (*v3.QueryPlanResponse, error) {
	item, found := k.GetPlan(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "plan does not exist for id %d", req.Id)
	}

	return &v3.QueryPlanResponse{Plan: item}, nil
}

func (k *Keeper) HandleQueryPlans(ctx sdk.Context, req *v3.QueryPlansRequest) (res *v3.QueryPlansResponse, err error) {
	var (
		items     []v3.Plan
		keyPrefix []byte
	)

	switch req.Status {
	case v1base.StatusActive:
		keyPrefix = types.ActivePlanKeyPrefix
	case v1base.StatusInactive:
		keyPrefix = types.InactivePlanKeyPrefix
	default:
		keyPrefix = types.PlanKeyPrefix
	}

	store := prefix.NewStore(k.Store(ctx), keyPrefix)
	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v3.Plan
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QueryPlansResponse{Plans: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryPlansForProvider(ctx sdk.Context, req *v3.QueryPlansForProviderRequest) (res *v3.QueryPlansForProviderResponse, err error) {
	addr, err := base.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []v3.Plan
		store = prefix.NewStore(k.Store(ctx), types.GetPlanForProviderKeyPrefix(addr))
	)

	pagination, err := sdkquery.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		if !accumulate {
			return false, nil
		}

		item, found := k.GetPlan(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return false, fmt.Errorf("plan for key %X does not exist", key)
		}

		if req.Status.Equal(v1base.StatusUnspecified) || item.Status.Equal(req.Status) {
			items = append(items, item)
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QueryPlansForProviderResponse{Plans: items, Pagination: pagination}, nil
}
