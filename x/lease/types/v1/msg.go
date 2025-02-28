package v1

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/lease/types"
)

var (
	_ sdk.Msg = (*MsgEndLeaseRequest)(nil)
	_ sdk.Msg = (*MsgRenewLeaseRequest)(nil)
	_ sdk.Msg = (*MsgStartLeaseRequest)(nil)
	_ sdk.Msg = (*MsgUpdateLeaseRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgEndLeaseRequest(from base.ProvAddress, id uint64) *MsgEndLeaseRequest {
	return &MsgEndLeaseRequest{
		From: from.String(),
		ID:   id,
	}
}

func (m *MsgEndLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgEndLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRenewLeaseRequest(from base.ProvAddress, id uint64, hours int64, denom string) *MsgRenewLeaseRequest {
	return &MsgRenewLeaseRequest{
		From:  from.String(),
		ID:    id,
		Hours: hours,
		Denom: denom,
	}
}

func (m *MsgRenewLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.Hours == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "hours cannot be zero")
	}
	if m.Hours < 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "hours cannot be negative")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
	}

	return nil
}

func (m *MsgRenewLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartLeaseRequest(from base.ProvAddress, nodeAddr base.NodeAddress, hours int64, denom string, renewalPricePolicy v1base.RenewalPricePolicy) *MsgStartLeaseRequest {
	return &MsgStartLeaseRequest{
		From:               from.String(),
		NodeAddress:        nodeAddr.String(),
		Hours:              hours,
		Denom:              denom,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

func (m *MsgStartLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Hours == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "hours cannot be zero")
	}
	if m.Hours < 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "hours cannot be negative")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
	}
	if !m.RenewalPricePolicy.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "renewal price policy must be valid")
	}

	return nil
}

func (m *MsgStartLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateLeaseRequest(from base.ProvAddress, id uint64, renewalPricePolicy v1base.RenewalPricePolicy) *MsgUpdateLeaseRequest {
	return &MsgUpdateLeaseRequest{
		From:               from.String(),
		ID:                 id,
		RenewalPricePolicy: renewalPricePolicy,
	}
}

func (m *MsgUpdateLeaseRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if !m.RenewalPricePolicy.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "renewal price policy must be valid")
	}

	return nil
}

func (m *MsgUpdateLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
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
