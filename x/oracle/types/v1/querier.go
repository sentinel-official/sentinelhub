package v1

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryAssetsRequest creates a new instance of QueryAssetsRequest.
func NewQueryAssetsRequest(pagination *query.PageRequest) *QueryAssetsRequest {
	return &QueryAssetsRequest{
		Pagination: pagination,
	}
}

// NewQueryAssetRequest creates a new instance of QueryAssetRequest.
func NewQueryAssetRequest(denom string) *QueryAssetRequest {
	return &QueryAssetRequest{
		Denom: denom,
	}
}

// NewQueryParamsRequest creates a new instance of QueryParamsRequest.
func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
