package v1

import (
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/sentinel-official/sentinelhub/v13/types"
)

func NewQueryLeaseRequest(id uint64) *QueryLeaseRequest {
	return &QueryLeaseRequest{
		Id: id,
	}
}

func NewQueryLeasesRequest(pagination *sdkquery.PageRequest) *QueryLeasesRequest {
	return &QueryLeasesRequest{
		Pagination: pagination,
	}
}

func NewQueryLeasesForNodeRequest(addr base.NodeAddress, pagination *sdkquery.PageRequest) *QueryLeasesForNodeRequest {
	return &QueryLeasesForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryLeasesForProviderRequest(addr base.ProvAddress, pagination *sdkquery.PageRequest) *QueryLeasesForProviderRequest {
	return &QueryLeasesForProviderRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
