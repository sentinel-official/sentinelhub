package v1

import (
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	"github.com/sentinel-official/sentinelhub/v13/x/swap/types"
)

func NewQuerySwapRequest(txHash types.EthereumHash) *QuerySwapRequest {
	return &QuerySwapRequest{
		TxHash: txHash.Bytes(),
	}
}

func NewQuerySwapsRequest(pagination *sdkquery.PageRequest) *QuerySwapsRequest {
	return &QuerySwapsRequest{
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
