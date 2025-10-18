package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	_ = 100 + iota
	ErrCodeDuplicateLease
	ErrCodeInvalidHours
	ErrCodeInvalidMessage
	ErrCodeInvalidNodeStatus
	ErrCodeInvalidPrice
	ErrCodeInvalidProviderStatus
	ErrCodeInvalidRenewalPolicy
	ErrCodeLeaseNotFound
	ErrCodeNodeNotFound
	ErrCodePriceNotFound
	ErrCodeProviderNotFound
	ErrCodeUnauthorized
)

var (
	ErrDuplicateLease        = sdkerrors.Register(ModuleName, ErrCodeDuplicateLease, "duplicate lease")
	ErrInvalidHours          = sdkerrors.Register(ModuleName, ErrCodeInvalidHours, "invalid hours")
	ErrInvalidMessage        = sdkerrors.Register(ModuleName, ErrCodeInvalidMessage, "invalid message")
	ErrInvalidNodeStatus     = sdkerrors.Register(ModuleName, ErrCodeInvalidNodeStatus, "invalid node status")
	ErrInvalidPrice          = sdkerrors.Register(ModuleName, ErrCodeInvalidPrice, "invalid price")
	ErrInvalidProviderStatus = sdkerrors.Register(ModuleName, ErrCodeInvalidProviderStatus, "invalid provider status")
	ErrInvalidRenewalPolicy  = sdkerrors.Register(ModuleName, ErrCodeInvalidRenewalPolicy, "invalid renewal policy")
	ErrLeaseNotFound         = sdkerrors.Register(ModuleName, ErrCodeLeaseNotFound, "lease not found")
	ErrNodeNotFound          = sdkerrors.Register(ModuleName, ErrCodeNodeNotFound, "node not found")
	ErrPriceNotFound         = sdkerrors.Register(ModuleName, ErrCodePriceNotFound, "price not found")
	ErrProviderNotFound      = sdkerrors.Register(ModuleName, ErrCodeProviderNotFound, "provider not found")
	ErrUnauthorized          = sdkerrors.Register(ModuleName, ErrCodeUnauthorized, "unauthorized")
)

// NewErrorDuplicateLease returns an error indicating that a lease for the specified node and provider already exists.
func NewErrorDuplicateLease(nodeAddr base.NodeAddress, provAddr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateLease, "lease already exists for node %s by provider %s", nodeAddr, provAddr)
}

// NewErrorInvalidHours returns an error indicating that the provided hours are invalid.
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

// NewErrorInvalidProviderStatus returns an error indicating that the provided status is invalid for the given provider.
func NewErrorInvalidProviderStatus(addr base.ProvAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidProviderStatus, "invalid status %s for provider %s", status, addr)
}

// NewErrorLeaseNotFound returns an error indicating that the specified lease does not exist.
func NewErrorLeaseNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrLeaseNotFound, "lease %d does not exist", id)
}

// NewErrorNodeNotFound returns an error indicating that the specified node does not exist.
func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNodeNotFound, "node %s does not exist", addr)
}

// NewErrorPriceNotFound returns an error indicating that the price for the specified denom does not exist.
func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrPriceNotFound, "price for denom %s does not exist", denom)
}

// NewErrorProviderNotFound returns an error indicating that the specified provider does not exist.
func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrProviderNotFound, "provider %s does not exist", addr)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
