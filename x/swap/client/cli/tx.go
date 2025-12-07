package cli

import (
	"encoding/hex"
	"strconv"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinelhub/v13/x/swap/types"
	"github.com/sentinel-official/sentinelhub/v13/x/swap/types/v1"
)

func txSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [tx-hash] [receiver-addr] [amount]",
		Short: "Swap coins from SENT to DVPN",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txHash, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			receiver, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			amount, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := v1.NewMsgSwapRequest(
				ctx.FromAddress.Bytes(),
				types.BytesToHash(txHash),
				receiver,
				sdkmath.NewInt(amount),
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
