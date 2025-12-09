package keeper

import (
	"fmt"
	"time"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/cosmos/gogoproto/types"

	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types"
	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types/v3"
)

// SetSubscription stores a subscription in the module's KVStore.
func (k *Keeper) SetSubscription(ctx sdk.Context, subscription v3.Subscription) {
	store := k.Store(ctx)
	key := types.SubscriptionKey(subscription.ID)
	value := k.cdc.MustMarshal(&subscription)

	store.Set(key, value)
}

// GetSubscription retrieves a subscription from the module's KVStore based on the subscription ID.
// Returns the subscription and a boolean indicating if it was found.
func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription v3.Subscription, found bool) {
	store := k.Store(ctx)
	key := types.SubscriptionKey(id)
	value := store.Get(key)

	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshal(value, &subscription)

	return subscription, true
}

// DeleteSubscription removes a subscription from the module's KVStore based on the subscription ID.
func (k *Keeper) DeleteSubscription(ctx sdk.Context, id uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionKey(id)

	store.Delete(key)
}

// GetSubscriptions retrieves all subscriptions from the module's KVStore.
func (k *Keeper) GetSubscriptions(ctx sdk.Context) (items []v3.Subscription) {
	store := k.Store(ctx)
	iterator := storetypes.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v3.Subscription
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

// IterateSubscriptions iterates over all subscriptions in the module's KVStore and calls the provided function for each subscription.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateSubscriptions(ctx sdk.Context, fn func(index int, item v3.Subscription) (stop bool)) {
	store := k.Store(ctx)
	iterator := storetypes.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v3.Subscription
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}

		i++
	}
}

// SetSubscriptionForAccount links a subscription ID to an account address in the module's KVStore.
func (k *Keeper) SetSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForAccountKey(addr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// HasSubscriptionForAccount checks if there is a subscription ID associated with a given account address.
func (k *Keeper) HasSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	store := k.Store(ctx)
	key := types.SubscriptionForAccountKey(addr, id)

	return store.Has(key)
}

// DeleteSubscriptionForAccount removes the association between a subscription ID and an account address from the module's KVStore.
func (k *Keeper) DeleteSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForAccountKey(addr, id)

	store.Delete(key)
}

// GetSubscriptionsForAccount retrieves all subscriptions associated with a specific account address.
func (k *Keeper) GetSubscriptionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items []v3.Subscription) {
	store := k.Store(ctx)
	iterator := storetypes.KVStorePrefixIterator(store, types.GetSubscriptionForAccountKeyPrefix(addr))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForAccountKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for account key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

// SetSubscriptionForPlan links a subscription ID to a plan ID in the module's KVStore.
func (k *Keeper) SetSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForPlanKey(planID, subscriptionID)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// HasSubscriptionForPlan checks if there is a subscription ID associated with a given plan ID.
func (k *Keeper) HasSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) bool {
	store := k.Store(ctx)
	key := types.SubscriptionForPlanKey(planID, subscriptionID)

	return store.Has(key)
}

// DeleteSubscriptionForPlan removes the association between a subscription ID and a plan ID from the module's KVStore.
func (k *Keeper) DeleteSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForPlanKey(planID, subscriptionID)

	store.Delete(key)
}

// GetSubscriptionsForPlan retrieves all subscriptions associated with a specific plan ID.
func (k *Keeper) GetSubscriptionsForPlan(ctx sdk.Context, id uint64) (items []v3.Subscription) {
	store := k.Store(ctx)
	iterator := storetypes.KVStorePrefixIterator(store, types.GetSubscriptionForPlanKeyPrefix(id))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForPlanKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for plan key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

// SetSubscriptionForInactiveAt sets a subscription to be inactive at a specified time in the module's KVStore.
func (k *Keeper) SetSubscriptionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSubscriptionForInactiveAt removes the inactive subscription record from the module's KVStore based on the specified time and subscription ID.
func (k *Keeper) DeleteSubscriptionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForInactiveAtKey(at, id)

	store.Delete(key)
}

// IterateSubscriptionsForInactiveAt iterates over all subscriptions that will be inactive before a specified time and calls the provided function for each subscription.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateSubscriptionsForInactiveAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Subscription) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.SubscriptionForInactiveAtKeyPrefix, storetypes.PrefixEndBytes(types.GetSubscriptionForInactiveAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}

		i++
	}
}

// SetSubscriptionForRenewalAt sets a subscription to be renewed at a specified time in the module's KVStore.
func (k *Keeper) SetSubscriptionForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForRenewalAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteSubscriptionForRenewalAt removes the renewal subscription record from the module's KVStore based on the specified time and subscription ID.
func (k *Keeper) DeleteSubscriptionForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForRenewalAtKey(at, id)

	store.Delete(key)
}

// IterateSubscriptionsForRenewalAt iterates over all subscriptions that will be renewed before a specified time and calls the provided function for each subscription.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateSubscriptionsForRenewalAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Subscription) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.SubscriptionForRenewalAtKeyPrefix, storetypes.PrefixEndBytes(types.GetSubscriptionForRenewalAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForRenewalAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for renewal at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}

		i++
	}
}
