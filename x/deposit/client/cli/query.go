package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinelhub/v12/x/deposit/types/v1"
)

func queryDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [acc-addr]",
		Short: "Query a deposit for an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			qc := v1.NewQueryServiceClient(ctx)

			res, err := qc.QueryDeposit(
				cmd.Context(),
				v1.NewQueryDepositRequest(addr),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryDeposits() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits",
		Short: "Query all deposits with optional pagination",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			qc := v1.NewQueryServiceClient(ctx)

			res, err := qc.QueryDeposits(
				cmd.Context(),
				v1.NewQueryDepositsRequest(pagination),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "deposits")

	return cmd
}
