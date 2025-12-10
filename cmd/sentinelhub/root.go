package main

import (
	"fmt"
	"os"

	"cosmossdk.io/log"
	confixcmd "cosmossdk.io/tools/confix/cmd"
	"github.com/CosmWasm/wasmd/x/wasm"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	tmdb "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	clientconfig "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/pruning"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/snapshot"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/sentinelhub/v13/app"
)

func moduleInitFlags(cmd *cobra.Command) {
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

	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func NewRootCmd(homeDir string) *cobra.Command {
	ac := appCreator{}

	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()

	tmpOpts := viper.New()
	tmpOpts.Set(flags.FlagHome, tmpDir)

	tmpApp := app.NewApp(tmpOpts, tmdb.NewMemDB(), tmpDir, true, log.NewNopLogger(), nil, nil, nil)

	defer func() {
		if err := tmpApp.Close(); err != nil {
			panic(err)
		}
	}()

	clientCtx := client.Context{}.
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithCodec(tmpApp.Codec).
		WithHomeDir(homeDir).
		WithInput(os.Stdin).
		WithInterfaceRegistry(tmpApp.InterfaceRegistry).
		WithLegacyAmino(tmpApp.Amino).
		WithTxConfig(tmpApp.TxConfig).
		WithViper("")

	cmd := &cobra.Command{
		Use:          "sentinelhub",
		Short:        "Sentinel Hub application",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			clientCtx = clientCtx.WithCmdContext(cmd.Context())

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
		confixcmd.ConfigCommand(),
		debug.Cmd(),
		genutilcli.GenesisCoreCommand(tmpApp.TxConfig, tmpApp.BasicManager, homeDir),
		genutilcli.InitCmd(tmpApp.BasicManager, homeDir),
		keys.Commands(),
		pruning.Cmd(ac.NewApp, homeDir),
		queryCommand(),
		server.StatusCommand(),
		snapshot.Cmd(ac.NewApp),
		tmcli.NewCompletionCmd(cmd, true),
		txCommand(),
	)

	server.AddCommands(cmd, homeDir, ac.NewApp, ac.AppExport, moduleInitFlags)

	autoCliOpts := tmpApp.AutoCliOpts()
	autoCliOpts.AddressCodec = address.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	autoCliOpts.ValidatorAddressCodec = address.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix())
	autoCliOpts.ConsensusAddressCodec = address.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix())
	autoCliOpts.ClientCtx = clientCtx

	if err := autoCliOpts.EnhanceRootCommand(cmd); err != nil {
		panic(err)
	}

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
