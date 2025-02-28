package v2

import (
	"crypto/rand"
	"testing"

	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func TestNewQueryProviderRequest(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 40; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		require.Equal(
			t,
			&QueryProviderRequest{
				Address: base.ProvAddress(address).String(),
			},
			NewQueryProviderRequest(address),
		)
	}
}

func TestNewQueryProvidersRequest(t *testing.T) {
	var (
		status     v1base.Status
		pagination *sdkquery.PageRequest
	)

	for i := 0; i < 20; i++ {
		status = v1base.Status(i % 4)
		pagination = &sdkquery.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryProvidersRequest{
				Status:     status,
				Pagination: pagination,
			},
			NewQueryProvidersRequest(status, pagination),
		)
	}
}
