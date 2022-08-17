package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferModuleOwnership{}

func NewMsgTransferModuleOwnership(sender sdk.AccAddress, newOwner string) *MsgTransferModuleOwnership {
	return &MsgTransferModuleOwnership{
		Sender:   sender.String(),
		NewOwner: newOwner,
	}
}

func (msg MsgTransferModuleOwnership) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgTransferModuleOwnership) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgWhitelistNftContracts{}

func NewMsgWhitelistNftContracts(sender sdk.AccAddress, contracts []string,
) *MsgWhitelistNftContracts {
	return &MsgWhitelistNftContracts{
		Sender:    sender.String(),
		Contracts: contracts,
	}
}

func (msg MsgWhitelistNftContracts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgWhitelistNftContracts) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgRemoveWhitelistedNftContracts{}

func NewMsgRemoveWhitelistedNftContracts(sender sdk.AccAddress, contracts []string,
) *MsgRemoveWhitelistedNftContracts {
	return &MsgRemoveWhitelistedNftContracts{
		Sender:    sender.String(),
		Contracts: contracts,
	}
}

func (msg MsgRemoveWhitelistedNftContracts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgRemoveWhitelistedNftContracts) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgDepositNft{}

func NewMsgDepositNft(sender sdk.AccAddress, contract string, tokenId uint64,
) *MsgDepositNft {
	return &MsgDepositNft{
		Sender:   sender.String(),
		Contract: contract,
		TokenId:  tokenId,
	}
}

func (msg MsgDepositNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgDepositNft) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgWithdrawUpdatedNft{}

func NewMsgWithdrawUpdatedNft(sender sdk.AccAddress, contract string, tokenId uint64, execMsg string, signature []byte,
) *MsgWithdrawUpdatedNft {
	return &MsgWithdrawUpdatedNft{
		Sender:    sender.String(),
		Contract:  contract,
		TokenId:   tokenId,
		ExecMsg:   execMsg,
		Signature: signature,
	}
}

func (msg MsgWithdrawUpdatedNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgWithdrawUpdatedNft) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func NewMsgSignerWithdrawUpdatedNft(sender sdk.AccAddress, contract string, tokenId uint64, execMsg string,
) *MsgSignerWithdrawUpdatedNft {
	return &MsgSignerWithdrawUpdatedNft{
		Sender:   sender.String(),
		Contract: contract,
		TokenId:  tokenId,
		ExecMsg:  execMsg,
	}
}

func (msg MsgSignerWithdrawUpdatedNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

func (msg MsgSignerWithdrawUpdatedNft) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

var _ sdk.Msg = &MsgDepositToken{}

func NewMsgDepositToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgDepositToken {
	return &MsgDepositToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgDepositToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgDepositToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgWithdrawToken{}

func NewMsgWithdrawToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgWithdrawToken {
	return &MsgWithdrawToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgWithdrawToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgWithdrawToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgStakeInGameToken{}

func NewMsgStakeInGameToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgStakeInGameToken {
	return &MsgStakeInGameToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgStakeInGameToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgStakeInGameToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgBeginUnstakeInGameToken{}

func NewMsgBeginUnstakeInGameToken(sender sdk.AccAddress, coin sdk.Coin,
) *MsgBeginUnstakeInGameToken {
	return &MsgBeginUnstakeInGameToken{
		Sender: sender.String(),
		Amount: coin,
	}
}

func (msg MsgBeginUnstakeInGameToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgBeginUnstakeInGameToken) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgClaimInGameStakingReward{}

func NewMsgClaimInGameStakingReward(sender sdk.AccAddress,
) *MsgClaimInGameStakingReward {
	return &MsgClaimInGameStakingReward{
		Sender: sender.String(),
	}
}

func (msg MsgClaimInGameStakingReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgClaimInGameStakingReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
