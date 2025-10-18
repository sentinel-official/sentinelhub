package v3

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types"
)

// Ensure the message types implement sdk.Msg interface.
var (
	_ sdk.Msg = (*MsgCreatePlanRequest)(nil)
	_ sdk.Msg = (*MsgLinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUnlinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUpdatePlanDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdatePlanStatusRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
)

// NewMsgCreatePlanRequest creates a new MsgCreatePlanRequest instance.
func NewMsgCreatePlanRequest(from base.ProvAddress, bytes sdkmath.Int, duration time.Duration, prices v1base.Prices, private bool) *MsgCreatePlanRequest {
	return &MsgCreatePlanRequest{
		From:     from.String(),
		Bytes:    bytes,
		Duration: duration,
		Prices:   prices,
		Private:  private,
	}
}

// GetPrices returns the plan prices.
func (m *MsgCreatePlanRequest) GetPrices() v1base.Prices {
	return m.Prices
}

// ValidateBasic performs basic validation checks on the MsgCreatePlanRequest.
func (m *MsgCreatePlanRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.Bytes.IsNil() {
		return types.NewErrorInvalidMessage("bytes cannot be nil")
	}

	if m.Bytes.IsZero() {
		return types.NewErrorInvalidMessage("bytes cannot be zero")
	}

	if m.Bytes.IsNegative() {
		return types.NewErrorInvalidMessage("bytes cannot be negative")
	}

	if m.Duration == 0 {
		return types.NewErrorInvalidMessage("duration cannot be zero")
	}

	if m.Duration < 0 {
		return types.NewErrorInvalidMessage("duration cannot be negative")
	}

	prices := m.GetPrices()
	if err := prices.Validate(); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid prices: %w", err))
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgCreatePlanRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgLinkNodeRequest creates a new MsgLinkNodeRequest instance.
func NewMsgLinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgLinkNodeRequest {
	return &MsgLinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

// ValidateBasic performs basic validation checks on the MsgLinkNodeRequest.
func (m *MsgLinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	if m.NodeAddress == "" {
		return types.NewErrorInvalidMessage("node_address cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid node_address: %w", err))
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgLinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUnlinkNodeRequest creates a new MsgUnlinkNodeRequest instance.
func NewMsgUnlinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgUnlinkNodeRequest {
	return &MsgUnlinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

// ValidateBasic performs basic validation checks on the MsgUnlinkNodeRequest.
func (m *MsgUnlinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	if m.NodeAddress == "" {
		return types.NewErrorInvalidMessage("node_address cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid node_address: %w", err))
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUnlinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdatePlanDetailsRequest creates a new MsgUpdatePlanDetailsRequest instance.
func NewMsgUpdatePlanDetailsRequest(from base.ProvAddress, id uint64, private bool) *MsgUpdatePlanDetailsRequest {
	return &MsgUpdatePlanDetailsRequest{
		From:    from.String(),
		ID:      id,
		Private: private,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdatePlanDetailsRequest.
func (m *MsgUpdatePlanDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdatePlanDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdatePlanStatusRequest creates a new MsgUpdatePlanStatusRequest instance.
func NewMsgUpdatePlanStatusRequest(from base.ProvAddress, id uint64, status v1base.Status) *MsgUpdatePlanStatusRequest {
	return &MsgUpdatePlanStatusRequest{
		From:   from.String(),
		ID:     id,
		Status: status,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdatePlanStatusRequest.
func (m *MsgUpdatePlanStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return types.NewErrorInvalidMessage("status must be one of [active, inactive]")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdatePlanStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgStartSessionRequest creates a new MsgStartSessionRequest instance.
func NewMsgStartSessionRequest(from sdk.AccAddress, id uint64, denom string, renewalPricePolicy v1base.RenewalPricePolicy, nodeAddr base.NodeAddress) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:               from.String(),
		ID:                 id,
		Denom:              denom,
		RenewalPricePolicy: renewalPricePolicy,
		NodeAddress:        nodeAddr.String(),
	}
}

// ValidateBasic performs basic validation checks on the MsgStartSessionRequest.
func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return types.NewErrorInvalidMessage(fmt.Errorf("invalid denom: %w", err))
		}
	}

	if !m.RenewalPricePolicy.IsValid() {
		return types.NewErrorInvalidMessage("renewal_price_policy must be valid")
	}

	if m.NodeAddress == "" {
		return types.NewErrorInvalidMessage("node_address cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid node_address: %w", err))
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
