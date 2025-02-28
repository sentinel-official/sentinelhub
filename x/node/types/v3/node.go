package v3

import (
	"errors"
	"fmt"
	"net/url"

	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Node) GetGigabytePrices() v1base.Prices {
	return m.GigabytePrices
}

func (m *Node) GetHourlyPrices() v1base.Prices {
	return m.HourlyPrices
}

func (m *Node) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrapf(err, "invalid address %s", m.Address)
	}
	if prices := m.GetGigabytePrices(); !prices.IsValid() {
		return errors.New("gigabyte_prices must be valid")
	}
	if prices := m.GetHourlyPrices(); !prices.IsValid() {
		return errors.New("hourly_prices must be valid")
	}
	if m.RemoteURL == "" {
		return fmt.Errorf("remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length cannot be greater than %d chars", 64)
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid remote_url %s", m.RemoteURL)
	}
	if remoteURL.Scheme != "https" {
		return fmt.Errorf("remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return fmt.Errorf("remote_url port cannot be empty")
	}

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
		return fmt.Errorf("status must be one of [active, inactive]")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

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
