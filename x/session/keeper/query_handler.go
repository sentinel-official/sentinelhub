package keeper

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) HandleQuerySession(ctx sdk.Context, req *v3.QuerySessionRequest) (*v3.QuerySessionResponse, error) {
	v, found := k.GetSession(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "session does not exist for id %d", req.Id)
	}

	item, err := codectypes.NewAnyWithValue(v)
	if err != nil {
		return nil, err
	}

	return &v3.QuerySessionResponse{Session: item}, nil
}

func (k *Keeper) HandleQuerySessions(ctx sdk.Context, req *v3.QuerySessionsRequest) (*v3.QuerySessionsResponse, error) {
	var (
		items []*codectypes.Any
		store = prefix.NewStore(k.Store(ctx), types.SessionKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var v v3.Session
		if err := k.cdc.UnmarshalInterface(value, &v); err != nil {
			return err
		}
		if err := k.UpdateMaxValues(ctx, v); err != nil {
			return err
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

	return &v3.QuerySessionsResponse{Sessions: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQuerySessionsForAccount(ctx sdk.Context, req *v3.QuerySessionsForAccountRequest) (*v3.QuerySessionsForAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForAccountKeyPrefix(addr))
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

func (k *Keeper) HandleQuerySessionsForNode(ctx sdk.Context, req *v3.QuerySessionsForNodeRequest) (*v3.QuerySessionsForNodeResponse, error) {
	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForNodeKeyPrefix(addr))
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

func (k *Keeper) HandleQuerySessionsForSubscription(ctx sdk.Context, req *v3.QuerySessionsForSubscriptionRequest) (*v3.QuerySessionsForSubscriptionResponse, error) {
	var (
		items []*codectypes.Any
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForSubscriptionKeyPrefix(req.Id))
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

func (k *Keeper) HandleQuerySessionsForAllocation(ctx sdk.Context, req *v3.QuerySessionsForAllocationRequest) (*v3.QuerySessionsForAllocationResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForAllocationKeyPrefix(req.Id, addr))
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

func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	params := k.GetParams(ctx)
	return &v3.QueryParamsResponse{Params: params}, nil
}
