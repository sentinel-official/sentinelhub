package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/swap/types/v1"
)

func (k *Keeper) SwapEnabled(ctx sdk.Context) (yes bool) {
	k.params.Get(ctx, v1.KeySwapEnabled, &yes)

	return
}

func (k *Keeper) SwapDenom(ctx sdk.Context) (denom string) {
	k.params.Get(ctx, v1.KeySwapDenom, &denom)

	return
}

func (k *Keeper) ApproveBy(ctx sdk.Context) (address string) {
	k.params.Get(ctx, v1.KeyApproveBy, &address)

	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params v1.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) v1.Params {
	return v1.NewParams(
		k.SwapEnabled(ctx),
		k.SwapDenom(ctx),
		k.ApproveBy(ctx),
	)
}
