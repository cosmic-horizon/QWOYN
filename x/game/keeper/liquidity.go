package keeper

import (
	"github.com/cosmic-horizon/coho/x/game/types"
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
