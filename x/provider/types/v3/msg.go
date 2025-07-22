package v3

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/provider/types"
)

// Ensure the message types implement sdk.Msg interface
var (
	_ sdk.Msg = (*MsgRegisterProviderRequest)(nil)
	_ sdk.Msg = (*MsgUpdateProviderDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdateProviderStatusRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

// NewMsgRegisterProviderRequest creates a new MsgRegisterProviderRequest instance.
func NewMsgRegisterProviderRequest(from sdk.AccAddress, name, identity, website, description string) *MsgRegisterProviderRequest {
	return &MsgRegisterProviderRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

// ValidateBasic performs basic validation checks on the MsgRegisterProviderRequest.
func (m *MsgRegisterProviderRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if m.Name == "" {
		return types.NewErrorInvalidMessage("name cannot be empty")
	}
	if len(m.Name) > 64 {
		return types.NewErrorInvalidMessage("name length cannot be greater than 64 chars")
	}
	if len(m.Identity) > 64 {
		return types.NewErrorInvalidMessage("identity length cannot be greater than 64 chars")
	}
	if len(m.Website) > 64 {
		return types.NewErrorInvalidMessage("website length cannot be greater than 64 chars")
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return types.NewErrorInvalidMessage(err)
		}
	}
	if len(m.Description) > 256 {
		return types.NewErrorInvalidMessage("description length cannot be greater than 256 chars")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgRegisterProviderRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateProviderDetailsRequest creates a new MsgUpdateProviderDetailsRequest instance.
func NewMsgUpdateProviderDetailsRequest(from base.ProvAddress, name, identity, website, description string) *MsgUpdateProviderDetailsRequest {
	return &MsgUpdateProviderDetailsRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateProviderDetailsRequest.
func (m *MsgUpdateProviderDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if len(m.Name) > 64 {
		return types.NewErrorInvalidMessage("name length cannot be greater than 64 chars")
	}
	if len(m.Identity) > 64 {
		return types.NewErrorInvalidMessage("identity length cannot be greater than 64 chars")
	}
	if len(m.Website) > 64 {
		return types.NewErrorInvalidMessage("website length cannot be greater than 64 chars")
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return types.NewErrorInvalidMessage(err)
		}
	}
	if len(m.Description) > 256 {
		return types.NewErrorInvalidMessage("description length cannot be greater than 256 chars")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdateProviderDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateProviderStatusRequest creates a new MsgUpdateProviderStatusRequest instance.
func NewMsgUpdateProviderStatusRequest(from base.ProvAddress, status v1base.Status) *MsgUpdateProviderStatusRequest {
	return &MsgUpdateProviderStatusRequest{
		From:   from.String(),
		Status: status,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateProviderStatusRequest.
func (m *MsgUpdateProviderStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return types.NewErrorInvalidMessage("status must be one of [active, inactive]")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdateProviderStatusRequest) GetSigners() []sdk.AccAddress {
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
