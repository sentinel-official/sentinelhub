package v3

import (
	"errors"
	"fmt"
	"net/url"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
)

// GetGigabytePrices returns the node's defined prices per gigabyte.
func (m *Node) GetGigabytePrices() v1base.Prices {
	return m.GigabytePrices
}

// GetHourlyPrices returns the node's defined prices per hour.
func (m *Node) GetHourlyPrices() v1base.Prices {
	return m.HourlyPrices
}

// Validate checks the integrity and validity of the Node's fields.
func (m *Node) Validate() error {
	if m.Address == "" {
		return errors.New("address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}
	if prices := m.GetGigabytePrices(); !prices.IsValid() {
		return errors.New("gigabyte_prices must be valid")
	}
	if prices := m.GetHourlyPrices(); !prices.IsValid() {
		return errors.New("hourly_prices must be valid")
	}
	if m.RemoteURL == "" {
		return errors.New("remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length cannot be greater than %d chars", 64)
	}

	// Parse and validate remote URL format and contents
	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return fmt.Errorf("invalid remote_url: %w", err)
	}
	if remoteURL.Scheme != "https" {
		return errors.New("remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return errors.New("remote_url port cannot be empty")
	}

	// Validate status vs. inactive timestamp logic
	if m.InactiveAt.IsZero() {
		if !m.Status.Equal(v1base.StatusInactive) {
			return fmt.Errorf("invalid inactive_at %s; expected positive", m.InactiveAt)
		}
	}
	if !m.InactiveAt.IsZero() {
		if !m.Status.Equal(v1base.StatusActive) {
			return fmt.Errorf("invalid inactive_at %s; expected zero", m.InactiveAt)
		}
	}

	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return errors.New("status must be one of [active, inactive]")
	}
	if m.StatusAt.IsZero() {
		return errors.New("status_at cannot be zero")
	}

	return nil
}

// GigabytePrice returns the price per gigabyte for the given denom, if found.
func (m *Node) GigabytePrice(denom string) (v1base.Price, bool) {
	prices := m.GetGigabytePrices()
	if prices.Len() == 0 {
		return v1base.ZeroPrice(denom), true
	}

	price, found := prices.Find(denom)
	if !found {
		return v1base.Price{}, false
	}

	return price, true
}

// HourlyPrice returns the price per hour for the given denom, if found.
func (m *Node) HourlyPrice(denom string) (v1base.Price, bool) {
	prices := m.GetHourlyPrices()
	if prices.Len() == 0 {
		return v1base.ZeroPrice(denom), true
	}

	price, found := prices.Find(denom)
	if !found {
		return v1base.Price{}, false
	}

	return price, true
}
