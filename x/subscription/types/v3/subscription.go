package v3

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

// RenewalAt returns the renewal time for the subscription.
func (m *Subscription) RenewalAt() time.Time {
	if m.RenewalPricePolicy.Equal(v1base.RenewalPricePolicyUnspecified) {
		return time.Time{}
	}

	return m.InactiveAt
}

// ValidateRenewalPolicies checks if the renewal policy is valid under the given quoted price.
func (m *Subscription) ValidateRenewalPolicies(price v1base.Price) error {
	if err := m.RenewalPricePolicy.Validate(price, m.Price); err != nil {
		return fmt.Errorf("invalid renewal price policy: %w", err)
	}

	return nil
}

// Validate performs basic validation checks on the subscription fields.
func (m *Subscription) Validate() error {
	if m.ID == 0 {
		return errors.New("id cannot be zero")
	}
	if m.AccAddress == "" {
		return errors.New("acc_address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.AccAddress); err != nil {
		return fmt.Errorf("invalid acc_address: %w", err)
	}
	if m.PlanID == 0 {
		return errors.New("plan_id cannot be zero")
	}
	if err := m.Price.Validate(); err != nil {
		return fmt.Errorf("invalid price: %w", err)
	}
	if !m.RenewalPricePolicy.IsValid() {
		return errors.New("renewal_price_policy must be valid")
	}
	if m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactivePending) {
		return errors.New("status must be one of [active, inactive_pending]")
	}
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
