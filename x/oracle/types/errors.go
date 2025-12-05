package types

import (
	sdkerrors "cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcerrors "github.com/cosmos/ibc-go/v7/modules/core/errors"
)

const (
	_ = 100 + iota
	ErrCodeAssetNotFound
	ErrCodeDenomForPacketNotFound
	ErrCodeDuplicateAsset
	ErrCodeInvalidMessage
)

var (
	ErrAssetNotFound          = sdkerrors.Register(ModuleName, ErrCodeAssetNotFound, "asset not found")
	ErrDenomForPacketNotFound = sdkerrors.Register(ModuleName, ErrCodeDenomForPacketNotFound, "denom for packet not found")
	ErrDuplicateAsset         = sdkerrors.Register(ModuleName, ErrCodeDuplicateAsset, "duplicate asset")
	ErrInvalidMessage         = sdkerrors.Register(ModuleName, ErrCodeInvalidMessage, "invalid message")
)

// NewErrorAssetNotFound returns an error indicating that the specified asset does not exist.
func NewErrorAssetNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrAssetNotFound, "asset %s does not exist", denom)
}

// NewErrorDenomForPacketNotFound returns an error indicating that the denom for the specified packet does not exist.
func NewErrorDenomForPacketNotFound(portID, channelID string, sequence uint64) error {
	return sdkerrors.Wrapf(ErrDenomForPacketNotFound, "denom for packet %s/%s/%d does not exist", portID, channelID, sequence)
}

// NewErrorDuplicateAsset returns an error indicating that the specified asset already exists.
func NewErrorDuplicateAsset(denom string) error {
	return sdkerrors.Wrapf(ErrDuplicateAsset, "asset %s already exists", denom)
}

// NewErrorInvalidMessage returns an error indicating that the provided message is invalid.
func NewErrorInvalidMessage(desc any) error {
	return sdkerrors.Wrapf(ErrInvalidMessage, "%v", desc)
}

// NewErrorInvalidCounterpartyVersion returns an error indicating that the counterparty version is invalid.
func NewErrorInvalidCounterpartyVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid counterparty version %s; expected %s", version, expected)
}

// NewErrorInvalidPort returns an error indicating that the provided port is invalid.
func NewErrorInvalidPort(portID, expected string) error {
	return sdkerrors.Wrapf(ibcporttypes.ErrInvalidPort, "invalid port %s; expected %s", portID, expected)
}

// NewErrorInvalidSigner returns an error indicating that the provided signer is invalid.
func NewErrorInvalidSigner(from, expected string) error {
	return sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority %s; expected %s", from, expected)
}

// NewErrorInvalidVersion returns an error indicating that the provided IBC version is invalid.
func NewErrorInvalidVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid version %s; expected %s", version, expected)
}
