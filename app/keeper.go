package app

import (
	"path/filepath"

	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramsproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcicacontroller "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/controller"
	ibcicacontrollerkeeper "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/controller/keeper"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/controller/types"
	ibcicahost "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/host"
	ibcicahostkeeper "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/host/keeper"
	ibcicahosttypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/host/types"
	ibcfee "github.com/cosmos/ibc-go/v10/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v10/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v10/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v10/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v10/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v10/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v10/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v10/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v10/modules/core/keeper"

	custommintkeeper "github.com/sentinel-official/sentinelhub/v13/x/mint/keeper"
	customminttypes "github.com/sentinel-official/sentinelhub/v13/x/mint/types"
	oraclekeeper "github.com/sentinel-official/sentinelhub/v13/x/oracle/keeper"
	oracletypes "github.com/sentinel-official/sentinelhub/v13/x/oracle/types"
	swapkeeper "github.com/sentinel-official/sentinelhub/v13/x/swap/keeper"
	swaptypes "github.com/sentinel-official/sentinelhub/v13/x/swap/types"
	vpnkeeper "github.com/sentinel-official/sentinelhub/v13/x/vpn/keeper"
	vpntypes "github.com/sentinel-official/sentinelhub/v13/x/vpn/types"
)

type Keepers struct {
	// Cosmos SDK keepers
	AccountKeeper      authkeeper.AccountKeeper
	AuthzKeeper        authzkeeper.Keeper
	BankKeeper         bankkeeper.Keeper
	CapabilityKeeper   *capabilitykeeper.Keeper
	ConsensusKeeper    consensuskeeper.Keeper
	CrisisKeeper       *crisiskeeper.Keeper
	DistributionKeeper distributionkeeper.Keeper
	EvidenceKeeper     evidencekeeper.Keeper
	FeeGrantKeeper     feegrantkeeper.Keeper
	GovKeeper          *govkeeper.Keeper
	GroupKeeper        groupkeeper.Keeper
	MintKeeper         mintkeeper.Keeper
	NFTKeeper          nftkeeper.Keeper
	ParamsKeeper       paramskeeper.Keeper
	SlashingKeeper     slashingkeeper.Keeper
	StakingKeeper      *stakingkeeper.Keeper
	UpgradeKeeper      *upgradekeeper.Keeper

	// Cosmos IBC keepers
	IBCKeeper              *ibckeeper.Keeper
	IBCFeeKeeper           ibcfeekeeper.Keeper
	IBCICAControllerKeeper ibcicacontrollerkeeper.Keeper
	IBCICAHostKeeper       ibcicahostkeeper.Keeper
	IBCTransferKeeper      ibctransferkeeper.Keeper

	// Sentinel Hub keepers
	CustomMintKeeper custommintkeeper.Keeper
	OracleKeeper     oraclekeeper.Keeper
	SwapKeeper       swapkeeper.Keeper
	VPNKeeper        vpnkeeper.Keeper

	// Other keepers
	ContractKeeper *wasmkeeper.PermissionedKeeper
	WasmKeeper     wasmkeeper.Keeper

	// Cosmos IBC scoped keepers
	ScopedIBCKeeper              capabilitykeeper.ScopedKeeper
	ScopedIBCFeeKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedIBCICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper      capabilitykeeper.ScopedKeeper

	// Sentinel scoped keepers
	ScopedOracleKeeper capabilitykeeper.ScopedKeeper

	// Other scoped keepers
	ScopedWasmKeeper capabilitykeeper.ScopedKeeper
}

func (k *Keepers) Subspace(v string) paramstypes.Subspace {
	subspace, _ := k.ParamsKeeper.GetSubspace(v)

	return subspace
}

func (k *Keepers) SetParamSubspaces(app *baseapp.BaseApp) {
	// Tendermint subspaces
	app.SetParamStore(&k.ConsensusKeeper)

	// Cosmos SDK subspaces
	k.ParamsKeeper.Subspace(authtypes.ModuleName)
	k.ParamsKeeper.Subspace(banktypes.ModuleName)
	k.ParamsKeeper.Subspace(crisistypes.ModuleName)
	k.ParamsKeeper.Subspace(distributiontypes.ModuleName)
	k.ParamsKeeper.Subspace(govtypes.ModuleName)
	k.ParamsKeeper.Subspace(minttypes.ModuleName)
	k.ParamsKeeper.Subspace(slashingtypes.ModuleName)
	k.ParamsKeeper.Subspace(stakingtypes.ModuleName)

	// Cosmos IBC subspaces
	k.ParamsKeeper.Subspace(ibcexported.ModuleName)
	k.ParamsKeeper.Subspace(ibcicacontrollertypes.SubModuleName)
	k.ParamsKeeper.Subspace(ibcicahosttypes.SubModuleName)
	k.ParamsKeeper.Subspace(ibctransfertypes.ModuleName)

	// Sentinel Hub subspaces
	k.ParamsKeeper.Subspace(swaptypes.ModuleName)

	// Other subspaces
	k.ParamsKeeper.Subspace(wasmtypes.ModuleName)
}

func NewKeepers(
	app *baseapp.BaseApp, bech32Prefix string, blockedAddrs map[string]bool, encCfg EncodingConfig, homeDir string,
	invCheckPeriod uint, keys StoreKeys, mAccPerms map[string][]string, skipUpgradeHeights map[int64]bool,
	wasmConfig wasmtypes.WasmConfig, wasmOpts []wasmkeeper.Option,
) (k Keepers) {
	govModuleAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	// Cosmos SDK keepers
	k.ConsensusKeeper = consensuskeeper.NewKeeper(encCfg.Codec, keys.KV(consensustypes.StoreKey), govModuleAddr)
	k.ParamsKeeper = paramskeeper.NewKeeper(
		encCfg.Codec, encCfg.Amino, keys.KV(paramstypes.StoreKey), keys.Transient(paramstypes.TStoreKey),
	)
	k.SetParamSubspaces(app)

	k.AccountKeeper = authkeeper.NewAccountKeeper(
		encCfg.Codec, keys.KV(authtypes.StoreKey), authtypes.ProtoBaseAccount, mAccPerms, bech32Prefix, govModuleAddr,
	)
	k.AuthzKeeper = authzkeeper.NewKeeper(
		keys.KV(authzkeeper.StoreKey), encCfg.Codec, app.MsgServiceRouter(), k.AccountKeeper,
	)
	k.BankKeeper = bankkeeper.NewBaseKeeper(
		encCfg.Codec, keys.KV(banktypes.StoreKey), k.AccountKeeper, blockedAddrs, govModuleAddr,
	)
	k.CapabilityKeeper = capabilitykeeper.NewKeeper(
		encCfg.Codec, keys.KV(capabilitytypes.StoreKey), keys.Memory(capabilitytypes.MemStoreKey),
	)

	k.CrisisKeeper = crisiskeeper.NewKeeper(
		encCfg.Codec, keys.KV(crisistypes.StoreKey), invCheckPeriod, k.BankKeeper, authtypes.FeeCollectorName,
		govModuleAddr,
	)

	k.StakingKeeper = stakingkeeper.NewKeeper(
		encCfg.Codec, keys.KV(stakingtypes.StoreKey), k.AccountKeeper, k.BankKeeper, govModuleAddr,
	)
	k.DistributionKeeper = distributionkeeper.NewKeeper(
		encCfg.Codec, keys.KV(distributiontypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
		authtypes.FeeCollectorName, govModuleAddr,
	)
	k.SlashingKeeper = slashingkeeper.NewKeeper(
		encCfg.Codec, encCfg.Amino, keys.KV(slashingtypes.StoreKey), k.StakingKeeper, govModuleAddr,
	)

	k.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(k.DistributionKeeper.Hooks(), k.SlashingKeeper.Hooks()),
	)

	k.EvidenceKeeper = *evidencekeeper.NewKeeper(
		encCfg.Codec, keys.KV(evidencetypes.StoreKey), k.StakingKeeper, k.SlashingKeeper,
	)

	evidenceRouter := evidencetypes.NewRouter()
	k.EvidenceKeeper.SetRouter(evidenceRouter)

	k.FeeGrantKeeper = feegrantkeeper.NewKeeper(encCfg.Codec, keys.KV(feegrant.StoreKey), k.AccountKeeper)

	groupConfig := group.DefaultConfig()
	k.GroupKeeper = groupkeeper.NewKeeper(
		keys.KV(group.StoreKey), encCfg.Codec, app.MsgServiceRouter(), k.AccountKeeper, groupConfig,
	)

	k.MintKeeper = mintkeeper.NewKeeper(
		encCfg.Codec, keys.KV(minttypes.StoreKey), k.StakingKeeper, k.AccountKeeper, k.BankKeeper,
		authtypes.FeeCollectorName, govModuleAddr,
	)
	k.NFTKeeper = nftkeeper.NewKeeper(keys.KV(nftkeeper.StoreKey), encCfg.Codec, k.AccountKeeper, k.BankKeeper)
	k.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights, keys.KV(upgradetypes.StoreKey), encCfg.Codec, homeDir, app, govModuleAddr,
	)

	// Cosmos IBC keepers
	k.ScopedIBCKeeper = k.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	k.ScopedIBCFeeKeeper = k.CapabilityKeeper.ScopeToModule(ibcfeetypes.ModuleName)
	k.ScopedIBCICAControllerKeeper = k.CapabilityKeeper.ScopeToModule(ibcicacontrollertypes.SubModuleName)
	k.ScopedIBCICAHostKeeper = k.CapabilityKeeper.ScopeToModule(ibcicahosttypes.SubModuleName)
	k.ScopedIBCTransferKeeper = k.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	k.IBCKeeper = ibckeeper.NewKeeper(
		encCfg.Codec, keys.KV(ibcexported.StoreKey), k.Subspace(ibcexported.ModuleName),
		k.StakingKeeper, k.UpgradeKeeper, k.ScopedIBCKeeper,
	)
	k.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		encCfg.Codec, keys.KV(ibcfeetypes.StoreKey), k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper,
		&k.IBCKeeper.PortKeeper, k.AccountKeeper, k.BankKeeper,
	)
	k.IBCICAControllerKeeper = ibcicacontrollerkeeper.NewKeeper(
		encCfg.Codec, keys.KV(ibcicacontrollertypes.StoreKey), k.Subspace(ibcicacontrollertypes.SubModuleName),
		k.IBCFeeKeeper, k.IBCKeeper.ChannelKeeper, &k.IBCKeeper.PortKeeper,
		k.ScopedIBCICAControllerKeeper, app.MsgServiceRouter(),
	)

	k.IBCICAHostKeeper = ibcicahostkeeper.NewKeeper(
		encCfg.Codec, keys.KV(ibcicahosttypes.StoreKey), k.Subspace(ibcicahosttypes.SubModuleName),
		k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper, &k.IBCKeeper.PortKeeper, k.AccountKeeper,
		k.ScopedIBCICAHostKeeper, app.MsgServiceRouter(),
	)
	k.IBCICAHostKeeper.WithQueryRouter(app.GRPCQueryRouter())

	k.IBCTransferKeeper = ibctransferkeeper.NewKeeper(
		encCfg.Codec, keys.KV(ibctransfertypes.StoreKey), k.Subspace(ibctransfertypes.ModuleName),
		k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper, &k.IBCKeeper.PortKeeper,
		k.AccountKeeper, k.BankKeeper, k.ScopedIBCTransferKeeper,
	)

	// Sentinel Hub keepers
	k.ScopedOracleKeeper = k.CapabilityKeeper.ScopeToModule(oracletypes.ModuleName)

	k.CustomMintKeeper = custommintkeeper.NewKeeper(encCfg.Codec, keys.KV(customminttypes.StoreKey), k.MintKeeper)
	k.OracleKeeper = oraclekeeper.NewKeeper(
		encCfg.Codec, keys.KV(oracletypes.StoreKey), k.ScopedOracleKeeper, k.IBCKeeper.ChannelKeeper,
		&k.IBCKeeper.PortKeeper, govModuleAddr,
	)
	k.SwapKeeper = swapkeeper.NewKeeper(
		encCfg.Codec, keys.KV(swaptypes.StoreKey), k.Subspace(swaptypes.ModuleName), k.AccountKeeper, k.BankKeeper,
	)
	k.VPNKeeper = vpnkeeper.NewKeeper(
		encCfg.Codec, keys.KV(vpntypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.DistributionKeeper, &k.OracleKeeper,
		app.MsgServiceRouter(), govModuleAddr, authtypes.FeeCollectorName,
	)

	// Other keepers
	k.ScopedWasmKeeper = k.CapabilityKeeper.ScopeToModule(wasmtypes.ModuleName)

	wasmCapabilities := "iterator,staking,stargate,cosmwasm_1_1,cosmwasm_1_2"
	wasmDir := filepath.Join(homeDir, "data")

	k.WasmKeeper = wasmkeeper.NewKeeper(
		encCfg.Codec, keys.KV(wasmtypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
		distributionkeeper.NewQuerier(k.DistributionKeeper), k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper,
		&k.IBCKeeper.PortKeeper, k.ScopedWasmKeeper, k.IBCTransferKeeper, app.MsgServiceRouter(), app.GRPCQueryRouter(),
		wasmDir, wasmConfig, wasmCapabilities, govModuleAddr, wasmOpts...,
	)

	govConfig := govtypes.DefaultConfig()
	k.GovKeeper = govkeeper.NewKeeper(
		encCfg.Codec, keys.KV(govtypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
		app.MsgServiceRouter(), govConfig, govModuleAddr,
	)

	// Cosmos SDK Governance router
	govRouter := govv1beta1types.NewRouter().
		AddRoute(govtypes.RouterKey, govv1beta1types.ProposalHandler).
		AddRoute(paramsproposal.RouterKey, params.NewParamChangeProposalHandler(k.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(k.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(k.IBCKeeper.ClientKeeper))
	k.GovKeeper.SetLegacyRouter(govRouter)

	// Cosmos IBC port router
	var ibcICAControllerIBCModule ibcporttypes.IBCModule

	ibcICAControllerIBCModule = ibcicacontroller.NewIBCMiddleware(ibcICAControllerIBCModule, k.IBCICAControllerKeeper)
	ibcICAControllerIBCModule = ibcfee.NewIBCMiddleware(ibcICAControllerIBCModule, k.IBCFeeKeeper)

	var ibcICAHostIBCModule ibcporttypes.IBCModule = ibcicahost.NewIBCModule(k.IBCICAHostKeeper)

	var ibcTransferIBCModule ibcporttypes.IBCModule

	ibcTransferIBCModule = ibctransfer.NewIBCModule(k.IBCTransferKeeper)
	ibcTransferIBCModule = ibcfee.NewIBCMiddleware(ibcTransferIBCModule, k.IBCFeeKeeper)

	var wasmIBCModule ibcporttypes.IBCModule

	wasmIBCModule = wasm.NewIBCHandler(k.WasmKeeper, k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper)
	wasmIBCModule = ibcfee.NewIBCMiddleware(wasmIBCModule, k.IBCFeeKeeper)

	ibcPortRouter := ibcporttypes.NewRouter().
		AddRoute(ibcicacontrollertypes.SubModuleName, ibcICAControllerIBCModule).
		AddRoute(ibcicahosttypes.SubModuleName, ibcICAHostIBCModule).
		AddRoute(ibctransfertypes.ModuleName, ibcTransferIBCModule).
		AddRoute(wasmtypes.ModuleName, wasmIBCModule)
	k.IBCKeeper.SetRouter(ibcPortRouter)

	k.CapabilityKeeper.Seal()

	return k
}
