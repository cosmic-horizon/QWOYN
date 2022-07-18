package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmic-horizon/coho/x/game/types"
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
	_ = ctx
	return &types.MsgTransferModuleOwnershipResponse{}, nil
}

func (m msgServer) WhitelistNftContracts(goCtx context.Context, msg *types.MsgWhitelistNftContracts) (*types.MsgWhitelistNftContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgWhitelistNftContractsResponse{}, nil
}

func (m msgServer) RemoveWhitelistedNftContracts(goCtx context.Context, msg *types.MsgRemoveWhitelistedNftContracts) (*types.MsgRemoveWhitelistedNftContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgRemoveWhitelistedNftContractsResponse{}, nil
}

func (m msgServer) DepositNft(goCtx context.Context, msg *types.MsgDepositNft) (*types.MsgDepositNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgDepositNftResponse{}, nil
}

func (m msgServer) WithdrawUpdatedNft(goCtx context.Context, msg *types.MsgWithdrawUpdatedNft) (*types.MsgWithdrawUpdatedNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgWithdrawUpdatedNftResponse{}, nil
}
