package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle",
		Short: "Querying commands for the Oracle module",
	}

	cmd.AddCommand(
		queryAsset(),
		queryAssets(),
		queryParams(),
	)

	return cmd
}

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle",
		Short: "Oracle transactions subcommands",
	}

	return cmd
}
