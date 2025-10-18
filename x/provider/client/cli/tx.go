package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/provider/types/v3"
)

func txRegisterProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-provider [name]",
		Short: "Register a new provider with a name and optional details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			identity, err := cmd.Flags().GetString(flagIdentity)
			if err != nil {
				return err
			}

			website, err := cmd.Flags().GetString(flagWebsite)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(flagDescription)
			if err != nil {
				return err
			}

			msg := v3.NewMsgRegisterProviderRequest(
				ctx.FromAddress.Bytes(),
				args[0],
				identity,
				website,
				description,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagIdentity, "", "unique identity of the provider")
	cmd.Flags().String(flagWebsite, "", "official website URL of the provider")
	cmd.Flags().String(flagDescription, "", "brief description of the provider's services or offerings")

	return cmd
}

func txUpdateProviderDetails() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-provider-details",
		Short: "Update the details of an existing provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString(flagName)
			if err != nil {
				return err
			}

			identity, err := cmd.Flags().GetString(flagIdentity)
			if err != nil {
				return err
			}

			website, err := cmd.Flags().GetString(flagWebsite)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(flagDescription)
			if err != nil {
				return err
			}

			msg := v3.NewMsgUpdateProviderDetailsRequest(
				ctx.FromAddress.Bytes(),
				name,
				identity,
				website,
				description,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagName, "", "name of the provider")
	cmd.Flags().String(flagIdentity, "", "unique identity of the provider")
	cmd.Flags().String(flagWebsite, "", "official website URL of the provider")
	cmd.Flags().String(flagDescription, "", "brief description of the provider's services or offerings")

	return cmd
}

func txUpdateProviderStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-provider-status [status]",
		Short: "Update the operational status of an existing provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v3.NewMsgUpdateProviderStatusRequest(
				ctx.FromAddress.Bytes(),
				v1base.StatusFromString(args[0]),
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
