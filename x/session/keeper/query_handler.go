package keeper

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/x/session/types"
	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

// HandleQuerySession handles a query to retrieve a single session by its ID.
// The session is fetched from store and wrapped into an Any type for interface-safe encoding.
// Returns a gRPC NotFound error if the session does not exist.
func (k *Keeper) HandleQuerySession(ctx sdk.Context, req *v3.QuerySessionRequest) (*v3.QuerySessionResponse, error) {
	v, found := k.GetSession(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "session does not exist for id %d", req.Id)
	}

	// Wrap the session in an Any type for dynamic message encoding
	item, err := codectypes.NewAnyWithValue(v)
	if err != nil {
		return nil, err
	}

	return &v3.QuerySessionResponse{Session: item}, nil
}

// HandleQuerySessions handles a paginated query to list all sessions in the store.
// Each session is decoded, updated with max values (if needed), and wrapped into an Any type.
func (k *Keeper) HandleQuerySessions(ctx sdk.Context, req *v3.QuerySessionsRequest) (*v3.QuerySessionsResponse, error) {
	var (
		items []*codectypes.Any                                       // Collected wrapped sessions
		store = prefix.NewStore(k.Store(ctx), types.SessionKeyPrefix) // Prefixed store containing all sessions
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var v v3.Session
		if err := k.cdc.UnmarshalInterface(value, &v); err != nil {
			return err
		}

		// Apply max value updates before returning the session
		if err := k.UpdateMaxValues(ctx, v); err != nil {
			return err
		}

		// Wrap the session in Any
		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySessionsResponse{Sessions: items, Pagination: pagination}, nil
}

// HandleQuerySessionsForAccount handles a paginated query for sessions tied to a specific account address.
// Validates the address, retrieves session keys, and fetches full session data.
func (k *Keeper) HandleQuerySessionsForAccount(ctx sdk.Context, req *v3.QuerySessionsForAccountRequest) (*v3.QuerySessionsForAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any                                                          // Collected session objects
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForAccountKeyPrefix(addr)) // Scoped store by account address
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySessionsForAccountResponse{Sessions: items, Pagination: pagination}, nil
}

// HandleQuerySessionsForNode handles a paginated query for sessions tied to a specific node address.
// Validates the node address and retrieves session data based on the node-session relationship.
func (k *Keeper) HandleQuerySessionsForNode(ctx sdk.Context, req *v3.QuerySessionsForNodeRequest) (*v3.QuerySessionsForNodeResponse, error) {
	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any                                                       // Collected sessions
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForNodeKeyPrefix(addr)) // Store scoped to node
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySessionsForNodeResponse{Sessions: items, Pagination: pagination}, nil
}

// HandleQuerySessionsForSubscription handles a paginated query to list all sessions under a specific subscription ID.
// Sessions are retrieved using the subscription ID key prefix.
func (k *Keeper) HandleQuerySessionsForSubscription(ctx sdk.Context, req *v3.QuerySessionsForSubscriptionRequest) (*v3.QuerySessionsForSubscriptionResponse, error) {
	var (
		items []*codectypes.Any                                                                 // Session list
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForSubscriptionKeyPrefix(req.Id)) // Store by subscription ID
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySessionsForSubscriptionResponse{Sessions: items, Pagination: pagination}, nil
}

// HandleQuerySessionsForAllocation handles a paginated query to list all sessions for a specific allocation.
// The allocation is defined by a subscription ID and an account address.
func (k *Keeper) HandleQuerySessionsForAllocation(ctx sdk.Context, req *v3.QuerySessionsForAllocationRequest) (*v3.QuerySessionsForAllocationResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any                                                                     // Sessions for allocation
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForAllocationKeyPrefix(req.Id, addr)) // Keyed by allocation (sub ID + address)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySessionsForAllocationResponse{Sessions: items, Pagination: pagination}, nil
}

// HandleQueryParams handles a request to fetch the current session module parameters.
// This function simply retrieves and returns the stored Params object.
func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	params := k.GetParams(ctx)

	return &v3.QueryParamsResponse{Params: params}, nil
}
