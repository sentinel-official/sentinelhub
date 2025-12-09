package main

import (
	"io"

	"cosmossdk.io/log"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	tmdb "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"

	"github.com/sentinel-official/sentinelhub/v13/app"
	base "github.com/sentinel-official/sentinelhub/v13/types"
)

type appCreator struct {
	encCfg app.EncodingConfig
}

func (ac appCreator) NewApp(
	logger log.Logger, db tmdb.DB, traceWriter io.Writer, appOpts servertypes.AppOptions,
) servertypes.Application {
	baseAppOpts := server.DefaultBaseappOptions(appOpts)

	skipUpgradeHeights := make(map[int64]bool)
	for _, height := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(height)] = true
	}

	var wasmOpts []wasmkeeper.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return app.NewApp(
		appOpts, base.Bech32MainPrefix, db, ac.encCfg, cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), true, logger,
		cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants)), traceWriter, version.Version,
		skipUpgradeHeights, wasmOpts, baseAppOpts...,
	)
}

func (ac appCreator) AppExport(
	logger log.Logger, db tmdb.DB, traceWriter io.Writer, height int64, forZeroHeight bool, jailWhitelist []string,
	appOpts servertypes.AppOptions, modulesToExport []string,
) (servertypes.ExportedApp, error) {
	v := app.NewApp(
		appOpts, base.Bech32MainPrefix, db, ac.encCfg, cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), height == -1, logger,
		cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants)), traceWriter, version.Version,
		map[int64]bool{}, nil,
	)

	if height != -1 {
		if err := v.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return v.ExportAppStateAndValidators(forZeroHeight, jailWhitelist, modulesToExport)
}
