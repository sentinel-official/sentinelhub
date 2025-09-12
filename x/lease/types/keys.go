package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

const (
	ModuleName = "lease"
)

var (
	CountKey                    = []byte{0x00}
	ParamsKey                   = []byte{0x01}
	LeaseKeyPrefix              = []byte{0x10}
	LeaseForNodeKeyPrefix       = []byte{0x11}
	LeaseForProviderKeyPrefix   = []byte{0x12}
	LeaseForInactiveAtKeyPrefix = []byte{0x13}
	LeaseForPayoutAtKeyPrefix   = []byte{0x14}
	LeaseForRenewalAtKeyPrefix  = []byte{0x15}
)

func LeaseKey(id uint64) []byte {
	return append(LeaseKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForNodeKeyPrefix(addr base.NodeAddress) []byte {
	return append(LeaseForNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetLeaseForNodeByProviderKeyPrefix(nodeAddr base.NodeAddress, provAddr base.ProvAddress) []byte {
	return append(GetLeaseForNodeKeyPrefix(nodeAddr), sdkaddress.MustLengthPrefix(provAddr.Bytes())...)
}

func LeaseForNodeByProviderKey(nodeAddr base.NodeAddress, provAddr base.ProvAddress, id uint64) []byte {
	return append(GetLeaseForNodeByProviderKeyPrefix(nodeAddr, provAddr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForProviderKeyPrefix(addr base.ProvAddress) []byte {
	return append(LeaseForProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func LeaseForProviderKey(addr base.ProvAddress, id uint64) []byte {
	return append(GetLeaseForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForInactiveAtKeyPrefix(timestamp time.Time) []byte {
	return append(LeaseForInactiveAtKeyPrefix, sdk.FormatTimeBytes(timestamp)...)
}

func LeaseForInactiveAtKey(timestamp time.Time, id uint64) []byte {
	return append(GetLeaseForInactiveAtKeyPrefix(timestamp), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForPayoutAtKeyPrefix(timestamp time.Time) []byte {
	return append(LeaseForPayoutAtKeyPrefix, sdk.FormatTimeBytes(timestamp)...)
}

func LeaseForPayoutAtKey(timestamp time.Time, id uint64) []byte {
	return append(GetLeaseForPayoutAtKeyPrefix(timestamp), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForRenewalAtKeyPrefix(timestamp time.Time) []byte {
	return append(LeaseForRenewalAtKeyPrefix, sdk.FormatTimeBytes(timestamp)...)
}

func LeaseForRenewalAtKey(timestamp time.Time, id uint64) []byte {
	return append(GetLeaseForRenewalAtKeyPrefix(timestamp), sdk.Uint64ToBigEndian(id)...)
}

func IDFromLeaseForNodeByProviderKey(key []byte) uint64 {
	// prefix (1 byte) | nodeAddrLen (1 byte) | nodeAddr (nodeAddrLen bytes) | provAddrLen (1 byte) | provAddr (provAddrLen bytes) | id (8 bytes)
	nodeAddrLen, provAddrLen := int(key[1]), int(key[2+int(key[1])])
	if len(key) != 11+nodeAddrLen+provAddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 11+nodeAddrLen+provAddrLen))
	}

	return sdk.BigEndianToUint64(key[3+nodeAddrLen+provAddrLen:])
}

func IDFromLeaseForProviderKey(key []byte) uint64 {
	// prefix (1 bytes) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)
	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromLeaseForInactiveAtKey(key []byte) uint64 {
	// prefix (1 byte) | timestamp (29 bytes) | id (8 bytes)
	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}

func IDFromLeaseForPayoutAtKey(key []byte) uint64 {
	// prefix (1 byte) | timestamp (29 bytes) | id (8 bytes)
	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}

func IDFromLeaseForRenewalAtKey(key []byte) uint64 {
	// prefix (1 byte) | timestamp (29 bytes) | id (8 bytes)
	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
