package v3

import (
	"errors"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

// GetBytes returns the plan's bandwidth as sdkmath.Int in bytes.
func (m *Plan) GetBytes() sdkmath.Int {
	return base.Gigabyte.MulRaw(m.Gigabytes)
}

// GetDuration returns the plan's duration as a time.Duration value.
func (m *Plan) GetDuration() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

// GetPrices returns the list of prices associated with the plan.
func (m *Plan) GetPrices() v1base.Prices {
	return m.Prices
}

// IsPrivate returns true if the plan has no prices and is therefore private.
func (m *Plan) IsPrivate() bool {
	return m.GetPrices().Len() == 0
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

	if m.Gigabytes < 0 {
		return errors.New("gigabytes cannot be negative")
	}

	if m.Gigabytes == 0 {
		return errors.New("gigabytes cannot be zero")
	}

	if m.Hours < 0 {
		return errors.New("hours cannot be negative")
	}

	if m.Hours == 0 {
		return errors.New("hours cannot be zero")
	}

	if prices := m.GetPrices(); !prices.IsValid() {
		return errors.New("prices must be valid")
	}

	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return errors.New("status must be one of [active, inactive]")
	}

	if m.StatusAt.IsZero() {
		return errors.New("status_at cannot be zero")
	}

	return nil
}
