package keeper

import (
	"context"

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
	// TODO: check contract owner is module account via wasm call
	// - Update permission
	// - Minter permission as well
	for _, contract := range msg.Contracts {
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
