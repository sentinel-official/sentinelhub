package v3

import (
	"net/url"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/provider/types"
)

var (
	_ sdk.Msg = (*MsgRegisterProviderRequest)(nil)
	_ sdk.Msg = (*MsgUpdateProviderDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdateProviderStatusRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgRegisterProviderRequest(from sdk.AccAddress, name, identity, website, description string) *MsgRegisterProviderRequest {
	return &MsgRegisterProviderRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

func (m *MsgRegisterProviderRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Name == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "name cannot be empty")
	}
	if len(m.Name) > 64 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "name length cannot be greater than %d chars", 64)
	}
	if len(m.Identity) > 64 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "identity length cannot be greater than %d chars", 64)
	}
	if len(m.Website) > 64 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "website length cannot be greater than %d chars", 64)
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
	}
	if len(m.Description) > 256 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "description length cannot be greater than %d chars", 256)
	}

	return nil
}

func (m *MsgRegisterProviderRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateProviderDetailsRequest(from base.ProvAddress, name, identity, website, description string) *MsgUpdateProviderDetailsRequest {
	return &MsgUpdateProviderDetailsRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

func (m *MsgUpdateProviderDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if len(m.Name) > 64 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "name length cannot be greater than %d chars", 64)
	}
	if len(m.Identity) > 64 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "identity length cannot be greater than %d chars", 64)
	}
	if len(m.Website) > 64 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "website length cannot be greater than %d chars", 64)
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
	}
	if len(m.Description) > 256 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "description length cannot be greater than %d chars", 256)
	}

	return nil
}

func (m *MsgUpdateProviderDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateProviderStatusRequest(from base.ProvAddress, status v1base.Status) *MsgUpdateProviderStatusRequest {
	return &MsgUpdateProviderStatusRequest{
		From:   from.String(),
		Status: status,
	}
}

func (m *MsgUpdateProviderStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

func (m *MsgUpdateProviderStatusRequest) GetSigners() []sdk.AccAddress {
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
