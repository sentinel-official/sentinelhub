package v1

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

// DepositAmount calculates the total amount to be deposited for the lease
func (m *Lease) DepositAmount() sdk.Coin {
	amount := m.Price.QuoteValue.MulRaw(m.MaxHours)
	return sdk.Coin{Denom: m.Price.Denom, Amount: amount}
}

// GetHours returns the duration the lease has been used in time.Duration.
func (m *Lease) GetHours() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

// GetMaxHours returns the total maximum duration of the lease in time.Duration.
func (m *Lease) GetMaxHours() time.Duration {
	return time.Duration(m.MaxHours) * time.Hour
}

// InactiveAt returns the time when the lease becomes inactive,
func (m *Lease) InactiveAt() time.Time {
	return m.StartAt.Add(m.GetMaxHours())
}

// PayoutAt returns the time when the lease's payout should occur.
func (m *Lease) PayoutAt() time.Time {
	if m.Hours < m.MaxHours {
		return m.StartAt.Add(m.GetHours())
	}

	return time.Time{}
}

// RefundAmount calculates the amount to be refunded based on unused hours.
func (m *Lease) RefundAmount() sdk.Coin {
	diff := m.MaxHours - m.Hours
	if diff < 0 {
		panic("invalid refund hours")
	}

	amount := m.Price.QuoteValue.MulRaw(diff)
	return sdk.Coin{Denom: m.Price.Denom, Amount: amount}
}

// RenewalAt returns the renewal time of the lease.
func (m *Lease) RenewalAt() time.Time {
	if m.RenewalPricePolicy.Equal(v1base.RenewalPricePolicyUnspecified) {
		return time.Time{}
	}

	return m.InactiveAt()
}

// ValidateRenewalPolicies validates the renewal pricing policy.
func (m *Lease) ValidateRenewalPolicies(price v1base.Price) error {
	if err := m.RenewalPricePolicy.Validate(price, m.Price); err != nil {
		return fmt.Errorf("invalid renewal price policy: %w", err)
	}

	return nil
}

// Validate checks the correctness of all Lease fields and returns an error if any are invalid.
func (m *Lease) Validate() error {
	if m.ID == 0 {
		return errors.New("id cannot be zero")
	}
	if m.ProvAddress == "" {
		return errors.New("prov_address cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.ProvAddress); err != nil {
		return fmt.Errorf("invalid prov_address %s: %w", m.ProvAddress, err)
	}
	if m.NodeAddress == "" {
		return errors.New("node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return fmt.Errorf("invalid node_address %s: %w", m.NodeAddress, err)
	}
	if err := m.Price.Validate(); err != nil {
		return fmt.Errorf("invalid price %s: %w", m.Price, err)
	}
	if m.Hours < 0 {
		return errors.New("hours cannot be negative")
	}
	if m.MaxHours < 0 {
		return errors.New("max_hours cannot be negative")
	}
	if m.MaxHours == 0 {
		return errors.New("max_hours cannot be zero")
	}
	if m.MaxHours < m.Hours {
		return errors.New("max_hours cannot be less than hours")
	}
	if !m.RenewalPricePolicy.IsValid() {
		return errors.New("renewal_price_policy must be valid")
	}
	if m.StartAt.IsZero() {
		return errors.New("start_at cannot be zero")
	}

	return nil
}
