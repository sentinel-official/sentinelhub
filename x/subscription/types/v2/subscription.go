package v2

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

type (
	Subscription interface {
		proto.Message
		Type() SubscriptionType
		Validate() error
		GetID() uint64
		GetAddress() sdk.AccAddress
		GetInactiveAt() time.Time
		GetStatus() v1base.Status
		GetStatusAt() time.Time
		SetInactiveAt(v time.Time)
		SetStatus(v v1base.Status)
		SetStatusAt(v time.Time)
	}
	Subscriptions []Subscription
)

var (
	_ Subscription = (*NodeSubscription)(nil)
	_ Subscription = (*PlanSubscription)(nil)
)

func (s *BaseSubscription) GetID() uint64            { return s.ID }
func (s *BaseSubscription) GetInactiveAt() time.Time { return s.InactiveAt }
func (s *BaseSubscription) GetStatus() v1base.Status { return s.Status }
func (s *BaseSubscription) GetStatusAt() time.Time   { return s.StatusAt }

func (s *BaseSubscription) GetAddress() sdk.AccAddress {
	if s.Address == "" {
		return nil
	}

	addr, err := sdk.AccAddressFromBech32(s.Address)
	if err != nil {
		panic(err)
	}

	return addr
}

func (s *BaseSubscription) SetInactiveAt(v time.Time) { s.InactiveAt = v }
func (s *BaseSubscription) SetStatus(v v1base.Status) { s.Status = v }
func (s *BaseSubscription) SetStatusAt(v time.Time)   { s.StatusAt = v }

func (s *BaseSubscription) Validate() error {
	return nil
}

func (s *NodeSubscription) Type() SubscriptionType {
	return TypeNode
}

func (s *NodeSubscription) Validate() error {
	return nil
}

func (s *PlanSubscription) Type() SubscriptionType {
	return TypePlan
}

func (s *PlanSubscription) Validate() error {
	return nil
}
