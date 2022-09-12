package keeper

import (
	"fmt"

	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) DepositNft(ctx sdk.Context, msg *types.MsgDepositNft) error {
	if !k.IsWhitelistedContract(ctx, msg.Contract) {
		return types.ErrNotWhitelistedContract
	}

	contractAddr, err := sdk.AccAddressFromBech32(msg.Contract)
	if err != nil {
		return err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	moduleAddr := k.AccountKeeper.GetModuleAddress(types.ModuleName)
	execMsg := fmt.Sprintf(`{"transfer_nft":{"token_id":"%d","recipient":"%s"}}`, msg.TokenId, moduleAddr.String())
	_, err = k.WasmKeeper.Execute(ctx, contractAddr, sender, []byte(execMsg), sdk.Coins{})
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) WithdrawUpdatedNft(ctx sdk.Context, msg *types.MsgWithdrawUpdatedNft) error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	signerMsg := types.NewMsgSignerWithdrawUpdatedNft(sender, msg.Contract, msg.TokenId, msg.ExecMsg)
	signBytes := signerMsg.GetSignBytes()

	signer := k.GetParamSet(ctx).Owner
	signerAcc, err := sdk.AccAddressFromBech32(signer)
	if err != nil {
		return err
	}
	acc := k.AccountKeeper.GetAccount(ctx, signerAcc)
	if acc == nil {
		return types.ErrSignerAccountNotRegistered
	}

	// retrieve pubkey
	pubKey := acc.GetPubKey()
	if pubKey == nil {
		return types.ErrSignerAccountPubKeyNotRegistered
	}

	if !pubKey.VerifySignature(signBytes, msg.Signature) {
		return fmt.Errorf("unable to verify signer signature")
	}

	contractAddr, err := sdk.AccAddressFromBech32(msg.Contract)
	if err != nil {
		return err
	}

	// execute update
	moduleAddr := k.AccountKeeper.GetModuleAddress(types.ModuleName)
	_, err = k.WasmKeeper.Execute(ctx, contractAddr, moduleAddr, []byte(msg.ExecMsg), sdk.Coins{})
	if err != nil {
		return err
	}

	// send nft to msg.Sender
	execMsg := fmt.Sprintf(`{"transfer_nft":{"token_id":"%d","recipient":"%s"}}`, msg.TokenId, msg.Sender)
	_, err = k.WasmKeeper.Execute(ctx, contractAddr, moduleAddr, []byte(execMsg), sdk.Coins{})
	if err != nil {
		return err
	}

	return nil
}
