package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

const (
	ModuleName = "session"
)

var (
	CountKey  = []byte{0x00}
	ParamsKey = []byte{0x01}

	SessionKeyPrefix                = []byte{0x10}
	SessionForAccountKeyPrefix      = []byte{0x11}
	SessionForAllocationKeyPrefix   = []byte{0x12}
	SessionForNodeKeyPrefix         = []byte{0x13}
	SessionForPlanKeyPrefix         = []byte{0x14}
	SessionForSubscriptionKeyPrefix = []byte{0x15}
	SessionForInactiveAtKeyPrefix   = []byte{0x16}
)

func SessionKey(id uint64) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForAccountKeyPrefix(addr sdk.AccAddress) []byte {
	return append(SessionForAccountKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func SessionForAccountKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetSessionForAccountKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForNodeKeyPrefix(addr base.NodeAddress) []byte {
	return append(SessionForNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func SessionForNodeKey(addr base.NodeAddress, id uint64) []byte {
	return append(GetSessionForNodeKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForPlanKeyPrefix(id uint64) []byte {
	return append(SessionForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForPlanByNodeKeyPrefix(id uint64, addr base.NodeAddress) []byte {
	return append(GetSessionForPlanKeyPrefix(id), sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func SessionForPlanByNodeKey(planID uint64, addr base.NodeAddress, sessionID uint64) []byte {
	return append(GetSessionForPlanByNodeKeyPrefix(planID, addr), sdk.Uint64ToBigEndian(sessionID)...)
}

func GetSessionForSubscriptionKeyPrefix(id uint64) []byte {
	return append(SessionForSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SessionForSubscriptionKey(subscriptionID, sessionID uint64) []byte {
	return append(GetSessionForSubscriptionKeyPrefix(subscriptionID), sdk.Uint64ToBigEndian(sessionID)...)
}

func GetSessionForAllocationKeyPrefix(id uint64, addr sdk.AccAddress) []byte {
	return append(SessionForAllocationKeyPrefix, append(sdk.Uint64ToBigEndian(id), sdkaddress.MustLengthPrefix(addr)...)...)
}

func SessionForAllocationKey(subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) []byte {
	return append(GetSessionForAllocationKeyPrefix(subscriptionID, addr), sdk.Uint64ToBigEndian(sessionID)...)
}

func GetSessionForInactiveAtKeyPrefix(timestamp time.Time) []byte {
	return append(SessionForInactiveAtKeyPrefix, sdk.FormatTimeBytes(timestamp)...)
}

func SessionForInactiveAtKey(timestamp time.Time, id uint64) []byte {
	return append(GetSessionForInactiveAtKeyPrefix(timestamp), sdk.Uint64ToBigEndian(id)...)
}

func IDFromSessionForAccountKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)
	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSessionForNodeKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)
	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSessionForPlanByNodeKey(key []byte) uint64 {
	// prefix (1 byte) | planID (8 bytes) | addrLen (1 byte) | addr (addrLen bytes) | sessionID (8 bytes)
	addrLen := int(key[9])
	if len(key) != 18+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 18+addrLen))
	}

	return sdk.BigEndianToUint64(key[10+addrLen:])
}

func IDFromSessionForSubscriptionKey(key []byte) uint64 {
	// prefix (1 byte) | subscriptionID (8 bytes) | sessionID (8 bytes)
	if len(key) != 17 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 17))
	}

	return sdk.BigEndianToUint64(key[9:])
}

func IDFromSessionForAllocationKey(key []byte) uint64 {
	// prefix (1 byte) | subscriptionID (8 bytes) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)
	addrLen := int(key[9])
	if len(key) != 18+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 18+addrLen))
	}

	return sdk.BigEndianToUint64(key[10+addrLen:])
}

func IDFromSessionForInactiveAtKey(key []byte) uint64 {
	// prefix (1 byte) | timestamp (29 bytes) | session (8 bytes)
	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
