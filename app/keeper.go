package app

import (
	"path/filepath"

	"cosmossdk.io/log"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
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
	ibctransfer "github.com/cosmos/ibc-go/v10/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v10/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v10/modules/apps/transfer/types"
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
	ConsensusKeeper    consensuskeeper.Keeper
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
}

func (k *Keepers) Subspace(v string) paramstypes.Subspace {
	subspace, _ := k.ParamsKeeper.GetSubspace(v)

	return subspace
}

func (k *Keepers) SetParamSubspaces(app *baseapp.BaseApp) {
	// Tendermint subspaces
	app.SetParamStore(k.ConsensusKeeper.ParamsStore)

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
	app *baseapp.BaseApp, blockedAddrs map[string]bool, encCfg EncodingConfig, homeDir string, keys StoreKeys,
	logger log.Logger, mAccPerms map[string][]string, skipUpgradeHeights map[int64]bool,
	wasmConfig wasmtypes.NodeConfig, wasmOpts []wasmkeeper.Option,
) (k Keepers) {
	govModuleAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	// Cosmos SDK keepers
	k.ConsensusKeeper = consensuskeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(consensustypes.StoreKey), govModuleAddr, runtime.EventService{},
	)
	k.ParamsKeeper = paramskeeper.NewKeeper(
		encCfg.Codec, encCfg.Amino, keys.KV(paramstypes.StoreKey), keys.Transient(paramstypes.TStoreKey),
	)
	k.SetParamSubspaces(app)

	k.AccountKeeper = authkeeper.NewAccountKeeper(
		encCfg.Codec, keys.KVStoreService(authtypes.StoreKey), authtypes.ProtoBaseAccount, mAccPerms,
		address.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.GetConfig().GetBech32AccountAddrPrefix(), govModuleAddr,
	)
	k.AuthzKeeper = authzkeeper.NewKeeper(
		keys.KVStoreService(authzkeeper.StoreKey), encCfg.Codec, app.MsgServiceRouter(), k.AccountKeeper,
	)
	k.BankKeeper = bankkeeper.NewBaseKeeper(
		encCfg.Codec, keys.KVStoreService(banktypes.StoreKey), k.AccountKeeper, blockedAddrs, govModuleAddr, logger,
	)

	k.StakingKeeper = stakingkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(stakingtypes.StoreKey), k.AccountKeeper, k.BankKeeper, govModuleAddr,
		address.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		address.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)
	k.DistributionKeeper = distributionkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(distributiontypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
		authtypes.FeeCollectorName, govModuleAddr,
	)
	k.SlashingKeeper = slashingkeeper.NewKeeper(
		encCfg.Codec, encCfg.Amino, keys.KVStoreService(slashingtypes.StoreKey), k.StakingKeeper, govModuleAddr,
	)

	k.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(k.DistributionKeeper.Hooks(), k.SlashingKeeper.Hooks()),
	)

	k.EvidenceKeeper = *evidencekeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(evidencetypes.StoreKey), k.StakingKeeper, k.SlashingKeeper,
		k.AccountKeeper.AddressCodec(), runtime.ProvideCometInfoService(),
	)

	evidenceRouter := evidencetypes.NewRouter()
	k.EvidenceKeeper.SetRouter(evidenceRouter)

	k.FeeGrantKeeper = feegrantkeeper.NewKeeper(encCfg.Codec, keys.KVStoreService(feegrant.StoreKey), k.AccountKeeper)

	groupConfig := group.DefaultConfig()
	k.GroupKeeper = groupkeeper.NewKeeper(
		keys.KV(group.StoreKey), encCfg.Codec, app.MsgServiceRouter(), k.AccountKeeper, groupConfig,
	)

	k.MintKeeper = mintkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(minttypes.StoreKey), k.StakingKeeper, k.AccountKeeper, k.BankKeeper,
		authtypes.FeeCollectorName, govModuleAddr,
	)
	k.NFTKeeper = nftkeeper.NewKeeper(keys.KVStoreService(nftkeeper.StoreKey), encCfg.Codec, k.AccountKeeper, k.BankKeeper)
	k.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights, keys.KVStoreService(upgradetypes.StoreKey), encCfg.Codec, homeDir, app, govModuleAddr,
	)

	// Cosmos IBC keepers
	k.IBCKeeper = ibckeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(ibcexported.StoreKey), k.Subspace(ibcexported.ModuleName),
		k.UpgradeKeeper, govModuleAddr,
	)
	k.IBCICAControllerKeeper = ibcicacontrollerkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(ibcicacontrollertypes.StoreKey), k.Subspace(ibcicacontrollertypes.SubModuleName),
		k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper, app.MsgServiceRouter(), govModuleAddr,
	)
	k.IBCICAHostKeeper = ibcicahostkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(ibcicahosttypes.StoreKey), k.Subspace(ibcicahosttypes.SubModuleName),
		k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper, k.AccountKeeper, app.MsgServiceRouter(),
		app.GRPCQueryRouter(), govModuleAddr,
	)
	k.IBCTransferKeeper = ibctransferkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(ibctransfertypes.StoreKey), k.Subspace(ibctransfertypes.ModuleName),
		k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper, app.MsgServiceRouter(), k.AccountKeeper, k.BankKeeper,
		govModuleAddr,
	)

	// Sentinel Hub keepers
	k.CustomMintKeeper = custommintkeeper.NewKeeper(encCfg.Codec, keys.KVStoreService(customminttypes.StoreKey), nil) // TODO: set minter
	k.OracleKeeper = oraclekeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(oracletypes.StoreKey), k.IBCKeeper.ChannelKeeper, govModuleAddr,
	)
	k.SwapKeeper = swapkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(swaptypes.StoreKey), k.Subspace(swaptypes.ModuleName), k.AccountKeeper, k.BankKeeper,
	)
	k.VPNKeeper = vpnkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(vpntypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.DistributionKeeper, &k.OracleKeeper,
		app.MsgServiceRouter(), govModuleAddr, authtypes.FeeCollectorName,
	)

	// Other keepers
	wasmDir := filepath.Join(homeDir, "data")
	k.WasmKeeper = wasmkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(wasmtypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
		distributionkeeper.NewQuerier(k.DistributionKeeper), k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper,
		k.IBCKeeper.ChannelKeeperV2, k.IBCTransferKeeper, app.MsgServiceRouter(), app.GRPCQueryRouter(),
		wasmDir, wasmConfig, wasmtypes.VMConfig{}, wasmkeeper.BuiltInCapabilities(), govModuleAddr, wasmOpts...,
	)

	govConfig := govtypes.DefaultConfig()
	k.GovKeeper = govkeeper.NewKeeper(
		encCfg.Codec, keys.KVStoreService(govtypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
		k.DistributionKeeper, app.MsgServiceRouter(), govConfig, govModuleAddr,
	)

	// Cosmos SDK Governance router
	govRouter := govv1beta1types.NewRouter().
		AddRoute(govtypes.RouterKey, govv1beta1types.ProposalHandler).
		AddRoute(paramsproposal.RouterKey, params.NewParamChangeProposalHandler(k.ParamsKeeper))
	k.GovKeeper.SetLegacyRouter(govRouter)

	// Cosmos IBC port router
	var ibcICAControllerIBCModule ibcporttypes.IBCModule
	var ibcICAHostIBCModule ibcporttypes.IBCModule
	var ibcTransferIBCModule ibcporttypes.IBCModule
	var wasmIBCModule ibcporttypes.IBCModule

	ibcICAControllerIBCModule = ibcicacontroller.NewIBCMiddleware(k.IBCICAControllerKeeper)

	ibcICAHostIBCModule = ibcicahost.NewIBCModule(k.IBCICAHostKeeper)

	ibcTransferIBCModule = ibctransfer.NewIBCModule(k.IBCTransferKeeper)

	wasmIBCModule = wasm.NewIBCHandler(
		k.WasmKeeper, k.IBCKeeper.ChannelKeeper, k.IBCTransferKeeper, k.IBCKeeper.ChannelKeeper,
	)

	ibcPortRouter := ibcporttypes.NewRouter().
		AddRoute(ibcicacontrollertypes.SubModuleName, ibcICAControllerIBCModule).
		AddRoute(ibcicahosttypes.SubModuleName, ibcICAHostIBCModule).
		AddRoute(ibctransfertypes.ModuleName, ibcTransferIBCModule).
		AddRoute(wasmtypes.ModuleName, wasmIBCModule)
	k.IBCKeeper.SetRouter(ibcPortRouter)

	return k
}
