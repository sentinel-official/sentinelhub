package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/x/lease/types/v1"
)

type LeaseKeeper interface {
	SetParams(ctx sdk.Context, params v1.Params)
}
