package keeper

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	depositkeeper "github.com/sentinel-official/hub/v12/x/deposit/keeper"
	leasekeeper "github.com/sentinel-official/hub/v12/x/lease/keeper"
	nodekeeper "github.com/sentinel-official/hub/v12/x/node/keeper"
	plankeeper "github.com/sentinel-official/hub/v12/x/plan/keeper"
	providerkeeper "github.com/sentinel-official/hub/v12/x/provider/keeper"
	sessionkeeper "github.com/sentinel-official/hub/v12/x/session/keeper"
	subscriptionkeeper "github.com/sentinel-official/hub/v12/x/subscription/keeper"
)

type Keeper struct {
	Deposit      depositkeeper.Keeper
	Lease        leasekeeper.Keeper
	Node         nodekeeper.Keeper
	Plan         plankeeper.Keeper
	Provider     providerkeeper.Keeper
	Session      sessionkeeper.Keeper
	Subscription subscriptionkeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, accountKeeper AccountKeeper,
	bankKeeper BankKeeper, distributionKeeper DistributionKeeper, oracleKeeper OracleKeeper,
	router *baseapp.MsgServiceRouter, authority, feeCollectorName string,
) Keeper {
	k := Keeper{
		Deposit:      depositkeeper.NewKeeper(cdc, key),
		Lease:        leasekeeper.NewKeeper(cdc, key, router, authority, feeCollectorName),
		Node:         nodekeeper.NewKeeper(cdc, key, router, authority, feeCollectorName),
		Plan:         plankeeper.NewKeeper(cdc, key, router),
		Provider:     providerkeeper.NewKeeper(cdc, key, router, authority),
		Session:      sessionkeeper.NewKeeper(cdc, key, router, authority, feeCollectorName),
		Subscription: subscriptionkeeper.NewKeeper(cdc, key, router, authority, feeCollectorName),
	}

	k.Deposit.WithBankKeeper(bankKeeper)

	k.Lease.WithDepositKeeper(&k.Deposit)
	k.Lease.WithNodeKeeper(&k.Node)
	k.Lease.WithPlanKeeper(&k.Plan)
	k.Lease.WithOracleKeeper(oracleKeeper)
	k.Lease.WithProviderKeeper(&k.Provider)

	k.Node.WithDepositKeeper(&k.Deposit)
	k.Node.WithDistributionKeeper(distributionKeeper)
	k.Node.WithLeaseKeeper(&k.Lease)
	k.Node.WithOracleKeeper(oracleKeeper)
	k.Node.WithSessionKeeper(&k.Session)

	k.Plan.WithLeaseKeeper(&k.Lease)
	k.Plan.WithNodeKeeper(&k.Node)
	k.Plan.WithProviderKeeper(&k.Provider)
	k.Plan.WithSessionKeeper(&k.Session)
	k.Plan.WithSubscriptionKeeper(&k.Subscription)

	k.Provider.WithDistributionKeeper(distributionKeeper)
	k.Provider.WithPlanKeeper(&k.Plan)
	k.Provider.WithLeaseKeeper(&k.Lease)

	k.Session.WithAccountKeeper(accountKeeper)
	k.Session.WithDepositKeeper(&k.Deposit)
	k.Session.WithNodeKeeper(&k.Node)
	k.Session.WithSubscriptionKeeper(&k.Subscription)

	k.Subscription.WithBankKeeper(bankKeeper)
	k.Subscription.WithNodeKeeper(&k.Node)
	k.Subscription.WithOracleKeeper(oracleKeeper)
	k.Subscription.WithPlanKeeper(&k.Plan)
	k.Subscription.WithSessionKeeper(&k.Session)

	return k
}
