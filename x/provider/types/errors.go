package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/sentinelhub/v12/types"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateProvider = sdkerrors.Register(ModuleName, 201, "duplicate provider")
	ErrProviderNotFound  = sdkerrors.Register(ModuleName, 202, "provider not found")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 203, "unauthorized")
)

// NewErrorDuplicateProvider returns an error indicating that the specified provider already exists.
func NewErrorDuplicateProvider(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateProvider, "provider %s already exists", addr)
}

// NewErrorProviderNotFound returns an error indicating that the specified provider does not exist.
func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrProviderNotFound, "provider %s does not exist", addr)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
