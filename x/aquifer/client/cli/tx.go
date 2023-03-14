package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdDepositIntoOutpostFundingPool(),
		GetCmdWithdrawFromOutpostFundingPool(),
	)

	return cmd
}

func GetCmdDepositIntoOutpostFundingPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit-outpost-funding [coin] [flags]",
		Long: "Deposit token into the outpost funding pool",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx deposit-outpost-funding [coin]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositIntoOutpostFunding(
				clientCtx.GetFromAddress(),
				coin,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdWithdrawFromOutpostFundingPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw-outpost-funding [coin] [flags]",
		Long: "Withdraw token from the outpost funding pool",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx withdraw-outpost-funding [coin]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawFromOutpostFunding(
				clientCtx.GetFromAddress(),
				coin,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
