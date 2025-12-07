package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"

	base "github.com/sentinel-official/sentinelhub/v13/types"
)

func TestActiveNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 1; i <= 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(ActiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr)...),
				ActiveNodeKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveNodeKey(addr)
		})
	}
}

func TestAddressFromNodeForInactiveAtKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = NodeForInactiveAtKey(base.TestTimeNow, addr)
		require.Equal(
			t,
			base.NodeAddress(addr),
			AddressFromNodeForInactiveAtKey(key),
		)
	}
}

func TestAddressFromNodeForPlanKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = NodeForPlanKey(uint64(i), addr)
		require.Equal(
			t,
			base.NodeAddress(addr),
			AddressFromNodeForPlanKey(key),
		)
	}
}

func TestInactiveNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 1; i <= 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(InactiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr)...),
				InactiveNodeKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveNodeKey(addr)
		})
	}
}

func TestNodeForInactiveAtKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 1; i <= 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(NodeForInactiveAtKeyPrefix, sdk.FormatTimeBytes(base.TestTimeNow)...), sdkaddress.MustLengthPrefix(addr)...),
				NodeForInactiveAtKey(base.TestTimeNow, addr),
			)

			continue
		}

		require.Panics(t, func() {
			NodeForInactiveAtKey(base.TestTimeNow, addr)
		})
	}
}

func TestNodeForPlanKey(t *testing.T) {
	var (
		addr []byte
		id   uint64
	)

	for i := 1; i <= 512; i += 64 {
		id = uint64(i)
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...), sdkaddress.MustLengthPrefix(addr)...),
				NodeForPlanKey(id, addr),
			)

			continue
		}

		require.Panics(t, func() {
			NodeForPlanKey(id, addr)
		})
	}
}
