package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName = "subscription"
)

var (
	CountKey  = []byte{0x00}
	ParamsKey = []byte{0x01}

	SubscriptionKeyPrefix              = []byte{0x10}
	SubscriptionForAccountKeyPrefix    = []byte{0x11}
	SubscriptionForPlanKeyPrefix       = []byte{0x12}
	SubscriptionForInactiveAtKeyPrefix = []byte{0x13}
	SubscriptionForRenewalAtKeyPrefix  = []byte{0x14}

	AllocationKeyPrefix = []byte{0x20}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForAccountKeyPrefix(addr sdk.AccAddress) []byte {
	return append(SubscriptionForAccountKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func SubscriptionForAccountKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetSubscriptionForAccountKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForPlanKeyPrefix(id uint64) []byte {
	return append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForPlanKey(planID, subscriptionID uint64) []byte {
	return append(GetSubscriptionForPlanKeyPrefix(planID), sdk.Uint64ToBigEndian(subscriptionID)...)
}

func GetSubscriptionForInactiveAtKeyPrefix(timestamp time.Time) []byte {
	return append(SubscriptionForInactiveAtKeyPrefix, sdk.FormatTimeBytes(timestamp)...)
}

func SubscriptionForInactiveAtKey(timestamp time.Time, id uint64) []byte {
	return append(GetSubscriptionForInactiveAtKeyPrefix(timestamp), sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForRenewalAtKeyPrefix(timestamp time.Time) []byte {
	return append(SubscriptionForRenewalAtKeyPrefix, sdk.FormatTimeBytes(timestamp)...)
}

func SubscriptionForRenewalAtKey(timestamp time.Time, id uint64) []byte {
	return append(GetSubscriptionForRenewalAtKeyPrefix(timestamp), sdk.Uint64ToBigEndian(id)...)
}

func GetAllocationForSubscriptionKeyPrefix(id uint64) []byte {
	return append(AllocationKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func AllocationKey(id uint64, addr sdk.AccAddress) []byte {
	return append(GetAllocationForSubscriptionKeyPrefix(id), sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func AccAddrFromSubscriptionForAccountKey(key []byte) sdk.AccAddress {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return key[2 : 2+addrLen]
}

func IDFromSubscriptionForAccountKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSubscriptionForPlanKey(key []byte) uint64 {
	// prefix (1 byte) | planID (8 bytes) | subscriptionID (8 bytes)

	if len(key) != 17 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 17))
	}

	return sdk.BigEndianToUint64(key[9:])
}

func IDFromSubscriptionForInactiveAtKey(key []byte) uint64 {
	// prefix (1 byte) | timestamp (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}

func IDFromSubscriptionForRenewalAtKey(key []byte) uint64 {
	// prefix (1 byte) | timestamp (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
