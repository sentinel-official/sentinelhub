package types

import (
	sdkmath "cosmossdk.io/math"
)

const (
	ModuleName = "swap"
	StoreKey   = ModuleName
)

var (
	PrecisionLoss = sdkmath.NewInt(100)
)

var (
	SwapKeyPrefix = []byte{0x10}
)

func SwapKey(hash EthereumHash) []byte {
	return append(SwapKeyPrefix, hash.Bytes()...)
}
