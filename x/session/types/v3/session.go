package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/gogoproto/proto"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

type Session interface {
	proto.Message

	Bytes() sdkmath.Int

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

func (m *BaseSession) Bytes() sdkmath.Int {
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
