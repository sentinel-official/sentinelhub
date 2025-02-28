package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

// SetSession stores a session in the module's KVStore.
func (k *Keeper) SetSession(ctx sdk.Context, session v3.Session) {
	store := k.Store(ctx)
	key := types.SessionKey(session.GetID())

	value, err := k.cdc.MarshalInterface(session)
	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}

// GetSession retrieves a session from the module's KVStore based on the session ID.
// If the session exists, it returns the session and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetSession(ctx sdk.Context, id uint64) (session v3.Session, found bool) {
	store := k.Store(ctx)
	key := types.SessionKey(id)
	value := store.Get(key)

	if value == nil {
		return session, false
	}
	if err := k.cdc.UnmarshalInterface(value, &session); err != nil {
		panic(err)
	}
	if err := k.UpdateMaxValues(ctx, session); err != nil {
		panic(err)
	}

	return session, true
}

// DeleteSession removes a session from the module's KVStore based on the session ID.
func (k *Keeper) DeleteSession(ctx sdk.Context, id uint64) {
	store := k.Store(ctx)
	key := types.SessionKey(id)

	store.Delete(key)
}

// GetSessions retrieves all sessions stored in the module's KVStore.
func (k *Keeper) GetSessions(ctx sdk.Context) (items []v3.Session) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v3.Session
		if err := k.cdc.UnmarshalInterface(iterator.Value(), &item); err != nil {
			panic(err)
		}
		if err := k.UpdateMaxValues(ctx, item); err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}

// IterateSessions iterates over all sessions stored in the module's KVStore and calls the provided function for each session.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateSessions(ctx sdk.Context, fn func(index int, item v3.Session) (stop bool)) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v3.Session
		if err := k.cdc.UnmarshalInterface(iterator.Value(), &item); err != nil {
			panic(err)
		}
		if err := k.UpdateMaxValues(ctx, item); err != nil {
			panic(err)
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetSessionForAccount links a session ID to an account address in the module's KVStore.
func (k *Keeper) SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SessionForAccountKey(addr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSessionForAccount removes the association between a session ID and an account address from the module's KVStore.
func (k *Keeper) DeleteSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SessionForAccountKey(addr, id)

	store.Delete(key)
}

// GetSessionsForAccount retrieves all sessions associated with a specific account address.
func (k *Keeper) GetSessionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items []v3.Session) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetSessionForAccountKeyPrefix(addr))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForAccountKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for account key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

// SetSessionForNode links a session ID to a node address in the module's KVStore.
func (k *Keeper) SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SessionForNodeKey(addr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSessionForNode removes the association between a session ID and a node address from the module's KVStore.
func (k *Keeper) DeleteSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SessionForNodeKey(addr, id)

	store.Delete(key)
}

// GetSessionsForNode retrieves all sessions associated with a specific node address.
func (k *Keeper) GetSessionsForNode(ctx sdk.Context, addr base.NodeAddress) (items []v3.Session) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetSessionForNodeKeyPrefix(addr))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for node key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

// IterateSessionsForNode iterates over all sessions for a specific node and calls the provided function for each session.
// The iteration stops when the provided function returns 'true' or an error occurs.
func (k *Keeper) IterateSessionsForNode(ctx sdk.Context, addr base.NodeAddress, fn func(int, v3.Session) (bool, error)) error {
	store := k.Store(ctx)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForNodeKeyPrefix(addr))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for node key %X does not exist", iterator.Key()))
		}

		stop, err := fn(i, item)
		if err != nil {
			return err
		}
		if stop {
			break
		}

		i++
	}

	return nil
}

// SetSessionForPlanByNode links a session ID to a plan ID and node address in the module's KVStore.
func (k *Keeper) SetSessionForPlanByNode(ctx sdk.Context, planID uint64, addr base.NodeAddress, sessionID uint64) {
	store := k.Store(ctx)
	key := types.SessionForPlanByNodeKey(planID, addr, sessionID)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSessionForPlanByNode removes the association between a session ID, a plan ID, and a node address from the module's KVStore.
func (k *Keeper) DeleteSessionForPlanByNode(ctx sdk.Context, planID uint64, addr base.NodeAddress, sessionID uint64) {
	store := k.Store(ctx)
	key := types.SessionForPlanByNodeKey(planID, addr, sessionID)

	store.Delete(key)
}

// IterateSessionsForPlanByNode iterates over all sessions for a specific plan ID and node address, calling the provided function for each session.
// The iteration stops when the provided function returns 'true' or an error occurs.
func (k *Keeper) IterateSessionsForPlanByNode(ctx sdk.Context, id uint64, addr base.NodeAddress, fn func(int, v3.Session) (bool, error)) error {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetSessionForPlanByNodeKeyPrefix(id, addr))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForPlanByNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for plan by node key %X does not exist", iterator.Key()))
		}

		stop, err := fn(i, item)
		if err != nil {
			return err
		}
		if stop {
			break
		}

		i++
	}

	return nil
}

// SetSessionForSubscription links a session ID to a subscription ID in the module's KVStore.
func (k *Keeper) SetSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	store := k.Store(ctx)
	key := types.SessionForSubscriptionKey(subscriptionID, sessionID)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSessionForSubscription removes the association between a session ID and a subscription ID from the module's KVStore.
func (k *Keeper) DeleteSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	store := k.Store(ctx)
	key := types.SessionForSubscriptionKey(subscriptionID, sessionID)

	store.Delete(key)
}

// GetSessionsForSubscription retrieves all sessions associated with a specific subscription ID.
func (k *Keeper) GetSessionsForSubscription(ctx sdk.Context, id uint64) (items []v3.Session) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetSessionForSubscriptionKeyPrefix(id))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForSubscriptionKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

// IterateSessionsForSubscription iterates over all sessions associated with a specific subscription ID and calls the provided function for each session.
// The iteration stops when the provided function returns 'true' or an error occurs.
func (k *Keeper) IterateSessionsForSubscription(ctx sdk.Context, id uint64, fn func(int, v3.Session) (bool, error)) error {
	store := k.Store(ctx)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForSubscriptionKeyPrefix(id))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForSubscriptionKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription key %X does not exist", iterator.Key()))
		}

		stop, err := fn(i, item)
		if err != nil {
			return err
		}
		if stop {
			break
		}

		i++
	}

	return nil
}

// SetSessionForAllocation links a session ID to a subscription ID and an account address in the module's KVStore.
func (k *Keeper) SetSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	store := k.Store(ctx)
	key := types.SessionForAllocationKey(subscriptionID, addr, sessionID)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSessionForAllocation removes the association between a session ID, a subscription ID, and an account address from the module's KVStore.
func (k *Keeper) DeleteSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	store := k.Store(ctx)
	key := types.SessionForAllocationKey(subscriptionID, addr, sessionID)

	store.Delete(key)
}

// IterateSessionsForAllocation iterates over all sessions associated with a specific subscription ID and account address and calls the provided function for each session.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateSessionsForAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress, fn func(index int, item v3.Session) (stop bool)) {
	store := k.Store(ctx)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForAllocationKeyPrefix(id, addr))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForAllocationKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription allocation key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetSessionForInactiveAt sets a session to be inactive at a specified time in the module's KVStore.
func (k *Keeper) SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	store := k.Store(ctx)
	key := types.SessionForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSessionForInactiveAt removes the inactive session record from the module's KVStore based on the specified time and session ID.
func (k *Keeper) DeleteSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	store := k.Store(ctx)
	key := types.SessionForInactiveAtKey(at, id)

	store.Delete(key)
}

// IterateSessionsForInactiveAt iterates over all sessions that will be inactive before a specified time and calls the provided function for each session.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateSessionsForInactiveAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Session) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.SessionForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetSessionForInactiveAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// GetLatestSessionForAllocation retrieves the latest session for a given subscription ID and account address.
func (k *Keeper) GetLatestSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress) (session v3.Session, found bool) {
	store := k.Store(ctx)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForAllocationKeyPrefix(subscriptionID, addr))

	defer iterator.Close()

	if iterator.Valid() {
		session, found = k.GetSession(ctx, types.IDFromSessionForAllocationKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for allocation key %X does not exist", iterator.Key()))
		}
	}

	return session, found
}
