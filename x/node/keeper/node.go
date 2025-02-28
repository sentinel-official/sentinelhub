package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

// SetActiveNode stores an active node in the module's KVStore.
func (k *Keeper) SetActiveNode(ctx sdk.Context, node v3.Node) {
	addr, err := base.NodeAddressFromBech32(node.Address)
	if err != nil {
		panic(err)
	}

	store := k.Store(ctx)
	key := types.ActiveNodeKey(addr)
	value := k.cdc.MustMarshal(&node)

	store.Set(key, value)
}

// HasActiveNode checks if an active node exists in the module's KVStore based on the node address.
func (k *Keeper) HasActiveNode(ctx sdk.Context, addr base.NodeAddress) bool {
	store := k.Store(ctx)
	key := types.ActiveNodeKey(addr)

	return store.Has(key)
}

// GetActiveNode retrieves an active node from the module's KVStore based on the node address.
// If the active node exists, it returns the node and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetActiveNode(ctx sdk.Context, addr base.NodeAddress) (v v3.Node, found bool) {
	store := k.Store(ctx)
	key := types.ActiveNodeKey(addr)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// DeleteActiveNode removes an active node from the module's KVStore based on the node address.
func (k *Keeper) DeleteActiveNode(ctx sdk.Context, addr base.NodeAddress) {
	store := k.Store(ctx)
	key := types.ActiveNodeKey(addr)

	store.Delete(key)
}

// SetInactiveNode stores an inactive node in the module's KVStore.
func (k *Keeper) SetInactiveNode(ctx sdk.Context, node v3.Node) {
	addr, err := base.NodeAddressFromBech32(node.Address)
	if err != nil {
		panic(err)
	}

	store := k.Store(ctx)
	key := types.InactiveNodeKey(addr)
	value := k.cdc.MustMarshal(&node)

	store.Set(key, value)
}

// HasInactiveNode checks if an inactive node exists in the module's KVStore based on the node address.
func (k *Keeper) HasInactiveNode(ctx sdk.Context, addr base.NodeAddress) bool {
	store := k.Store(ctx)
	key := types.InactiveNodeKey(addr)

	return store.Has(key)
}

// GetInactiveNode retrieves an inactive node from the module's KVStore based on the node address.
// If the inactive node exists, it returns the node and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetInactiveNode(ctx sdk.Context, addr base.NodeAddress) (v v3.Node, found bool) {
	store := k.Store(ctx)
	key := types.InactiveNodeKey(addr)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// DeleteInactiveNode removes an inactive node from the module's KVStore based on the node address.
func (k *Keeper) DeleteInactiveNode(ctx sdk.Context, addr base.NodeAddress) {
	store := k.Store(ctx)
	key := types.InactiveNodeKey(addr)

	store.Delete(key)
}

// SetNode stores a node in the module's KVStore based on its status (active or inactive).
func (k *Keeper) SetNode(ctx sdk.Context, node v3.Node) {
	switch node.Status {
	case v1base.StatusActive:
		k.SetActiveNode(ctx, node)
	case v1base.StatusInactive:
		k.SetInactiveNode(ctx, node)
	default:
		panic(fmt.Errorf("failed to set the node %v", node))
	}
}

// HasNode checks if a node exists in the module's KVStore based on the node address.
func (k *Keeper) HasNode(ctx sdk.Context, addr base.NodeAddress) bool {
	return k.HasActiveNode(ctx, addr) || k.HasInactiveNode(ctx, addr)
}

// GetNode retrieves a node from the module's KVStore based on the node address.
// It checks both active and inactive nodes and returns the node if found.
func (k *Keeper) GetNode(ctx sdk.Context, addr base.NodeAddress) (node v3.Node, found bool) {
	node, found = k.GetActiveNode(ctx, addr)
	if found {
		return
	}

	node, found = k.GetInactiveNode(ctx, addr)
	if found {
		return
	}

	return node, false
}

// GetNodes retrieves all nodes stored in the module's KVStore.
func (k *Keeper) GetNodes(ctx sdk.Context) (items []v3.Node) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v3.Node
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

// IterateNodes iterates over all nodes stored in the module's KVStore and calls the provided function for each node.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateNodes(ctx sdk.Context, fn func(index int, item v3.Node) (stop bool)) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var node v3.Node
		k.cdc.MustUnmarshal(iterator.Value(), &node)

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}

// SetNodeForInactiveAt stores a node's inactivity timestamp in the module's KVStore.
func (k *Keeper) SetNodeForInactiveAt(ctx sdk.Context, at time.Time, addr base.NodeAddress) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.NodeForInactiveAtKey(at, addr)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteNodeForInactiveAt removes a node's inactivity timestamp from the module's KVStore.
func (k *Keeper) DeleteNodeForInactiveAt(ctx sdk.Context, at time.Time, addr base.NodeAddress) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.NodeForInactiveAtKey(at, addr)

	store.Delete(key)
}

// IterateNodesForInactiveAt iterates over all nodes with inactivity timestamps stored in the module's KVStore and calls the provided function for each node.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateNodesForInactiveAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Node) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.NodeForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetNodeForInactiveAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		node, found := k.GetNode(ctx, types.AddressFromNodeForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("node for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}

// SetNodeForPlan stores a node associated with a plan in the module's KVStore.
func (k *Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	store := k.Store(ctx)
	key := types.NodeForPlanKey(id, addr)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// HasNodeForPlan checks if a node associated with a plan exists in the module's KVStore based on the plan ID and node address.
func (k *Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) bool {
	store := k.Store(ctx)
	key := types.NodeForPlanKey(id, addr)

	return store.Has(key)
}

// DeleteNodeForPlan removes a node associated with a plan from the module's KVStore based on the plan ID and node address.
func (k *Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	store := k.Store(ctx)
	key := types.NodeForPlanKey(id, addr)

	store.Delete(key)
}

// GetNodesForPlan retrieves all nodes associated with a plan stored in the module's KVStore based on the plan ID.
func (k *Keeper) GetNodesForPlan(ctx sdk.Context, id uint64) (items []v3.Node) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetNodeForPlanKeyPrefix(id))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetNode(ctx, types.AddressFromNodeForPlanKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("node for plan key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}
