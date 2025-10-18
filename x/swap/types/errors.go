package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrorInvalidField    = sdkerrors.Register(ModuleName, 101, "invalid field")
	ErrorInvalidFrom     = sdkerrors.Register(ModuleName, 102, "invalid from")
	ErrorInvalidReceiver = sdkerrors.Register(ModuleName, 103, "invalid receiver")
	ErrorInvalidTxHash   = sdkerrors.Register(ModuleName, 104, "invalid tx_hash")
	ErrorInvalidAmount   = sdkerrors.Register(ModuleName, 105, "invalid amount")
)

var (
	ErrorSwapIsDisabled = sdkerrors.Register(ModuleName, 201, "swap is disabled")
	ErrorUnauthorized   = sdkerrors.Register(ModuleName, 202, "unauthorized")
	ErrorDuplicateSwap  = sdkerrors.Register(ModuleName, 203, "duplicate swap")
)
