package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	deposit "github.com/sentinel-official/sentinelhub/v13/x/deposit/services"
	lease "github.com/sentinel-official/sentinelhub/v13/x/lease/services"
	node "github.com/sentinel-official/sentinelhub/v13/x/node/services"
	plan "github.com/sentinel-official/sentinelhub/v13/x/plan/services"
	provider "github.com/sentinel-official/sentinelhub/v13/x/provider/services"
	session "github.com/sentinel-official/sentinelhub/v13/x/session/services"
	subscription "github.com/sentinel-official/sentinelhub/v13/x/subscription/services"
	"github.com/sentinel-official/sentinelhub/v13/x/vpn/keeper"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	deposit.RegisterServices(configurator, k.Deposit)
	lease.RegisterServices(configurator, k.Lease)
	node.RegisterServices(configurator, k.Node)
	plan.RegisterServices(configurator, k.Plan)
	provider.RegisterServices(configurator, k.Provider)
	session.RegisterServices(configurator, k.Session)
	subscription.RegisterServices(configurator, k.Subscription)
}
