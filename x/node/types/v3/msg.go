package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types"
)

// Ensure the message types implement sdk.Msg interface.
var (
	_ sdk.Msg = (*MsgRegisterNodeRequest)(nil)
	_ sdk.Msg = (*MsgUpdateNodeDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdateNodeStatusRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

// NewMsgRegisterNodeRequest creates a MsgRegisterNodeRequest with pricing and remote URL for a new node registration.
func NewMsgRegisterNodeRequest(from sdk.AccAddress, gigabytePrices, hourlyPrices v1base.Prices, remoteAddrs []string) *MsgRegisterNodeRequest {
	return &MsgRegisterNodeRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteAddrs:    remoteAddrs,
	}
}

// GetGigabytePrices returns the gigabyte pricing from the MsgRegisterNodeRequest.
func (m *MsgRegisterNodeRequest) GetGigabytePrices() v1base.Prices {
	return m.GigabytePrices
}

// GetHourlyPrices returns the hourly pricing from the MsgRegisterNodeRequest.
func (m *MsgRegisterNodeRequest) GetHourlyPrices() v1base.Prices {
	return m.HourlyPrices
}

// ValidateBasic performs basic validation of MsgRegisterNodeRequest fields such as address, pricing, and URL formatting.
func (m *MsgRegisterNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if prices := m.GetGigabytePrices(); !prices.IsValid() {
		return types.NewErrorInvalidMessage("gigabyte_prices must be valid")
	}

	if prices := m.GetHourlyPrices(); !prices.IsValid() {
		return types.NewErrorInvalidMessage("hourly_prices must be valid")
	}

	if err := validateRemoteAddrs(m.RemoteAddrs); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	return nil
}

// GetSigners returns the account address that must sign the MsgRegisterNodeRequest.
func (m *MsgRegisterNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateNodeDetailsRequest creates a MsgUpdateNodeDetailsRequest for updating a node’s prices or remote URL.
func NewMsgUpdateNodeDetailsRequest(from base.NodeAddress, gigabytePrices, hourlyPrices v1base.Prices, remoteAddrs []string) *MsgUpdateNodeDetailsRequest {
	return &MsgUpdateNodeDetailsRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteAddrs:    remoteAddrs,
	}
}

// GetGigabytePrices returns the updated gigabyte pricing in MsgUpdateNodeDetailsRequest.
func (m *MsgUpdateNodeDetailsRequest) GetGigabytePrices() v1base.Prices {
	return m.GigabytePrices
}

// GetHourlyPrices returns the updated hourly pricing in MsgUpdateNodeDetailsRequest.
func (m *MsgUpdateNodeDetailsRequest) GetHourlyPrices() v1base.Prices {
	return m.HourlyPrices
}

// ValidateBasic performs basic validation for MsgUpdateNodeDetailsRequest including address, pricing, and remote URL if set.
func (m *MsgUpdateNodeDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if prices := m.GetGigabytePrices(); !prices.IsValid() {
		return types.NewErrorInvalidMessage("gigabyte_prices must be valid")
	}

	if prices := m.GetHourlyPrices(); !prices.IsValid() {
		return types.NewErrorInvalidMessage("hourly_prices must be valid")
	}

	if len(m.RemoteAddrs) > 0 {
		if err := validateRemoteAddrs(m.RemoteAddrs); err != nil {
			return types.NewErrorInvalidMessage(err)
		}
	}

	return nil
}

// GetSigners returns the node address that must sign the MsgUpdateNodeDetailsRequest.
func (m *MsgUpdateNodeDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateNodeStatusRequest creates a MsgUpdateNodeStatusRequest to change a node's active/inactive status.
func NewMsgUpdateNodeStatusRequest(from base.NodeAddress, status v1base.Status) *MsgUpdateNodeStatusRequest {
	return &MsgUpdateNodeStatusRequest{
		From:   from.String(),
		Status: status,
	}
}

// ValidateBasic checks MsgUpdateNodeStatusRequest for valid address and that the status is active or inactive.
func (m *MsgUpdateNodeStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return types.NewErrorInvalidMessage("status must be one of [active, inactive]")
	}

	return nil
}

// GetSigners returns the node address that must sign the MsgUpdateNodeStatusRequest.
func (m *MsgUpdateNodeStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgStartSessionRequest creates a MsgStartSessionRequest to initiate a usage session with a node.
func NewMsgStartSessionRequest(from sdk.AccAddress, nodeAddr base.NodeAddress, gigabytes, hours int64, maxPrice v1base.Price) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		NodeAddress: nodeAddr.String(),
		Gigabytes:   gigabytes,
		Hours:       hours,
		MaxPrice:    maxPrice,
	}
}

// GetGigabytes returns the total data usage in bytes for MsgStartSessionRequest.
func (m *MsgStartSessionRequest) GetGigabytes() sdkmath.Int {
	return base.Gigabyte.MulRaw(m.Gigabytes)
}

// GetHours returns the session duration in time.Duration from MsgStartSessionRequest.
func (m *MsgStartSessionRequest) GetHours() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

// ValidateBasic checks MsgStartSessionRequest for valid sender, node, usage mode (hour/data), and max price constraints.
func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if m.NodeAddress == "" {
		return types.NewErrorInvalidMessage("node_address cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return types.NewErrorInvalidMessage(err)
	}

	if m.Gigabytes == 0 && m.Hours == 0 {
		return types.NewErrorInvalidMessage("[gigabytes, hours] cannot be zero")
	}

	if m.Gigabytes != 0 && m.Hours != 0 {
		return types.NewErrorInvalidMessage("[gigabytes, hours] cannot be non-zero")
	}

	if m.Gigabytes < 0 {
		return types.NewErrorInvalidMessage("gigabytes cannot be negative")
	}

	if m.Hours < 0 {
		return types.NewErrorInvalidMessage("hours cannot be negative")
	}

	if m.MaxPrice.Denom != "" {
		if err := m.MaxPrice.Validate(); err != nil {
			return types.NewErrorInvalidMessage(err)
		}
	}

	return nil
}

// GetSigners returns the account address that must sign the MsgStartSessionRequest.
func (m *MsgStartSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateParamsRequest creates a MsgUpdateParamsRequest to update module parameters by governance or admin.
func NewMsgUpdateParamsRequest(from sdk.AccAddress, params Params) *MsgUpdateParamsRequest {
	return &MsgUpdateParamsRequest{
		From:   from.String(),
		Params: params,
	}
}

// ValidateBasic checks MsgUpdateParamsRequest for valid address and parameter content.
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

// GetSigners returns the account address that must sign the MsgUpdateParamsRequest.
func (m *MsgUpdateParamsRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
