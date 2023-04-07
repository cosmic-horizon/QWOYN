package types

import (
	"github.com/cosmic-horizon/qwoyn/osmosis/balancer"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPutAllocationToken{}

func NewMsgPutAllocationToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgPutAllocationToken {
	return &MsgPutAllocationToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgPutAllocationToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgPutAllocationToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgBuyAllocationToken{}

func NewMsgBuyAllocationToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgBuyAllocationToken {
	return &MsgBuyAllocationToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgBuyAllocationToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgBuyAllocationToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgSetDepositEndTime{}

func NewMsgSetDepositEndTime(sender sdk.AccAddress, endTime uint64) *MsgSetDepositEndTime {
	return &MsgSetDepositEndTime{
		Sender:  sender.String(),
		EndTime: endTime,
	}
}

func (msg MsgSetDepositEndTime) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgSetDepositEndTime) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgInitICA{}

func NewMsgInitICA(sender sdk.AccAddress, connectionId string) *MsgInitICA {
	return &MsgInitICA{
		Sender:       sender.String(),
		ConnectionId: connectionId,
	}
}

func (msg MsgInitICA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgInitICA) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgExecTransfer{}

func NewMsgExecTransfer(sender sdk.AccAddress, channelId string, timeoutNanoSecond uint64) *MsgExecTransfer {
	return &MsgExecTransfer{
		Sender:            sender.String(),
		TransferChannelId: channelId,
		TimeoutNanoSecond: timeoutNanoSecond,
	}
}

func (msg MsgExecTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgExecTransfer) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgExecAddLiquidity{}

func NewMsgExecAddLiquidity(sender sdk.AccAddress, msg balancer.MsgCreateBalancerPool) *MsgExecAddLiquidity {
	return &MsgExecAddLiquidity{
		Sender: sender.String(),
		Msg:    msg,
	}
}

func (msg MsgExecAddLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgExecAddLiquidity) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
