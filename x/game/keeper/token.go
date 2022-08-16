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

func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	bz := k.cdc.MustMarshal(&deposit)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AccountDepositKey(deposit.Address), bz)
}

func (k Keeper) DeleteDeposit(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AccountDepositKey(addr))
}

func (k Keeper) GetDeposit(ctx sdk.Context, addr sdk.AccAddress) types.Deposit {
	deposit := types.Deposit{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AccountDepositKey(addr))
	if bz == nil {
		return deposit
	}
	k.cdc.MustUnmarshal(bz, &deposit)
	return deposit
}

func (k Keeper) IncreaseDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	deposit := k.GetDeposit(ctx, addr)
	if deposit.Address == "" {
		deposit.LastRewardClaimTime = ctx.BlockTime()
		deposit.Address = addr.String()
		deposit.Amount = amount
	} else {
		deposit.Amount = deposit.Amount.Add(amount)
	}
	k.SetDeposit(ctx, addr, deposit)
}

func (k Keeper) DecreaseDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	deposit := k.GetDeposit(ctx, addr)
	// TODO: check deposit
	// TODO: check staking balance
	deposit.Amount = deposit.Amount.Sub(amount)
	k.SetDeposit(ctx, addr, deposit)
}

func (k Keeper) IncreaseStaking(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	// TODO: check deposit balance
	// TODO: ClaimInGameStakingReward
}

func (k Keeper) DecreaseStaking(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	// TODO: check staking balance
	// TODO: ClaimInGameStakingReward
}

func (k Keeper) IncreaseUnbonding(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	// TODO: check staking balance
	// TODO: check already unbonding balance
	// TODO: ClaimInGameStakingReward
}

func (k Keeper) DecreaseUnbonding(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) {
	// TODO: check unbonding balance
}

func (k Keeper) ClaimInGameStakingReward(ctx sdk.Context, addr sdk.AccAddress) {
	// TODO: calculate reward amount
	// TODO: mint coins and send rewards
	// TODO: set last claim time
}
