package keeper

import (
	"fmt"
	"time"

	appparams "github.com/cosmic-horizon/qwoyn/app/params"
	"github.com/cosmic-horizon/qwoyn/x/stimulus/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	outpostPoolAddr := k.ak.GetModuleAddress(types.OutpostFundingPoolName)
	qwoynBalance := k.bk.GetBalance(ctx, outpostPoolAddr, appparams.BondDenom)

	cacheCtx, write := ctx.CacheContext()
	err := k.gk.Swap(cacheCtx, outpostPoolAddr, qwoynBalance)
	if err == nil {
		write()
	} else {
		fmt.Println("automatic swap error", err)
	}
}
