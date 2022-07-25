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
