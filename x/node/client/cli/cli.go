package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryNode(),
		queryNodes(),
		queryParams(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Node module sub-commands",
	}

	cmd.AddCommand(
		txRegisterNode(),
		txUpdateNodeDetails(),
		txUpdateNodeStatus(),
		txStartSession(),
	)

	return []*cobra.Command{cmd}
}
