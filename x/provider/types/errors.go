package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

const (
	_ = 100 + iota
	ErrCodeDuplicateProvider
	ErrCodeInvalidMessage
	ErrCodeProviderNotFound
	ErrCodeUnauthorized
)

var (
	ErrDuplicateProvider = sdkerrors.Register(ModuleName, ErrCodeDuplicateProvider, "duplicate provider")
	ErrInvalidMessage    = sdkerrors.Register(ModuleName, ErrCodeInvalidMessage, "invalid message")
	ErrProviderNotFound  = sdkerrors.Register(ModuleName, ErrCodeProviderNotFound, "provider not found")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, ErrCodeUnauthorized, "unauthorized")
)

// NewErrorDuplicateProvider returns an error indicating that the specified provider already exists.
func NewErrorDuplicateProvider(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateProvider, "provider %s already exists", addr)
}

// NewErrorInvalidMessage returns an error indicating that the provided message is invalid.
func NewErrorInvalidMessage(desc any) error {
	return sdkerrors.Wrapf(ErrInvalidMessage, "%v", desc)
}

// NewErrorProviderNotFound returns an error indicating that the specified provider does not exist.
func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrProviderNotFound, "provider %s does not exist", addr)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
