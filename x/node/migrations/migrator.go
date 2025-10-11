package migrations

import (
	"net"
	"net/url"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v2"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
)

type Migrator struct {
	cdc  codec.BinaryCodec
	node NodeKeeper
}

func NewMigrator(cdc codec.BinaryCodec, node NodeKeeper) Migrator {
	return Migrator{
		cdc:  cdc,
		node: node,
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	k.setParams(ctx)
	k.migrateNodes(ctx)

	nodeForInactiveAtKeys := k.deleteKeys(ctx, []byte{0x11})
	nodeForPlanKeys := k.deleteKeys(ctx, []byte{0x12})

	k.setNodeForPlanKeys(ctx, nodeForPlanKeys...)
	k.setNodeForInactiveAtKeys(ctx, nodeForInactiveAtKeys...)

	return nil
}

func (k *Migrator) deleteKeys(ctx sdk.Context, keyPrefix []byte) (keys [][]byte) {
	store := prefix.NewStore(k.node.Store(ctx), keyPrefix)

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
		Deposit:           sdk.NewInt64Coin("udvpn", 0),
		MinGigabytePrices: []v1.Price{},
		MinHourlyPrices:   []v1.Price{},
		StatusTimeout:     1 * time.Hour,
	}

	k.node.SetParams(ctx, params)
}

func (k *Migrator) migrateNodes(ctx sdk.Context) {
	store := prefix.NewStore(k.node.Store(ctx), []byte{0x10})

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())

		var item v2.Node
		k.cdc.MustUnmarshal(it.Value(), &item)

		gigabytePrices, err := v1.NewPricesFromCoins(item.GigabytePrices...)
		if err != nil {
			panic(err)
		}

		hourlyPrices, err := v1.NewPricesFromCoins(item.HourlyPrices...)
		if err != nil {
			panic(err)
		}

		remoteURL, err := url.ParseRequestURI(item.RemoteURL)
		if err != nil {
			panic(err)
		}

		remoteAddr := net.JoinHostPort(remoteURL.Hostname(), remoteURL.Port())

		node := v3.Node{
			Address:        item.Address,
			GigabytePrices: gigabytePrices,
			HourlyPrices:   hourlyPrices,
			RemoteAddrs:    []string{remoteAddr},
			InactiveAt:     item.InactiveAt,
			Status:         item.Status,
			StatusAt:       item.StatusAt,
		}

		k.node.SetNode(ctx, node)
	}
}

func (k *Migrator) setNodeForPlanKeys(ctx sdk.Context, keys ...[]byte) {
	for _, key := range keys {
		id := sdk.BigEndianToUint64(key[:8])
		addrLen := int(key[8])
		addr := types.NodeAddress(key[9 : 9+addrLen])

		k.node.SetNodeForPlan(ctx, id, addr)
	}
}

func (k *Migrator) setNodeForInactiveAtKeys(ctx sdk.Context, keys ...[]byte) {
	for _, key := range keys {
		inactiveAt, err := sdk.ParseTimeBytes(key[:29])
		if err != nil {
			panic(err)
		}

		addrLen := int(key[29])
		addr := types.NodeAddress(key[30 : 30+addrLen])

		k.node.SetNodeForInactiveAt(ctx, inactiveAt, addr)
	}
}
