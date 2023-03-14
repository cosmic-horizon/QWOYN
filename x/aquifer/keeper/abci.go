package keeper

import (
	"fmt"
	"time"

	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	outpostPoolAddr := k.ak.GetModuleAddress(types.OutpostFundingPoolName)
	mintDenom := k.mk.GetParams(ctx).MintDenom
	mintedBalance := k.bk.GetBalance(ctx, outpostPoolAddr, mintDenom)

	cacheCtx, write := ctx.CacheContext()
	err := k.gk.SwapFromModule(cacheCtx, types.OutpostFundingPoolName, mintedBalance)
	if err == nil {
		write()
	} else {
		fmt.Println("automatic swap error", mintedBalance, err)
	}
}
