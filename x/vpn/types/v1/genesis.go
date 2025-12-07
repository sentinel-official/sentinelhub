package v1

import (
	"fmt"

	deposittypes "github.com/sentinel-official/sentinelhub/v13/x/deposit/types/v1"
	leasetypes "github.com/sentinel-official/sentinelhub/v13/x/lease/types/v1"
	nodetypes "github.com/sentinel-official/sentinelhub/v13/x/node/types/v3"
	plantypes "github.com/sentinel-official/sentinelhub/v13/x/plan/types/v3"
	providertypes "github.com/sentinel-official/sentinelhub/v13/x/provider/types/v3"
	sessiontypes "github.com/sentinel-official/sentinelhub/v13/x/session/types/v3"
	subscriptiontypes "github.com/sentinel-official/sentinelhub/v13/x/subscription/types/v3"
)

func NewGenesisState(
	deposit *deposittypes.GenesisState,
	lease *leasetypes.GenesisState,
	node *nodetypes.GenesisState,
	plan *plantypes.GenesisState,
	provider *providertypes.GenesisState,
	session *sessiontypes.GenesisState,
	subscription *subscriptiontypes.GenesisState,
) *GenesisState {
	return &GenesisState{
		Deposit:      deposit,
		Lease:        lease,
		Node:         node,
		Plan:         plan,
		Provider:     provider,
		Session:      session,
		Subscription: subscription,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		deposittypes.DefaultGenesisState(),
		leasetypes.DefaultGenesisState(),
		nodetypes.DefaultGenesisState(),
		plantypes.DefaultGenesisState(),
		providertypes.DefaultGenesisState(),
		sessiontypes.DefaultGenesisState(),
		subscriptiontypes.DefaultGenesisState(),
	)
}

func (m *GenesisState) Validate() error {
	if err := deposittypes.ValidateGenesisState(m.Deposit); err != nil {
		return fmt.Errorf("invalid deposit genesis state: %w", err)
	}

	if err := leasetypes.ValidateGenesis(m.Lease); err != nil {
		return fmt.Errorf("invalid lease genesis state: %w", err)
	}

	if err := nodetypes.ValidateGenesis(m.Node); err != nil {
		return fmt.Errorf("invalid node genesis state: %w", err)
	}

	if err := plantypes.ValidateGenesis(m.Plan); err != nil {
		return fmt.Errorf("invalid plan genesis state: %w", err)
	}

	if err := providertypes.ValidateGenesis(m.Provider); err != nil {
		return fmt.Errorf("invalid provider genesis state: %w", err)
	}

	if err := sessiontypes.ValidateGenesis(m.Session); err != nil {
		return fmt.Errorf("invalid session genesis state: %w", err)
	}

	if err := subscriptiontypes.ValidateGenesis(m.Subscription); err != nil {
		return fmt.Errorf("invalid subscription genesis state: %w", err)
	}

	return nil
}
