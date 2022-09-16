package keeper

import (
	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
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
		k.DeleteUnbonding(ctx, unbonding)

		// emit event
		ctx.EventManager().EmitTypedEvent(&types.EventCompleteUnstakeInGameToken{
			User:           unbonding.StakerAddress,
			Amount:         unbonding.Amount.String(),
			CompletionTime: uint64(ctx.BlockTime().Unix()),
			UnbondingId:    unbonding.Id,
		})
	}
}
