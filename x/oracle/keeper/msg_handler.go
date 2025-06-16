package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
)

func (k *Keeper) HandleMsgCreateAsset(ctx sdk.Context, msg *v1.MsgCreateAssetRequest) (*v1.MsgCreateAssetResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	if k.HasAsset(ctx, msg.Denom) {
		return nil, types.NewErrorDuplicateAsset(msg.Denom)
	}

	asset := v1.Asset{
		Denom:           msg.Denom,
		Decimals:        msg.Decimals,
		BaseAssetDenom:  msg.BaseAssetDenom,
		QuoteAssetDenom: msg.QuoteAssetDenom,
		Price:           sdkmath.ZeroInt(),
		Height:          0,
	}

	k.SetAsset(ctx, asset)
	ctx.EventManager().EmitTypedEvent(
		&v1.EventCreate{
			Denom:           asset.Denom,
			Decimals:        asset.Decimals,
			BaseAssetDenom:  asset.BaseAssetDenom,
			QuoteAssetDenom: asset.QuoteAssetDenom,
		},
	)

	return &v1.MsgCreateAssetResponse{}, nil
}

func (k *Keeper) HandleMsgDeleteAsset(ctx sdk.Context, msg *v1.MsgDeleteAssetRequest) (*v1.MsgDeleteAssetResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	if !k.HasAsset(ctx, msg.Denom) {
		return nil, types.NewErrorAssetNotFound(msg.Denom)
	}

	k.DeleteAsset(ctx, msg.Denom)
	ctx.EventManager().EmitTypedEvent(
		&v1.EventDelete{
			Denom: msg.Denom,
		},
	)

	return &v1.MsgDeleteAssetResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateAsset(ctx sdk.Context, msg *v1.MsgUpdateAssetRequest) (*v1.MsgUpdateAssetResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	asset, found := k.GetAsset(ctx, msg.Denom)
	if !found {
		return nil, types.NewErrorAssetNotFound(msg.Denom)
	}

	asset.Decimals = msg.Decimals
	asset.BaseAssetDenom = msg.BaseAssetDenom
	asset.QuoteAssetDenom = msg.QuoteAssetDenom

	k.SetAsset(ctx, asset)
	ctx.EventManager().EmitTypedEvent(
		&v1.EventUpdate{
			Denom:           asset.Denom,
			Decimals:        asset.Decimals,
			BaseAssetDenom:  asset.BaseAssetDenom,
			QuoteAssetDenom: asset.QuoteAssetDenom,
		},
	)

	return &v1.MsgUpdateAssetResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	k.SetParams(ctx, msg.Params)
	return &v1.MsgUpdateParamsResponse{}, nil
}
