package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

const (
	ModuleName = "plan"
)

var (
	CountKey = []byte{0x00}

	PlanKeyPrefix            = []byte{0x10}
	ActivePlanKeyPrefix      = append(PlanKeyPrefix, 0x01)
	InactivePlanKeyPrefix    = append(PlanKeyPrefix, 0x02)
	PlanForNodeKeyPrefix     = []byte{0x11}
	PlanForProviderKeyPrefix = []byte{0x12}
)

func ActivePlanKey(id uint64) []byte {
	return append(ActivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func InactivePlanKey(id uint64) []byte {
	return append(InactivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetPlanForNodeKeyPrefix(addr base.NodeAddress) []byte {
	return append(PlanForNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetPlanForNodeByProviderKeyPrefix(nodeAddr base.NodeAddress, provAddr base.ProvAddress) []byte {
	return append(GetPlanForNodeKeyPrefix(nodeAddr), sdkaddress.MustLengthPrefix(provAddr.Bytes())...)
}

func PlanForNodeByProviderKey(nodeAddr base.NodeAddress, provAddr base.ProvAddress, id uint64) []byte {
	return append(GetPlanForNodeByProviderKeyPrefix(nodeAddr, provAddr), sdk.Uint64ToBigEndian(id)...)
}

func GetPlanForProviderKeyPrefix(addr base.ProvAddress) []byte {
	return append(PlanForProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func PlanForProviderKey(addr base.ProvAddress, id uint64) []byte {
	return append(GetPlanForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func IDFromPlanForNodeByProviderKey(key []byte) uint64 {
	// prefix (1 byte) | nodeAddrLen (1 byte) | nodeAddr (nodeAddrLen bytes) | provAddrLen (1 byte) | provAddr (provAddrLen bytes) | id (8 bytes)

	nodeAddrLen, provAddrLen := int(key[1]), int(key[2+int(key[1])])
	if len(key) != 11+nodeAddrLen+provAddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 11+nodeAddrLen+provAddrLen))
	}

	return sdk.BigEndianToUint64(key[3+nodeAddrLen+provAddrLen:])
}

func IDFromPlanForProviderKey(key []byte) uint64 {
	// prefix (1 bytes) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}
