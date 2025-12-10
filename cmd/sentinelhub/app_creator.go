package main

import (
	"io"

	"cosmossdk.io/log"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	tmdb "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"

	"github.com/sentinel-official/sentinelhub/v13/app"
)

type appCreator struct{}

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
		appOpts, db, cast.ToString(appOpts.Get(flags.FlagHome)), true, logger, traceWriter,
		skipUpgradeHeights, wasmOpts, baseAppOpts...,
	)
}

func (ac appCreator) AppExport(
	logger log.Logger, db tmdb.DB, traceWriter io.Writer, height int64, forZeroHeight bool, jailWhitelist []string,
	appOpts servertypes.AppOptions, modulesToExport []string,
) (servertypes.ExportedApp, error) {
	v := app.NewApp(
		appOpts, db, cast.ToString(appOpts.Get(flags.FlagHome)), height == -1, logger, traceWriter,
		map[int64]bool{}, nil,
	)

	defer func() {
		if err := v.Close(); err != nil {
			panic(err)
		}
	}()

	if height != -1 {
		if err := v.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return v.ExportAppStateAndValidators(forZeroHeight, jailWhitelist, modulesToExport)
}
