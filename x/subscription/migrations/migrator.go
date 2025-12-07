package migrations

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v13/types"
	v1base "github.com/sentinel-official/sentinelhub/v13/types/v1"
	"github.com/sentinel-official/sentinelhub/v13/x/lease/types/v1"
	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types/v2"
	"github.com/sentinel-official/sentinelhub/v13/x/subscription/types/v3"
)

type Migrator struct {
	cdc          codec.BinaryCodec
	deposit      DepositKeeper
	lease        LeaseKeeper
	node         NodeKeeper
	plan         PlanKeeper
	provider     ProviderKeeper
	subscription SubscriptionKeeper
}

func NewMigrator(
	cdc codec.BinaryCodec, deposit DepositKeeper, lease LeaseKeeper, node NodeKeeper, plan PlanKeeper,
	provider ProviderKeeper, subscription SubscriptionKeeper,
) Migrator {
	return Migrator{
		cdc:          cdc,
		deposit:      deposit,
		lease:        lease,
		node:         node,
		plan:         plan,
		provider:     provider,
		subscription: subscription,
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	k.setParams(ctx)

	_ = k.deleteKeys(ctx, []byte{0x11}) // SubscriptionForInactiveAtKeyPrefix
	_ = k.deleteKeys(ctx, []byte{0x12}) // SubscriptionForAccountKeyPrefix
	_ = k.deleteKeys(ctx, []byte{0x13}) // SubscriptionForNodeKeyPrefix
	_ = k.deleteKeys(ctx, []byte{0x14}) // SubscriptionForPlanKeyPrefix

	_ = k.deleteKeys(ctx, []byte{0x31}) // PayoutForNextAtKeyPrefix
	_ = k.deleteKeys(ctx, []byte{0x32}) // PayoutForAccountKeyPrefix
	_ = k.deleteKeys(ctx, []byte{0x33}) // PayoutForNodeKeyPrefix
	_ = k.deleteKeys(ctx, []byte{0x34}) // PayoutForAccountByNodeKeyPrefix

	k.migrateSubscriptions(ctx)

	_ = k.deleteKeys(ctx, []byte{0x30}) // PayoutKeyPrefix

	return nil
}

func (k *Migrator) deleteKeys(ctx sdk.Context, keyPrefix []byte) (keys [][]byte) {
	store := prefix.NewStore(k.subscription.Store(ctx), keyPrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())

		keys = append(keys, it.Key())
	}

	return keys
}

func (k *Migrator) setParams(ctx sdk.Context) {
	params := v3.Params{
		MaxAllocations: 8,
		StakingShare:   sdkmath.LegacyMustNewDecFromStr("0.2"),
		StatusTimeout:  4 * time.Hour,
	}

	k.subscription.SetParams(ctx, params)
}

func (k *Migrator) getPayout(ctx sdk.Context, id uint64) v2.Payout {
	key := append([]byte{0x30}, sdk.Uint64ToBigEndian(id)...)
	value := k.subscription.Store(ctx).Get(key)

	if value == nil {
		panic(fmt.Errorf("payout %d does not exist", id))
	}

	var payout v2.Payout
	k.cdc.MustUnmarshal(value, &payout)

	return payout
}

func (k *Migrator) migrateSubscriptions(ctx sdk.Context) {
	store := prefix.NewStore(k.subscription.Store(ctx), []byte{0x10})

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())

		var item v2.Subscription
		if err := k.cdc.UnmarshalInterface(it.Value(), &item); err != nil {
			panic(err)
		}

		accAddr := item.GetAddress()

		if item, ok := item.(*v2.NodeSubscription); ok {
			if item.Gigabytes != 0 {
				alloc, found := k.subscription.GetAllocation(ctx, item.GetID(), accAddr)
				if !found {
					panic(fmt.Errorf("subscription allocation %d/%s does not exist", item.GetID(), accAddr))
				}

				gigabytePrice := sdk.NewCoin(item.Deposit.Denom, item.Deposit.Amount.QuoRaw(item.Gigabytes))
				bytePrice := gigabytePrice.Amount.ToLegacyDec().QuoInt(types.Gigabyte)
				paidAmount := alloc.UtilisedBytes.ToLegacyDec().Mul(bytePrice).Ceil().TruncateInt()
				refund := sdk.NewCoin(item.Deposit.Denom, item.Deposit.Amount.Sub(paidAmount))

				if !refund.IsZero() {
					if err := k.deposit.SubtractDeposit(ctx, accAddr, sdk.NewCoins(refund)); err != nil {
						panic(err)
					}
				}

				k.subscription.DeleteAllocation(ctx, item.GetID(), accAddr)
			}

			if item.Hours != 0 {
				payout := k.getPayout(ctx, item.ID)
				provAddr := types.ProvAddress(accAddr.Bytes())

				if !k.provider.HasProvider(ctx, provAddr) {
					refund := sdk.NewCoin(payout.Price.Denom, payout.Price.Amount.MulRaw(payout.Hours))

					if !refund.IsZero() {
						if err := k.deposit.SubtractDeposit(ctx, accAddr, sdk.NewCoins(refund)); err != nil {
							panic(err)
						}
					}
				} else {
					nodeAddr, err := types.NodeAddressFromBech32(item.NodeAddress)
					if err != nil {
						panic(err)
					}

					refund := false

					if !item.Status.Equal(v1base.StatusActive) {
						refund = true
					}

					if k.lease.HasAnyLeaseForNodeByProvider(ctx, nodeAddr, provAddr) {
						refund = true
					}

					node, found := k.node.GetNode(ctx, nodeAddr)
					if !found {
						panic(fmt.Errorf("node %s does not exist", nodeAddr))
					}

					if !node.Status.Equal(v1base.StatusActive) {
						refund = true
					}

					provider, found := k.provider.GetProvider(ctx, provAddr)
					if !found {
						panic(fmt.Errorf("provider %s does not exist", provAddr))
					}

					if !provider.Status.Equal(v1base.StatusActive) {
						refund = true
					}

					if refund {
						amount := sdk.NewCoin(payout.Price.Denom, payout.Price.Amount.MulRaw(payout.Hours))

						if !amount.IsZero() {
							if err := k.deposit.SubtractDeposit(ctx, accAddr, sdk.NewCoins(amount)); err != nil {
								panic(err)
							}
						}
					} else {
						count := k.lease.GetLeaseCount(ctx)
						lease := v1.Lease{
							ID:                 count + 1,
							ProvAddress:        provAddr.String(),
							NodeAddress:        item.NodeAddress,
							Price:              v1base.NewPriceFromCoin(payout.Price),
							Hours:              item.Hours - payout.Hours,
							MaxHours:           item.Hours,
							RenewalPricePolicy: v1base.RenewalPricePolicyUnspecified,
							StartAt:            item.StatusAt,
						}

						k.lease.SetLeaseCount(ctx, count+1)
						k.lease.SetLease(ctx, lease)
						k.lease.SetLeaseForNodeByProvider(ctx, nodeAddr, provAddr, lease.ID)
						k.lease.SetLeaseForProvider(ctx, provAddr, lease.ID)
						k.lease.SetLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
						k.lease.SetLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
						k.lease.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)
					}
				}
			}
		}

		if item, ok := item.(*v2.PlanSubscription); ok {
			if !item.Status.Equal(v1base.StatusActive) {
				k.subscription.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) (stop bool) {
					addr, err := sdk.AccAddressFromBech32(item.Address)
					if err != nil {
						panic(err)
					}

					k.subscription.DeleteAllocation(ctx, item.ID, addr)

					return false
				})
			} else {
				plan, found := k.plan.GetPlan(ctx, item.PlanID)
				if !found {
					panic(fmt.Errorf("plan %d does not exist", item.PlanID))
				}

				price, found := plan.Price(item.Denom)
				if !found {
					panic(fmt.Errorf("denom %s for plan %d does not exist", item.Denom, item.PlanID))
				}

				subscription := v3.Subscription{
					ID:                 item.ID,
					AccAddress:         item.Address,
					PlanID:             item.PlanID,
					Price:              price,
					RenewalPricePolicy: v1base.RenewalPricePolicyUnspecified,
					Status:             item.Status,
					InactiveAt:         item.InactiveAt,
					StartAt:            item.StatusAt,
					StatusAt:           item.StatusAt,
				}

				k.subscription.SetSubscription(ctx, subscription)
				k.subscription.SetSubscriptionForAccount(ctx, accAddr, subscription.ID)
				k.subscription.SetSubscriptionForPlan(ctx, subscription.PlanID, subscription.ID)
				k.subscription.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
				k.subscription.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt(), subscription.ID)

				k.subscription.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) bool {
					accAddr, err := sdk.AccAddressFromBech32(item.Address)
					if err != nil {
						panic(err)
					}

					k.subscription.SetSubscriptionForAccount(ctx, accAddr, subscription.ID)

					return false
				})
			}
		}
	}
}
