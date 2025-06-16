package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateNodeForPlan           = sdkerrors.Register(ModuleName, 201, "duplicate node for plan")
	ErrInvalidNodeStatus              = sdkerrors.Register(ModuleName, 202, "invalid node status")
	ErrLeaseForNodeByProviderNotFound = sdkerrors.Register(ModuleName, 203, "lease for node by provider not found")
	ErrLeaseNotFound                  = sdkerrors.Register(ModuleName, 204, "lease not found")
	ErrNodeForPlanNotFound            = sdkerrors.Register(ModuleName, 205, "node for plan not found")
	ErrNodeNotFound                   = sdkerrors.Register(ModuleName, 206, "node not found")
	ErrPlanNotFound                   = sdkerrors.Register(ModuleName, 207, "plan not found")
	ErrProviderNotFound               = sdkerrors.Register(ModuleName, 208, "provider not found")
	ErrUnauthorized                   = sdkerrors.Register(ModuleName, 209, "unauthorized")
)

// NewErrorDuplicateNodeForPlan returns an error indicating that a node already exists for the specified plan.
func NewErrorDuplicateNodeForPlan(id uint64, addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateNodeForPlan, "node %s for plan %d already exists", addr, id)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the given node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorLeaseForNodeByProviderNotFound returns an error indicating that the specified lease does not exist.
func NewErrorLeaseForNodeByProviderNotFound(nodeAddr base.NodeAddress, provAddr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrLeaseForNodeByProviderNotFound, "lease for node %s by provider %s does not exist", nodeAddr, provAddr)
}

// NewErrorLeaseNotFound returns an error indicating that the specified lease does not exist.
func NewErrorLeaseNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrLeaseNotFound, "lease %d does not exist", id)
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

// NewErrorProviderNotFound returns an error indicating that the specified provider does not exist.
func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrProviderNotFound, "provider %s does not exist", addr)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
