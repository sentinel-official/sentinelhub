package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	leasetypes "github.com/sentinel-official/sentinelhub/v12/x/lease/types/v1"
	nodetypes "github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
)

func (k *Keeper) GetLease(ctx sdk.Context, id uint64) (leasetypes.Lease, bool) {
	return k.lease.GetLease(ctx, id)
}

func (k *Keeper) IterateLeasesForNodeByProvider(ctx sdk.Context, nodeAddr base.NodeAddress, provAddr base.ProvAddress, fn func(index int, item leasetypes.Lease) (stop bool)) {
	k.lease.IterateLeasesForNodeByProvider(ctx, nodeAddr, provAddr, fn)
}

func (k *Keeper) GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, addr)
}

func (k *Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	k.node.SetNodeForPlan(ctx, id, addr)
}

func (k *Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	k.node.DeleteNodeForPlan(ctx, id, addr)
}

func (k *Keeper) GetNodesForPlan(ctx sdk.Context, id uint64) []nodetypes.Node {
	return k.node.GetNodesForPlan(ctx, id)
}

func (k *Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) bool {
	return k.node.HasNodeForPlan(ctx, id, addr)
}

func (k *Keeper) HasProvider(ctx sdk.Context, addr base.ProvAddress) bool {
	return k.provider.HasProvider(ctx, addr)
}

func (k *Keeper) PlanUnlinkNodePreHook(ctx sdk.Context, id uint64, addr base.NodeAddress) error {
	if err := k.session.PlanUnlinkNodePreHook(ctx, id, addr); err != nil {
		return err
	}

	return nil
}
