package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	_ = 100 + iota
	ErrCodeDuplicateNode
	ErrCodeInvalidGigabytes
	ErrCodeInvalidHours
	ErrCodeInvalidMessage
	ErrCodeInvalidNodeStatus
	ErrCodeInvalidPrice
	ErrCodeInvalidPrices
	ErrCodeInvalidSessionStatus
	ErrCodeNodeNotFound
	ErrCodePriceNotFound
	ErrCodeSessionNotFound
	ErrCodeUnauthorized
)

var (
	ErrDuplicateNode        = sdkerrors.Register(ModuleName, ErrCodeDuplicateNode, "duplicate node")
	ErrInvalidGigabytes     = sdkerrors.Register(ModuleName, ErrCodeInvalidGigabytes, "invalid gigabytes")
	ErrInvalidHours         = sdkerrors.Register(ModuleName, ErrCodeInvalidHours, "invalid hours")
	ErrInvalidMessage       = sdkerrors.Register(ModuleName, ErrCodeInvalidMessage, "invalid message")
	ErrInvalidNodeStatus    = sdkerrors.Register(ModuleName, ErrCodeInvalidNodeStatus, "invalid node status")
	ErrInvalidPrice         = sdkerrors.Register(ModuleName, ErrCodeInvalidPrice, "invalid price")
	ErrInvalidPrices        = sdkerrors.Register(ModuleName, ErrCodeInvalidPrices, "invalid prices")
	ErrInvalidSessionStatus = sdkerrors.Register(ModuleName, ErrCodeInvalidSessionStatus, "invalid session status")
	ErrNodeNotFound         = sdkerrors.Register(ModuleName, ErrCodeNodeNotFound, "node not found")
	ErrPriceNotFound        = sdkerrors.Register(ModuleName, ErrCodePriceNotFound, "price not found")
	ErrSessionNotFound      = sdkerrors.Register(ModuleName, ErrCodeSessionNotFound, "session not found")
	ErrUnauthorized         = sdkerrors.Register(ModuleName, ErrCodeUnauthorized, "unauthorized")
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

// NewErrorInvalidMessage returns an error indicating that the provided message is invalid.
func NewErrorInvalidMessage(desc interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidMessage, "%v", desc)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the given node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorInvalidPrice returns an error indicating that the price is invalid.
func NewErrorInvalidPrice(price v1base.Price) error {
	return sdkerrors.Wrapf(ErrInvalidPrice, "invalid price %s", price)
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
