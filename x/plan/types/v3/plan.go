package v3

import (
	"fmt"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Plan) GetGigabytes() sdkmath.Int {
	return base.Gigabyte.MulRaw(m.Gigabytes)
}

func (m *Plan) GetHours() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

func (m *Plan) GetPrices() v1base.Prices {
	return m.Prices
}

func (m *Plan) IsPrivate() bool {
	return m.GetPrices().Len() == 0
}

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

func (m *Plan) Validate() error {
	if m.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.ProvAddress == "" {
		return fmt.Errorf("prov_address cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.ProvAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid prov_address %s", m.ProvAddress)
	}
	if m.Gigabytes < 0 {
		return fmt.Errorf("gigabytes cannot be negative")
	}
	if m.Gigabytes == 0 {
		return fmt.Errorf("gigabytes cannot be zero")
	}
	if m.Hours < 0 {
		return fmt.Errorf("hours cannot be negative")
	}
	if m.Hours == 0 {
		return fmt.Errorf("hours cannot be zero")
	}
	if prices := m.GetPrices(); !prices.IsValid() {
		return fmt.Errorf("prices must be valid")
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return fmt.Errorf("status must be one of [active, inactive]")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}
