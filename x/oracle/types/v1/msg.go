package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/third_party/osmosis/x/poolmanager/client/queryproto"
	protorev "github.com/sentinel-official/sentinelhub/v12/third_party/osmosis/x/protorev/types"
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
func NewMsgCreateAssetRequest(
	from sdk.AccAddress, denom string, decimals int64, protoRevPoolRequest protorev.QueryGetProtoRevPoolRequest,
	spotPriceRequest queryproto.SpotPriceRequest,
) *MsgCreateAssetRequest {
	return &MsgCreateAssetRequest{
		From:                from.String(),
		Denom:               denom,
		Decimals:            decimals,
		ProtoRevPoolRequest: protoRevPoolRequest,
		SpotPriceRequest:    spotPriceRequest,
	}
}

// ValidateBasic performs basic validation checks on the MsgCreateAssetRequest.
func (m *MsgCreateAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.Denom == "" {
		return types.NewErrorInvalidMessage("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid denom: %w", err))
	}

	if m.Decimals < 0 {
		return types.NewErrorInvalidMessage("decimals cannot be negative")
	}

	if m.ProtoRevPoolRequest.BaseDenom == "" {
		return types.NewErrorInvalidMessage("base_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.ProtoRevPoolRequest.BaseDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid base_denom: %w", err))
	}

	if m.ProtoRevPoolRequest.OtherDenom == "" {
		return types.NewErrorInvalidMessage("other_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.ProtoRevPoolRequest.OtherDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid other_denom: %w", err))
	}

	if m.ProtoRevPoolRequest.BaseDenom == m.ProtoRevPoolRequest.OtherDenom {
		return types.NewErrorInvalidMessage("base_denom and other_denom cannot be the same")
	}

	if m.SpotPriceRequest.BaseAssetDenom == "" {
		return types.NewErrorInvalidMessage("base_asset_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.SpotPriceRequest.BaseAssetDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid base_asset_denom: %w", err))
	}

	if m.SpotPriceRequest.QuoteAssetDenom == "" {
		return types.NewErrorInvalidMessage("quote_asset_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.SpotPriceRequest.QuoteAssetDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid quote_asset_denom: %w", err))
	}

	if m.SpotPriceRequest.BaseAssetDenom == m.SpotPriceRequest.QuoteAssetDenom {
		return types.NewErrorInvalidMessage("base_asset_denom and quote_asset_denom cannot be the same")
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
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.Denom == "" {
		return types.NewErrorInvalidMessage("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid denom: %w", err))
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
func NewMsgUpdateAssetRequest(
	from sdk.AccAddress, denom string, decimals int64, protoRevPoolRequest protorev.QueryGetProtoRevPoolRequest,
	spotPriceRequest queryproto.SpotPriceRequest,
) *MsgUpdateAssetRequest {
	return &MsgUpdateAssetRequest{
		From:                from.String(),
		Denom:               denom,
		Decimals:            decimals,
		ProtoRevPoolRequest: protoRevPoolRequest,
		SpotPriceRequest:    spotPriceRequest,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateAssetRequest.
func (m *MsgUpdateAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.Denom == "" {
		return types.NewErrorInvalidMessage("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid denom: %w", err))
	}

	if m.Decimals < 0 {
		return types.NewErrorInvalidMessage("decimals cannot be negative")
	}

	if m.ProtoRevPoolRequest.BaseDenom == "" {
		return types.NewErrorInvalidMessage("base_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.ProtoRevPoolRequest.BaseDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid base_denom: %w", err))
	}

	if m.ProtoRevPoolRequest.OtherDenom == "" {
		return types.NewErrorInvalidMessage("other_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.ProtoRevPoolRequest.OtherDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid other_denom: %w", err))
	}

	if m.ProtoRevPoolRequest.BaseDenom == m.ProtoRevPoolRequest.OtherDenom {
		return types.NewErrorInvalidMessage("base_denom and other_denom cannot be the same")
	}

	if m.SpotPriceRequest.BaseAssetDenom == "" {
		return types.NewErrorInvalidMessage("base_asset_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.SpotPriceRequest.BaseAssetDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid base_asset_denom: %w", err))
	}

	if m.SpotPriceRequest.QuoteAssetDenom == "" {
		return types.NewErrorInvalidMessage("quote_asset_denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.SpotPriceRequest.QuoteAssetDenom); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid quote_asset_denom: %w", err))
	}

	if m.SpotPriceRequest.BaseAssetDenom == m.SpotPriceRequest.QuoteAssetDenom {
		return types.NewErrorInvalidMessage("base_asset_denom and quote_asset_denom cannot be the same")
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
