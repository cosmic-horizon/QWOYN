package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmic-horizon/coho/x/game/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
)

// GetTxCmd returns the transaction commands for the game module.
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "game transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdTransferModuleOwnership(),
		GetCmdWhitelistNftContracts(),
		GetCmdRemoveWhitelistedNftContracts(),
		GetCmdDepositNft(),
		GetCmdSignWithdrawUpdatedNft(),
		GetCmdWithdrawUpdatedNft(),
	)

	return txCmd
}

func GetCmdTransferModuleOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "transfer-module-ownership [newOwner] [flags]",
		Long: "Transfer module ownership to new address",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx transfer-module-ownership [newOwner]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferModuleOwnership(
				clientCtx.GetFromAddress(),
				args[0],
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

func GetCmdWhitelistNftContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "whitelist-contracts [contracts] [flags]",
		Long: "Whitelist contracts by module owner",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx whitelist-contracts [contracts]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contracts := strings.Split(args[0], ",")
			msg := types.NewMsgWhitelistNftContracts(
				clientCtx.GetFromAddress(),
				contracts,
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

func GetCmdRemoveWhitelistedNftContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "remove-whitelisted-contracts [contracts] [flags]",
		Long: "Remove whitelisted contracts by module owner",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx whitelist-contracts [contracts]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contracts := strings.Split(args[0], ",")
			msg := types.NewMsgRemoveWhitelistedNftContracts(
				clientCtx.GetFromAddress(),
				contracts,
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

func GetCmdDepositNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit-nft [contract] [tokenId] [flags]",
		Long: "Deposit nft to the game",
		Args: cobra.ExactArgs(2),
		Example: fmt.Sprintf(
			`$ %s tx deposit-nft [contract] [tokenId]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenId, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositNft(
				clientCtx.GetFromAddress(),
				args[0],
				uint64(tokenId),
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

// GetCmdSignWithdrawUpdatedNft returns signature from signer
func GetCmdSignWithdrawUpdatedNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "sign-withdraw-updated-nft [contract] [tokenId] [execMsg] [flags]",
		Long: "Sign withdraw updated nft",
		Args: cobra.ExactArgs(3),
		Example: fmt.Sprintf(
			`$ %s tx withdraw-updated-nft [contract] [tokenId] [execMsg]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenId, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgSignerWithdrawUpdatedNft(
				clientCtx.GetFromAddress(),
				args[0],
				uint64(tokenId),
				args[2],
			)

			bytesToSign := msg.GetSignBytes()
			if err != nil {
				return err
			}

			from, _ := cmd.Flags().GetString(flags.FlagFrom)
			txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			sigBytes, _, err := txFactory.Keybase().Sign(from, bytesToSign)
			if err != nil {
				return err
			}

			fmt.Println(hex.EncodeToString(sigBytes))
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdWithdrawUpdatedNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw-updated-nft [contract] [tokenId] [execMsg] [flags]",
		Long: "Withdraw updated nft",
		Args: cobra.ExactArgs(4),
		Example: fmt.Sprintf(
			`$ %s tx withdraw-updated-nft [contract] [tokenId] [execMsg] [signature]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenId, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			signature, err := hex.DecodeString(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawUpdatedNft(
				clientCtx.GetFromAddress(),
				args[0],
				uint64(tokenId),
				args[2],
				signature,
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
