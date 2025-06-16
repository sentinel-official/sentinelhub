package v3

import (
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

func NewQueryPlanRequest(id uint64) *QueryPlanRequest {
	return &QueryPlanRequest{
		Id: id,
	}
}

func NewQueryPlansRequest(status v1base.Status, pagination *sdkquery.PageRequest) *QueryPlansRequest {
	return &QueryPlansRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryPlansForProviderRequest(addr base.ProvAddress, status v1base.Status, pagination *sdkquery.PageRequest) *QueryPlansForProviderRequest {
	return &QueryPlansForProviderRequest{
		Address:    addr.String(),
		Status:     status,
		Pagination: pagination,
	}
}
