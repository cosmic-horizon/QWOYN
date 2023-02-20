package keeper

import (
	"github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLiquidity(ctx sdk.Context, liquidity types.Liquidity) {
	bz := k.cdc.MustMarshal(&liquidity)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLiquidity, bz)
}

func (k Keeper) DeleteLiquidity(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyLiquidity)
}

func (k Keeper) GetLiquidity(ctx sdk.Context) types.Liquidity {
	liquidity := types.Liquidity{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLiquidity)
	if bz == nil {
		return liquidity
	}
	k.cdc.MustUnmarshal(bz, &liquidity)
	return liquidity
}

func (k Keeper) IncreaseLiquidity(ctx sdk.Context, amounts sdk.Coins) {
	liquidity := k.GetLiquidity(ctx)
	liquidity.Amounts = sdk.Coins(liquidity.Amounts).Add(amounts...)
	k.SetLiquidity(ctx, liquidity)
}

func (k Keeper) DecreaseLiquidity(ctx sdk.Context, amounts sdk.Coins) error {
	liquidity := k.GetLiquidity(ctx)
	liqAmounts := sdk.Coins(liquidity.Amounts)
	if !liqAmounts.IsAllGTE(amounts) {
		return types.ErrInsufficientLiquidityAmount
	}
	liquidity.Amounts = liqAmounts.Sub(amounts)
	k.SetLiquidity(ctx, liquidity)
	return nil
}

func (k Keeper) SwapOutAmount(ctx sdk.Context, amount sdk.Coin) (sdk.Coin, error) {
	liquidity := k.GetLiquidity(ctx)
	if len(liquidity.Amounts) != 2 {
		return sdk.Coin{}, types.ErrLiquidityShouldHoldTwoTokens
	}

	srcLiq := liquidity.Amounts[0]
	tarLiq := liquidity.Amounts[1]
	if liquidity.Amounts[1].Denom == amount.Denom {
		srcLiq = liquidity.Amounts[1]
		tarLiq = liquidity.Amounts[0]
	}

	srcLiqAmount := srcLiq.Amount
	tarLiqAmount := tarLiq.Amount
	constantK := srcLiqAmount.Mul(tarLiqAmount)
	tarLiqRemaining := constantK.Quo(srcLiqAmount.Add(amount.Amount))
	tarAmount := tarLiqAmount.Sub(tarLiqRemaining)
	return sdk.NewCoin(tarLiq.Denom, tarAmount), nil
}

func (k Keeper) Swap(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coin) error {
	// deposit coins into module and increase liquidity
	err := k.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{amount})
	if err != nil {
		return err
	}

	// withdraw coins from module and decrease liquidity
	tarCoin, err := k.SwapOutAmount(ctx, amount)
	if err != nil {
		return err
	}

	tarCoins := sdk.Coins{tarCoin}
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, tarCoins)
	if err != nil {
		return err
	}

	k.IncreaseLiquidity(ctx, sdk.Coins{amount})
	err = k.DecreaseLiquidity(ctx, tarCoins)
	if err != nil {
		return err
	}

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventSwap{
		Sender:    sender.String(),
		InAmount:  amount.String(),
		OutAmount: tarCoin.String(),
	})
	return nil
}

func (k Keeper) SwapFromModule(ctx sdk.Context, moduleName string, amount sdk.Coin) error {
	// deposit coins into module and increase liquidity
	err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, moduleName, types.ModuleName, sdk.Coins{amount})
	if err != nil {
		return err
	}

	// withdraw coins from module and decrease liquidity
	tarCoin, err := k.SwapOutAmount(ctx, amount)
	if err != nil {
		return err
	}

	tarCoins := sdk.Coins{tarCoin}
	err = k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, moduleName, tarCoins)
	if err != nil {
		return err
	}

	k.IncreaseLiquidity(ctx, sdk.Coins{amount})
	err = k.DecreaseLiquidity(ctx, tarCoins)
	if err != nil {
		return err
	}

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventSwap{
		Sender:    moduleName,
		InAmount:  amount.String(),
		OutAmount: tarCoin.String(),
	})
	return nil
}
