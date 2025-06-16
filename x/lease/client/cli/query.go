package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/x/lease/types/v1"
)

func queryLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease [id]",
		Short: "Query a lease by ID",
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

			qc := v1.NewQueryServiceClient(ctx)

			res, err := qc.QueryLease(
				cmd.Context(),
				v1.NewQueryLeaseRequest(id),
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

func queryLeases() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leases",
		Short: "Query all leases with optional filters and pagination",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nodeAddr, err := base.NodeAddrFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			provAddr, err := base.ProvAddrFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			qc := v1.NewQueryServiceClient(ctx)

			switch {
			case nodeAddr != nil:
				res, err := qc.QueryLeasesForNode(
					cmd.Context(),
					v1.NewQueryLeasesForNodeRequest(nodeAddr, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			case provAddr != nil:
				res, err := qc.QueryLeasesForProvider(
					cmd.Context(),
					v1.NewQueryLeasesForProviderRequest(provAddr, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			default:
				res, err := qc.QueryLeases(
					cmd.Context(),
					v1.NewQueryLeasesRequest(pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "leases")
	cmd.Flags().String(base.FlagNodeAddr, "", "filter the leases by node address")
	cmd.Flags().String(base.FlagProvAddr, "", "filter the leases by provider address")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease-params",
		Short: "Query the lease module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			qc := v1.NewQueryServiceClient(ctx)

			res, err := qc.QueryParams(
				cmd.Context(),
				v1.NewQueryParamsRequest(),
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
