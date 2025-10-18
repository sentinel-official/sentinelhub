package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinelhub/v12/x/oracle/types/v1"
)

func queryAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "asset [denom]",
		Short: "Query an asset by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			qc := v1.NewQueryServiceClient(ctx)

			res, err := qc.QueryAsset(
				cmd.Context(),
				v1.NewQueryAssetRequest(args[0]),
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

func queryAssets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assets",
		Short: "Query the list of all assets",
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

			res, err := qc.QueryAssets(
				cmd.Context(),
				v1.NewQueryAssetsRequest(pagination),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "assets")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the oracle module parameters",
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
