package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryDeposit(),
		queryDeposits(),
	}
}
