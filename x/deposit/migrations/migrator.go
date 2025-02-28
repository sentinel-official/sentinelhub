package migrations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/deposit/types"
	"github.com/sentinel-official/hub/v12/x/deposit/types/v1"
)

type Migrator struct {
	cdc     codec.BinaryCodec
	deposit DepositKeeper
}

func NewMigrator(cdc codec.BinaryCodec, deposit DepositKeeper) Migrator {
	return Migrator{
		cdc:     cdc,
		deposit: deposit,
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	if err := k.resetDeposits(ctx); err != nil {
		return err
	}

	return nil
}

func (k *Migrator) resetDeposits(ctx sdk.Context) error {
	store := prefix.NewStore(k.deposit.Store(ctx), types.DepositKeyPrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var item v1.Deposit
		k.cdc.MustUnmarshal(it.Value(), &item)

		k.deposit.SetDeposit(ctx, item)
	}

	return nil
}
