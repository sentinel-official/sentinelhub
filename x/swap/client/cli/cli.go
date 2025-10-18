package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap",
		Short: "Querying commands for the swap module",
	}

	cmd.AddCommand(
		querySwap(),
		querySwaps(),
		queryParams(),
	)

	return cmd
}

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap",
		Short: "Swap module sub-commands",
	}

	cmd.AddCommand(
		txSwap(),
	)

	return cmd
}
