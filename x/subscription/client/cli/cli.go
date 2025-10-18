package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryAllocation(),
		queryAllocations(),
		querySubscription(),
		querySubscriptions(),
		queryParams(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Subscription module sub-commands",
	}

	cmd.AddCommand(
		txCancelSubscription(),
		txRenewSubscription(),
		txShareSubscription(),
		txStartSubscription(),
		txUpdateSubscription(),
		txStartSession(),
	)

	return []*cobra.Command{cmd}
}
