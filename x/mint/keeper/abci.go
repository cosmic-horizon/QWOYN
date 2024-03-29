package keeper

import (
	"time"

	"github.com/cosmic-horizon/qwoyn/x/mint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)
	bondedRatio := k.BondedRatio(ctx)
	minter.Inflation = minter.NextInflationRate(params, bondedRatio)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
	k.SetMinter(ctx, minter)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	// prevent mint if total supply + mint amount is bigger than max cap
	supply := k.bankKeeper.GetSupply(ctx, params.MintDenom)
	if supply.Amount.Add(mintedCoin.Amount).GT(params.MaxCap) {
		return
	}

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	outpostFundingCoin := sdk.NewCoin(mintedCoin.Denom, sdk.NewDecFromInt(mintedCoin.Amount).Mul(params.OutpostFundingPoolPortion).RoundInt())
	if outpostFundingCoin.IsPositive() {
		outpostFundingCoins := sdk.NewCoins(outpostFundingCoin)
		err = k.AddToOutpostFundingPool(ctx, outpostFundingCoins)
		if err != nil {
			panic(err)
		}
		mintedCoins = mintedCoins.Sub(outpostFundingCoins...)
	}

	if mintedCoins.IsAllPositive() {
		// send the minted coins to the fee collector account
		err = k.AddCollectedFees(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
