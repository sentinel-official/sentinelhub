package types

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrInvalidDownloadBytes = sdkerrors.Register(ModuleName, 201, "invalid download bytes")
	ErrInvalidDuration      = sdkerrors.Register(ModuleName, 202, "invalid duration")
	ErrInvalidSessionStatus = sdkerrors.Register(ModuleName, 203, "invalid session status")
	ErrInvalidSignature     = sdkerrors.Register(ModuleName, 204, "invalid signature")
	ErrInvalidUploadBytes   = sdkerrors.Register(ModuleName, 205, "invalid upload bytes")
	ErrSessionNotFound      = sdkerrors.Register(ModuleName, 206, "session not found")
	ErrUnauthorized         = sdkerrors.Register(ModuleName, 207, "unauthorized")
)

// NewErrorInvalidDownloadBytes returns an error indicating that the download bytes are invalid.
func NewErrorInvalidDownloadBytes(bytes sdkmath.Int) error {
	return sdkerrors.Wrapf(ErrInvalidDownloadBytes, "invalid download bytes %s", bytes)
}

// NewErrorInvalidDuration returns an error indicating that the specified duration is invalid.
func NewErrorInvalidDuration(duration time.Duration) error {
	return sdkerrors.Wrapf(ErrInvalidDuration, "invalid duration %d", duration)
}

// NewErrorInvalidSessionStatus returns an error indicating that the provided status is invalid for the session.
func NewErrorInvalidSessionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidSessionStatus, "invalid status %s for session %d", status, id)
}

// NewErrorInvalidSignature returns an error indicating that the provided signature is invalid.
func NewErrorInvalidSignature(signature []byte) error {
	return sdkerrors.Wrapf(ErrInvalidSignature, "invalid signature %X", signature)
}

// NewErrorInvalidUploadBytes returns an error indicating that the upload bytes are invalid.
func NewErrorInvalidUploadBytes(bytes sdkmath.Int) error {
	return sdkerrors.Wrapf(ErrInvalidUploadBytes, "invalid upload bytes %s", bytes)
}

// NewErrorSessionNotFound returns an error indicating that the specified session does not exist.
func NewErrorSessionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrSessionNotFound, "session %d does not exist", id)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
