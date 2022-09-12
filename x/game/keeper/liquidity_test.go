package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestLiquidityGetSetDelete() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// get not available liquidity
	liq := suite.app.GameKeeper.GetLiquidity(suite.ctx)
	suite.Require().Equal(liq, types.Liquidity{})

	// set liquidity
	liquidity := types.Liquidity{
		Amounts: sdk.Coins{sdk.NewInt64Coin("ucoho", 1000), sdk.NewInt64Coin("qwoyn", 1000)},
	}

	suite.app.GameKeeper.SetLiquidity(suite.ctx, liquidity)

	liq = suite.app.GameKeeper.GetLiquidity(suite.ctx)
	suite.Require().Equal(liquidity, liq)

	suite.app.GameKeeper.DeleteLiquidity(suite.ctx)

	liq = suite.app.GameKeeper.GetLiquidity(suite.ctx)
	suite.Require().Equal(liq, types.Liquidity{})
}

func (suite *KeeperTestSuite) TestIncreaseLiquidity() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// increase liquidity
	coins := sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)}
	suite.app.GameKeeper.IncreaseLiquidity(suite.ctx, coins)

	liq := suite.app.GameKeeper.GetLiquidity(suite.ctx)
	suite.Require().Equal(types.Liquidity{
		Amounts: coins,
	}, liq)
}

func (suite *KeeperTestSuite) TestDecreaseLiquidity() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// decrease liquidity when its empty
	decreaseCoins := sdk.Coins{sdk.NewInt64Coin("qwoyn", 500), sdk.NewInt64Coin("ucoho", 500)}
	err := suite.app.GameKeeper.DecreaseLiquidity(suite.ctx, decreaseCoins)
	suite.Require().Error(err)

	// increase liquidity and decrease
	increaseCoins := sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)}
	suite.app.GameKeeper.IncreaseLiquidity(suite.ctx, increaseCoins)
	err = suite.app.GameKeeper.DecreaseLiquidity(suite.ctx, decreaseCoins)
	suite.Require().NoError(err)

	liq := suite.app.GameKeeper.GetLiquidity(suite.ctx)
	suite.Require().Equal(types.Liquidity{
		Amounts: increaseCoins.Sub(decreaseCoins),
	}, liq)
}

func (suite *KeeperTestSuite) TestSwapOutAmount() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// estimate swap out amount when its empty
	inAmount := sdk.NewInt64Coin("qwoyn", 1000)
	_, err := suite.app.GameKeeper.SwapOutAmount(suite.ctx, inAmount)
	suite.Require().Error(err)

	// increase liquidity
	increaseCoins := sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)}
	suite.app.GameKeeper.IncreaseLiquidity(suite.ctx, increaseCoins)

	// estimatate from qwoyn -> ucoho
	outAmount, err := suite.app.GameKeeper.SwapOutAmount(suite.ctx, inAmount)
	suite.Require().NoError(err)
	suite.Require().Equal(outAmount, sdk.NewInt64Coin("ucoho", 500))

	// estimatate from ucoho -> qwoyn
	inAmount = sdk.NewInt64Coin("ucoho", 1000)
	outAmount, err = suite.app.GameKeeper.SwapOutAmount(suite.ctx, inAmount)
	suite.Require().NoError(err)
	suite.Require().Equal(outAmount, sdk.NewInt64Coin("qwoyn", 500))
}
