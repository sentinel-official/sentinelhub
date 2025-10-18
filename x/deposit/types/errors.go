package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	_ = 100 + iota
	ErrCodeDepositNotFound
	ErrCodeInsufficientDeposit
	ErrCodeInsufficientFunds
	ErrCodeInvalidMessage
)

var (
	ErrDepositNotFound     = sdkerrors.Register(ModuleName, ErrCodeDepositNotFound, "deposit not found")
	ErrInsufficientDeposit = sdkerrors.Register(ModuleName, ErrCodeInsufficientDeposit, "insufficient deposit")
	ErrInsufficientFunds   = sdkerrors.Register(ModuleName, ErrCodeInsufficientFunds, "insufficient funds")
	ErrInvalidMessage      = sdkerrors.Register(ModuleName, ErrCodeInvalidMessage, "invalid message")
)

// NewErrorDepositNotFound wraps ErrDepositNotFound with the provided address context for a missing deposit.
func NewErrorDepositNotFound(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrDepositNotFound, "deposit for address %s does not exist", addr)
}

// NewErrorInsufficientDeposit wraps ErrInsufficientDeposit with the provided address context for a deposit issue.
func NewErrorInsufficientDeposit(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInsufficientDeposit, "insufficient deposit for address %s", addr)
}

// NewErrorInsufficientFunds wraps ErrInsufficientFunds with the provided address context for a funds issue.
func NewErrorInsufficientFunds(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInsufficientFunds, "insufficient funds for address %s", addr)
}

// NewErrorInvalidMessage returns an error indicating that the provided message is invalid.
func NewErrorInvalidMessage(desc interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidMessage, "%v", desc)
}
