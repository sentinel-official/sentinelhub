package v3

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
)

var (
	_ sdk.Msg = (*MsgCancelSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgRenewSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgShareSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgStartSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgCancelSubscriptionRequest(from sdk.AccAddress, id uint64) *MsgCancelSubscriptionRequest {
	return &MsgCancelSubscriptionRequest{
		From: from.String(),
		ID:   id,
	}
}

func (m *MsgCancelSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgCancelSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRenewSubscriptionRequest(from sdk.AccAddress, id uint64, denom string) *MsgRenewSubscriptionRequest {
	return &MsgRenewSubscriptionRequest{
		From:  from.String(),
		ID:    id,
		Denom: denom,
	}
}

func (m *MsgRenewSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
	}

	return nil
}

func (m *MsgRenewSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgShareSubscriptionRequest(from sdk.AccAddress, id uint64, accAddr sdk.AccAddress, bytes sdkmath.Int) *MsgShareSubscriptionRequest {
	return &MsgShareSubscriptionRequest{
		From:       from.String(),
		ID:         id,
		AccAddress: accAddr.String(),
		Bytes:      bytes,
	}
}

func (m *MsgShareSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.AccAddress == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "acc_address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.AccAddress); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "invalid acc_address %s", err)
	}
	if m.Bytes.IsNil() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "bytes cannot be nil")
	}
	if m.Bytes.IsNegative() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "bytes cannot be negative")
	}

	return nil
}

func (m *MsgShareSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSubscriptionRequest(from sdk.AccAddress, id uint64, denom string, renewalPricePolicy v1base.RenewalPricePolicy) *MsgStartSubscriptionRequest {
	return &MsgStartSubscriptionRequest{
		From:               from.String(),
		ID:                 id,
		Denom:              denom,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

func (m *MsgStartSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
	}
	if !m.RenewalPricePolicy.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "renewal_price_policy must be valid")
	}

	return nil
}

func (m *MsgStartSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateSubscriptionRequest(from sdk.AccAddress, id uint64, renewalPricePolicy v1base.RenewalPricePolicy) *MsgUpdateSubscriptionRequest {
	return &MsgUpdateSubscriptionRequest{
		From:               from.String(),
		ID:                 id,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

func (m *MsgUpdateSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if !m.RenewalPricePolicy.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "renewal_price_policy must be valid")
	}

	return nil
}

func (m *MsgUpdateSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(from sdk.AccAddress, id uint64, nodeAddr base.NodeAddress) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: nodeAddr.String(),
	}
}

func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgStartSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateParamsRequest(from sdk.AccAddress, params Params) *MsgUpdateParamsRequest {
	return &MsgUpdateParamsRequest{
		From:   from.String(),
		Params: params,
	}
}

func (m *MsgUpdateParamsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if err := m.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (m *MsgUpdateParamsRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
