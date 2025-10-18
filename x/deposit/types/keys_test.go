package types

import (
	"crypto/rand"
	"testing"

	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestDepositKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(DepositKeyPrefix, sdkaddress.MustLengthPrefix(addr)...),
				DepositKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			DepositKey(addr)
		})
	}
}
