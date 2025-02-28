package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

type SessionKeeper interface {
	SetParams(ctx sdk.Context, params v3.Params)
	Store(ctx sdk.Context) sdk.KVStore
}
