package v3

import (
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

var _ sessiontypes.Session = (*Session)(nil)
