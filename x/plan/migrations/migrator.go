package migrations

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/types/v1"
	nodetypes "github.com/sentinel-official/sentinelhub/v12/x/node/types"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types/v2"
	"github.com/sentinel-official/sentinelhub/v12/x/plan/types/v3"
)

type Migrator struct {
	cdc   codec.BinaryCodec
	lease LeaseKeeper
	node  NodeKeeper
	plan  PlanKeeper
}

func NewMigrator(cdc codec.BinaryCodec, lease LeaseKeeper, node NodeKeeper, plan PlanKeeper) Migrator {
	return Migrator{
		cdc:   cdc,
		lease: lease,
		node:  node,
		plan:  plan,
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	k.migratePlans(ctx)

	planForProviderKeys := k.deleteKeys(ctx, []byte{0x11})

	k.setPlanForNodeByProviderKeys(ctx)
	k.setPlanForProviderKeys(ctx, planForProviderKeys...)

	return nil
}

func (k *Migrator) deleteKeys(ctx sdk.Context, keyPrefix []byte) (keys [][]byte) {
	store := prefix.NewStore(k.plan.Store(ctx), keyPrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())

		keys = append(keys, it.Key())
	}

	return keys
}

func (k *Migrator) migratePlans(ctx sdk.Context) {
	store := prefix.NewStore(k.plan.Store(ctx), []byte{0x10})

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())

		var item v2.Plan
		k.cdc.MustUnmarshal(it.Value(), &item)

		prices, err := v1.NewPricesFromCoins(item.Prices...)
		if err != nil {
			panic(err)
		}

		plan := v3.Plan{
			ID:          item.ID,
			ProvAddress: item.ProviderAddress,
			Bytes:       base.Gigabyte.MulRaw(item.Gigabytes),
			Duration:    item.Duration,
			Prices:      prices,
			Private:     false,
			Status:      item.Status,
			StatusAt:    item.StatusAt,
		}

		k.plan.SetPlan(ctx, plan)
	}
}

func (k *Migrator) setPlanForNodeByProviderKeys(ctx sdk.Context) {
	store := prefix.NewStore(k.node.Store(ctx), nodetypes.NodeForPlanKeyPrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		key := it.Key()
		id := sdk.BigEndianToUint64(key[:8])
		addrLen := int(key[8])
		nodeAddr := base.NodeAddress(key[9 : 9+addrLen])

		plan, found := k.plan.GetPlan(ctx, id)
		if !found {
			panic(fmt.Errorf("plan %d not found", id))
		}

		provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
		if err != nil {
			panic(err)
		}

		k.plan.SetPlanForNodeByProvider(ctx, nodeAddr, provAddr, id)
	}
}

func (k *Migrator) setPlanForProviderKeys(ctx sdk.Context, keys ...[]byte) {
	for _, key := range keys {
		addrLen := int(key[0])
		addr := base.ProvAddress(key[1 : 1+addrLen])
		id := sdk.BigEndianToUint64(key[1+addrLen : 1+addrLen+8])

		k.plan.SetPlanForProvider(ctx, addr, id)
	}
}

func (k *Migrator) PostMigrate(ctx sdk.Context) error {
	store := prefix.NewStore(k.plan.Store(ctx), types.PlanForNodeKeyPrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		key := it.Key()
		nodeAddrLen := int(key[0])
		nodeAddr := base.NodeAddress(key[1 : 1+nodeAddrLen])
		provAddrLen := int(key[1+nodeAddrLen])
		provAddr := base.ProvAddress(key[1+nodeAddrLen+1 : 1+nodeAddrLen+1+provAddrLen])
		id := binary.BigEndian.Uint64(key[1+nodeAddrLen+1+provAddrLen:])

		if !k.lease.HasAnyLeaseForNodeByProvider(ctx, nodeAddr, provAddr) {
			k.node.DeleteNodeForPlan(ctx, id, nodeAddr)
			k.plan.DeletePlanForNodeByProvider(ctx, nodeAddr, provAddr, id)
		}
	}

	return nil
}
