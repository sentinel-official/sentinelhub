package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	ibcicatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v7/modules/apps/29-fee"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibcsolomachine "github.com/cosmos/ibc-go/v7/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	deposittypes "github.com/sentinel-official/sentinelhub/v13/x/deposit/types"
	custommint "github.com/sentinel-official/sentinelhub/v13/x/mint"
	customminttypes "github.com/sentinel-official/sentinelhub/v13/x/mint/types"
	"github.com/sentinel-official/sentinelhub/v13/x/oracle"
	oracletypes "github.com/sentinel-official/sentinelhub/v13/x/oracle/types"
	"github.com/sentinel-official/sentinelhub/v13/x/swap"
	swaptypes "github.com/sentinel-official/sentinelhub/v13/x/swap/types"
	"github.com/sentinel-official/sentinelhub/v13/x/vpn"
	vpntypes "github.com/sentinel-official/sentinelhub/v13/x/vpn/types"
)

var (
	ModuleBasics = sdkmodule.NewBasicManager(
		// Cosmos SDK module basics
		auth.AppModuleBasic{},
		authvesting.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		consensus.AppModuleBasic{},
		crisis.AppModuleBasic{},
		distribution.AppModuleBasic{},
		evidence.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				// Cosmos SDK proposal handlers
				paramsclient.ProposalHandler,
				upgradeclient.LegacyProposalHandler,
				upgradeclient.LegacyCancelProposalHandler,

				// Cosmos IBC proposal handlers
				ibcclientclient.UpdateClientProposalHandler,
				ibcclientclient.UpgradeProposalHandler,
			},
		),
		groupmodule.AppModuleBasic{},
		mint.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		staking.AppModuleBasic{},
		upgrade.AppModuleBasic{},

		// Cosmos IBC module basics
		ibc.AppModuleBasic{},
		ibcsolomachine.AppModuleBasic{},
		ibcfee.AppModuleBasic{},
		ibcica.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},

		// Sentinel Hub module basics
		custommint.AppModuleBasic{},
		oracle.AppModuleBasic{},
		swap.AppModuleBasic{},
		vpn.AppModuleBasic{},

		// Other module basics
		wasm.AppModuleBasic{},
	)
)

func ModuleAccPerms() map[string][]string {
	return map[string][]string{
		// Cosmos SDK module account permissions
		authtypes.FeeCollectorName:     nil,
		distributiontypes.ModuleName:   nil,
		govtypes.ModuleName:            {authtypes.Burner},
		minttypes.ModuleName:           {authtypes.Minter},
		nft.ModuleName:                 nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},

		// Cosmos IBC module account permissions
		ibcicatypes.ModuleName:      nil,
		ibcfeetypes.ModuleName:      nil,
		ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},

		// Sentinel Hub module account permissions
		customminttypes.ModuleName: nil,
		oracletypes.ModuleName:     nil,
		deposittypes.ModuleName:    nil,
		swaptypes.ModuleName:       {authtypes.Minter},

		// Other module account permissions
		wasmtypes.ModuleName: {authtypes.Burner},
	}
}

func BlockedAccAddrs() map[string]bool {
	accAddrs := make(map[string]bool)

	for v := range ModuleAccPerms() {
		accAddr := authtypes.NewModuleAddress(v)
		accAddrs[accAddr.String()] = true
	}

	return accAddrs
}

func NewModuleManager(
	deliverTxFunc func(abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx, encCfg EncodingConfig, k Keepers,
	msgRouter *baseapp.MsgServiceRouter, skipGenesisInvariants bool,
) *sdkmodule.Manager {
	manager := sdkmodule.NewManager(
		// Cosmos SDK modules
		auth.NewAppModule(encCfg.Codec, k.AccountKeeper, nil, k.Subspace(authtypes.ModuleName)),
		authvesting.NewAppModule(k.AccountKeeper, k.BankKeeper),
		authzmodule.NewAppModule(encCfg.Codec, k.AuthzKeeper, k.AccountKeeper, k.BankKeeper, encCfg.InterfaceRegistry),
		bank.NewAppModule(encCfg.Codec, k.BankKeeper, k.AccountKeeper, k.Subspace(banktypes.ModuleName)),
		capability.NewAppModule(encCfg.Codec, *k.CapabilityKeeper, false),
		consensus.NewAppModule(encCfg.Codec, k.ConsensusKeeper),
		crisis.NewAppModule(k.CrisisKeeper, skipGenesisInvariants, k.Subspace(crisistypes.ModuleName)),
		distribution.NewAppModule(
			encCfg.Codec, k.DistributionKeeper, k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
			k.Subspace(distributiontypes.ModuleName),
		),
		evidence.NewAppModule(k.EvidenceKeeper),
		feegrantmodule.NewAppModule(encCfg.Codec, k.AccountKeeper, k.BankKeeper, k.FeeGrantKeeper, encCfg.InterfaceRegistry),
		genutil.NewAppModule(k.AccountKeeper, k.StakingKeeper, deliverTxFunc, encCfg.TxConfig),
		gov.NewAppModule(encCfg.Codec, k.GovKeeper, k.AccountKeeper, k.BankKeeper, k.Subspace(govtypes.ModuleName)),
		groupmodule.NewAppModule(encCfg.Codec, k.GroupKeeper, k.AccountKeeper, k.BankKeeper, encCfg.InterfaceRegistry),
		mint.NewAppModule(encCfg.Codec, k.MintKeeper, k.AccountKeeper, nil, k.Subspace(minttypes.ModuleName)),
		nftmodule.NewAppModule(encCfg.Codec, k.NFTKeeper, k.AccountKeeper, k.BankKeeper, encCfg.InterfaceRegistry),
		params.NewAppModule(k.ParamsKeeper),
		slashing.NewAppModule(
			encCfg.Codec, k.SlashingKeeper, k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
			k.Subspace(slashingtypes.ModuleName),
		),
		staking.NewAppModule(
			encCfg.Codec, k.StakingKeeper, k.AccountKeeper, k.BankKeeper, k.Subspace(stakingtypes.ModuleName),
		),
		upgrade.NewAppModule(k.UpgradeKeeper),

		// Cosmos IBC modules
		ibcfee.NewAppModule(k.IBCFeeKeeper),
		ibcica.NewAppModule(&k.IBCICAControllerKeeper, &k.IBCICAHostKeeper),
		ibc.NewAppModule(k.IBCKeeper),
		ibctransfer.NewAppModule(k.IBCTransferKeeper),

		// Sentinel Hub modules
		custommint.NewAppModule(encCfg.Codec, k.CustomMintKeeper),
		oracle.NewAppModule(encCfg.Codec, k.OracleKeeper),
		swap.NewAppModule(encCfg.Codec, k.SwapKeeper),
		vpn.NewAppModule(encCfg.Codec, k.AccountKeeper, k.BankKeeper, k.VPNKeeper),

		// Other modules
		wasm.NewAppModule(
			encCfg.Codec, &k.WasmKeeper, k.StakingKeeper, k.AccountKeeper, k.BankKeeper, msgRouter,
			k.Subspace(wasmtypes.ModuleName),
		),
	)

	manager.SetOrderBeginBlockers(
		// Cosmos SDK modules
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		customminttypes.ModuleName, // Sentinel Hub module
		minttypes.ModuleName,
		distributiontypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		nft.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		authvestingtypes.ModuleName,
		consensustypes.ModuleName,

		// Cosmos IBC modules
		ibctransfertypes.ModuleName,
		ibcexported.ModuleName,
		ibcicatypes.ModuleName,
		ibcfeetypes.ModuleName,

		// Sentinel Hub modules
		swaptypes.ModuleName,
		oracletypes.ModuleName,
		vpntypes.ModuleName,

		// Other modules
		wasmtypes.ModuleName,
	)
	manager.SetOrderEndBlockers(
		// Cosmos SDK modules
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distributiontypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		nft.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		authvestingtypes.ModuleName,
		consensustypes.ModuleName,

		// Cosmos IBC modules
		ibctransfertypes.ModuleName,
		ibcexported.ModuleName,
		ibcicatypes.ModuleName,
		ibcfeetypes.ModuleName,

		// Sentinel Hub modules
		customminttypes.ModuleName,
		oracletypes.ModuleName,
		swaptypes.ModuleName,
		vpntypes.ModuleName,

		// Other modules
		wasmtypes.ModuleName,
	)
	manager.SetOrderInitGenesis(
		// Cosmos SDK modules
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distributiontypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		nft.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		authvestingtypes.ModuleName,
		consensustypes.ModuleName,

		// Cosmos IBC modules
		ibctransfertypes.ModuleName,
		ibcexported.ModuleName,
		ibcicatypes.ModuleName,
		ibcfeetypes.ModuleName,

		// Sentinel Hub modules
		customminttypes.ModuleName,
		oracletypes.ModuleName,
		swaptypes.ModuleName,
		vpntypes.ModuleName,

		// Other modules
		wasmtypes.ModuleName,
	)

	return manager
}

func NewSimulationManager(
	encCfg EncodingConfig, k Keepers, msgRouter *baseapp.MsgServiceRouter,
) *sdkmodule.SimulationManager {
	return sdkmodule.NewSimulationManager(
		// Cosmos SDK modules
		auth.NewAppModule(encCfg.Codec, k.AccountKeeper, nil, k.Subspace(authtypes.ModuleName)),
		authzmodule.NewAppModule(encCfg.Codec, k.AuthzKeeper, k.AccountKeeper, k.BankKeeper, encCfg.InterfaceRegistry),
		bank.NewAppModule(encCfg.Codec, k.BankKeeper, k.AccountKeeper, k.Subspace(banktypes.ModuleName)),
		capability.NewAppModule(encCfg.Codec, *k.CapabilityKeeper, false),
		distribution.NewAppModule(
			encCfg.Codec, k.DistributionKeeper, k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
			k.Subspace(distributiontypes.ModuleName),
		),
		evidence.NewAppModule(k.EvidenceKeeper),
		feegrantmodule.NewAppModule(encCfg.Codec, k.AccountKeeper, k.BankKeeper, k.FeeGrantKeeper, encCfg.InterfaceRegistry),
		gov.NewAppModule(encCfg.Codec, k.GovKeeper, k.AccountKeeper, k.BankKeeper, k.Subspace(govtypes.ModuleName)),
		groupmodule.NewAppModule(encCfg.Codec, k.GroupKeeper, k.AccountKeeper, k.BankKeeper, encCfg.InterfaceRegistry),
		mint.NewAppModule(encCfg.Codec, k.MintKeeper, k.AccountKeeper, nil, k.Subspace(minttypes.ModuleName)),
		nftmodule.NewAppModule(encCfg.Codec, k.NFTKeeper, k.AccountKeeper, k.BankKeeper, encCfg.InterfaceRegistry),
		params.NewAppModule(k.ParamsKeeper),
		slashing.NewAppModule(
			encCfg.Codec, k.SlashingKeeper, k.AccountKeeper, k.BankKeeper, k.StakingKeeper,
			k.Subspace(slashingtypes.ModuleName),
		),
		staking.NewAppModule(
			encCfg.Codec, k.StakingKeeper, k.AccountKeeper, k.BankKeeper, k.Subspace(stakingtypes.ModuleName),
		),

		// Cosmos IBC modules
		ibcfee.NewAppModule(k.IBCFeeKeeper),
		ibc.NewAppModule(k.IBCKeeper),
		ibctransfer.NewAppModule(k.IBCTransferKeeper),

		// Sentinel Hub modules
		custommint.NewAppModule(encCfg.Codec, k.CustomMintKeeper),
		oracle.NewAppModule(encCfg.Codec, k.OracleKeeper),
		swap.NewAppModule(encCfg.Codec, k.SwapKeeper),
		vpn.NewAppModule(encCfg.Codec, k.AccountKeeper, k.BankKeeper, k.VPNKeeper),

		// Other modules
		wasm.NewAppModule(
			encCfg.Codec, &k.WasmKeeper, k.StakingKeeper, k.AccountKeeper, k.BankKeeper, msgRouter,
			k.Subspace(wasmtypes.ModuleName),
		),
	)
}
