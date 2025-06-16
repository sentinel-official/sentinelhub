package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateLease        = sdkerrors.Register(ModuleName, 201, "duplicate lease")
	ErrInvalidHours          = sdkerrors.Register(ModuleName, 202, "invalid hours")
	ErrInvalidNodeStatus     = sdkerrors.Register(ModuleName, 203, "invalid node status")
	ErrInvalidProviderStatus = sdkerrors.Register(ModuleName, 204, "invalid provider status")
	ErrInvalidRenewalPolicy  = sdkerrors.Register(ModuleName, 205, "invalid renewal policy")
	ErrLeaseNotFound         = sdkerrors.Register(ModuleName, 206, "lease not found")
	ErrNodeNotFound          = sdkerrors.Register(ModuleName, 207, "node not found")
	ErrPriceNotFound         = sdkerrors.Register(ModuleName, 208, "price not found")
	ErrProviderNotFound      = sdkerrors.Register(ModuleName, 209, "provider not found")
	ErrUnauthorized          = sdkerrors.Register(ModuleName, 210, "unauthorized")
)

// NewErrorDuplicateLease returns an error indicating that a lease for the specified node and provider already exists.
func NewErrorDuplicateLease(nodeAddr base.NodeAddress, provAddr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateLease, "lease already exists for node %s by provider %s", nodeAddr, provAddr)
}

// NewErrorInvalidHours returns an error indicating that the provided hours are invalid.
func NewErrorInvalidHours(hours int64) error {
	return sdkerrors.Wrapf(ErrInvalidHours, "invalid hours %d", hours)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the given node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidNodeStatus, "invalid status %s for node %s", status, addr)
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
