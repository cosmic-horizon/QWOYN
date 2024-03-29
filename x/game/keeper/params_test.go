package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestParamsGetSet() {
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params := types.Params{
		Owner:            addr.String(),
		DepositDenom:     "stake",
		StakingInflation: sdk.NewDec(10),
		UnstakingTime:    time.Hour * 24,
		SwapFeeCollector: addr.String(),
		SwapFee:          sdk.NewDecWithPrec(1, 1),
	}

	suite.app.GameKeeper.SetParamSet(suite.ctx, params)
	newParams := suite.app.GameKeeper.GetParamSet(suite.ctx)
	suite.Require().Equal(params, newParams)
}
