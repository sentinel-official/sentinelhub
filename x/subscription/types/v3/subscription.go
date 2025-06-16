package v3

import (
	"fmt"
	"time"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

func (m *Subscription) RenewalAt() time.Time {
	if m.RenewalPricePolicy.Equal(v1base.RenewalPricePolicyUnspecified) {
		return time.Time{}
	}

	return m.InactiveAt
}

func (m *Subscription) ValidateRenewalPolicies(price v1base.Price) error {
	if err := m.RenewalPricePolicy.Validate(price, m.Price); err != nil {
		return fmt.Errorf("invalid renewal price policy: %w", err)
	}

	return nil
}

func (m *Subscription) Validate() error {
	return nil
}
