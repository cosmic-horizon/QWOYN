package keeper

import (
	"time"

	"github.com/cosmic-horizon/qwoyn/x/game/types"
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
	addr, err := sdk.AccAddressFromBech32(deposit.Address)
	if err != nil {
		panic(err)
	}
	store.Set(types.AccountDepositKey(addr), bz)
}

func (k Keeper) DeleteDeposit(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AccountDepositKey(addr))
}

func (k Keeper) GetDeposit(ctx sdk.Context, addr sdk.AccAddress) types.Deposit {
	deposit := types.Deposit{
		Amount:    sdk.ZeroInt(),
		Staking:   sdk.ZeroInt(),
		Unbonding: sdk.ZeroInt(),
	}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AccountDepositKey(addr))
	if bz == nil {
		return deposit
	}
	k.cdc.MustUnmarshal(bz, &deposit)
	return deposit
}

func (k Keeper) IncreaseDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int) {
	deposit := k.GetDeposit(ctx, addr)
	if deposit.Address == "" {
		deposit.Address = addr.String()
		deposit.Amount = amount
	} else {
		deposit.Amount = deposit.Amount.Add(amount)
	}
	k.SetDeposit(ctx, deposit)
}

func (k Keeper) DecreaseDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int) error {
	deposit := k.GetDeposit(ctx, addr)
	if amount.Add(deposit.Staking).GT(deposit.Amount) {
		return types.ErrInsufficientDepositAmount
	}
	deposit.Amount = deposit.Amount.Sub(amount)
	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) IncreaseStaking(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int) error {
	deposit := k.GetDeposit(ctx, addr)
	if deposit.Staking.Add(amount).GT(deposit.Amount) {
		return types.ErrInsufficientDepositAmount
	}
	deposit.Staking = deposit.Staking.Add(amount)
	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) DecreaseStaking(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int) error {
	deposit := k.GetDeposit(ctx, addr)
	if deposit.Unbonding.Add(amount).GT(deposit.Staking) {
		return types.ErrInsufficientStakingAmount
	}
	deposit.Staking = deposit.Staking.Sub(amount)
	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) IncreaseUnbonding(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int) error {
	deposit := k.GetDeposit(ctx, addr)
	if deposit.Unbonding.Add(amount).GT(deposit.Staking) {
		return types.ErrInsufficientStakingAmount
	}
	deposit.Unbonding = deposit.Unbonding.Add(amount)
	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) DecreaseUnbonding(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int) error {
	deposit := k.GetDeposit(ctx, addr)
	if deposit.Unbonding.LT(amount) {
		return types.ErrInsufficientUnbondingAmount
	}
	deposit.Unbonding = deposit.Unbonding.Sub(amount)
	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) ClaimInGameStakingReward(ctx sdk.Context, addr sdk.AccAddress) error {
	deposit := k.GetDeposit(ctx, addr)
	duration := ctx.BlockTime().Sub(deposit.RewardClaimTime)
	params := k.GetParamSet(ctx)

	// calculate reward amount
	rewardAmount := deposit.Staking.Sub(deposit.Unbonding).
		Mul(sdk.NewInt(int64(params.StakingInflation))).
		Mul(sdk.NewInt(int64(duration))).
		Quo(sdk.NewInt(int64(24 * time.Hour * 365)))

	// mint coins and send rewards
	if rewardAmount.IsPositive() {
		rewardCoin := sdk.NewCoin(params.DepositDenom, rewardAmount)
		err := k.BankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{rewardCoin})
		if err != nil {
			return err
		}
		deposit.Amount = deposit.Amount.Add(rewardAmount)
	}

	// set last claim time
	deposit.RewardClaimTime = ctx.BlockTime()
	k.SetDeposit(ctx, deposit)

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventClaimInGameStakingReward{
		Sender:          addr.String(),
		Amount:          rewardAmount.String(),
		RewardClaimTime: uint64(ctx.BlockTime().Unix()),
	})

	return nil
}
