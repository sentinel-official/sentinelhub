package app

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

func (a *App) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string, modulesToExport []string,
) (servertypes.ExportedApp, error) {
	return servertypes.ExportedApp{}, nil
}
