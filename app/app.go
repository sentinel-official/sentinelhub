package app

import (
	"encoding/json"
	"io"

	"cosmossdk.io/client/v2/autocli"
	"cosmossdk.io/core/appmodule"
	tmlog "cosmossdk.io/log"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmdb "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sentinel-official/sentinelhub/v13/app/ante"
)

const (
	appName = "Sentinel Hub"
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

type App struct {
	*baseapp.BaseApp
	sdkmodule.BasicManager
	EncodingConfig
	Keepers
	StoreKeys

	mm *sdkmodule.Manager
}

func NewApp(
	appOpts servertypes.AppOptions, db tmdb.DB, homeDir string, loadLatest bool, logger tmlog.Logger,
	traceWriter io.Writer, skipUpgradeHeights map[int64]bool, wasmOpts []wasmkeeper.Option,
	baseAppOpts ...func(*baseapp.BaseApp),
) *App {
	encCfg := DefaultEncodingConfig()

	baseApp := baseapp.NewBaseApp(appName, logger, db, encCfg.TxConfig.TxDecoder(), baseAppOpts...)
	baseApp.SetCommitMultiStoreTracer(traceWriter)
	baseApp.SetVersion(version.Version)
	baseApp.SetInterfaceRegistry(encCfg.InterfaceRegistry)

	wasmConfig, err := wasm.ReadNodeConfig(appOpts)
	if err != nil {
		panic("failed to read the wasm config: " + err.Error())
	}

	var (
		storeKeys = NewStoreKeys()
		keepers   = NewKeepers(
			baseApp, BlockedAccAddrs(), encCfg, homeDir, storeKeys, logger, ModuleAccPerms(), skipUpgradeHeights,
			wasmConfig, wasmOpts,
		)
		mm = NewModuleManager(baseApp, encCfg, keepers, baseApp.MsgServiceRouter())
		bm = NewModuleBasicManager(encCfg, mm)
	)

	app := &App{
		BaseApp:        baseApp,
		BasicManager:   bm,
		EncodingConfig: encCfg,
		Keepers:        keepers,
		StoreKeys:      storeKeys,
		mm:             mm,
	}

	configurator := sdkmodule.NewConfigurator(encCfg.Codec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	if err := app.mm.RegisterServices(configurator); err != nil {
		panic("registering services: " + err.Error())
	}

	app.MountKVStores(app.KVKeys())
	app.MountMemoryStores(app.MemoryKeys())
	app.MountTransientStores(app.TransientKeys())

	app.SetupAnteHandler(wasmConfig)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetInitChainer(app.InitChainer)
	app.SetUpgradeHandler(configurator)
	app.SetUpgradeStoreLoader()
	app.RegisterSnapshotExtensions()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit("failed to load the latest version: " + err.Error())
		}

		ctx := app.NewUncachedContext(true, tmproto.Header{})
		if err := app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit("failed to initialize the pinned codes: " + err.Error())
		}
	}

	return app
}

func (a *App) LegacyAmino() *codec.LegacyAmino {
	return a.Amino
}

func (a *App) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return a.mm.BeginBlock(ctx)
}

func (a *App) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return a.mm.EndBlock(ctx)
}

func (a *App) InitChainer(ctx sdk.Context, req *abcitypes.RequestInitChain) (*abcitypes.ResponseInitChain, error) {
	var state map[string]json.RawMessage
	if err := tmjson.Unmarshal(req.AppStateBytes, &state); err != nil {
		panic("failed to unmarshal the app state: " + err.Error())
	}

	a.UpgradeKeeper.SetModuleVersionMap(ctx, a.mm.GetVersionMap())

	return a.mm.InitGenesis(ctx, a.Codec, state)
}

func (a *App) LoadHeight(height int64) error {
	return a.LoadVersion(height)
}

func (a *App) SimulationManager() *sdkmodule.SimulationManager {
	return nil
}

func (a *App) RegisterAPIRoutes(server *api.Server, _ serverconfig.APIConfig) {
	authtx.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
	cmtservice.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
	node.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
	a.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
}

func (a *App) RegisterTxService(ctx client.Context) {
	authtx.RegisterTxService(a.GRPCQueryRouter(), ctx, a.Simulate, a.InterfaceRegistry)
}

func (a *App) RegisterTendermintService(ctx client.Context) {
	cmtservice.RegisterTendermintService(ctx, a.GRPCQueryRouter(), a.InterfaceRegistry, a.Query)
}

func (a *App) RegisterNodeService(ctx client.Context, cfg serverconfig.Config) {
	node.RegisterNodeService(ctx, a.GRPCQueryRouter(), cfg)
}

func (a *App) ModuleAccountAddrs() map[string]bool {
	addrs := make(map[string]bool)

	for v := range ModuleAccPerms() {
		addr := authtypes.NewModuleAddress(v)
		addrs[addr.String()] = true
	}

	return addrs
}

func (a *App) SetupAnteHandler(wasmConfig wasmtypes.NodeConfig) {
	handler, err := ante.NewHandler(
		ante.HandlerOptions{
			HandlerOptions: authante.HandlerOptions{
				AccountKeeper:   a.AccountKeeper,
				BankKeeper:      a.BankKeeper,
				FeegrantKeeper:  a.FeeGrantKeeper,
				SignModeHandler: a.TxConfig.SignModeHandler(),
				SigGasConsumer:  authante.DefaultSigVerificationGasConsumer,
			},
			TxCounterStoreKey: a.KVStoreService(wasmtypes.StoreKey),
			IBCKeeper:         a.IBCKeeper,
			WasmConfig:        wasmConfig,
		},
	)
	if err != nil {
		panic("failed to create the ante handler: " + err.Error())
	}

	a.SetAnteHandler(handler)
}

func (a *App) SetUpgradeStoreLoader() {
	upgradeInfo, err := a.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic("failed to read the upgrade info from disk: " + err.Error())
	}

	if a.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	if upgradeInfo.Name == UpgradeName {
		a.SetStoreLoader(
			upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, StoreUpgrades),
		)
	}
}

func (a *App) SetUpgradeHandler(configurator sdkmodule.Configurator) {
	a.UpgradeKeeper.SetUpgradeHandler(
		UpgradeName,
		UpgradeHandler(a.mm, configurator),
	)
}

func (a *App) RegisterSnapshotExtensions() {
	if m := a.SnapshotManager(); m != nil {
		if err := m.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(a.CommitMultiStore(), &a.WasmKeeper),
		); err != nil {
			panic("failed to register the snapshot extension: " + err.Error())
		}
	}
}

func (a *App) AutoCliOpts() autocli.AppOptions {
	modules := make(map[string]appmodule.AppModule, 0)

	for _, m := range a.mm.Modules {
		if moduleWithName, ok := m.(sdkmodule.HasName); ok {
			moduleName := moduleWithName.Name()
			if appModule, ok := moduleWithName.(appmodule.AppModule); ok {
				modules[moduleName] = appModule
			}
		}
	}

	return autocli.AppOptions{
		Modules:               modules,
		AddressCodec:          authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		ValidatorAddressCodec: authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		ConsensusAddressCodec: authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	}
}
