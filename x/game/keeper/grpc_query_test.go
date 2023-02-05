package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestGRPCQueryLiquidity() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// get not available liquidity
	resp, err := suite.app.GameKeeper.Liquidity(sdk.WrapSDKContext(suite.ctx), &types.QueryLiquidityRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.Liquidity, types.Liquidity{})

	// set liquidity
	liquidity := types.Liquidity{
		Amounts: sdk.Coins{sdk.NewInt64Coin("ucoho", 1000), sdk.NewInt64Coin("qwoyn", 1000)},
	}
	suite.app.GameKeeper.SetLiquidity(suite.ctx, liquidity)

	resp, err = suite.app.GameKeeper.Liquidity(sdk.WrapSDKContext(suite.ctx), &types.QueryLiquidityRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.Liquidity, liquidity)
}

func (suite *KeeperTestSuite) TestGRPCQueryEstimatedSwapOut() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// get not available liquidity
	resp, err := suite.app.GameKeeper.EstimatedSwapOut(sdk.WrapSDKContext(suite.ctx), &types.QueryEstimatedSwapOutRequest{})
	suite.Require().Error(err)

	// set liquidity
	liquidity := types.Liquidity{
		Amounts: sdk.Coins{sdk.NewInt64Coin("ucoho", 1000), sdk.NewInt64Coin("qwoyn", 1000)},
	}
	suite.app.GameKeeper.SetLiquidity(suite.ctx, liquidity)

	resp, err = suite.app.GameKeeper.EstimatedSwapOut(sdk.WrapSDKContext(suite.ctx), &types.QueryEstimatedSwapOutRequest{
		Amount: "1000ucoho",
	})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.Amount.String(), "500qwoyn")

	resp, err = suite.app.GameKeeper.EstimatedSwapOut(sdk.WrapSDKContext(suite.ctx), &types.QueryEstimatedSwapOutRequest{
		Amount: "1000qwoyn",
	})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.Amount.String(), "500ucoho")
}

func (suite *KeeperTestSuite) TestGRPCQuerySwapRate() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// get not available liquidity
	resp, err := suite.app.GameKeeper.SwapRate(sdk.WrapSDKContext(suite.ctx), &types.QuerySwapRateRequest{})
	suite.Require().Error(err)

	// set liquidity
	liquidity := types.Liquidity{
		Amounts: sdk.Coins{sdk.NewInt64Coin("ucoho", 1000), sdk.NewInt64Coin("qwoyn", 1000)},
	}
	suite.app.GameKeeper.SetLiquidity(suite.ctx, liquidity)

	resp, err = suite.app.GameKeeper.SwapRate(sdk.WrapSDKContext(suite.ctx), &types.QuerySwapRateRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.Rate, sdk.OneDec())
	suite.Require().Equal(resp.SrcDenom, "ucoho")
	suite.Require().Equal(resp.TarDenom, "qwoyn")
}
