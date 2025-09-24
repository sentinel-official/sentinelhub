package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types"
)

// Ensure the message types implement sdk.Msg interface.
var (
	_ sdk.Msg = (*MsgEndLeaseRequest)(nil)
	_ sdk.Msg = (*MsgRenewLeaseRequest)(nil)
	_ sdk.Msg = (*MsgStartLeaseRequest)(nil)
	_ sdk.Msg = (*MsgUpdateLeaseRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

// NewMsgEndLeaseRequest creates a new MsgEndLeaseRequest instance.
func NewMsgEndLeaseRequest(from base.ProvAddress, id uint64) *MsgEndLeaseRequest {
	return &MsgEndLeaseRequest{
		From: from.String(),
		ID:   id,
	}
}

// ValidateBasic performs basic validation checks on the MsgEndLeaseRequest.
func (m *MsgEndLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgEndLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgRenewLeaseRequest creates a new MsgRenewLeaseRequest instance.
func NewMsgRenewLeaseRequest(from base.ProvAddress, id uint64, hours int64, maxPrice v1base.Price) *MsgRenewLeaseRequest {
	return &MsgRenewLeaseRequest{
		From:     from.String(),
		ID:       id,
		Hours:    hours,
		MaxPrice: maxPrice,
	}
}

// ValidateBasic performs basic validation checks on the MsgRenewLeaseRequest.
func (m *MsgRenewLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	if m.Hours == 0 {
		return types.NewErrorInvalidMessage("hours cannot be zero")
	}

	if m.Hours < 0 {
		return types.NewErrorInvalidMessage("hours cannot be negative")
	}

	if m.MaxPrice.Denom != "" {
		if err := m.MaxPrice.Validate(); err != nil {
			return types.NewErrorInvalidMessage(fmt.Errorf("invalid max_price: %w", err))
		}
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgRenewLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgStartLeaseRequest creates a new MsgStartLeaseRequest instance.
func NewMsgStartLeaseRequest(from base.ProvAddress, nodeAddr base.NodeAddress, hours int64, maxPrice v1base.Price, renewalPricePolicy v1base.RenewalPricePolicy) *MsgStartLeaseRequest {
	return &MsgStartLeaseRequest{
		From:               from.String(),
		NodeAddress:        nodeAddr.String(),
		Hours:              hours,
		MaxPrice:           maxPrice,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

// ValidateBasic performs basic validation checks on the MsgStartLeaseRequest.
func (m *MsgStartLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.NodeAddress == "" {
		return types.NewErrorInvalidMessage("node_address cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid node_address: %w", err))
	}

	if m.Hours == 0 {
		return types.NewErrorInvalidMessage("hours cannot be zero")
	}

	if m.Hours < 0 {
		return types.NewErrorInvalidMessage("hours cannot be negative")
	}

	if m.MaxPrice.Denom != "" {
		if err := m.MaxPrice.Validate(); err != nil {
			return types.NewErrorInvalidMessage(fmt.Errorf("invalid max_price: %w", err))
		}
	}

	if !m.RenewalPricePolicy.IsValid() {
		return types.NewErrorInvalidMessage("renewal_price_policy must be valid")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgStartLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateLeaseRequest creates a new MsgUpdateLeaseRequest instance.
func NewMsgUpdateLeaseRequest(from base.ProvAddress, id uint64, renewalPricePolicy v1base.RenewalPricePolicy) *MsgUpdateLeaseRequest {
	return &MsgUpdateLeaseRequest{
		From:               from.String(),
		ID:                 id,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateLeaseRequest.
func (m *MsgUpdateLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	if !m.RenewalPricePolicy.IsValid() {
		return types.NewErrorInvalidMessage("renewal_price_policy must be valid")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdateLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateParamsRequest creates a new MsgUpdateParamsRequest instance.
func NewMsgUpdateParamsRequest(from sdk.AccAddress, params Params) *MsgUpdateParamsRequest {
	return &MsgUpdateParamsRequest{
		From:   from.String(),
		Params: params,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateParamsRequest.
func (m *MsgUpdateParamsRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if err := m.Params.Validate(); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid params: %w", err))
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdateParamsRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
