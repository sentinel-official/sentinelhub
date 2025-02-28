package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/node/types/v3"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

// UpdateSessionMaxValues checks the session type and updates its values if needed.
func (k *Keeper) UpdateSessionMaxValues(_ sdk.Context, session sessiontypes.Session) error {
	// Check if the session is of type v3.Session.
	_, ok := session.(*v3.Session)
	if !ok {
		return nil
	}

	// Return nil as no update needed.
	return nil
}
