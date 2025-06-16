package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
)

func (k *Keeper) HandleQueryAssets(ctx sdk.Context, req *v1.QueryAssetsRequest) (res *v1.QueryAssetsResponse, err error) {
	var (
		items []v1.Asset
		store = prefix.NewStore(k.Store(ctx), types.AssetKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v1.Asset
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryAssetsResponse{Assets: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryAsset(ctx sdk.Context, req *v1.QueryAssetRequest) (*v1.QueryAssetResponse, error) {
	item, found := k.GetAsset(ctx, req.Denom)
	if !found {
		return nil, status.Errorf(codes.NotFound, "asset %s does not exist", req.Denom)
	}

	return &v1.QueryAssetResponse{Asset: item}, nil
}

func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	params := k.GetParams(ctx)
	return &v1.QueryParamsResponse{Params: params}, nil
}
