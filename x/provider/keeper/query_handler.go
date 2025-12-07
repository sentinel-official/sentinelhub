package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	v1base "github.com/sentinel-official/sentinelhub/v13/types/v1"
	"github.com/sentinel-official/sentinelhub/v13/x/provider/types"
	"github.com/sentinel-official/sentinelhub/v13/x/provider/types/v2"
	"github.com/sentinel-official/sentinelhub/v13/x/provider/types/v3"
)

// HandleQueryProvider handles a query to fetch a single provider by its address.
// Validates the Bech32 address and returns a gRPC NotFound error if the provider does not exist.
func (k *Keeper) HandleQueryProvider(ctx sdk.Context, req *v2.QueryProviderRequest) (*v2.QueryProviderResponse, error) {
	addr, err := base.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	// Retrieve the provider record associated with the address
	item, found := k.GetProvider(ctx, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "provider %s does not exist", req.Address)
	}

	return &v2.QueryProviderResponse{Provider: item}, nil
}

// HandleQueryProviders handles a paginated query to list all providers.
// Supports optional filtering based on provider status (active/inactive).
func (k *Keeper) HandleQueryProviders(ctx sdk.Context, req *v2.QueryProvidersRequest) (*v2.QueryProvidersResponse, error) {
	var (
		items     v2.Providers // Collected provider entries
		keyPrefix []byte       // Key prefix based on requested status
	)

	// Determine the prefix based on status filter
	switch req.Status {
	case v1base.StatusActive:
		keyPrefix = types.ActiveProviderKeyPrefix
	case v1base.StatusInactive:
		keyPrefix = types.InactiveProviderKeyPrefix
	default:
		keyPrefix = types.ProviderKeyPrefix
	}

	store := prefix.NewStore(k.Store(ctx), keyPrefix)

	// Paginate through the provider store and decode each entry
	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Provider
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryProvidersResponse{Providers: items, Pagination: pagination}, nil
}

// HandleQueryParams handles a request to retrieve the module's current parameter settings.
// Returns the Params object directly from the keeper.
func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	params := k.GetParams(ctx)

	return &v3.QueryParamsResponse{Params: params}, nil
}
