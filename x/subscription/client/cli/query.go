package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func querySubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription [id]",
		Short: "Query a subscription by ID",
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

			res, err := qc.QuerySubscription(
				cmd.Context(),
				v3.NewQuerySubscriptionRequest(id),
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

func querySubscriptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Query all subscriptions with optional filters and pagination",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			accAddr, err := base.AccAddrFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			id, err := base.PlanIDFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			switch {
			case accAddr != nil:
				res, err := qc.QuerySubscriptionsForAccount(
					cmd.Context(),
					v3.NewQuerySubscriptionsForAccountRequest(accAddr, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			case id != 0:
				res, err := qc.QuerySubscriptionsForPlan(
					cmd.Context(),
					v3.NewQuerySubscriptionsForPlanRequest(id, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			default:
				res, err := qc.QuerySubscriptions(
					cmd.Context(),
					v3.NewQuerySubscriptionsRequest(pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "subscriptions")
	cmd.Flags().String(base.FlagAccAddr, "", "filter the subscriptions by account address")
	cmd.Flags().Uint64(base.FlagPlanID, 0, "filter the subscriptions by subscription plan ID")

	return cmd
}

func queryAllocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allocation [id] [acc-addr]",
		Short: "Query a allocation by subscription ID and account address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			qc := v2.NewQueryServiceClient(ctx)

			res, err := qc.QueryAllocation(
				cmd.Context(),
				v2.NewQueryAllocationRequest(id, addr),
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

func queryAllocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allocations [id]",
		Short: "Query all allocations of a subscription with optional filters and pagination",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			qc := v2.NewQueryServiceClient(ctx)

			res, err := qc.QueryAllocations(
				cmd.Context(),
				v2.NewQueryAllocationsRequest(id, pagination),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "allocations")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription-params",
		Short: "Query the subscription module parameters",
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
