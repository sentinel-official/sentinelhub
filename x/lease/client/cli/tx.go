package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types/v1"
)

func txEndLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end-lease [id]",
		Short: "End an existing lease",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := v1.NewMsgEndLeaseRequest(
				ctx.FromAddress.Bytes(),
				id,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txRenewLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew-lease [id] [hours]",
		Short: "Renew an existing lease for a specified duration",
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

			hours, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			maxPrice, err := GetMaxPrice(cmd.Flags())
			if err != nil {
				return err
			}

			msg := v1.NewMsgRenewLeaseRequest(
				ctx.FromAddress.Bytes(),
				id,
				hours,
				maxPrice,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagMaxPrice, "", "Specify the maximum hourly price for the lease")

	return cmd
}

func txStartLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-lease [node-addr] [hours]",
		Short: "Start a lease with a node for the specified duration",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nodeAddr, err := base.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			hours, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			maxPrice, err := GetMaxPrice(cmd.Flags())
			if err != nil {
				return err
			}

			renewalPricePolicy, err := GetRenewalPricePolicy(cmd.Flags())
			if err != nil {
				return err
			}

			msg := v1.NewMsgStartLeaseRequest(
				ctx.FromAddress.Bytes(),
				nodeAddr,
				hours,
				maxPrice,
				renewalPricePolicy,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagMaxPrice, "", "Specify the maximum hourly price for the lease")
	cmd.Flags().String(flagRenewalPricePolicy, "", "Specify the lease renewal price policy")

	return cmd
}

func txUpdateLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-lease [id]",
		Short: "Update the details of an existing lease",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			renewalPricePolicy, err := GetRenewalPricePolicy(cmd.Flags())
			if err != nil {
				return err
			}

			msg := v1.NewMsgUpdateLeaseRequest(
				ctx.FromAddress.Bytes(),
				id,
				renewalPricePolicy,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagRenewalPricePolicy, "", "Specify the lease renewal price policy")

	return cmd
}
