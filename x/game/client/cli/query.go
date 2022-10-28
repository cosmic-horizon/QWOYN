package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/cosmic-horizon/qwoyn/x/game/types"
)

// GetQueryCmd returns the query commands for the game module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                types.ModuleName,
		Short:              "Querying commands for the game module",
		DisableFlagParsing: true,
	}

	queryCmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdWhitelistedContracts(),
		GetCmdInGameNfts(),
		GetCmdDepositBalance(),
		GetCmdAllDepositBalances(),
		GetCmdUserUnbondings(),
		GetCmdAllUnbondings(),
		GetCmdQueryLiquidity(),
		GetCmdQueryEstimatedSwapOut(),
		GetCmdQuerySwapRate(),
	)

	return queryCmd
}

func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "params [flags]",
		Long: "Query params.",
		Example: fmt.Sprintf(
			`$ %s query game params`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdWhitelistedContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "whitelisted-contracts [flags]",
		Long: "Query whitelisted contracts.",
		Example: fmt.Sprintf(
			`$ %s query game whitelisted-contracts`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.WhitelistedContracts(context.Background(), &types.QueryWhitelistedContractsRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdInGameNfts() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ingame-nfts [flags]",
		Long: "Query in-game nfts.",
		Example: fmt.Sprintf(
			`$ %s query game ingame-nfts`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.InGameNfts(context.Background(), &types.QueryInGameNftsRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdDepositBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit-balance [address] [flags]",
		Long: "Query in-game nfts.",
		Example: fmt.Sprintf(
			`$ %s query game deposit-balance [address]`, version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DepositBalance(context.Background(), &types.QueryDepositBalanceRequest{
				Address: args[0],
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdAllDepositBalances() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "all-deposit-balances [flags]",
		Long: "Query in-game nfts.",
		Example: fmt.Sprintf(
			`$ %s query game all-deposit-balances`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllDepositBalance(context.Background(), &types.QueryAllDepositBalancesRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdAllUnbondings() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "all-unbondings [flags]",
		Long: "Query all unbondings.",
		Example: fmt.Sprintf(
			`$ %s query game all-unbondings`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllUnbondings(context.Background(), &types.QueryAllUnbondingsRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdUserUnbondings() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "user-unbondings [user] [flags]",
		Long: "Query user unbondings.",
		Example: fmt.Sprintf(
			`$ %s query game user-unbondings [user]`, version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UserUnbondings(context.Background(), &types.QueryUserUnbondingsRequest{
				Address: args[0],
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "liquidity [flags]",
		Long: "Query total liquidity.",
		Example: fmt.Sprintf(
			`$ %s query game liquidity`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Liquidity(context.Background(), &types.QueryLiquidityRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryEstimatedSwapOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "estimated-swap-out [amount] [flags]",
		Long: "Query estimated swap out.",
		Example: fmt.Sprintf(
			`$ %s query game estimated-swap-out 1000ucoho`, version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.EstimatedSwapOut(context.Background(), &types.QueryEstimatedSwapOutRequest{
				Amount: args[0],
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQuerySwapRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "swap-rate [flags]",
		Long: "Query swap rate.",
		Example: fmt.Sprintf(
			`$ %s query game swap-rate`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.SwapRate(context.Background(), &types.QuerySwapRateRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
