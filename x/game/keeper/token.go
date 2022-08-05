package keeper

import (
	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetAllDeposits(ctx sdk.Context) []types.Deposit {
	store := ctx.KVStore(k.storeKey)

	deposits := []types.Deposit{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixAccountDeposit))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		deposit := types.Deposit{}
		k.cdc.MustUnmarshal(it.Value(), &deposit)

		deposits = append(deposits, deposit)
	}
	return deposits
}

func (k Keeper) SetDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	bz := k.cdc.MustMarshal(&types.Deposit{
		Address: addr.String(),
		Amount:  amount,
	})
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AccountDepositKey(addr), bz)
}

func (k Keeper) DeleteDeposit(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AccountDepositKey(addr))
}

func (k Keeper) GetDeposit(ctx sdk.Context, addr sdk.AccAddress) sdk.Coin {
	deposit := types.Deposit{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AccountDepositKey(addr))
	if bz == nil {
		return sdk.Coin{}
	}
	k.cdc.MustUnmarshal(bz, &deposit)
	return deposit.Amount
}

func (k Keeper) IncreaseDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	deposit := k.GetDeposit(ctx, addr)
	if deposit.Denom == "" {
		deposit = amount
	} else {
		deposit = deposit.Add(amount)
	}
	k.SetDeposit(ctx, addr, deposit)
}

func (k Keeper) DecreaseDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	deposit := k.GetDeposit(ctx, addr)
	deposit.Denom = amount.Denom
	deposit = deposit.Sub(amount)
	k.SetDeposit(ctx, addr, deposit)
}
