package migrations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	deposit "github.com/sentinel-official/hub/v12/x/deposit/migrations"
	lease "github.com/sentinel-official/hub/v12/x/lease/migrations"
	node "github.com/sentinel-official/hub/v12/x/node/migrations"
	plan "github.com/sentinel-official/hub/v12/x/plan/migrations"
	provider "github.com/sentinel-official/hub/v12/x/provider/migrations"
	session "github.com/sentinel-official/hub/v12/x/session/migrations"
	subscription "github.com/sentinel-official/hub/v12/x/subscription/migrations"
	"github.com/sentinel-official/hub/v12/x/vpn/keeper"
)

type Migrator struct {
	deposit      deposit.Migrator
	lease        lease.Migrator
	provider     provider.Migrator
	node         node.Migrator
	plan         plan.Migrator
	subscription subscription.Migrator
	session      session.Migrator
}

func NewMigrator(cdc codec.BinaryCodec, k keeper.Keeper) Migrator {
	return Migrator{
		deposit:      deposit.NewMigrator(cdc, &k.Deposit),
		lease:        lease.NewMigrator(cdc, &k.Lease),
		provider:     provider.NewMigrator(cdc, &k.Provider),
		node:         node.NewMigrator(cdc, &k.Node),
		plan:         plan.NewMigrator(cdc, &k.Node, &k.Plan),
		subscription: subscription.NewMigrator(cdc, &k.Deposit, &k.Lease, &k.Plan, &k.Provider, &k.Subscription),
		session:      session.NewMigrator(cdc, &k.Session),
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	if err := k.deposit.Migrate(ctx); err != nil {
		panic(err)
	}
	if err := k.lease.Migrate(ctx); err != nil {
		return err
	}
	if err := k.node.Migrate(ctx); err != nil {
		return err
	}
	if err := k.plan.Migrate(ctx); err != nil {
		return err
	}
	if err := k.provider.Migrate(ctx); err != nil {
		return err
	}
	if err := k.session.Migrate(ctx); err != nil {
		return err
	}
	if err := k.subscription.Migrate(ctx); err != nil {
		return err
	}

	return nil
}
