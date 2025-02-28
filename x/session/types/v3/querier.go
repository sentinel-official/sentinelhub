package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/sentinel-official/hub/v12/types"
)

func NewQuerySessionsRequest(pagination *sdkquery.PageRequest) *QuerySessionsRequest {
	return &QuerySessionsRequest{
		Pagination: pagination,
	}
}

func NewQuerySessionsForAccountRequest(addr sdk.AccAddress, pagination *sdkquery.PageRequest) *QuerySessionsForAccountRequest {
	return &QuerySessionsForAccountRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySessionsForNodeRequest(addr base.NodeAddress, pagination *sdkquery.PageRequest) *QuerySessionsForNodeRequest {
	return &QuerySessionsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySessionsForSubscriptionRequest(id uint64, pagination *sdkquery.PageRequest) *QuerySessionsForSubscriptionRequest {
	return &QuerySessionsForSubscriptionRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQuerySessionsForAllocationRequest(id uint64, addr sdk.AccAddress, pagination *sdkquery.PageRequest) *QuerySessionsForAllocationRequest {
	return &QuerySessionsForAllocationRequest{
		Id:         id,
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySessionRequest(id uint64) *QuerySessionRequest {
	return &QuerySessionRequest{
		Id: id,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
