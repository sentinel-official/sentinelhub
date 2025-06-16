package v1

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
)

var (
	_ sdk.Msg = (*MsgCreateAssetRequest)(nil)
	_ sdk.Msg = (*MsgDeleteAssetRequest)(nil)
	_ sdk.Msg = (*MsgUpdateAssetRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func (m *MsgCreateAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Decimals < 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "decimals cannot be negative")
	}
	if m.BaseAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "base_asset_denom cannot be empty")
	}
	if m.QuoteAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "quote_asset_denom cannot be empty")
	}

	return nil
}

func (m *MsgCreateAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func (m *MsgDeleteAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgDeleteAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func (m *MsgUpdateAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Decimals < 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "decimals cannot be negative")
	}
	if m.BaseAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "base_asset_denom cannot be empty")
	}
	if m.QuoteAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "quote_asset_denom cannot be empty")
	}

	return nil
}

func (m *MsgUpdateAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
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
