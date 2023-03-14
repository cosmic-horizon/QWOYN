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

func (m msgServer) DepositIntoOutpostFunding(goCtx context.Context, msg *types.MsgDepositIntoOutpostFunding) (*types.MsgDepositIntoOutpostFundingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.gk.GetParamSet(ctx)
	if msg.Amount.Denom != params.DepositDenom {
		return nil, gametypes.ErrInvalidDepositDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.bk.SendCoinsFromAccountToModule(ctx, sender, types.OutpostFundingPoolName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventDepositIntoOutpostFunding{
		Sender: msg.Sender,
		Amount: msg.Amount.String(),
	})

	return &types.MsgDepositIntoOutpostFundingResponse{}, nil
}
