package keeper

import (
	"context"
	"time"

	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	gametypes "github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// TODO: set discount on allocation amount
	allocationAmount := msg.Amount.Amount.ToDec().Quo(params.InitLiquidityPrice).RoundInt()
	allocationCoins := sdk.Coins{sdk.NewCoin(params.AllocationToken, allocationAmount)}
	account := m.ak.GetAccount(ctx, sender)
	switch account.(type) {
	case *authtypes.BaseAccount:
		baseVestingAccount := vestingtypes.NewBaseVestingAccount(account.(*authtypes.BaseAccount), allocationCoins, ctx.BlockTime().Unix()+int64(params.VestingDuration))
		m.ak.SetAccount(ctx, vestingtypes.NewDelayedVestingAccountRaw(baseVestingAccount))
		err = m.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, allocationCoins)
		if err != nil {
			return nil, err
		}
	default:
		return nil, types.ErrNotBaseAccount
	}

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
	params.IcsConnectionId = msg.ConnectionId
	m.SetParams(ctx, params)

	if err := m.icaControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, types.ModuleName); err != nil {
		return nil, err
	}

	return &types.MsgInitICAResponse{}, nil
}

func (m msgServer) ExecTransfer(goCtx context.Context, msg *types.MsgExecTransfer) (*types.MsgExecTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)

	portID, err := icatypes.NewControllerPortID(types.ModuleName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	icaAddr, found := m.icaControllerKeeper.GetInterchainAccountAddress(ctx, params.IcsConnectionId, portID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
	}

	moduleAddr := authtypes.NewModuleAddress(types.ModuleName)
	depositAmount := m.bk.GetBalance(ctx, moduleAddr, params.DepositToken)

	// TODO: transfer tokens to Osmosis network (this won't be direct IBC transfer in case it's USDC from Axelar)
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + msg.TimeoutNanoSecond
	_, err = m.TransferKeeper.Transfer(goCtx, ibctransfertypes.NewMsgTransfer(
		ibctransfertypes.PortID,
		msg.TransferChannelId,
		depositAmount,
		moduleAddr.String(),
		icaAddr,
		clienttypes.Height{},
		timeoutTimestamp))
	if err != nil {
		return nil, err
	}
	return &types.MsgExecTransferResponse{}, nil
}

func (m msgServer) ExecAddLiquidity(goCtx context.Context, msg *types.MsgExecAddLiquidity) (*types.MsgExecAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)

	portID, err := icatypes.NewControllerPortID(types.ModuleName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	channelID, found := m.icaControllerKeeper.GetActiveChannelID(ctx, params.IcsConnectionId, portID)
	if !found {
		return nil, icatypes.ErrActiveChannelNotFound.Wrapf("failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := m.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return nil, channeltypes.ErrChannelCapabilityNotFound.Wrap("module does not own channel capability")
	}

	addr, found := m.icaControllerKeeper.GetInterchainAccountAddress(ctx, params.IcsConnectionId, portID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
	}

	msg.Msg.Sender = addr

	data, err := icatypes.SerializeCosmosTx(m.cdc, []sdk.Msg{&msg.Msg})
	if err != nil {
		return nil, err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	_, err = m.icaControllerKeeper.SendTx(ctx, chanCap, params.IcsConnectionId, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	return &types.MsgExecAddLiquidityResponse{}, nil
}
