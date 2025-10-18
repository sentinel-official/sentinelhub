package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestActivePlanKey(t *testing.T) {
	require.Equal(
		t,
		append(append(PlanKeyPrefix, 0x01), sdk.Uint64ToBigEndian(1000)...),
		ActivePlanKey(1000),
	)
}

func TestInactivePlanKey(t *testing.T) {
	require.Equal(
		t,
		append(append(PlanKeyPrefix, 0x02), sdk.Uint64ToBigEndian(1000)...),
		InactivePlanKey(1000),
	)
}

func TestPlanForProviderKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(PlanForProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(1000)...),
				PlanForProviderKey(addr, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			PlanForProviderKey(addr, 1000)
		})
	}
}

func TestIDFromPlanForProviderKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = PlanForProviderKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromPlanForProviderKey(key),
		)
	}
}
