package keeper_test

import (
	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
)

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.aquiferKeeper.GetParams(suite.ctx)

	params = types.Params{}

	suite.app.aquiferKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.aquiferKeeper.GetParams(suite.ctx)
	suite.Require().Equal(params, newParams)
}
