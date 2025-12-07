package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	v1base "github.com/sentinel-official/sentinelhub/v13/types/v1"
	"github.com/sentinel-official/sentinelhub/v13/x/provider/types"
	"github.com/sentinel-official/sentinelhub/v13/x/provider/types/v2"
	"github.com/sentinel-official/sentinelhub/v13/x/provider/types/v3"
)

// HandleMsgRegisterProvider handles a request to register a new provider.
// It validates the input, deducts the registration deposit, creates the provider, and emits a creation event.
func (k *Keeper) HandleMsgRegisterProvider(ctx sdk.Context, msg *v3.MsgRegisterProviderRequest) (*v3.MsgRegisterProviderResponse, error) {
	// Parse and validate the provider's account address
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Derive provider address and check for duplication
	provAddr := base.ProvAddress(accAddr.Bytes())
	if k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorDuplicateProvider(provAddr)
	}

	// Deduct registration deposit and transfer to the community pool
	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	// Construct a new provider record with inactive status
	provider := v2.Provider{
		Address:     provAddr.String(),
		Name:        msg.Name,
		Identity:    msg.Identity,
		Website:     msg.Website,
		Description: msg.Description,
		Status:      v1base.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	// Store the provider in state
	k.SetProvider(ctx, provider)

	// Emit event to indicate provider registration
	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreate{
			ProvAddress: provider.Address,
			Name:        provider.Name,
			Identity:    provider.Identity,
			Website:     provider.Website,
			Description: provider.Description,
		},
	)

	return &v3.MsgRegisterProviderResponse{}, nil
}

// HandleMsgUpdateProviderDetails handles a request to update metadata of an existing provider.
// It updates fields like name, identity, website, and description, then emits an update event.
func (k *Keeper) HandleMsgUpdateProviderDetails(ctx sdk.Context, msg *v3.MsgUpdateProviderDetailsRequest) (*v3.MsgUpdateProviderDetailsResponse, error) {
	// Parse and validate provider address
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Fetch the provider and ensure it exists
	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	// Apply updates to individual fields if present
	if msg.Name != "" {
		provider.Name = msg.Name
	}

	provider.Identity = msg.Identity
	provider.Website = msg.Website
	provider.Description = msg.Description

	// Persist the updated provider record
	k.SetProvider(ctx, provider)

	// Emit event indicating metadata update
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateDetails{
			ProvAddress: provider.Address,
			Name:        provider.Name,
			Identity:    provider.Identity,
			Website:     provider.Website,
			Description: provider.Description,
		},
	)

	return &v3.MsgUpdateProviderDetailsResponse{}, nil
}

// HandleMsgUpdateProviderStatus handles a request to change a provider's status.
// It applies pre-hooks, cleans up old status indices, updates the record, and emits a status update event.
func (k *Keeper) HandleMsgUpdateProviderStatus(ctx sdk.Context, msg *v3.MsgUpdateProviderStatusRequest) (*v3.MsgUpdateProviderStatusResponse, error) {
	// Parse and validate provider address
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Fetch the provider and ensure it exists
	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	// Execute pre-hook if transitioning to inactive
	if msg.Status.Equal(v1base.StatusInactive) {
		if err := k.ProviderInactivePreHook(ctx, provAddr); err != nil {
			return nil, err
		}
	}

	// Remove previous status index if changing state
	if msg.Status.Equal(v1base.StatusActive) {
		if provider.Status.Equal(v1base.StatusInactive) {
			k.DeleteInactiveProvider(ctx, provAddr)
		}
	}

	if msg.Status.Equal(v1base.StatusInactive) {
		if provider.Status.Equal(v1base.StatusActive) {
			k.DeleteActiveProvider(ctx, provAddr)
		}
	}

	// Update the status and timestamp
	provider.Status = msg.Status
	provider.StatusAt = ctx.BlockTime()

	// Store the updated provider
	k.SetProvider(ctx, provider)

	// Emit event indicating status change
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateStatus{
			ProvAddress: provider.Address,
			Status:      provider.Status.String(),
		},
	)

	return &v3.MsgUpdateProviderStatusResponse{}, nil
}

// HandleMsgUpdateParams allows the module authority to update provider module parameters.
// It checks authorization and persists the new parameter values.
func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	// Restrict access to the designated authority account
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Store the updated parameter set
	k.SetParams(ctx, msg.Params)

	return &v3.MsgUpdateParamsResponse{}, nil
}
