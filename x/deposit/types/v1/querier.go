package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
)

func NewQueryDepositRequest(addr sdk.AccAddress) *QueryDepositRequest {
	return &QueryDepositRequest{
		Address: addr.String(),
	}
}

func NewQueryDepositsRequest(pagination *sdkquery.PageRequest) *QueryDepositsRequest {
	return &QueryDepositsRequest{
		Pagination: pagination,
	}
}
