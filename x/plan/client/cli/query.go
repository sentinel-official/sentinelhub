package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	"github.com/sentinel-official/sentinelhub/v13/x/plan/types/v3"
)

func queryPlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan [id]",
		Short: "Query a plan by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			res, err := qc.QueryPlan(
				cmd.Context(),
				v3.NewQueryPlanRequest(id),
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

func queryPlans() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plans",
		Short: "Query all plans with optional filters and pagination",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			provAddr, err := base.ProvAddrFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			status, err := base.StatusFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			if provAddr != nil {
				res, err := qc.QueryPlansForProvider(
					cmd.Context(),
					v3.NewQueryPlansForProviderRequest(provAddr, status, pagination),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QueryPlans(
				cmd.Context(),
				v3.NewQueryPlansRequest(status, pagination),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "plans")
	cmd.Flags().String(base.FlagStatus, "", "filter the plans by status (active|inactive)")
	cmd.Flags().String(base.FlagProvAddr, "", "filter the plans by provider address")

	return cmd
}
