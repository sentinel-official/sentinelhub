package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
)

// Ensure the message types implement sdk.Msg interface.
var (
	_ sdk.Msg = (*MsgCreateAssetRequest)(nil)
	_ sdk.Msg = (*MsgDeleteAssetRequest)(nil)
	_ sdk.Msg = (*MsgUpdateAssetRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

// NewMsgCreateAssetRequest creates a new MsgCreateAssetRequest instance.
func NewMsgCreateAssetRequest(from sdk.AccAddress, denom string, decimals int64, baseAssetDenom, quoteAssetDenom string) *MsgCreateAssetRequest {
	return &MsgCreateAssetRequest{
		From:            from.String(),
		Denom:           denom,
		Decimals:        decimals,
		BaseAssetDenom:  baseAssetDenom,
		QuoteAssetDenom: quoteAssetDenom,
	}
}

// ValidateBasic performs basic validation checks on the MsgCreateAssetRequest.
func (m *MsgCreateAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if m.Denom == "" {
		return types.NewErrorInvalidMessage("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if m.Decimals < 0 {
		return types.NewErrorInvalidMessage("decimals cannot be negative")
	}

	if m.BaseAssetDenom == "" {
		return types.NewErrorInvalidMessage("base_asset_denom cannot be empty")
	}

	if m.QuoteAssetDenom == "" {
		return types.NewErrorInvalidMessage("quote_asset_denom cannot be empty")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgCreateAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgDeleteAssetRequest creates a new MsgDeleteAssetRequest instance.
func NewMsgDeleteAssetRequest(from sdk.AccAddress, denom string) *MsgDeleteAssetRequest {
	return &MsgDeleteAssetRequest{
		From:  from.String(),
		Denom: denom,
	}
}

// ValidateBasic performs basic validation checks on the MsgDeleteAssetRequest.
func (m *MsgDeleteAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if m.Denom == "" {
		return types.NewErrorInvalidMessage("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgDeleteAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateAssetRequest creates a new MsgUpdateAssetRequest instance.
func NewMsgUpdateAssetRequest(from sdk.AccAddress, denom string, decimals int64, baseAssetDenom, quoteAssetDenom string) *MsgUpdateAssetRequest {
	return &MsgUpdateAssetRequest{
		From:            from.String(),
		Denom:           denom,
		Decimals:        decimals,
		BaseAssetDenom:  baseAssetDenom,
		QuoteAssetDenom: quoteAssetDenom,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateAssetRequest.
func (m *MsgUpdateAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if m.Denom == "" {
		return types.NewErrorInvalidMessage("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if m.Decimals < 0 {
		return types.NewErrorInvalidMessage("decimals cannot be negative")
	}

	if m.BaseAssetDenom == "" {
		return types.NewErrorInvalidMessage("base_asset_denom cannot be empty")
	}

	if m.QuoteAssetDenom == "" {
		return types.NewErrorInvalidMessage("quote_asset_denom cannot be empty")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdateAssetRequest) GetSigners() []sdk.AccAddress {
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
