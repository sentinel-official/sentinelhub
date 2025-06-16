package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types/v1"
)

func (k *Keeper) HandleQueryLease(ctx sdk.Context, req *v1.QueryLeaseRequest) (*v1.QueryLeaseResponse, error) {
	item, found := k.GetLease(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "lease %d does not exist", req.Id)
	}

	return &v1.QueryLeaseResponse{Lease: item}, nil
}

func (k *Keeper) HandleQueryLeases(ctx sdk.Context, req *v1.QueryLeasesRequest) (*v1.QueryLeasesResponse, error) {
	var (
		items []v1.Lease
		store = prefix.NewStore(k.Store(ctx), types.LeaseKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v1.Lease
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryLeasesResponse{Leases: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryLeasesForProvider(ctx sdk.Context, req *v1.QueryLeasesForProviderRequest) (*v1.QueryLeasesForProviderResponse, error) {
	addr, err := base.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items     []v1.Lease
		keyPrefix = types.GetLeaseForProviderKeyPrefix(addr)
		store     = prefix.NewStore(k.Store(ctx), keyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetLease(ctx, types.IDFromLeaseForProviderKey(append(keyPrefix, key...)))
		if !found {
			return fmt.Errorf("lease for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryLeasesForProviderResponse{Leases: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryLeasesForNode(ctx sdk.Context, req *v1.QueryLeasesForNodeRequest) (*v1.QueryLeasesForNodeResponse, error) {
	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items     []v1.Lease
		keyPrefix = types.GetLeaseForNodeKeyPrefix(addr)
		store     = prefix.NewStore(k.Store(ctx), keyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetLease(ctx, types.IDFromLeaseForNodeByProviderKey(append(keyPrefix, key...)))
		if !found {
			return fmt.Errorf("lease for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryLeasesForNodeResponse{Leases: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	params := k.GetParams(ctx)
	return &v1.QueryParamsResponse{Params: params}, nil
}
