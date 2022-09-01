package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmic-horizon/coho/x/game/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryParamsResponse{
		Params: k.GetParamSet(ctx),
	}, nil
}

func (k Keeper) WhitelistedContracts(c context.Context, req *types.QueryWhitelistedContractsRequest) (*types.QueryWhitelistedContractsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryWhitelistedContractsResponse{
		Contracts: k.GetAllWhitelistedContracts(ctx),
	}, nil
}

func (k Keeper) InGameNfts(c context.Context, req *types.QueryInGameNftsRequest) (*types.QueryInGameNftsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryInGameNftsResponse{}, nil
}

func (k Keeper) DepositBalance(c context.Context, req *types.QueryDepositBalanceRequest) (*types.QueryDepositBalanceResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	return &types.QueryDepositBalanceResponse{
		Deposit: k.GetDeposit(ctx, address),
	}, nil
}

func (k Keeper) AllDepositBalance(c context.Context, req *types.QueryAllDepositBalancesRequest) (*types.QueryAllDepositBalanceResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryAllDepositBalanceResponse{
		Deposits: k.GetAllDeposits(ctx),
	}, nil
}

func (k Keeper) AllUnbondings(c context.Context, req *types.QueryAllUnbondingsRequest) (*types.QueryAllUnbondingsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryAllUnbondingsResponse{
		Unbondings: k.GetAllUnbondings(ctx),
	}, nil
}

func (k Keeper) UserUnbondings(c context.Context, req *types.QueryUserUnbondingsRequest) (*types.QueryUserUnbondingsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	return &types.QueryUserUnbondingsResponse{
		Unbondings: k.GetUserUnbondings(ctx, addr),
	}, nil
}

func (k Keeper) Liquidity(c context.Context, req *types.QueryLiquidityRequest) (*types.QueryLiquidityResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryLiquidityResponse{
		Liquidity: k.GetLiquidity(ctx),
	}, nil
}

func (k Keeper) EstimatedSwapOut(c context.Context, req *types.QueryEstimatedSwapOutRequest) (*types.QueryEstimatedSwapOutResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	coin, err := sdk.ParseCoinNormalized(req.Amount)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	outAmount, err := k.SwapOutAmount(ctx, coin)
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimatedSwapOutResponse{
		Amount: outAmount,
	}, nil
}

func (k Keeper) SwapRate(c context.Context, req *types.QuerySwapRateRequest) (*types.QuerySwapRateResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	liquidity := k.GetLiquidity(ctx)
	if len(liquidity.Amounts) != 2 {
		return nil, types.ErrLiquidityShouldHoldTwoTokens
	}

	srcLiq := liquidity.Amounts[0]
	tarLiq := liquidity.Amounts[1]

	rate := tarLiq.Amount.ToDec().Quo(srcLiq.Amount.ToDec())

	return &types.QuerySwapRateResponse{
		Rate:     rate,
		SrcDenom: srcLiq.Denom,
		TarDenom: tarLiq.Denom,
	}, nil
}
