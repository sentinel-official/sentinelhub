package migrations

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

type Migrator struct {
	cdc   codec.BinaryCodec
	lease LeaseKeeper
}

func NewMigrator(cdc codec.BinaryCodec, lease LeaseKeeper) Migrator {
	return Migrator{
		cdc:   cdc,
		lease: lease,
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	k.setParams(ctx)

	return nil
}

func (k *Migrator) setParams(ctx sdk.Context) {
	params := v1.Params{
		MaxHours:     720,
		MinHours:     1,
		StakingShare: math.LegacyMustNewDecFromStr("0.2"),
	}

	k.lease.SetParams(ctx, params)
}
