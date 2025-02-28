package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/deposit/types/v1"
)

type DepositKeeper interface {
	SetDeposit(ctx sdk.Context, deposit v1.Deposit)
	Store(ctx sdk.Context) sdk.KVStore
}
