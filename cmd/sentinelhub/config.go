package main

import (
	"os"
	"path/filepath"
	"time"

	tmcfg "github.com/cometbft/cometbft/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/spf13/viper"
)

const flagSkipOverwriteConfig = "skip-overwrite-config"

// applyRecommendedValues sets default values for specific configuration types.
func applyRecommendedValues(cfg any) {
	switch c := cfg.(type) {
	case *serverconfig.Config:
		c.MinGasPrices = "0.1udvpn"
	case *tmcfg.Config:
		c.Consensus.TimeoutCommit = 3 * time.Second
	}
}

// initAppConfig initializes the application configuration with defaults.
func initAppConfig() (string, any) {
	cfg := serverconfig.DefaultConfig()
	cfgTemplate := serverconfig.DefaultConfigTemplate

	applyRecommendedValues(cfg)

	return cfgTemplate, cfg
}

// initTendermintConfig initializes the Tendermint configuration with defaults.
func initTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()
	applyRecommendedValues(cfg)

	return cfg
}

// overwriteConfig reads, updates, and writes a configuration file.
func overwriteConfig(name string, cfg any, write func(string, any) error) error {
	homeDir := viper.GetString(flags.FlagHome)
	cfgPath := filepath.Join(homeDir, "config", name)

	if _, err := os.Stat(cfgPath); err != nil {
		return nil
	}

	v := viper.New()
	v.SetConfigFile(cfgPath)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(cfg); err != nil {
		return err
	}

	applyRecommendedValues(cfg)

	return write(cfgPath, cfg)
}

// overwriteAppConfig updates and writes the app configuration.
func overwriteAppConfig() error {
	return overwriteConfig("app.toml", serverconfig.DefaultConfig(), func(cfgPath string, cfg any) error {
		serverconfig.WriteConfigFile(cfgPath, cfg.(*serverconfig.Config))

		return nil
	})
}

// overwriteTendermintConfig updates and writes the Tendermint configuration.
func overwriteTendermintConfig() error {
	return overwriteConfig("config.toml", tmcfg.DefaultConfig(), func(cfgPath string, cfg any) error {
		tmcfg.WriteConfigFile(cfgPath, cfg.(*tmcfg.Config))

		return nil
	})
}
