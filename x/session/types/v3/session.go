package v3

import (
	"errors"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	v1base "github.com/sentinel-official/sentinelhub/v13/types/v1"
)

// Session defines the behavior of a session, including bandwidth, timing, and status-related accessors and mutators.
type Session interface {
	proto.Message

	TotalBytes() sdkmath.Int

	GetID() uint64
	GetAccAddress() string
	GetNodeAddress() string
	GetDownloadBytes() sdkmath.Int
	GetUploadBytes() sdkmath.Int
	GetMaxBytes() sdkmath.Int
	GetDuration() time.Duration
	GetMaxDuration() time.Duration
	GetStatus() v1base.Status
	GetInactiveAt() time.Time
	GetStartAt() time.Time
	GetStatusAt() time.Time

	SetID(v uint64)
	SetAccAddress(v string)
	SetNodeAddress(v string)
	SetDownloadBytes(v sdkmath.Int)
	SetUploadBytes(v sdkmath.Int)
	SetMaxBytes(v sdkmath.Int)
	SetDuration(v time.Duration)
	SetMaxDuration(v time.Duration)
	SetStatus(v v1base.Status)
	SetInactiveAt(v time.Time)
	SetStartAt(v time.Time)
	SetStatusAt(v time.Time)
}

// TotalBytes returns the total data usage (download + upload) for the session.
func (m *BaseSession) TotalBytes() sdkmath.Int {
	return m.GetDownloadBytes().Add(m.GetUploadBytes())
}

func (m *BaseSession) GetID() uint64                 { return m.ID }
func (m *BaseSession) GetAccAddress() string         { return m.AccAddress }
func (m *BaseSession) GetNodeAddress() string        { return m.NodeAddress }
func (m *BaseSession) GetDownloadBytes() sdkmath.Int { return m.DownloadBytes }
func (m *BaseSession) GetUploadBytes() sdkmath.Int   { return m.UploadBytes }
func (m *BaseSession) GetMaxBytes() sdkmath.Int      { return m.MaxBytes }
func (m *BaseSession) GetDuration() time.Duration    { return m.Duration }
func (m *BaseSession) GetMaxDuration() time.Duration { return m.MaxDuration }
func (m *BaseSession) GetStatus() v1base.Status      { return m.Status }
func (m *BaseSession) GetInactiveAt() time.Time      { return m.InactiveAt }
func (m *BaseSession) GetStartAt() time.Time         { return m.StartAt }
func (m *BaseSession) GetStatusAt() time.Time        { return m.StatusAt }

func (m *BaseSession) SetID(v uint64)                 { m.ID = v }
func (m *BaseSession) SetAccAddress(v string)         { m.AccAddress = v }
func (m *BaseSession) SetNodeAddress(v string)        { m.NodeAddress = v }
func (m *BaseSession) SetDownloadBytes(v sdkmath.Int) { m.DownloadBytes = v }
func (m *BaseSession) SetUploadBytes(v sdkmath.Int)   { m.UploadBytes = v }
func (m *BaseSession) SetMaxBytes(v sdkmath.Int)      { m.MaxBytes = v }
func (m *BaseSession) SetDuration(v time.Duration)    { m.Duration = v }
func (m *BaseSession) SetMaxDuration(v time.Duration) { m.MaxDuration = v }
func (m *BaseSession) SetStatus(v v1base.Status)      { m.Status = v }
func (m *BaseSession) SetInactiveAt(v time.Time)      { m.InactiveAt = v }
func (m *BaseSession) SetStartAt(v time.Time)         { m.StartAt = v }
func (m *BaseSession) SetStatusAt(v time.Time)        { m.StatusAt = v }

// Validate checks whether the session's fields are properly initialized and logically consistent.
func (m *BaseSession) Validate() error {
	// Ensure ID is non-zero
	if m.ID == 0 {
		return errors.New("id cannot be zero")
	}

	// Ensure account address is non-empty and valid
	if m.AccAddress == "" {
		return errors.New("acc_address cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.AccAddress); err != nil {
		return fmt.Errorf("invalid acc_address %s: %w", m.AccAddress, err)
	}

	// Ensure node address is non-empty and valid
	if m.NodeAddress == "" {
		return errors.New("node_address cannot be empty")
	}

	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return fmt.Errorf("invalid node_address %s: %w", m.NodeAddress, err)
	}

	// DownloadBytes must be set and non-negative
	if m.DownloadBytes.IsNil() {
		return errors.New("download_bytes cannot be nil")
	}

	if m.DownloadBytes.IsNegative() {
		return errors.New("download_bytes cannot be negative")
	}

	// UploadBytes must be set and non-negative
	if m.UploadBytes.IsNil() {
		return errors.New("upload_bytes cannot be nil")
	}

	if m.UploadBytes.IsNegative() {
		return errors.New("upload_bytes cannot be negative")
	}

	// MaxBytes must be set and non-negative
	if m.MaxBytes.IsNil() {
		return errors.New("max_bytes cannot be nil")
	}

	if m.MaxBytes.IsNegative() {
		return errors.New("max_bytes cannot be negative")
	}

	// Ensure total usage doesn't exceed max limit
	if v := m.TotalBytes(); v.GT(m.MaxBytes) {
		return errors.New("total_bytes cannot be greater than max_bytes")
	}

	// Durations must be valid and within logical bounds
	if m.Duration < 0 {
		return errors.New("duration cannot be negative")
	}

	if m.MaxDuration < 0 {
		return errors.New("max_duration cannot be negative")
	}

	if m.Duration > m.MaxDuration {
		return errors.New("duration cannot be greater than max_duration")
	}

	// Status must be either active or inactive pending
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactivePending) {
		return errors.New("status must be one of [active, inactive_pending]")
	}

	// Time fields must be properly set
	if m.InactiveAt.IsZero() {
		return errors.New("inactive_at cannot be zero")
	}

	if m.StartAt.IsZero() {
		return errors.New("start_at cannot be zero")
	}

	if !m.StartAt.Before(m.InactiveAt) {
		return errors.New("start_at must be less than inactive_at")
	}

	if m.StatusAt.IsZero() {
		return errors.New("status_at cannot be zero")
	}

	return nil
}
