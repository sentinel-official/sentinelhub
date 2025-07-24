package v3

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

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
	if err := validateRemoteAddrs(m.RemoteAddrs); err != nil {
		return fmt.Errorf("invalid remote_addrs: %w", err)
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

// validateRemoteAddrs validates a slice of network addresses.
//   - The slice contains between 1 and 4 addresses.
//   - No duplicate addresses are present.
//   - Each address must be in "host:port" format.
//   - The host may be a domain name, IPv4, or IPv6 address.
//   - The host part must not exceed 64 characters.
//   - IPv6 addresses must be enclosed in square brackets.
//   - The port must be a numeric value between 1 and 65535.
func validateRemoteAddrs(addrs []string) error {
	if len(addrs) < 1 || len(addrs) > 4 {
		return errors.New("must contain between 1 and 4 addrs")
	}

	seen := make(map[string]bool)
	for _, addr := range addrs {
		if seen[addr] {
			return errors.New("duplicate addr found")
		}

		seen[addr] = true

		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return fmt.Errorf("invalid format: %w", err)
		}

		if port == "" {
			return errors.New("missing port")
		}

		portNum, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("invalid port: %w", err)
		}
		if portNum < 1 || portNum > 65535 {
			return errors.New("invalid port range")
		}

		host = strings.Trim(host, "[]")
		if len(host) > 64 {
			return errors.New("host exceeds 64 characters")
		}
	}

	return nil
}
