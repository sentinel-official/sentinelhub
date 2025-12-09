package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/mint/types/v1"
)

func (k *Keeper) BeginBlock(c context.Context) error {
	ctx := sdk.UnwrapSDKContext(c)

	k.IterateInflations(ctx, func(_ int, item v1.Inflation) bool {
		if item.Timestamp.After(ctx.BlockTime()) {
			return true
		}

		params := k.GetParams(ctx)
		params.InflationMax = item.Max
		params.InflationMin = item.Min

		params.InflationRateChange = item.RateChange
		if err := k.SetParams(ctx, params); err != nil {
			panic(err)
		}

		minter := k.GetMinter(ctx)
		minter.Inflation = item.Min
		k.SetMinter(ctx, minter)

		k.DeleteInflation(ctx, item.Timestamp)

		return false
	})

	return nil
}
