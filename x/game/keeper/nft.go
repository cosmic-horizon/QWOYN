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
	// moduleAddr := m.AccountKeeper.GetModuleAddress(types.ModuleName)
	// TODO: verify signature of mint / update
	// if mint, mint an nft and send it to the sender
	// if update, update nft and transfer it to the sender

	// execMsg := fmt.Sprintf(`{"mint":{"token_id":"1","owner":"%s","extension":{"ship_type":10,"owner":"100"}}}`, moduleAddr.String())
	// execMsg := fmt.Sprintf(`{"update_nft":{"token_id":"1","extension":{"ship_type":20,"owner":"200"}}}`)
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

	fmt.Println("events", ctx.EventManager().Events())

	// send nft to msg.Sender
	execMsg := fmt.Sprintf(`{"transfer_nft":{"token_id":"%d","recipient":"%s"}}`, msg.TokenId, msg.Sender)
	_, err = k.WasmKeeper.Execute(ctx, contractAddr, moduleAddr, []byte(execMsg), sdk.Coins{})
	if err != nil {
		return err
	}

	return nil
}
