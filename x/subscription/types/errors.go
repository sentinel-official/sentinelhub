package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

const (
	_ = 100 + iota
	ErrCodeAllocationNotFound
	ErrCodeInsufficientBytes
	ErrCodeInvalidAllocation
	ErrCodeInvalidMessage
	ErrCodeInvalidNodeStatus
	ErrCodeInvalidPlanStatus
	ErrCodeInvalidRenewalPolicy
	ErrCodeInvalidSessionStatus
	ErrCodeInvalidSubscriptionStatus
	ErrCodeNodeForPlanNotFound
	ErrCodeNodeNotFound
	ErrCodePlanNotFound
	ErrCodePriceNotFound
	ErrCodeSessionNotFound
	ErrCodeSubscriptionNotFound
	ErrCodeUnauthorized
)

var (
	ErrAllocationNotFound        = sdkerrors.Register(ModuleName, ErrCodeAllocationNotFound, "allocation not found")
	ErrInsufficientBytes         = sdkerrors.Register(ModuleName, ErrCodeInsufficientBytes, "insufficient bytes")
	ErrInvalidAllocation         = sdkerrors.Register(ModuleName, ErrCodeInvalidAllocation, "invalid allocation")
	ErrInvalidMessage            = sdkerrors.Register(ModuleName, ErrCodeInvalidMessage, "invalid message")
	ErrInvalidNodeStatus         = sdkerrors.Register(ModuleName, ErrCodeInvalidNodeStatus, "invalid node status")
	ErrInvalidPlanStatus         = sdkerrors.Register(ModuleName, ErrCodeInvalidPlanStatus, "invalid plan status")
	ErrInvalidRenewalPolicy      = sdkerrors.Register(ModuleName, ErrCodeInvalidRenewalPolicy, "invalid renewal policy")
	ErrInvalidSessionStatus      = sdkerrors.Register(ModuleName, ErrCodeInvalidSessionStatus, "invalid session status")
	ErrInvalidSubscriptionStatus = sdkerrors.Register(ModuleName, ErrCodeInvalidSubscriptionStatus, "invalid subscription status")
	ErrNodeForPlanNotFound       = sdkerrors.Register(ModuleName, ErrCodeNodeForPlanNotFound, "node for plan not found")
	ErrNodeNotFound              = sdkerrors.Register(ModuleName, ErrCodeNodeNotFound, "node not found")
	ErrPlanNotFound              = sdkerrors.Register(ModuleName, ErrCodePlanNotFound, "plan not found")
	ErrPriceNotFound             = sdkerrors.Register(ModuleName, ErrCodePriceNotFound, "price not found")
	ErrSessionNotFound           = sdkerrors.Register(ModuleName, ErrCodeSessionNotFound, "session not found")
	ErrSubscriptionNotFound      = sdkerrors.Register(ModuleName, ErrCodeSubscriptionNotFound, "subscription not found")
	ErrUnauthorized              = sdkerrors.Register(ModuleName, ErrCodeUnauthorized, "unauthorized")
)

// NewErrorAllocationNotFound returns an error indicating that the specified allocation does not exist.
func NewErrorAllocationNotFound(id uint64, addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrAllocationNotFound, "allocation %d/%s does not exist", id, addr)
}

// NewErrorInsufficientBytes returns an error indicating that there are insufficient bytes for the specified subscription.
func NewErrorInsufficientBytes(id uint64, bytes sdkmath.Int) error {
	return sdkerrors.Wrapf(ErrInsufficientBytes, "insufficient bytes %s for subscription %d", bytes, id)
}

// NewErrorInvalidAllocation returns an error indicating that the allocation is invalid.
func NewErrorInvalidAllocation(id uint64, addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

// NewErrorInvalidMessage returns an error indicating that the provided message is invalid.
func NewErrorInvalidMessage(desc interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidMessage, "%v", desc)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorInvalidPlanStatus returns an error indicating that the provided status is invalid for the plan.
func NewErrorInvalidPlanStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidPlanStatus, "invalid status %s for plan %d", status, id)
}

// NewErrorInvalidSessionStatus returns an error indicating that the provided status is invalid for the session.
func NewErrorInvalidSessionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidSessionStatus, "invalid status %s for session %d", status, id)
}

// NewErrorInvalidSubscriptionStatus returns an error indicating that the provided status is invalid for the subscription.
func NewErrorInvalidSubscriptionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidSubscriptionStatus, "invalid status %s for subscription %d", status, id)
}

// NewErrorNodeForPlanNotFound returns an error indicating that the specified node does not exist for the plan.
func NewErrorNodeForPlanNotFound(id uint64, addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNodeForPlanNotFound, "node %s for plan %d does not exist", addr, id)
}

// NewErrorNodeNotFound returns an error indicating that the specified node does not exist.
func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNodeNotFound, "node %s does not exist", addr)
}

// NewErrorPlanNotFound returns an error indicating that the specified plan does not exist.
func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrPlanNotFound, "plan %d does not exist", id)
}

// NewErrorPriceNotFound returns an error indicating that the price for the specified denomination does not exist.
func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrPriceNotFound, "price for denom %s does not exist", denom)
}

// NewErrorSessionNotFound returns an error indicating that the specified session does not exist.
func NewErrorSessionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrSessionNotFound, "session %d does not exist", id)
}

// NewErrorSubscriptionNotFound returns an error indicating that the specified subscription does not exist.
func NewErrorSubscriptionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrSubscriptionNotFound, "subscription %d does not exist", id)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
