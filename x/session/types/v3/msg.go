package v3

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/x/session/types"
)

// Ensure the message types implement sdk.Msg interface.
var (
	_ sdk.Msg = (*MsgCancelSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

// NewMsgCancelSessionRequest creates a new MsgCancelSessionRequest instance.
func NewMsgCancelSessionRequest(from sdk.AccAddress, id uint64) *MsgCancelSessionRequest {
	return &MsgCancelSessionRequest{
		From: from.String(),
		ID:   id,
	}
}

// ValidateBasic performs basic validation checks on the MsgCancelSessionRequest.
func (m *MsgCancelSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgCancelSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateSessionRequest creates a new MsgUpdateSessionRequest instance.
func NewMsgUpdateSessionRequest(from base.NodeAddress, id uint64, downloadBytes, uploadBytes sdkmath.Int, duration time.Duration, signature []byte) *MsgUpdateSessionRequest {
	return &MsgUpdateSessionRequest{
		From:          from.String(),
		ID:            id,
		DownloadBytes: downloadBytes,
		UploadBytes:   uploadBytes,
		Duration:      duration,
		Signature:     signature,
	}
}

// Bytes returns the total transferred bytes.
func (m *MsgUpdateSessionRequest) Bytes() sdkmath.Int {
	return m.DownloadBytes.Add(m.UploadBytes)
}

// Proof returns a session proof object.
func (m *MsgUpdateSessionRequest) Proof() *Proof {
	return &Proof{
		ID:            m.ID,
		DownloadBytes: m.DownloadBytes,
		UploadBytes:   m.UploadBytes,
		Duration:      m.Duration,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateSessionRequest.
func (m *MsgUpdateSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return types.NewErrorInvalidMessage("from cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return types.NewErrorInvalidMessage(fmt.Errorf("invalid from: %w", err))
	}

	if m.ID == 0 {
		return types.NewErrorInvalidMessage("id cannot be zero")
	}

	if m.DownloadBytes.IsNil() {
		return types.NewErrorInvalidMessage("download_bytes cannot be nil")
	}

	if m.DownloadBytes.IsNegative() {
		return types.NewErrorInvalidMessage("download_bytes cannot be negative")
	}

	if m.UploadBytes.IsNil() {
		return types.NewErrorInvalidMessage("upload_bytes cannot be nil")
	}

	if m.UploadBytes.IsNegative() {
		return types.NewErrorInvalidMessage("upload_bytes cannot be negative")
	}

	if m.Duration < 0 {
		return types.NewErrorInvalidMessage("duration cannot be negative")
	}

	if m.Signature != nil {
		if len(m.Signature) != 64 {
			return types.NewErrorInvalidMessage("signature length must be 64 bytes")
		}
	}

	return nil
}

// GetSigners returns the account addresses that must sign the message.
func (m *MsgUpdateSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
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
