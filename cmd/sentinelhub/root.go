package main

import (
	"os"

	"github.com/CosmWasm/wasmd/x/wasm"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client"
	clientconfig "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/hub/v12/app"
)

func moduleInitFlags(cmd *cobra.Command) {
	crisis.AddModuleInitFlags(cmd)
	wasm.AddModuleInitFlags(cmd)
	cmd.Flags().Bool(flagOverwriteConfigWithDefaults, true, "If set to true, recommended default values will overwrite any existing settings in config.toml and app.toml.")
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcli.GetAccountCmd(),
		authcli.QueryTxCmd(),
		authcli.QueryTxsByEventsCmd(),
		rpc.BlockCommand(),
		rpc.ValidatorCommand(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcli.GetBroadcastCommand(),
		authcli.GetDecodeCommand(),
		authcli.GetEncodeCommand(),
		authcli.GetMultiSignBatchCmd(),
		authcli.GetMultiSignCommand(),
		authcli.GetSignBatchCommand(),
		authcli.GetSignCommand(),
		authcli.GetValidateSignaturesCommand(),
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func NewRootCmd(homeDir string) *cobra.Command {
	encCfg := app.DefaultEncodingConfig()
	cmd := &cobra.Command{
		Use:          "sentinelhub",
		Short:        "Sentinel Hub application",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			if viper.GetBool(flagOverwriteConfigWithDefaults) {
				if err := overwriteTendermintConfig(); err != nil {
					return err
				}
				if err := overwriteAppConfig(); err != nil {
					return err
				}
			}

			clientCtx := client.Context{}.
				WithAccountRetriever(authtypes.AccountRetriever{}).
				WithCodec(encCfg.Codec).
				WithHomeDir(homeDir).
				WithInput(os.Stdin).
				WithInterfaceRegistry(encCfg.InterfaceRegistry).
				WithLegacyAmino(encCfg.Amino).
				WithTxConfig(encCfg.TxConfig).
				WithViper("")

			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			clientCtx, err = clientconfig.ReadFromClientConfig(clientCtx)
			if err != nil {
				return err
			}

			if err = client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			appConfigTemplate, appConfig := initAppConfig()
			tmConfig := initTendermintConfig()

			return server.InterceptConfigsPreRunHandler(cmd, appConfigTemplate, appConfig, tmConfig)
		},
	}

	cmd.AddCommand(
		clientconfig.Cmd(),
		debug.Cmd(),
		genutilcli.InitCmd(app.ModuleBasics, homeDir),
		genutilcli.GenesisCoreCommand(encCfg.TxConfig, app.ModuleBasics, homeDir),
		keys.Commands(homeDir),
		queryCommand(),
		rpc.StatusCommand(),
		tmcli.NewCompletionCmd(cmd, true),
		txCommand(),
	)

	creator := appCreator{encCfg: encCfg}
	server.AddCommands(cmd, homeDir, creator.NewApp, creator.AppExport, moduleInitFlags)

	return cmd
}
