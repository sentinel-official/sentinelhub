package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

// UpdateSessionMaxValues calculates and updates the maximum allowable duration and bytes for a session.
func (k *Keeper) UpdateSessionMaxValues(ctx sdk.Context, session sessiontypes.Session) error {
	// Ensure the session is of type v3.Session
	s, ok := session.(*v3.Session)
	if !ok {
		return nil
	}

	// Convert the account address from Bech32 format
	accAddr, err := sdk.AccAddressFromBech32(s.AccAddress)
	if err != nil {
		return err
	}

	// Retrieve the subscription associated with the session
	subscription, found := k.GetSubscription(ctx, s.SubscriptionID)
	if !found {
		return nil
	}

	// Retrieve the allocation details for the account
	alloc, found := k.GetAllocation(ctx, subscription.ID, accAddr)
	if !found {
		return nil
	}

	// Calculate the maximum allowable session bytes and duration
	diffBytes := alloc.GrantedBytes.Sub(alloc.UtilisedBytes)
	maxBytes := s.Bytes().Add(diffBytes)
	maxDuration := subscription.InactiveAt.Sub(s.StartAt)

	// Update the session with calculated max bytes and duration
	s.SetMaxBytes(maxBytes)
	s.SetMaxDuration(maxDuration)

	return nil
}
