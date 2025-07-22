package v3

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/subscription/types"
)

// Ensure the message types implement sdk.Msg interface
var (
	_ sdk.Msg = (*MsgCancelSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgRenewSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgShareSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgStartSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

// NewMsgCancelSubscriptionRequest creates a new MsgCancelSubscriptionRequest instance.
func NewMsgCancelSubscriptionRequest(from sdk.AccAddress, id uint64) *MsgCancelSubscriptionRequest {
	return &MsgCancelSubscriptionRequest{
		From: from.String(),
		ID:   id,
	}
}

// ValidateBasic performs basic validation checks on the MsgCancelSubscriptionRequest.
func (m *MsgCancelSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgCancelSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgRenewSubscriptionRequest creates a new MsgRenewSubscriptionRequest instance.
func NewMsgRenewSubscriptionRequest(from sdk.AccAddress, id uint64, denom string) *MsgRenewSubscriptionRequest {
	return &MsgRenewSubscriptionRequest{
		From:  from.String(),
		ID:    id,
		Denom: denom,
	}
}

// ValidateBasic performs basic validation checks on the MsgRenewSubscriptionRequest.
func (m *MsgRenewSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return types.NewErrorInvalidMessage(err)
		}
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgRenewSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgShareSubscriptionRequest creates a new MsgShareSubscriptionRequest instance.
func NewMsgShareSubscriptionRequest(from sdk.AccAddress, id uint64, accAddr sdk.AccAddress, bytes sdkmath.Int) *MsgShareSubscriptionRequest {
	return &MsgShareSubscriptionRequest{
		From:       from.String(),
		ID:         id,
		AccAddress: accAddr.String(),
		Bytes:      bytes,
	}
}

// ValidateBasic performs basic validation checks on the MsgShareSubscriptionRequest.
func (m *MsgShareSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}
	if m.AccAddress == "" {
		return types.NewErrorInvalidMessage("acc_address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.AccAddress); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if m.Bytes.IsNil() {
		return types.NewErrorInvalidMessage("bytes cannot be nil")
	}
	if m.Bytes.IsNegative() {
		return types.NewErrorInvalidMessage("bytes cannot be negative")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgShareSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgStartSubscriptionRequest creates a new MsgStartSubscriptionRequest instance.
func NewMsgStartSubscriptionRequest(from sdk.AccAddress, id uint64, denom string, renewalPricePolicy v1base.RenewalPricePolicy) *MsgStartSubscriptionRequest {
	return &MsgStartSubscriptionRequest{
		From:               from.String(),
		ID:                 id,
		Denom:              denom,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

// ValidateBasic performs basic validation checks on the MsgStartSubscriptionRequest.
func (m *MsgStartSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return types.NewErrorInvalidMessage(err)
		}
	}
	if !m.RenewalPricePolicy.IsValid() {
		return types.NewErrorInvalidMessage("renewal_price_policy must be valid")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgStartSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateSubscriptionRequest creates a new MsgUpdateSubscriptionRequest instance.
func NewMsgUpdateSubscriptionRequest(from sdk.AccAddress, id uint64, renewalPricePolicy v1base.RenewalPricePolicy) *MsgUpdateSubscriptionRequest {
	return &MsgUpdateSubscriptionRequest{
		From:               from.String(),
		ID:                 id,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateSubscriptionRequest.
func (m *MsgUpdateSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
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
func (m *MsgUpdateSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgStartSessionRequest creates a new MsgStartSessionRequest instance.
func NewMsgStartSessionRequest(from sdk.AccAddress, id uint64, nodeAddr base.NodeAddress) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: nodeAddr.String(),
	}
}

// ValidateBasic performs basic validation checks on the MsgStartSessionRequest.
func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}
	if m.NodeAddress == "" {
		return types.NewErrorInvalidMessage("node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgStartSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
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
		return types.NewErrorInvalidMessage(err)
	}
	if err := m.Params.Validate(); err != nil {
		return types.NewErrorInvalidMessage(err)
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
