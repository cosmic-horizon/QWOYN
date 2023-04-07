package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
	params.DepositToken = "uusdc"
	params.AllocationToken = "nqwoyn"
	params.VestingDuration = 86400
	params.DepositEndTime = 10000
	params.InitLiquidityPrice = sdk.NewDecWithPrec(2, 1)
	params.LiquidityBootstrapping = true
	params.LiquidityBootstrapped = true
	params.IcsConnectionId = "connection-1"

	suite.app.AquiferKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.AquiferKeeper.GetParams(suite.ctx)
	suite.Require().Equal(params, newParams)
}
