package app

import (
	"encoding/json"
	"io"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	tmdb "github.com/cometbft/cometbft-db"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	tmlog "github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

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
	EncodingConfig
	Keepers
	StoreKeys

	mm *sdkmodule.Manager
	sm *sdkmodule.SimulationManager
}

func NewApp(
	appOpts servertypes.AppOptions, bech32Prefix string, db tmdb.DB, encCfg EncodingConfig, homeDir string,
	invCheckPeriod uint, loadLatest bool, logger tmlog.Logger, skipGenesisInvariants bool,
	traceWriter io.Writer, version string, skipUpgradeHeights map[int64]bool, wasmOpts []wasmkeeper.Option,
	baseAppOpts ...func(*baseapp.BaseApp),
) *App {
	baseApp := baseapp.NewBaseApp(appName, logger, db, encCfg.TxConfig.TxDecoder(), baseAppOpts...)
	baseApp.SetCommitMultiStoreTracer(traceWriter)
	baseApp.SetVersion(version)
	baseApp.SetInterfaceRegistry(encCfg.InterfaceRegistry)

	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("failed to read the wasm config: " + err.Error())
	}

	var (
		storeKeys = NewStoreKeys()
		keepers   = NewKeepers(
			baseApp, bech32Prefix, BlockedAccAddrs(), encCfg, homeDir, invCheckPeriod, storeKeys,
			ModuleAccPerms(), skipUpgradeHeights, wasmConfig, wasmOpts,
		)
		mm = NewModuleManager(baseApp.DeliverTx, encCfg, keepers, baseApp.MsgServiceRouter(), skipGenesisInvariants)
		sm = NewSimulationManager(encCfg, keepers, baseApp.MsgServiceRouter())
	)

	app := &App{
		BaseApp:        baseApp,
		EncodingConfig: encCfg,
		Keepers:        keepers,
		StoreKeys:      storeKeys,
		mm:             mm,
		sm:             sm,
	}

	app.mm.RegisterInvariants(keepers.CrisisKeeper)

	configurator := sdkmodule.NewConfigurator(encCfg.Codec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(configurator)

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
		if err = app.LoadLatestVersion(); err != nil {
			tmos.Exit("failed to load the latest version: " + err.Error())
		}

		ctx := app.NewUncachedContext(true, tmproto.Header{})
		if err = app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit("failed to initialize the pinned codes: " + err.Error())
		}
	}

	return app
}

func (a *App) LegacyAmino() *codec.LegacyAmino {
	return a.Amino
}

func (a *App) BeginBlocker(ctx sdk.Context, req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return a.mm.BeginBlock(ctx, req)
}

func (a *App) EndBlocker(ctx sdk.Context, req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return a.mm.EndBlock(ctx, req)
}

func (a *App) InitChainer(ctx sdk.Context, req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
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
	return a.sm
}

func (a *App) RegisterAPIRoutes(server *api.Server, _ serverconfig.APIConfig) {
	authtx.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
	tmservice.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
	node.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
	ModuleBasics.RegisterGRPCGatewayRoutes(server.ClientCtx, server.GRPCGatewayRouter)
}

func (a *App) RegisterTxService(ctx client.Context) {
	authtx.RegisterTxService(a.GRPCQueryRouter(), ctx, a.Simulate, a.InterfaceRegistry)
}

func (a *App) RegisterTendermintService(ctx client.Context) {
	tmservice.RegisterTendermintService(ctx, a.GRPCQueryRouter(), a.InterfaceRegistry, a.Query)
}

func (a *App) RegisterNodeService(ctx client.Context) {
	node.RegisterNodeService(ctx, a.GRPCQueryRouter())
}

func (a *App) ModuleAccountAddrs() map[string]bool {
	addrs := make(map[string]bool)

	for v := range ModuleAccPerms() {
		addr := authtypes.NewModuleAddress(v)
		addrs[addr.String()] = true
	}

	return addrs
}

func (a *App) SetupAnteHandler(wasmConfig wasmtypes.WasmConfig) {
	handler, err := ante.NewHandler(
		ante.HandlerOptions{
			HandlerOptions: authante.HandlerOptions{
				AccountKeeper:   a.AccountKeeper,
				BankKeeper:      a.BankKeeper,
				FeegrantKeeper:  a.FeeGrantKeeper,
				SignModeHandler: a.TxConfig.SignModeHandler(),
				SigGasConsumer:  authante.DefaultSigVerificationGasConsumer,
			},
			TxCounterStoreKey: a.KV(wasmtypes.StoreKey),
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
		UpgradeHandler(a.Codec, a.mm, configurator, a.Keepers, a.KV(upgradetypes.StoreKey)),
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
