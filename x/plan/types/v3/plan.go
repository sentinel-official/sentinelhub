package v3

import (
	"errors"
	"fmt"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

// GetPrices returns the list of prices associated with the plan.
func (m *Plan) GetPrices() v1base.Prices {
	return m.Prices
}

// Price returns the price for the given denom, or false if not found.
func (m *Plan) Price(denom string) (v1base.Price, bool) {
	prices := m.GetPrices()
	if prices.Len() == 0 {
		return v1base.ZeroPrice(denom), true
	}

	price, found := prices.Find(denom)
	if !found {
		return v1base.Price{}, false
	}

	return price, true
}

// Validate validates the Plan fields for basic correctness.
func (m *Plan) Validate() error {
	if m.ID == 0 {
		return errors.New("id cannot be zero")
	}

	if m.ProvAddress == "" {
		return errors.New("prov_address cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.ProvAddress); err != nil {
		return fmt.Errorf("invalid prov_address: %w", err)
	}

	if m.Bytes.IsNil() {
		return errors.New("bytes cannot be nil")
	}

	if m.Bytes.IsZero() {
		return errors.New("bytes cannot be zero")
	}

	if m.Bytes.IsNegative() {
		return errors.New("bytes cannot be negative")
	}

	if m.Duration == 0 {
		return errors.New("duration cannot be zero")
	}

	if m.Duration < 0 {
		return errors.New("duration cannot be negative")
	}

	prices := m.GetPrices()
	if err := prices.Validate(); err != nil {
		return fmt.Errorf("invalid prices: %w", err)
	}

	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return errors.New("status must be one of [active, inactive]")
	}

	if m.StatusAt.IsZero() {
		return errors.New("status_at cannot be zero")
	}

	return nil
}
