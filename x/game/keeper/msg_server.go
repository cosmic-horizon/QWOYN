package keeper

import (
	"context"
	"fmt"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
	WasmKeeper    wasmtypes.ContractOpsKeeper
	AccountKeeper types.AccountKeeper
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
	if !m.IsWhitelistedContract(ctx, msg.Contract) {
		return nil, types.ErrNotWhitelistedContract
	}

	contractAddr, err := sdk.AccAddressFromBech32(msg.Contract)
	if err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	moduleAddr := m.AccountKeeper.GetModuleAddress(types.ModuleName)

	// TODO: to be replaced with transfer
	execMsg := fmt.Sprintf(`{"mint":{"token_id":"1","owner":"%s","extension":{"ship_type":10,"owner":"100"}}}`, moduleAddr.String())
	_, err = m.WasmKeeper.Execute(ctx, contractAddr, sender, []byte(execMsg), sdk.Coins{})
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositNftResponse{}, nil
}

func (m msgServer) WithdrawUpdatedNft(goCtx context.Context, msg *types.MsgWithdrawUpdatedNft) (*types.MsgWithdrawUpdatedNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	// moduleAddr := m.AccountKeeper.GetModuleAddress(types.ModuleName)
	// TODO: verify signature of mint / update
	// if mint, mint an nft and send it to the sender
	// if update, update nft and transfer it to the sender

	return &types.MsgWithdrawUpdatedNftResponse{}, nil
}
