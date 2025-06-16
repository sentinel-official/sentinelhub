package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/sentinelhub/v12/x/deposit/types"
	"github.com/sentinel-official/sentinelhub/v12/x/deposit/types/v1"
)

func (k *Keeper) HandleQueryDeposit(ctx sdk.Context, req *v1.QueryDepositRequest) (*v1.QueryDepositResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	item, found := k.GetDeposit(ctx, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "deposit does not exist for address %s", req.Address)
	}

	return &v1.QueryDepositResponse{Deposit: item}, nil
}

func (k *Keeper) HandleQueryDeposits(ctx sdk.Context, req *v1.QueryDepositsRequest) (*v1.QueryDepositsResponse, error) {
	var (
		items v1.Deposits
		store = prefix.NewStore(k.Store(ctx), types.DepositKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_ []byte, value []byte) error {
		var item v1.Deposit
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryDepositsResponse{Deposits: items, Pagination: pagination}, nil
}
