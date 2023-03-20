package keeper

import (
	"context"

	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	gametypes "github.com/cosmic-horizon/qwoyn/x/game/types"
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

func (m msgServer) PutAllocationToken(goCtx context.Context, msg *types.MsgPutAllocationToken) (*types.MsgPutAllocationTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)
	if msg.Amount.Denom != params.AllocationToken {
		return nil, gametypes.ErrInvalidDepositDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventPutAllocationToken{
		Sender: msg.Sender,
		Amount: msg.Amount.String(),
	})

	return &types.MsgPutAllocationTokenResponse{}, nil
}

func (m msgServer) BuyAllocationToken(goCtx context.Context, msg *types.MsgBuyAllocationToken) (*types.MsgBuyAllocationTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)
	if msg.Amount.Denom != params.DepositToken {
		return nil, gametypes.ErrInvalidDepositDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// TODO: allocate vesting to the account based on price

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventBuyAllocationToken{
		Sender: msg.Sender,
		Amount: msg.Amount.String(),
	})

	return &types.MsgBuyAllocationTokenResponse{}, nil
}

func (m msgServer) SetDepositEndTime(goCtx context.Context, msg *types.MsgSetDepositEndTime) (*types.MsgSetDepositEndTimeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)
	params.DepositEndTime = msg.EndTime

	// TODO: owner verification

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventSetDepositEndTime{
		Time: msg.EndTime,
	})

	return &types.MsgSetDepositEndTimeResponse{}, nil
}

func (m msgServer) InitICA(goCtx context.Context, msg *types.MsgInitICA) (*types.MsgInitICAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)
	_ = params

	// TODO: execute what intertx does

	return &types.MsgInitICAResponse{}, nil
}

func (m msgServer) ExecAddLiquidity(goCtx context.Context, msg *types.MsgExecAddLiquidity) (*types.MsgExecAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)
	_ = params

	// TODO: transfer tokens to Osmosis network (this won't be direct IBC transfer in case it's USDC from Axelar)
	// TODO: initiate add liquidity operation

	return &types.MsgExecAddLiquidityResponse{}, nil
}
