package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/sentinelhub/v13/types"
	"github.com/sentinel-official/sentinelhub/v13/x/node/types/v3"
)

func queryNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node [node-addr]",
		Short: "Query a node by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := base.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			res, err := qc.QueryNode(
				cmd.Context(),
				v3.NewQueryNodeRequest(addr),
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

func queryNodes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Query all nodes with optional filters and pagination",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := base.PlanIDFromFlags(cmd.Flags())
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

			switch {
			case id != 0:
				res, err := qc.QueryNodesForPlan(
					cmd.Context(),
					v3.NewQueryNodesForPlanRequest(id, status, pagination),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			default:
				res, err := qc.QueryNodes(
					cmd.Context(),
					v3.NewQueryNodesRequest(status, pagination),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nodes")
	cmd.Flags().String(base.FlagStatus, "", "filter the nodes by status (active|inactive)")
	cmd.Flags().Uint64(base.FlagPlanID, 0, "filter the nodes by subscription plan ID")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node-params",
		Short: "Query the node module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			res, err := qc.QueryParams(
				cmd.Context(),
				v3.NewQueryParamsRequest(),
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
