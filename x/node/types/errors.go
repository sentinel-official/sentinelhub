package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateNode        = sdkerrors.Register(ModuleName, 201, "duplicate node")
	ErrInvalidGigabytes     = sdkerrors.Register(ModuleName, 202, "invalid gigabytes")
	ErrInvalidHours         = sdkerrors.Register(ModuleName, 203, "invalid hours")
	ErrInvalidNodeStatus    = sdkerrors.Register(ModuleName, 204, "invalid node status")
	ErrInvalidPrices        = sdkerrors.Register(ModuleName, 205, "invalid prices")
	ErrInvalidSessionStatus = sdkerrors.Register(ModuleName, 206, "invalid session status")
	ErrNodeNotFound         = sdkerrors.Register(ModuleName, 207, "node not found")
	ErrPriceNotFound        = sdkerrors.Register(ModuleName, 208, "price not found")
	ErrSessionNotFound      = sdkerrors.Register(ModuleName, 209, "session not found")
	ErrUnauthorized         = sdkerrors.Register(ModuleName, 210, "unauthorized")
)

// NewErrorDuplicateNode returns an error indicating that the specified node already exists.
func NewErrorDuplicateNode(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateNode, "node %s already exists", addr)
}

// NewErrorInvalidGigabytes returns an error indicating that the provided gigabytes value is invalid.
func NewErrorInvalidGigabytes(gigabytes int64) error {
	return sdkerrors.Wrapf(ErrInvalidGigabytes, "invalid gigabytes %d", gigabytes)
}

// NewErrorInvalidHours returns an error indicating that the provided hours value is invalid.
func NewErrorInvalidHours(hours int64) error {
	return sdkerrors.Wrapf(ErrInvalidHours, "invalid hours %d", hours)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the given node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorInvalidPrices returns an error indicating that the provided prices are invalid.
func NewErrorInvalidPrices(prices v1base.Prices) error {
	return sdkerrors.Wrapf(ErrInvalidPrices, "invalid prices %s", prices)
}

// NewErrorInvalidSessionStatus returns an error indicating that the provided status is invalid for the session.
func NewErrorInvalidSessionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidSessionStatus, "invalid status %s for session %d", status, id)
}

// NewErrorNodeNotFound returns an error indicating that the specified node does not exist.
func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNodeNotFound, "node %s does not exist", addr)
}

// NewErrorPriceNotFound returns an error indicating that the price for the specified denom does not exist.
func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrPriceNotFound, "price for denom %s does not exist", denom)
}

// NewErrorSessionNotFound returns an error indicating that the specified session does not exist.
func NewErrorSessionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrSessionNotFound, "session %d does not exist", id)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
