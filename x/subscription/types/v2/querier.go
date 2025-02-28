package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
)

func NewQueryAllocationRequest(id uint64, addr sdk.AccAddress) *QueryAllocationRequest {
	return &QueryAllocationRequest{
		Id:      id,
		Address: addr.String(),
	}
}

func NewQueryAllocationsRequest(id uint64, pagination *sdkquery.PageRequest) *QueryAllocationsRequest {
	return &QueryAllocationsRequest{
		Id:         id,
		Pagination: pagination,
	}
}
