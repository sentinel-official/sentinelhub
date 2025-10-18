package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName = "deposit"
)

var (
	DepositKeyPrefix = []byte{0x10}
)

func DepositKey(addr sdk.AccAddress) []byte {
	return append(DepositKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}
