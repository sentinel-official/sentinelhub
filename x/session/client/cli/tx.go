package cli

import (
	"strconv"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

func txCancelSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-session [id]",
		Short: "Cancel an active session",
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

			msg := v3.NewMsgCancelSessionRequest(
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

func txUpdateSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-session [id] [download-bytes] [upload-bytes] [duration]",
		Short: "Update the details of an existing session",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			downloadBytes, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			uploadBytes, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			duration, err := time.ParseDuration(args[3])
			if err != nil {
				return err
			}

			signature, err := GetSignature(cmd.Flags())
			if err != nil {
				return err
			}

			msg := v3.NewMsgUpdateSessionRequest(
				ctx.FromAddress.Bytes(),
				id,
				sdkmath.NewInt(downloadBytes),
				sdkmath.NewInt(uploadBytes),
				duration,
				signature,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagSignature, "", "Provide the Base64-encoded client signature to authorize")

	return cmd
}
