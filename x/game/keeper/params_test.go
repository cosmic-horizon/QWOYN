package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.GameKeeper.GetParamSet(suite.ctx)

	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params = types.Params{
		Owner:            addr.String(),
		DepositDenom:     "stake",
		StakingInflation: 10,
		UnstakingTime:    time.Hour * 24,
	}

	suite.app.GameKeeper.SetParamSet(suite.ctx, params)
	newParams := suite.app.GameKeeper.GetParamSet(suite.ctx)
	suite.Require().Equal(params, newParams)
}
