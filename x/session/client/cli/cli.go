package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		querySession(),
		querySessions(),
		queryParams(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session module sub-commands",
	}

	cmd.AddCommand(
		txCancelSession(),
		txUpdateSession(),
	)

	return []*cobra.Command{cmd}
}
