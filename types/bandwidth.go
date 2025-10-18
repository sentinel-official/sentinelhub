package types

import (
	sdkmath "cosmossdk.io/math"
)

var (
	Kilobyte = sdkmath.NewInt(1000)
	Megabyte = sdkmath.NewInt(1000).Mul(Kilobyte)
	Gigabyte = sdkmath.NewInt(1000).Mul(Megabyte)
)
