package types

import (
	"crypto/rand"
	"testing"

	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestActiveProviderKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(ActiveProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr)...),
				ActiveProviderKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveProviderKey(addr)
		})
	}
}

func TestInactiveProviderKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(InactiveProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr)...),
				InactiveProviderKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveProviderKey(addr)
		})
	}
}
