package v1

import (
	"fmt"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Lease) DepositAmount() sdk.Coin {
	amount := m.Price.QuoteValue.MulRaw(m.MaxHours)
	return sdk.Coin{Denom: m.Price.Denom, Amount: amount}
}

func (m *Lease) GetHours() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

func (m *Lease) GetMaxHours() time.Duration {
	return time.Duration(m.MaxHours) * time.Hour
}

func (m *Lease) InactiveAt() time.Time {
	return m.StartAt.Add(m.GetMaxHours())
}

func (m *Lease) PayoutAt() time.Time {
	if m.Hours < m.MaxHours {
		return m.StartAt.Add(m.GetHours())
	}

	return time.Time{}
}

func (m *Lease) RefundAmount() sdk.Coin {
	amount := m.Price.QuoteValue.MulRaw(m.MaxHours - m.Hours)
	return sdk.Coin{Denom: m.Price.Denom, Amount: amount}
}

func (m *Lease) RenewalAt() time.Time {
	if m.RenewalPricePolicy.Equal(v1base.RenewalPricePolicyUnspecified) {
		return time.Time{}
	}

	return m.InactiveAt()
}

func (m *Lease) ValidateRenewalPolicies(price v1base.Price) error {
	if err := m.RenewalPricePolicy.Validate(price, m.Price); err != nil {
		return fmt.Errorf("invalid renewal price policy: %w", err)
	}

	return nil
}

func (m *Lease) Validate() error {
	if m.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.ProvAddress == "" {
		return fmt.Errorf("prov_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.ProvAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid prov_address %s", m.ProvAddress)
	}
	if m.NodeAddress == "" {
		return fmt.Errorf("node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid node_address %s", m.NodeAddress)
	}
	if !m.Price.IsValid() {
		return fmt.Errorf("price must be valid")
	}
	if m.Hours <= 0 {
		return fmt.Errorf("hours must be greater than zero")
	}
	if m.MaxHours <= 0 {
		return fmt.Errorf("max_hours must be greater than zero")
	}
	if m.MaxHours < m.Hours {
		return fmt.Errorf("max_hours cannot be less than hours")
	}

	return nil
}
