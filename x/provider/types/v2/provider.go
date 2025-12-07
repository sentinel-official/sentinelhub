package v2

import (
	"errors"
	"fmt"
	"net/url"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	v1base "github.com/sentinel-official/sentinelhub/v13/types/v1"
)

func (m *Provider) Validate() error {
	if m.Address == "" {
		return errors.New("address cannot be empty")
	}

	if _, err := base.ProvAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("invalid address %s: %w", m.Address, err)
	}

	if m.Name == "" {
		return errors.New("name cannot be empty")
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
			return fmt.Errorf("invalid website %s: %w", m.Website, err)
		}
	}

	if len(m.Description) > 256 {
		return fmt.Errorf("description length cannot be greater than %d chars", 256)
	}

	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return errors.New("status must be one of [active, inactive]")
	}

	return nil
}

type (
	Providers []Provider
)
