package game

import (
	"github.com/cosmic-horizon/coho/x/game/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	endingUnbondings := k.GetCompletedUnbondingsAt(ctx, ctx.BlockTime())
	for _, unbonding := range endingUnbondings {
		addr, err := sdk.AccAddressFromBech32(unbonding.StakerAddress)
		if err != nil {
			continue
		}
		err = k.DecreaseUnbonding(ctx, addr, unbonding.Amount)
		if err != nil {
			continue
		}
		err = k.DecreaseStaking(ctx, addr, unbonding.Amount)
		if err != nil {
			continue
		}
		k.IncreaseDeposit(ctx, addr, unbonding.Amount)
		k.DeleteUnbonding(ctx, unbonding)
	}
}
