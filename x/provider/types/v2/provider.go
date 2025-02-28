package v2

import (
	"fmt"
	"net/url"

	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Provider) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if len(m.Name) > 64 {
		return fmt.Errorf("name length cannot be greater than %d chars", 64)
	}
	if len(m.Identity) > 64 {
		return fmt.Errorf("identity length cannot be greater than %d chars", 64)
	}
	if len(m.Website) > 64 {
		return fmt.Errorf("website length cannot be greater than %d chars", 64)
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return sdkerrors.Wrapf(err, "invalid website %s", m.Website)
		}
	}
	if len(m.Description) > 256 {
		return fmt.Errorf("description length cannot be greater than %d chars", 256)
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return fmt.Errorf("status must be one of [active, inactive]")
	}

	return nil
}

type (
	Providers []Provider
)
