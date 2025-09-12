package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types/v1"
)

// HandleMsgEndLease handles a request to end an active lease.
// It performs authorization, refunds the remaining deposit, deletes lease state, and emits relevant events.
func (k *Keeper) HandleMsgEndLease(ctx sdk.Context, msg *v1.MsgEndLeaseRequest) (*v1.MsgEndLeaseResponse, error) {
	// Fetch lease and validate existence and ownership
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}

	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Run the pre-hook for inactivating a lease, validating readiness for termination
	if err := k.LeaseInactivePreHook(ctx, lease.ID); err != nil {
		return nil, err
	}

	// Convert provider address to internal format; return error on invalid input
	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return nil, err
	}

	// Calculate refundable amount and subtract it from provider’s deposit
	refund := lease.RefundAmount()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), refund); err != nil {
		return nil, err
	}

	// Emit an event indicating refund details to subscribers
	ctx.EventManager().EmitTypedEvent(
		&v1.EventRefund{
			ID:          lease.ID,
			ProvAddress: lease.ProvAddress,
			Amount:      refund.String(),
		},
	)

	// Convert node address from string to internal format
	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Clean up all associated state entries related to this lease
	k.DeleteLease(ctx, lease.ID)
	k.DeleteLeaseForNodeByProvider(ctx, nodeAddr, provAddr, lease.ID)
	k.DeleteLeaseForProvider(ctx, provAddr, lease.ID)
	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	// Emit an event indicating lease termination
	ctx.EventManager().EmitTypedEvent(
		&v1.EventEnd{
			ID:          lease.ID,
			NodeAddress: lease.NodeAddress,
			ProvAddress: lease.ProvAddress,
		},
	)

	return &v1.MsgEndLeaseResponse{}, nil
}

// HandleMsgRenewLease handles a request to renew an existing lease.
// It validates inputs, ensures pricing constraints are satisfied, refunds remaining deposit, and restarts the lease.
func (k *Keeper) HandleMsgRenewLease(ctx sdk.Context, msg *v1.MsgRenewLeaseRequest) (*v1.MsgRenewLeaseResponse, error) {
	// Validate the renewal hours requested against allowed bounds
	if !k.IsValidHours(ctx, msg.Hours) {
		return nil, types.NewErrorInvalidHours(msg.Hours)
	}

	// Fetch lease and validate existence and ownership
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}

	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Convert and validate the provider address
	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return nil, err
	}

	// Verify the provider exists and is in an active state
	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	if !provider.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
	}

	// Convert and verify the node address
	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Validate node existence and ensure it's active
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Retrieve and validate the node's hourly price for the given denomination
	price, found := node.HourlyPrice(msg.MaxPrice.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.MaxPrice.Denom)
	}

	// Convert price to current quote value; fail on error
	price, err = price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	// Ensure quoted price does not exceed the client's maximum acceptable price
	if price.IsGT(msg.MaxPrice) {
		return nil, types.NewErrorInvalidPrice(price)
	}

	// Ensure renewal is allowed by the lease's renewal policy at the current price
	if err := lease.ValidateRenewalPolicies(price); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidRenewalPolicy, err.Error())
	}

	// Refund unused lease value and update deposit ledger
	refund := lease.RefundAmount()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), refund); err != nil {
		return nil, err
	}

	// Emit refund event to subscribers
	ctx.EventManager().EmitTypedEvent(
		&v1.EventRefund{
			ID:          lease.ID,
			ProvAddress: lease.ProvAddress,
			Amount:      refund.String(),
		},
	)

	// Clear old indexing references before overwriting lease
	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	// Create a renewed lease with updated pricing and time, retaining other original fields
	lease = v1.Lease{
		ID:                 lease.ID,
		ProvAddress:        lease.ProvAddress,
		NodeAddress:        lease.NodeAddress,
		Price:              price,
		Hours:              0,
		MaxHours:           msg.Hours,
		RenewalPricePolicy: lease.RenewalPricePolicy,
		StartAt:            ctx.BlockTime(),
	}

	// Charge the deposit for the renewed lease
	deposit := lease.DepositAmount()
	if err := k.AddDeposit(ctx, provAddr.Bytes(), deposit); err != nil {
		return nil, err
	}

	// Persist the renewed lease and update all relevant index references
	k.SetLease(ctx, lease)
	k.SetLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	// Emit lease renewal event
	ctx.EventManager().EmitTypedEvent(
		&v1.EventRenew{
			ID:          lease.ID,
			NodeAddress: lease.NodeAddress,
			ProvAddress: lease.ProvAddress,
			MaxHours:    lease.MaxHours,
			Price:       lease.Price.String(),
		},
	)

	return &v1.MsgRenewLeaseResponse{}, nil
}

// HandleMsgStartLease handles a request to start a new lease.
// It validates the provider, node, and pricing details, ensures uniqueness, and creates the lease state.
func (k *Keeper) HandleMsgStartLease(ctx sdk.Context, msg *v1.MsgStartLeaseRequest) (*v1.MsgStartLeaseResponse, error) {
	// Validate lease duration request before proceeding
	if !k.IsValidHours(ctx, msg.Hours) {
		return nil, types.NewErrorInvalidHours(msg.Hours)
	}

	// Parse and validate the provider's Bech32 address
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Ensure provider exists and is currently active
	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	if !provider.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
	}

	// Parse and validate the node address
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Confirm the node exists and is active
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Fetch node's hourly price for the specified denomination
	price, found := node.HourlyPrice(msg.MaxPrice.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.MaxPrice.Denom)
	}

	// Adjust the price using a quote function; handle any error
	price, err = price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	// Ensure the effective price is within client-provided bounds
	if price.IsGT(msg.MaxPrice) {
		return nil, types.NewErrorInvalidPrice(price)
	}

	// Check if a lease already exists between this provider and node
	leaseExists := false

	k.IterateLeasesForNodeByProvider(ctx, nodeAddr, provAddr, func(_ int, _ v1.Lease) bool {
		leaseExists = true

		return true
	})

	if leaseExists {
		return nil, types.NewErrorDuplicateLease(nodeAddr, provAddr)
	}

	// Generate a new lease ID and construct the lease object
	count := k.GetLeaseCount(ctx)
	lease := v1.Lease{
		ID:                 count + 1,
		ProvAddress:        provAddr.String(),
		NodeAddress:        nodeAddr.String(),
		Price:              price,
		Hours:              0,
		MaxHours:           msg.Hours,
		RenewalPricePolicy: msg.RenewalPricePolicy,
		StartAt:            ctx.BlockTime(),
	}

	// Charge the initial deposit from the provider for the new lease
	deposit := lease.DepositAmount()
	if err := k.AddDeposit(ctx, provAddr.Bytes(), deposit); err != nil {
		return nil, err
	}

	// Persist the lease and update related indexes
	k.SetLeaseCount(ctx, count+1)
	k.SetLease(ctx, lease)
	k.SetLeaseForNodeByProvider(ctx, nodeAddr, provAddr, lease.ID)
	k.SetLeaseForProvider(ctx, provAddr, lease.ID)
	k.SetLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	// Emit event indicating successful lease creation
	ctx.EventManager().EmitTypedEvent(
		&v1.EventCreate{
			ID:                 lease.ID,
			NodeAddress:        lease.NodeAddress,
			ProvAddress:        lease.ProvAddress,
			MaxHours:           lease.MaxHours,
			Price:              lease.Price.String(),
			RenewalPricePolicy: lease.RenewalPricePolicy.String(),
		},
	)

	return &v1.MsgStartLeaseResponse{
		ID: lease.ID,
	}, nil
}

// HandleMsgUpdateLease handles a request to update the renewal price policy of an existing lease.
// It verifies lease ownership, updates the policy, and emits a lease update event.
func (k *Keeper) HandleMsgUpdateLease(ctx sdk.Context, msg *v1.MsgUpdateLeaseRequest) (*v1.MsgUpdateLeaseResponse, error) {
	// Fetch lease and validate existence and permission to update
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}

	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Remove existing renewal index before updating policy
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	// Apply new renewal price policy to the lease
	lease.RenewalPricePolicy = msg.RenewalPricePolicy

	// Persist lease and reindex it for renewal tracking
	k.SetLease(ctx, lease)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	// Emit event to signal lease update
	ctx.EventManager().EmitTypedEvent(
		&v1.EventUpdate{
			ID:                 lease.ID,
			NodeAddress:        lease.NodeAddress,
			ProvAddress:        lease.ProvAddress,
			RenewalPricePolicy: lease.RenewalPricePolicy.String(),
		},
	)

	return &v1.MsgUpdateLeaseResponse{}, nil
}

// HandleMsgUpdateParams allows the authority to update module-wide parameters.
// It restricts access to the authority address and saves the new parameters.
func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	// Ensure only the authority account is allowed to update module parameters
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Store the new parameter configuration in module state
	k.SetParams(ctx, msg.Params)

	return &v1.MsgUpdateParamsResponse{}, nil
}
