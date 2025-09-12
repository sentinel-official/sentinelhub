package v1

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestNewQueryDepositRequest(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		require.Equal(
			t,
			&QueryDepositRequest{
				Address: sdk.AccAddress(addr).String(),
			},
			NewQueryDepositRequest(addr),
		)
	}
}

func TestNewQueryDepositsRequest(t *testing.T) {
	var (
		pagination *sdkquery.PageRequest
	)

	for i := range 20 {
		pagination = &sdkquery.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryDepositsRequest{
				Pagination: pagination,
			},
			NewQueryDepositsRequest(pagination),
		)
	}
}
