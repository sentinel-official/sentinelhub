package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

func TestSessionForAccountKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(SessionForAccountKeyPrefix, sdkaddress.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i))...),
				SessionForAccountKey(addr, uint64(i)),
			)

			continue
		}

		require.Panics(t, func() {
			SessionForAccountKey(addr, uint64(i))
		})
	}
}

func TestSessionForInactiveAtKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(append(SessionForInactiveAtKeyPrefix, sdk.FormatTimeBytes(base.TestTimeNow)...), sdk.Uint64ToBigEndian(uint64(i))...),
			SessionForInactiveAtKey(base.TestTimeNow, uint64(i)),
		)
	}
}

func TestSessionKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(SessionKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...),
			SessionKey(uint64(i)),
		)
	}
}

func TestIDFromSessionForAccountKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SessionForAccountKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForAccountKey(key),
		)
	}
}

func TestIDFromSessionForInactiveAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SessionForInactiveAtKey(base.TestTimeNow, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForInactiveAtKey(key),
		)
	}
}

func TestIDFromSessionForNodeKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SessionForNodeKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForNodeKey(key),
		)
	}
}

func TestIDFromSessionForAllocationKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SessionForAllocationKey(uint64(i+64), addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForAllocationKey(key),
		)
	}
}

func TestIDFromSessionForSubscriptionKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SessionForSubscriptionKey(uint64(i+64), uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForSubscriptionKey(key),
		)
	}
}

func TestSessionForNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(SessionForNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i))...),
				SessionForNodeKey(addr, uint64(i)),
			)

			continue
		}

		require.Panics(t, func() {
			SessionForNodeKey(addr, uint64(i))
		})
	}
}

func TestSessionForAllocationKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(append(SessionForAllocationKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), sdkaddress.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i+64))...),
				SessionForAllocationKey(uint64(i), addr, uint64(i+64)),
			)

			continue
		}

		require.Panics(t, func() {
			SessionForAllocationKey(uint64(i), addr, uint64(i+64))
		})
	}
}

func TestSessionForSubscriptionKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(append(SessionForSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), sdk.Uint64ToBigEndian(uint64(i+64))...),
			SessionForSubscriptionKey(uint64(i), uint64(i+64)),
		)
	}
}
