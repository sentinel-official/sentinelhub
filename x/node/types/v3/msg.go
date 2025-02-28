package v3

import (
	"net/url"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
)

var (
	_ sdk.Msg = (*MsgRegisterNodeRequest)(nil)
	_ sdk.Msg = (*MsgUpdateNodeDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdateNodeStatusRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgRegisterNodeRequest(from sdk.AccAddress, gigabytePrices, hourlyPrices v1base.Prices, remoteURL string) *MsgRegisterNodeRequest {
	return &MsgRegisterNodeRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

func (m *MsgRegisterNodeRequest) GetGigabytePrices() v1base.Prices {
	return m.GigabytePrices
}

func (m *MsgRegisterNodeRequest) GetHourlyPrices() v1base.Prices {
	return m.HourlyPrices
}

func (m *MsgRegisterNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if prices := m.GetGigabytePrices(); !prices.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "gigabyte_prices must be valid")
	}
	if prices := m.GetHourlyPrices(); !prices.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "hourly_prices must be valid")
	}
	if m.RemoteURL == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "remote_url length cannot be greater than %d chars", 64)
	}

	s, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if s.Scheme != "https" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "remote_url scheme must be https")
	}
	if s.Port() == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "remote_url port cannot be empty")
	}

	return nil
}

func (m *MsgRegisterNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateNodeDetailsRequest(from base.NodeAddress, gigabytePrices, hourlyPrices v1base.Prices, remoteURL string) *MsgUpdateNodeDetailsRequest {
	return &MsgUpdateNodeDetailsRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

func (m *MsgUpdateNodeDetailsRequest) GetGigabytePrices() v1base.Prices {
	return m.GigabytePrices
}

func (m *MsgUpdateNodeDetailsRequest) GetHourlyPrices() v1base.Prices {
	return m.HourlyPrices
}

func (m *MsgUpdateNodeDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if prices := m.GetGigabytePrices(); !prices.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "gigabyte_prices must be valid")
	}
	if prices := m.GetHourlyPrices(); !prices.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "hourly_prices must be valid")
	}
	if m.RemoteURL != "" {
		if len(m.RemoteURL) > 64 {
			return sdkerrors.Wrapf(types.ErrInvalidMessage, "remote_url length cannot be greater than %d chars", 64)
		}

		s, err := url.ParseRequestURI(m.RemoteURL)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
		if s.Scheme != "https" {
			return sdkerrors.Wrap(types.ErrInvalidMessage, "remote_url scheme must be https")
		}
		if s.Port() == "" {
			return sdkerrors.Wrap(types.ErrInvalidMessage, "remote_url port cannot be empty")
		}
	}

	return nil
}

func (m *MsgUpdateNodeDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateNodeStatusRequest(from base.NodeAddress, status v1base.Status) *MsgUpdateNodeStatusRequest {
	return &MsgUpdateNodeStatusRequest{
		From:   from.String(),
		Status: status,
	}
}

func (m *MsgUpdateNodeStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

func (m *MsgUpdateNodeStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(from sdk.AccAddress, nodeAddr base.NodeAddress, gigabytes, hours int64, denom string) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		NodeAddress: nodeAddr.String(),
		Gigabytes:   gigabytes,
		Hours:       hours,
		Denom:       denom,
	}
}

func (m *MsgStartSessionRequest) GetGigabytes() sdkmath.Int {
	return base.Gigabyte.MulRaw(m.Gigabytes)
}

func (m *MsgStartSessionRequest) GetHours() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Gigabytes == 0 && m.Hours == 0 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "[gigabytes, hours] cannot be zero")
	}
	if m.Gigabytes != 0 && m.Hours != 0 {
		return sdkerrors.Wrapf(types.ErrInvalidMessage, "[gigabytes, hours] cannot be non-zero")
	}
	if m.Gigabytes < 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "gigabytes cannot be negative")
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
