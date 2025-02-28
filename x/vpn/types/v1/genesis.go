package v1

import (
	sdkerrors "cosmossdk.io/errors"

	deposittypes "github.com/sentinel-official/hub/v12/x/deposit/types/v1"
	leasetypes "github.com/sentinel-official/hub/v12/x/lease/types/v1"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v3"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v3"
	providertypes "github.com/sentinel-official/hub/v12/x/provider/types/v3"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v3"
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
		return sdkerrors.Wrapf(err, "invalid deposit genesis state")
	}
	if err := leasetypes.ValidateGenesis(m.Lease); err != nil {
		return sdkerrors.Wrapf(err, "invalid lease genesis state")
	}
	if err := nodetypes.ValidateGenesis(m.Node); err != nil {
		return sdkerrors.Wrapf(err, "invalid node genesis state")
	}
	if err := plantypes.ValidateGenesis(m.Plan); err != nil {
		return sdkerrors.Wrapf(err, "invalid plan genesis state")
	}
	if err := providertypes.ValidateGenesis(m.Provider); err != nil {
		return sdkerrors.Wrapf(err, "invalid provider genesis state")
	}
	if err := sessiontypes.ValidateGenesis(m.Session); err != nil {
		return sdkerrors.Wrapf(err, "invalid session genesis state")
	}
	if err := subscriptiontypes.ValidateGenesis(m.Subscription); err != nil {
		return sdkerrors.Wrapf(err, "invalid subscription genesis state")
	}

	return nil
}
