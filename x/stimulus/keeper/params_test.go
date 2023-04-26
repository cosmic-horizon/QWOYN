package keeper_test

import (
	"github.com/cosmic-horizon/qwoyn/x/stimulus/types"
)

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := types.Params{}

	suite.app.StimulusKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.StimulusKeeper.GetParams(suite.ctx)
	suite.Require().Equal(params, newParams)
}
