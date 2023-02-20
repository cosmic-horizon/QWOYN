package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDepositIntoOutpostFunding{}

func NewMsgDepositIntoOutpostFunding(sender sdk.AccAddress, coin sdk.Coin,
) *MsgDepositIntoOutpostFunding {
	return &MsgDepositIntoOutpostFunding{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgDepositIntoOutpostFunding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgDepositIntoOutpostFunding) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgWithdrawFromOutpostFunding{}

func NewMsgWithdrawFromOutpostFunding(sender sdk.AccAddress, coin sdk.Coin,
) *MsgWithdrawFromOutpostFunding {
	return &MsgWithdrawFromOutpostFunding{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgWithdrawFromOutpostFunding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgWithdrawFromOutpostFunding) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
