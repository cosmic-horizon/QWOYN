package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/qwoyn/x/game/keeper"
	"github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestEndBlocker() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)
	params := suite.app.GameKeeper.GetParamSet(suite.ctx)

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:         addr1.String(),
		Amount:          sdk.NewInt(2000),
		Staking:         sdk.NewInt(1000),
		Unbonding:       sdk.NewInt(0),
		RewardClaimTime: now,
	})

	future := now.Add(365 * 24 * time.Hour)
	suite.ctx = suite.ctx.WithBlockTime(future)

	// check reward is claimed
	msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
	_, err := msgServer.BeginUnstakeInGameToken(sdk.WrapSDKContext(suite.ctx), &types.MsgBeginUnstakeInGameToken{
		Sender: addr1.String(),
		Amount: sdk.NewInt64Coin(params.DepositDenom, 500),
	})
	suite.Require().NoError(err)

	future2 := future.Add(365 * 24 * time.Hour)
	suite.ctx = suite.ctx.WithBlockTime(future2)

	suite.app.GameKeeper.EndBlocker(suite.ctx)
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	increaseAmount := sdk.NewInt(int64(params.StakingInflation * 1000))
	suite.Require().Equal(deposit, types.Deposit{
		Address:         addr1.String(),
		Amount:          sdk.NewInt(2000).Add(increaseAmount),
		Staking:         sdk.NewInt(500),
		Unbonding:       sdk.NewInt(0),
		RewardClaimTime: future,
	})

	// get all unbondings after endblocker
	allUnbondings := suite.app.GameKeeper.GetAllUnbondings(suite.ctx)
	suite.Require().Len(allUnbondings, 0)
}
