package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func querySession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session [id]",
		Short: "Query a session by ID",
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

			res, err := qc.QuerySession(
				cmd.Context(),
				v3.NewQuerySessionRequest(id),
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

func querySessions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sessions",
		Short: "Query all sessions with optional filters and pagination",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			accAddr, err := base.AccAddrFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			nodeAddr, err := base.NodeAddrFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			id, err := base.SubscriptionIDFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			qc := v3.NewQueryServiceClient(ctx)

			switch {
			case id != 0 && accAddr != nil:
				res, err := qc.QuerySessionsForAllocation(
					cmd.Context(),
					v3.NewQuerySessionsForAllocationRequest(id, accAddr, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			case id != 0:
				res, err := qc.QuerySessionsForSubscription(
					cmd.Context(),
					v3.NewQuerySessionsForSubscriptionRequest(id, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			case accAddr != nil:
				res, err := qc.QuerySessionsForAccount(
					cmd.Context(),
					v3.NewQuerySessionsForAccountRequest(accAddr, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			case nodeAddr != nil:
				res, err := qc.QuerySessionsForNode(
					cmd.Context(),
					v3.NewQuerySessionsForNodeRequest(nodeAddr, pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			default:
				res, err := qc.QuerySessions(
					cmd.Context(),
					v3.NewQuerySessionsRequest(pagination),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "sessions")
	cmd.Flags().String(base.FlagAccAddr, "", "filter the sessions by account address")
	cmd.Flags().String(base.FlagNodeAddr, "", "filter the sessions by node address")
	cmd.Flags().Uint64(base.FlagSubscriptionID, 0, "filter the sessions by subscription ID")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session-params",
		Short: "Query the session module parameters",
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
