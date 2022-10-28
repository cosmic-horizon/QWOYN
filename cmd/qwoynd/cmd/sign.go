package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

// GetSignBytesCommand returns the string sign command.
func GetSignBytesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign_bytes [sign_data_string]",
		Short: "Sign a string with an account",
		Long:  `Sign a string with an account.`,
		RunE:  makeSignCmd(),
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().String(flags.FlagOutputDocument, "", "The document will be written to the given file instead of STDOUT")
	cmd.Flags().String(flags.FlagChainID, "", "The network chain ID")
	flags.AddTxFlagsToCmd(cmd)

	cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func makeSignCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) (err error) {
		var clientCtx client.Context

		clientCtx, err = client.GetClientTxContext(cmd)
		if err != nil {
			return err
		}

		from, _ := cmd.Flags().GetString(flags.FlagFrom)

		// Sign those bytes
		txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())
		bytesToSign, err := hex.DecodeString(args[0])
		if err != nil {
			return err
		}
		sigBytes, _, err := txFactory.Keybase().Sign(from, bytesToSign)
		if err != nil {
			return err
		}

		fmt.Println(hex.EncodeToString(sigBytes))
		return nil
	}
}
