package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) HandleMsgEndLease(ctx sdk.Context, msg *v1.MsgEndLeaseRequest) (*v1.MsgEndLeaseResponse, error) {
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if err := k.LeaseInactivePreHook(ctx, lease.ID); err != nil {
		return nil, err
	}

	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return nil, err
	}

	refund := lease.RefundAmount()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), refund); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&v1.EventRefund{
			ID:          lease.ID,
			ProvAddress: lease.ProvAddress,
			Amount:      refund.String(),
		},
	)

	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return nil, err
	}

	k.DeleteLease(ctx, lease.ID)
	k.DeleteLeaseForNodeByProvider(ctx, nodeAddr, provAddr, lease.ID)
	k.DeleteLeaseForProvider(ctx, provAddr, lease.ID)
	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	ctx.EventManager().EmitTypedEvent(
		&v1.EventEnd{
			ID:          lease.ID,
			NodeAddress: lease.NodeAddress,
			ProvAddress: lease.ProvAddress,
		},
	)

	return &v1.MsgEndLeaseResponse{}, nil
}

func (k *Keeper) HandleMsgRenewLease(ctx sdk.Context, msg *v1.MsgRenewLeaseRequest) (*v1.MsgRenewLeaseResponse, error) {
	if !k.IsValidHours(ctx, msg.Hours) {
		return nil, types.NewErrorInvalidHours(msg.Hours)
	}

	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}
	if !provider.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
	}

	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	price, found := node.HourlyPrice(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	price, err = price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	if err := lease.ValidateRenewalPolicies(price); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidRenewalPolicy, err.Error())
	}

	refund := lease.RefundAmount()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), refund); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&v1.EventRefund{
			ID:          lease.ID,
			ProvAddress: lease.ProvAddress,
			Amount:      refund.String(),
		},
	)

	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

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

	deposit := lease.DepositAmount()
	if err := k.AddDeposit(ctx, provAddr.Bytes(), deposit); err != nil {
		return nil, err
	}

	k.SetLease(ctx, lease)
	k.SetLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

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

func (k *Keeper) HandleMsgStartLease(ctx sdk.Context, msg *v1.MsgStartLeaseRequest) (*v1.MsgStartLeaseResponse, error) {
	if !k.IsValidHours(ctx, msg.Hours) {
		return nil, types.NewErrorInvalidHours(msg.Hours)
	}

	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}
	if !provider.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	price, found := node.HourlyPrice(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	price, err = price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	leaseExists := false
	k.IterateLeasesForNodeByProvider(ctx, nodeAddr, provAddr, func(_ int, _ v1.Lease) bool {
		leaseExists = true
		return true
	})

	if leaseExists {
		return nil, types.NewErrorDuplicateLease(nodeAddr, provAddr)
	}

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

	deposit := lease.DepositAmount()
	if err := k.AddDeposit(ctx, provAddr.Bytes(), deposit); err != nil {
		return nil, err
	}

	k.SetLeaseCount(ctx, count+1)
	k.SetLease(ctx, lease)
	k.SetLeaseForNodeByProvider(ctx, nodeAddr, provAddr, lease.ID)
	k.SetLeaseForProvider(ctx, provAddr, lease.ID)
	k.SetLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

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

func (k *Keeper) HandleMsgUpdateLease(ctx sdk.Context, msg *v1.MsgUpdateLeaseRequest) (*v1.MsgUpdateLeaseResponse, error) {
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	lease.RenewalPricePolicy = msg.RenewalPricePolicy

	k.SetLease(ctx, lease)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

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

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.SetParams(ctx, msg.Params)
	return &v1.MsgUpdateParamsResponse{}, nil
}
