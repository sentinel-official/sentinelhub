package types

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	_ = 100 + iota
	ErrCodeInvalidDownloadBytes
	ErrCodeInvalidDuration
	ErrCodeInvalidMessage
	ErrCodeInvalidSessionStatus
	ErrCodeInvalidSignature
	ErrCodeInvalidUploadBytes
	ErrCodeSessionNotFound
	ErrCodeUnauthorized
)

var (
	ErrInvalidDownloadBytes = sdkerrors.Register(ModuleName, ErrCodeInvalidDownloadBytes, "invalid download bytes")
	ErrInvalidDuration      = sdkerrors.Register(ModuleName, ErrCodeInvalidDuration, "invalid duration")
	ErrInvalidMessage       = sdkerrors.Register(ModuleName, ErrCodeInvalidMessage, "invalid message")
	ErrInvalidSessionStatus = sdkerrors.Register(ModuleName, ErrCodeInvalidSessionStatus, "invalid session status")
	ErrInvalidSignature     = sdkerrors.Register(ModuleName, ErrCodeInvalidSignature, "invalid signature")
	ErrInvalidUploadBytes   = sdkerrors.Register(ModuleName, ErrCodeInvalidUploadBytes, "invalid upload bytes")
	ErrSessionNotFound      = sdkerrors.Register(ModuleName, ErrCodeSessionNotFound, "session not found")
	ErrUnauthorized         = sdkerrors.Register(ModuleName, ErrCodeUnauthorized, "unauthorized")
)

// NewErrorInvalidDownloadBytes returns an error indicating that the download bytes are invalid.
func NewErrorInvalidDownloadBytes(bytes sdkmath.Int) error {
	return sdkerrors.Wrapf(ErrInvalidDownloadBytes, "invalid download bytes %s", bytes)
}

// NewErrorInvalidDuration returns an error indicating that the specified duration is invalid.
func NewErrorInvalidDuration(duration time.Duration) error {
	return sdkerrors.Wrapf(ErrInvalidDuration, "invalid duration %d", duration)
}

// NewErrorInvalidMessage returns an error indicating that the provided message is invalid.
func NewErrorInvalidMessage(desc interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidMessage, "%v", desc)
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
