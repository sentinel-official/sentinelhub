package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

func TestSubscriptionForInactiveAtKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(append(SubscriptionForInactiveAtKeyPrefix, sdk.FormatTimeBytes(base.TestTimeNow)...), sdk.Uint64ToBigEndian(uint64(i))...),
			SubscriptionForInactiveAtKey(base.TestTimeNow, uint64(i)),
		)
	}
}

func TestAllocationKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(AllocationKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), sdkaddress.MustLengthPrefix(addr)...),
				AllocationKey(uint64(i), addr),
			)

			continue
		}

		require.Panics(t, func() {
			AllocationKey(uint64(i), addr)
		})
	}
}

func TestSubscriptionForAccountKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(SubscriptionForAccountKeyPrefix, sdkaddress.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i))...),
				SubscriptionForAccountKey(addr, uint64(i)),
			)

			continue
		}

		require.Panics(t, func() {
			SubscriptionForAccountKey(addr, uint64(i))
		})
	}
}

func TestSubscriptionForPlanKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), sdk.Uint64ToBigEndian(uint64(i))...),
			SubscriptionForPlanKey(uint64(i), uint64(i)),
		)
	}
}

func TestSubscriptionKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...),
			SubscriptionKey(uint64(i)),
		)
	}
}

func TestAccAddrFromSubscriptionForAccountKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SubscriptionForAccountKey(addr, uint64(i))
		require.Equal(
			t,
			sdk.AccAddress(addr),
			AccAddrFromSubscriptionForAccountKey(key),
		)
	}
}

func TestIDFromSubscriptionForAccountKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SubscriptionForAccountKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForAccountKey(key),
		)
	}
}

func TestIDFromSubscriptionForInactiveAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SubscriptionForInactiveAtKey(base.TestTimeNow, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForInactiveAtKey(key),
		)
	}
}

func TestIDFromSubscriptionForPlanKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SubscriptionForPlanKey(uint64(i+64), uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForPlanKey(key),
		)
	}
}
