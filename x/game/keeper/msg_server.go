package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) TransferModuleOwnership(goCtx context.Context, msg *types.MsgTransferModuleOwnership) (*types.MsgTransferModuleOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}
	params.Owner = msg.NewOwner
	m.SetParamSet(ctx, params)
	return &types.MsgTransferModuleOwnershipResponse{}, nil
}

func (m msgServer) WhitelistNftContracts(goCtx context.Context, msg *types.MsgWhitelistNftContracts) (*types.MsgWhitelistNftContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}

	moduleAddr := m.AccountKeeper.GetModuleAddress(types.ModuleName)
	for _, contract := range msg.Contracts {
		contractAddr, err := sdk.AccAddressFromBech32(contract)
		if err != nil {
			return nil, err
		}

		minterJSON, err := m.WasmViewer.QuerySmart(ctx, contractAddr, []byte(`{"minter": {}}`))

		var parsed map[string]string
		err = json.Unmarshal(minterJSON, &parsed)
		if err != nil {
			return nil, err
		}
		if parsed["minter"] != moduleAddr.String() {
			fmt.Println("minter", parsed["minter"])
			return nil, types.ErrMinterIsNotModuleAddress
		}

		contractInfoJSON, err := m.WasmViewer.QuerySmart(ctx, contractAddr, []byte(`{"contract_info": {}}`))
		err = json.Unmarshal(contractInfoJSON, &parsed)
		if err != nil {
			return nil, err
		}
		if parsed["owner"] != moduleAddr.String() {
			fmt.Println("owner", parsed["owner"])
			return nil, types.ErrOwnerIsNotModuleAddress
		}

		m.SetWhitelistedContract(ctx, contract)
	}
	return &types.MsgWhitelistNftContractsResponse{}, nil
}

func (m msgServer) RemoveWhitelistedNftContracts(goCtx context.Context, msg *types.MsgRemoveWhitelistedNftContracts) (*types.MsgRemoveWhitelistedNftContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}
	for _, contract := range msg.Contracts {
		m.DeleteWhitelistedContract(ctx, contract)
	}
	return &types.MsgRemoveWhitelistedNftContractsResponse{}, nil
}

func (m msgServer) DepositNft(goCtx context.Context, msg *types.MsgDepositNft) (*types.MsgDepositNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.DepositNft(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositNftResponse{}, nil
}

func (m msgServer) WithdrawUpdatedNft(goCtx context.Context, msg *types.MsgWithdrawUpdatedNft) (*types.MsgWithdrawUpdatedNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.WithdrawUpdatedNft(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawUpdatedNftResponse{}, nil
}

func (m msgServer) DepositToken(goCtx context.Context, msg *types.MsgDepositToken) (*types.MsgDepositTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Amount.Denom != params.DepositDenom {
		return nil, types.ErrInvalidDepositDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	m.IncreaseDeposit(ctx, sender, msg.Amount.Amount)

	return &types.MsgDepositTokenResponse{}, nil
}

func (m msgServer) WithdrawToken(goCtx context.Context, msg *types.MsgWithdrawToken) (*types.MsgWithdrawTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Amount.Denom != params.DepositDenom {
		return nil, types.ErrInvalidWithdrawDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	m.DecreaseDeposit(ctx, sender, msg.Amount.Amount)

	return &types.MsgWithdrawTokenResponse{}, nil
}

func (m msgServer) StakeInGameToken(goCtx context.Context, msg *types.MsgStakeInGameToken) (*types.MsgStakeInGameTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Amount.Denom != params.DepositDenom {
		return nil, types.ErrInvalidWithdrawDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.ClaimInGameStakingReward(ctx, sender)
	if err != nil {
		return nil, err
	}

	err = m.IncreaseStaking(ctx, sender, msg.Amount.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgStakeInGameTokenResponse{}, nil
}

func (m msgServer) BeginUnstakeInGameToken(goCtx context.Context, msg *types.MsgBeginUnstakeInGameToken) (*types.MsgBeginUnstakeInGameTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Amount.Denom != params.DepositDenom {
		return nil, types.ErrInvalidWithdrawDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.ClaimInGameStakingReward(ctx, sender)
	if err != nil {
		return nil, err
	}

	err = m.IncreaseUnbonding(ctx, sender, msg.Amount.Amount)
	if err != nil {
		return nil, err
	}

	lastUnbondingId := m.GetLastUnbondingId(ctx)
	lastUnbondingId++
	m.SetLastUnbondingId(ctx, lastUnbondingId)

	m.SetUnbonding(ctx, types.Unbonding{
		Id:             lastUnbondingId,
		StakerAddress:  msg.Sender,
		CreationHeight: ctx.BlockHeight(),
		CompletionTime: ctx.BlockTime().Add(params.UnstakingTime),
		Amount:         msg.Amount.Amount,
	})

	return &types.MsgBeginUnstakeInGameTokenResponse{}, nil
}

func (m msgServer) ClaimInGameStakingReward(goCtx context.Context, msg *types.MsgClaimInGameStakingReward) (*types.MsgClaimInGameStakingRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.ClaimInGameStakingReward(ctx, sender)
	if err != nil {
		return nil, err
	}

	return &types.MsgClaimInGameStakingRewardResponse{}, nil
}

func (m msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}

	m.Keeper.IncreaseLiquidity(ctx, msg.Amounts)
	liquidity := m.Keeper.GetLiquidity(ctx)
	if len(liquidity.Amounts) != 2 {
		return nil, types.ErrLiquidityShouldHoldTwoTokens
	}

	err = m.Keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, msg.Amounts)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddLiquidityResponse{}, nil
}

func (m msgServer) RemoveLiquidity(goCtx context.Context, msg *types.MsgRemoveLiquidity) (*types.MsgRemoveLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}

	err = m.Keeper.DecreaseLiquidity(ctx, msg.Amounts)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, msg.Amounts)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveLiquidityResponse{}, nil
}

func (m msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// deposit coins into module and increase liquidity
	err = m.Keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// withdraw coins from module and decrease liquidity
	tarCoin, err := m.Keeper.SwapOutAmount(ctx, msg.Amount)
	if err != nil {
		return nil, err
	}

	tarCoins := sdk.Coins{tarCoin}
	err = m.Keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, tarCoins)
	if err != nil {
		return nil, err
	}

	m.Keeper.IncreaseLiquidity(ctx, sdk.Coins{msg.Amount})
	err = m.Keeper.DecreaseLiquidity(ctx, tarCoins)
	if err != nil {
		return nil, err
	}

	return &types.MsgSwapResponse{}, nil
}
