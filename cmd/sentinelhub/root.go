package main

import (
	"fmt"
	"os"

	confixcmd "cosmossdk.io/tools/confix/cmd"
	"github.com/CosmWasm/wasmd/x/wasm"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client"
	clientconfig "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/pruning"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/snapshot"
	"github.com/cosmos/cosmos-sdk/server"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/sentinelhub/v13/app"
)

func moduleInitFlags(cmd *cobra.Command) {
	crisis.AddModuleInitFlags(cmd)
	wasm.AddModuleInitFlags(cmd)
	cmd.Flags().Bool(flagSkipOverwriteConfig, false, "Skip overwriting config with recommended values")
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
		rpc.ValidatorCommand(),
		server.QueryBlocksCmd(),
		server.QueryBlockCmd(),
		server.QueryBlockResultsCmd(),
		authcli.QueryTxsByEventsCmd(),
		authcli.QueryTxCmd(),
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

	creator := appCreator{encCfg: encCfg}

	cmd.AddCommand(
		confixcmd.ConfigCommand(),
		debug.Cmd(),
		genutilcli.GenesisCoreCommand(encCfg.TxConfig, app.ModuleBasics, homeDir),
		genutilcli.InitCmd(app.ModuleBasics, homeDir),
		keys.Commands(),
		pruning.Cmd(creator.NewApp, homeDir),
		queryCommand(),
		server.StatusCommand(),
		snapshot.Cmd(creator.NewApp),
		tmcli.NewCompletionCmd(cmd, true),
		txCommand(),
	)

	server.AddCommands(cmd, homeDir, creator.NewApp, creator.AppExport, moduleInitFlags)

	startCmd, _, err := cmd.Find([]string{"start"})
	if err != nil {
		panic(fmt.Errorf("start command does not exist: %w", err))
	}

	startCmdRunE := startCmd.RunE
	startCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if !viper.GetBool(flagSkipOverwriteConfig) {
			if err := overwriteTendermintConfig(); err != nil {
				return fmt.Errorf("overwriting tendermint config: %w", err)
			}

			if err := overwriteAppConfig(); err != nil {
				return fmt.Errorf("overwriting app config: %w", err)
			}
		}

		return startCmdRunE(cmd, args)
	}

	return cmd
}
