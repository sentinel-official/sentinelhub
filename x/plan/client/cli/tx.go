package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types/v3"
)

func txCreatePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-plan [gigabytes] [hours]",
		Short: "Create a new subscription plan with gigabytes, hours and pricing details",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gigabytes, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			hours, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			prices, err := GetPrices(cmd.Flags())
			if err != nil {
				return err
			}

			msg := v3.NewMsgCreatePlanRequest(
				ctx.FromAddress.Bytes(),
				gigabytes,
				hours,
				prices,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagPrices, "", "specify the list of prices (e.g., 1000token)")

	return cmd
}

func txLinkNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-node [id] [node-addr]",
		Short: "Link a node to a subscription plan",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			addr, err := base.NodeAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := v3.NewMsgLinkNodeRequest(
				ctx.FromAddress.Bytes(),
				id,
				addr,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txUnlinkNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlink-node [id] [node-addr]",
		Short: "Unlink a node from a subscription plan",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			addr, err := base.NodeAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := v3.NewMsgUnlinkNodeRequest(
				ctx.FromAddress.Bytes(),
				id,
				addr,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txUpdatePlanStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-plan-status [id] [status]",
		Short: "Update the status of an existing subscription plan",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := v3.NewMsgUpdatePlanStatusRequest(
				ctx.FromAddress.Bytes(),
				id,
				v1base.StatusFromString(args[1]),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
