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

// HandleQueryLease handles a query to fetch a single lease by ID.
// Returns a gRPC NotFound error if the lease does not exist.
func (k *Keeper) HandleQueryLease(ctx sdk.Context, req *v1.QueryLeaseRequest) (*v1.QueryLeaseResponse, error) {
	item, found := k.GetLease(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "lease %d does not exist", req.Id)
	}

	return &v1.QueryLeaseResponse{Lease: item}, nil
}

// HandleQueryLeases handles a paginated query to list all leases in the store.
// Uses the LeaseKeyPrefix to iterate through the lease entries.
func (k *Keeper) HandleQueryLeases(ctx sdk.Context, req *v1.QueryLeasesRequest) (*v1.QueryLeasesResponse, error) {
	var (
		items []v1.Lease                                            // Collected leases
		store = prefix.NewStore(k.Store(ctx), types.LeaseKeyPrefix) // Prefixed store for all leases
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

// HandleQueryLeasesForNode handles a paginated query for all leases associated with a specific node.
// The node address is validated and used to derive a prefix for lease lookup.
func (k *Keeper) HandleQueryLeasesForNode(ctx sdk.Context, req *v1.QueryLeasesForNodeRequest) (*v1.QueryLeasesForNodeResponse, error) {
	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items     []v1.Lease                                 // Collected leases for the node
		keyPrefix = types.GetLeaseForNodeKeyPrefix(addr)     // Store prefix for node leases
		store     = prefix.NewStore(k.Store(ctx), keyPrefix) // Scoped store for the node's leases
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		// Use composite key to extract lease ID and retrieve full lease data
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

// HandleQueryLeasesForProvider handles a paginated query for all leases tied to a specific provider address.
// The address is validated and converted before deriving a prefix for filtering.
func (k *Keeper) HandleQueryLeasesForProvider(ctx sdk.Context, req *v1.QueryLeasesForProviderRequest) (*v1.QueryLeasesForProviderResponse, error) {
	addr, err := base.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items     []v1.Lease                                 // Collected leases for provider
		keyPrefix = types.GetLeaseForProviderKeyPrefix(addr) // Store prefix for provider leases
		store     = prefix.NewStore(k.Store(ctx), keyPrefix) // Scoped store for provider's leases
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		// Lookup full lease using composite key from prefixed store
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

// HandleQueryParams handles a request to fetch the module's current parameter configuration.
func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	params := k.GetParams(ctx)

	return &v1.QueryParamsResponse{Params: params}, nil
}
