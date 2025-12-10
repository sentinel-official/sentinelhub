package types

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var (
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)

func SerializeCosmosQuery(reqs []abcitypes.RequestQuery) (bz []byte, err error) {
	q := &CosmosQuery{
		Requests: reqs,
	}

	return ModuleCdc.Marshal(q)
}

func DeserializeCosmosQuery(bz []byte) (reqs []abcitypes.RequestQuery, err error) {
	var q CosmosQuery

	err = ModuleCdc.Unmarshal(bz, &q)

	return q.Requests, err
}

func DeserializeCosmosResponse(bz []byte) (resps []abcitypes.ResponseQuery, err error) {
	var r CosmosResponse

	err = ModuleCdc.Unmarshal(bz, &r)

	return r.Responses, err
}
