package cli

import (
	"github.com/spf13/cobra"

	depositcli "github.com/sentinel-official/sentinelhub/v12/x/deposit/client/cli"
	leasecli "github.com/sentinel-official/sentinelhub/v12/x/lease/client/cli"
	nodecli "github.com/sentinel-official/sentinelhub/v12/x/node/client/cli"
	plancli "github.com/sentinel-official/sentinelhub/v12/x/plan/client/cli"
	providercli "github.com/sentinel-official/sentinelhub/v12/x/provider/client/cli"
	sessioncli "github.com/sentinel-official/sentinelhub/v12/x/session/client/cli"
	subscriptioncli "github.com/sentinel-official/sentinelhub/v12/x/subscription/client/cli"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "Querying commands for the VPN module",
	}

	cmd.AddCommand(depositcli.GetQueryCommands()...)
	cmd.AddCommand(leasecli.GetQueryCommands()...)
	cmd.AddCommand(nodecli.GetQueryCommands()...)
	cmd.AddCommand(plancli.GetQueryCommands()...)
	cmd.AddCommand(providercli.GetQueryCommands()...)
	cmd.AddCommand(sessioncli.GetQueryCommands()...)
	cmd.AddCommand(subscriptioncli.GetQueryCommands()...)

	return cmd
}

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "VPN transactions subcommands",
	}

	cmd.AddCommand(leasecli.GetTxCommands()...)
	cmd.AddCommand(nodecli.GetTxCommands()...)
	cmd.AddCommand(plancli.GetTxCommands()...)
	cmd.AddCommand(providercli.GetTxCommands()...)
	cmd.AddCommand(sessioncli.GetTxCommands()...)
	cmd.AddCommand(subscriptioncli.GetTxCommands()...)

	return cmd
}
