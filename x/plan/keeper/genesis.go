package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	"github.com/sentinel-official/sentinelhub/v13/x/plan/types/v3"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v3.GenesisState) {
	for _, item := range state.Plans {
		addr, err := base.ProvAddressFromBech32(item.Plan.ProvAddress)
		if err != nil {
			panic(err)
		}

		k.SetPlan(ctx, item.Plan)
		k.SetPlanForProvider(ctx, addr, item.Plan.ID)

		for _, node := range item.Nodes {
			addr, err := base.NodeAddressFromBech32(node)
			if err != nil {
				panic(err)
			}

			k.SetNodeForPlan(ctx, item.Plan.ID, addr)
		}
	}

	count := uint64(0)
	for _, item := range state.Plans {
		if item.Plan.ID > count {
			count = item.Plan.ID
		}
	}

	k.SetPlanCount(ctx, count)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v3.GenesisState {
	var (
		plans = k.GetPlans(ctx)
		items = make(v3.GenesisPlans, 0, len(plans))
	)

	for _, plan := range plans {
		item := v3.GenesisPlan{
			Plan:  plan,
			Nodes: []string{},
		}

		nodes := k.GetNodesForPlan(ctx, plan.ID)
		for _, node := range nodes {
			item.Nodes = append(item.Nodes, node.Address)
		}

		items = append(items, item)
	}

	return v3.NewGenesisState(items)
}
