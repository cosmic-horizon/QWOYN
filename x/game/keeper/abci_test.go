package keeper_test

// TODO:
// func (k Keeper) EndBlocker(ctx sdk.Context) {
// 	endingUnbondings := k.GetCompletedUnbondingsAt(ctx, ctx.BlockTime())
// 	for _, unbonding := range endingUnbondings {
// 		addr, err := sdk.AccAddressFromBech32(unbonding.StakerAddress)
// 		if err != nil {
// 			continue
// 		}
// 		err = k.DecreaseUnbonding(ctx, addr, unbonding.Amount)
// 		if err != nil {
// 			continue
// 		}
// 		err = k.DecreaseStaking(ctx, addr, unbonding.Amount)
// 		if err != nil {
// 			continue
// 		}
// 		k.DeleteUnbonding(ctx, unbonding)
// 	}
// }
