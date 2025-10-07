package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types"
	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
)

// HandleMsgCreateAsset handles the creation of a new asset by validating authority and uniqueness,
// initializing it with zero price and zero height, then storing it.
func (k *Keeper) HandleMsgCreateAsset(ctx sdk.Context, msg *v1.MsgCreateAssetRequest) (*v1.MsgCreateAssetResponse, error) {
	// Reject if message signer is not the module authority
	if msg.From != k.authority {
		return nil, types.NewErrorInvalidSigner(msg.From, k.authority)
	}

	// Prevent creation if an asset with the same denom already exists
	if k.HasAsset(ctx, msg.Denom) {
		return nil, types.NewErrorDuplicateAsset(msg.Denom)
	}

	// Initialize asset with provided metadata, zero price and zero height
	asset := v1.Asset{
		Denom:               msg.Denom,
		Decimals:            msg.Decimals,
		ProtoRevPoolRequest: msg.ProtoRevPoolRequest,
		SpotPriceRequest:    msg.SpotPriceRequest,
		Height:              0,
		SpotPrice:           sdkmath.LegacyZeroDec(),
	}

	k.SetAsset(ctx, asset)

	// Emit event to signal asset creation
	ctx.EventManager().EmitTypedEvent(
		&v1.EventCreate{
			Denom:           asset.Denom,
			Decimals:        asset.Decimals,
			BaseAssetDenom:  asset.SpotPriceRequest.BaseAssetDenom,
			QuoteAssetDenom: asset.SpotPriceRequest.QuoteAssetDenom,
		},
	)

	return &v1.MsgCreateAssetResponse{}, nil
}

// HandleMsgDeleteAsset handles deletion of an existing asset by verifying authority and existence,
// then removing it from the store.
func (k *Keeper) HandleMsgDeleteAsset(ctx sdk.Context, msg *v1.MsgDeleteAssetRequest) (*v1.MsgDeleteAssetResponse, error) {
	// Reject if message signer is not the module authority
	if msg.From != k.authority {
		return nil, types.NewErrorInvalidSigner(msg.From, k.authority)
	}

	// Reject deletion if the asset does not exist
	if !k.HasAsset(ctx, msg.Denom) {
		return nil, types.NewErrorAssetNotFound(msg.Denom)
	}

	k.DeleteAsset(ctx, msg.Denom)

	// Emit event to signal asset deletion
	ctx.EventManager().EmitTypedEvent(
		&v1.EventDelete{
			Denom: msg.Denom,
		},
	)

	return &v1.MsgDeleteAssetResponse{}, nil
}

// HandleMsgUpdateAsset handles metadata updates for an existing asset by validating authority and existence,
// then applying the changes.
func (k *Keeper) HandleMsgUpdateAsset(ctx sdk.Context, msg *v1.MsgUpdateAssetRequest) (*v1.MsgUpdateAssetResponse, error) {
	// Reject if message signer is not the module authority
	if msg.From != k.authority {
		return nil, types.NewErrorInvalidSigner(msg.From, k.authority)
	}

	// Reject update if the asset does not exist
	asset, found := k.GetAsset(ctx, msg.Denom)
	if !found {
		return nil, types.NewErrorAssetNotFound(msg.Denom)
	}

	// Apply updated metadata to the asset
	asset.Decimals = msg.Decimals
	asset.ProtoRevPoolRequest = msg.ProtoRevPoolRequest
	asset.SpotPriceRequest = msg.SpotPriceRequest

	k.SetAsset(ctx, asset)

	// Emit event to signal asset update
	ctx.EventManager().EmitTypedEvent(
		&v1.EventUpdateDetails{
			Denom:           asset.Denom,
			Decimals:        asset.Decimals,
			BaseAssetDenom:  asset.SpotPriceRequest.BaseAssetDenom,
			QuoteAssetDenom: asset.SpotPriceRequest.QuoteAssetDenom,
		},
	)

	return &v1.MsgUpdateAssetResponse{}, nil
}

// HandleMsgUpdateParams updates module parameters after verifying the message authority.
func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	// Reject if message signer is not the module authority
	if msg.From != k.authority {
		return nil, types.NewErrorInvalidSigner(msg.From, k.authority)
	}

	k.SetParams(ctx, msg.Params)

	return &v1.MsgUpdateParamsResponse{}, nil
}
