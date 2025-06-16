package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/provider/types"
	"github.com/sentinel-official/sentinelhub/v12/x/provider/types/v2"
	"github.com/sentinel-official/sentinelhub/v12/x/provider/types/v3"
)

func (k *Keeper) HandleMsgRegisterProvider(ctx sdk.Context, msg *v3.MsgRegisterProviderRequest) (*v3.MsgRegisterProviderResponse, error) {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provAddr := base.ProvAddress(accAddr.Bytes())
	if k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorDuplicateProvider(provAddr)
	}

	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	provider := v2.Provider{
		Address:     provAddr.String(),
		Name:        msg.Name,
		Identity:    msg.Identity,
		Website:     msg.Website,
		Description: msg.Description,
		Status:      v1base.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	k.SetProvider(ctx, provider)
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

func (k *Keeper) HandleMsgUpdateProviderDetails(ctx sdk.Context, msg *v3.MsgUpdateProviderDetailsRequest) (*v3.MsgUpdateProviderDetailsResponse, error) {
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	if msg.Name != "" {
		provider.Name = msg.Name
	}
	provider.Identity = msg.Identity
	provider.Website = msg.Website
	provider.Description = msg.Description

	k.SetProvider(ctx, provider)
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

func (k *Keeper) HandleMsgUpdateProviderStatus(ctx sdk.Context, msg *v3.MsgUpdateProviderStatusRequest) (*v3.MsgUpdateProviderStatusResponse, error) {
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	if msg.Status.Equal(v1base.StatusInactive) {
		if err := k.ProviderInactivePreHook(ctx, provAddr); err != nil {
			return nil, err
		}
	}

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

	provider.Status = msg.Status
	provider.StatusAt = ctx.BlockTime()

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateStatus{
			ProvAddress: provider.Address,
			Status:      provider.Status,
		},
	)

	return &v3.MsgUpdateProviderStatusResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.SetParams(ctx, msg.Params)
	return &v3.MsgUpdateParamsResponse{}, nil
}
