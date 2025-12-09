package app

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
)

const (
	UpgradeName = "v13_0_0"
)

var (
	StoreUpgrades = &storetypes.StoreUpgrades{
		Added: []string{},
	}
)

func UpgradeHandler(mm *sdkmodule.Manager, configurator sdkmodule.Configurator) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM sdkmodule.VersionMap) (sdkmodule.VersionMap, error) {
		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		return newVM, nil
	}
}
