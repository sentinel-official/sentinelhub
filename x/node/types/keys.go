package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

const (
	ModuleName = "node"
)

var (
	ParamsKey = []byte{0x00}

	NodeKeyPrefix              = []byte{0x10}
	ActiveNodeKeyPrefix        = append(NodeKeyPrefix, 0x01)
	InactiveNodeKeyPrefix      = append(NodeKeyPrefix, 0x02)
	NodeForPlanKeyPrefix       = []byte{0x11}
	NodeForInactiveAtKeyPrefix = []byte{0x12}
)

func ActiveNodeKey(addr base.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func InactiveNodeKey(addr base.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetNodeForPlanKeyPrefix(id uint64) []byte {
	return append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func NodeForPlanKey(id uint64, addr base.NodeAddress) []byte {
	return append(GetNodeForPlanKeyPrefix(id), sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetNodeForInactiveAtKeyPrefix(at time.Time) []byte {
	return append(NodeForInactiveAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func NodeForInactiveAtKey(at time.Time, addr base.NodeAddress) []byte {
	return append(GetNodeForInactiveAtKeyPrefix(at), sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func AddressFromNodeForPlanKey(key []byte) base.NodeAddress {
	// prefix (1 byte) | id (8 bytes) | addrLen (1 byte) | addr (addrLen bytes)

	addrLen := int(key[9])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return key[10:]
}

func AddressFromNodeForInactiveAtKey(key []byte) base.NodeAddress {
	// prefix (1 byte) | at (29 bytes) | addrLen (1 byte) | addr (addrLen bytes)

	addrLen := int(key[30])
	if len(key) != 31+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 31+addrLen))
	}

	return key[31:]
}
