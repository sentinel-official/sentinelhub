package app

import (
	storetypes "cosmossdk.io/store/types"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/controller/types"
	ibcicahosttypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v10/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v10/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v10/modules/core/exported"

	customminttypes "github.com/sentinel-official/sentinelhub/v13/x/mint/types"
	oracletypes "github.com/sentinel-official/sentinelhub/v13/x/oracle/types"
	swaptypes "github.com/sentinel-official/sentinelhub/v13/x/swap/types"
	vpntypes "github.com/sentinel-official/sentinelhub/v13/x/vpn/types"
)

type StoreKeys struct {
	kv        map[string]*storetypes.KVStoreKey
	memory    map[string]*storetypes.MemoryStoreKey
	transient map[string]*storetypes.TransientStoreKey
}

func NewStoreKeys() StoreKeys {
	var (
		kv = sdk.NewKVStoreKeys(
			// Cosmos SDK keys
			authtypes.StoreKey,
			authzkeeper.StoreKey,
			banktypes.StoreKey,
			capabilitytypes.StoreKey,
			consensustypes.StoreKey,
			crisistypes.StoreKey,
			distributiontypes.StoreKey,
			evidencetypes.StoreKey,
			feegrant.StoreKey,
			govtypes.StoreKey,
			group.StoreKey,
			minttypes.StoreKey,
			nftkeeper.StoreKey,
			paramstypes.StoreKey,
			slashingtypes.StoreKey,
			stakingtypes.StoreKey,
			upgradetypes.StoreKey,

			// Cosmos IBC keys
			ibcexported.StoreKey,
			ibcfeetypes.StoreKey,
			ibcicacontrollertypes.StoreKey,
			ibcicahosttypes.StoreKey,
			ibctransfertypes.StoreKey,

			// Sentinel Hub keys
			customminttypes.StoreKey,
			oracletypes.StoreKey,
			swaptypes.StoreKey,
			vpntypes.StoreKey,

			// Other keys
			wasmtypes.StoreKey,
		)
		memory = sdk.NewMemoryStoreKeys(
			capabilitytypes.MemStoreKey,
		)
		transient = sdk.NewTransientStoreKeys(
			paramstypes.TStoreKey,
		)
	)

	return StoreKeys{
		kv:        kv,
		memory:    memory,
		transient: transient,
	}
}

func (sk *StoreKeys) KVKeys() map[string]*storetypes.KVStoreKey               { return sk.kv }
func (sk *StoreKeys) MemoryKeys() map[string]*storetypes.MemoryStoreKey       { return sk.memory }
func (sk *StoreKeys) TransientKeys() map[string]*storetypes.TransientStoreKey { return sk.transient }

func (sk *StoreKeys) KV(v string) *storetypes.KVStoreKey               { return sk.kv[v] }
func (sk *StoreKeys) Memory(v string) *storetypes.MemoryStoreKey       { return sk.memory[v] }
func (sk *StoreKeys) Transient(v string) *storetypes.TransientStoreKey { return sk.transient[v] }
