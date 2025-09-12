package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/sentinelhub/v12/x/swap/types"
	"github.com/sentinel-official/sentinelhub/v12/x/swap/types/v1"
)

func (k *Keeper) HandleQuerySwap(ctx sdk.Context, req *v1.QuerySwapRequest) (*v1.QuerySwapResponse, error) {
	item, found := k.GetSwap(ctx, types.BytesToHash(req.TxHash))
	if !found {
		return nil, status.Errorf(codes.NotFound, "swap does not exist for hash %X", req.TxHash)
	}

	return &v1.QuerySwapResponse{Swap: item}, nil
}

func (k *Keeper) HandleQuerySwaps(ctx sdk.Context, req *v1.QuerySwapsRequest) (*v1.QuerySwapsResponse, error) {
	var (
		items v1.Swaps
		store = prefix.NewStore(k.Store(ctx), types.SwapKeyPrefix)
	)

	pagination, err := sdkquery.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			var item v1.Swap
			if err := k.cdc.Unmarshal(value, &item); err != nil {
				return false, err
			}

			items = append(items, item)
		}

		return true, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QuerySwapsResponse{Swaps: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	params := k.GetParams(ctx)

	return &v1.QueryParamsResponse{Params: params}, nil
}
