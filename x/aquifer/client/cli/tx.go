package cli

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cosmic-horizon/qwoyn/osmosis/balancer"
	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
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
		GetCmdPutAllocationToken(),
		GetCmdTakeOutAllocationToken(),
		GetCmdBuyAllocationToken(),
		GetCmdSetDepositEndTime(),
		GetCmdInitICA(),
		GetCmdExecTransfer(),
		GetCmdExecAddLiquidity(),
	)

	return cmd
}

func GetCmdPutAllocationToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "put-allocation-token [coin] [flags]",
		Long: "Put allocation token into aquifer pool",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx put-allocation-token [coin]`,
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

			msg := types.NewMsgPutAllocationToken(
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

func GetCmdTakeOutAllocationToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "take-out-allocation-token [coin] [flags]",
		Long: "Take out allocation token from aquifer pool",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx take-out-allocation-token [coin]`,
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

			msg := types.NewMsgTakeOutAllocationToken(
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

func GetCmdBuyAllocationToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "buy-allocation-token [coin] [flags]",
		Long: "Buy allocation token from aquifer pool",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx buy-allocation-token [coin]`,
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

			msg := types.NewMsgBuyAllocationToken(
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

func GetCmdSetDepositEndTime() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "set-deposit-endtime [endTime] [flags]",
		Long: "Set deposit end time",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx set-deposit-endtime [coin]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			endTime, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgSetDepositEndTime(
				clientCtx.GetFromAddress(),
				uint64(endTime),
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

func GetCmdInitICA() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "init-ica [connectionId] [flags]",
		Long: "Initialize interchain account",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(
			`$ %s tx init-ica [connectionId]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitICA(
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

func GetCmdExecTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "exec-transfer [channelId] [timeoutNanoSecond] [flags]",
		Long: "Execute ibc transfer to target network",
		Args: cobra.ExactArgs(2),
		Example: fmt.Sprintf(
			`$ %s tx exec-transfer [channelId] [timeoutNanoSecond]`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			timeoutNanoSecond, err := strconv.Atoi(args[1])
			msg := types.NewMsgExecTransfer(
				clientCtx.GetFromAddress(),
				args[0],
				uint64(timeoutNanoSecond),
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

func GetCmdExecAddLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "exec-add-liquidity [flags]",
		Long: "Execute ibc transfer to target network",
		Args: cobra.ExactArgs(0),
		Example: fmt.Sprintf(
			`$ %s tx exec-add-liquidity --pool-file="pool.json"
			Sample pool JSON file contents:
			{
    "weights": "4uatom,4osmo,2uakt",
    "initial-deposit": "100uatom,5osmo,20uakt",
    "swap-fee": "0.01",
    "exit-fee": "0.01",
    "future-governor": "168h"
}`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf, _ := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			txf = txf.WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)
			var msg sdk.Msg
			txf, balancerPoolMsg, err := NewBuildCreateBalancerPoolMsg(clientCtx, txf, cmd.Flags())
			if err != nil {
				return err
			}

			msg = types.NewMsgExecAddLiquidity(clientCtx.GetFromAddress(), *balancerPoolMsg)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetCreatePool())
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagPoolFile)

	return cmd
}

func NewBuildCreateBalancerPoolMsg(clientCtx client.Context, txf tx.Factory, fs *flag.FlagSet) (tx.Factory, *balancer.MsgCreateBalancerPool, error) {
	pool, err := parseCreateBalancerPoolFlags(fs)
	if err != nil {
		return txf, nil, fmt.Errorf("failed to parse pool: %w", err)
	}

	deposit, err := sdk.ParseCoinsNormalized(pool.InitialDeposit)
	if err != nil {
		return txf, nil, err
	}

	poolAssetCoins, err := sdk.ParseDecCoins(pool.Weights)
	if err != nil {
		return txf, nil, err
	}

	if len(deposit) != len(poolAssetCoins) {
		return txf, nil, errors.New("deposit tokens and token weights should have same length")
	}

	swapFee, err := sdk.NewDecFromStr(pool.SwapFee)
	if err != nil {
		return txf, nil, err
	}

	exitFee, err := sdk.NewDecFromStr(pool.ExitFee)
	if err != nil {
		return txf, nil, err
	}

	var poolAssets []balancer.PoolAsset
	for i := 0; i < len(poolAssetCoins); i++ {
		if poolAssetCoins[i].Denom != deposit[i].Denom {
			return txf, nil, errors.New("deposit tokens and token weights should have same denom order")
		}

		poolAssets = append(poolAssets, balancer.PoolAsset{
			Weight: poolAssetCoins[i].Amount.RoundInt(),
			Token:  deposit[i],
		})
	}

	poolParams := &balancer.PoolParams{
		SwapFee: swapFee,
		ExitFee: exitFee,
	}

	msg := &balancer.MsgCreateBalancerPool{
		Sender:             clientCtx.GetFromAddress().String(),
		PoolParams:         poolParams,
		PoolAssets:         poolAssets,
		FuturePoolGovernor: pool.FutureGovernor,
	}

	if (pool.SmoothWeightChangeParams != smoothWeightChangeParamsInputs{}) {
		duration, err := time.ParseDuration(pool.SmoothWeightChangeParams.Duration)
		if err != nil {
			return txf, nil, fmt.Errorf("could not parse duration: %w", err)
		}

		targetPoolAssetCoins, err := sdk.ParseDecCoins(pool.SmoothWeightChangeParams.TargetPoolWeights)
		if err != nil {
			return txf, nil, err
		}

		var targetPoolAssets []balancer.PoolAsset
		for i := 0; i < len(targetPoolAssetCoins); i++ {
			if targetPoolAssetCoins[i].Denom != poolAssetCoins[i].Denom {
				return txf, nil, errors.New("initial pool weights and target pool weights should have same denom order")
			}

			targetPoolAssets = append(targetPoolAssets, balancer.PoolAsset{
				Weight: targetPoolAssetCoins[i].Amount.RoundInt(),
				Token:  deposit[i],
			})
		}

		smoothWeightParams := balancer.SmoothWeightChangeParams{
			Duration:           duration,
			InitialPoolWeights: poolAssets,
			TargetPoolWeights:  targetPoolAssets,
		}

		if pool.SmoothWeightChangeParams.StartTime != "" {
			startTime, err := time.Parse(time.RFC3339, pool.SmoothWeightChangeParams.StartTime)
			if err != nil {
				return txf, nil, fmt.Errorf("could not parse time: %w", err)
			}

			smoothWeightParams.StartTime = startTime
		}

		msg.PoolParams.SmoothWeightChangeParams = &smoothWeightParams
	}

	return txf, msg, nil
}
