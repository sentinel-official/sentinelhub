package v3

import (
	"errors"

	sessiontypes "github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

var _ sessiontypes.Session = (*Session)(nil)

// Validate performs basic validation on the session.
func (m *Session) Validate() error {
	if err := m.BaseSession.Validate(); err != nil {
		return err
	}
	if m.SubscriptionID == 0 {
		return errors.New("subscription id cannot be zero")
	}

	return nil
}
