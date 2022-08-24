package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/coho/x/game/keeper"
	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestMsgServerDepositToken() {
	params := suite.app.GameKeeper.GetParamSet(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc    string
		balance sdk.Coins
		deposit sdk.Coin
		expPass bool
	}{
		{
			desc:    "invalid denom deposit",
			balance: sdk.Coins{sdk.NewInt64Coin("badtoken", 100000)},
			deposit: sdk.NewInt64Coin("badtoken", 100000),
			expPass: false,
		},
		{
			desc:    "not enough balance",
			balance: sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 10000)},
			deposit: sdk.NewInt64Coin(params.DepositDenom, 100000),
			expPass: false,
		},
		{
			desc:    "successful deposit",
			balance: sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 100000)},
			deposit: sdk.NewInt64Coin(params.DepositDenom, 100000),
			expPass: true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, tc.balance)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, addr, tc.balance)
			suite.Require().NoError(err)

			msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
			_, err = msgServer.DepositToken(sdk.WrapSDKContext(suite.ctx), &types.MsgDepositToken{
				Sender: addr.String(),
				Amount: tc.deposit,
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check balance has decreased
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr, params.DepositDenom)
				suite.Require().Equal(balance.Amount, tc.balance.Sub(sdk.Coins{tc.deposit}).AmountOf(params.DepositDenom))

				// check module balance has increased
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
				balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, params.DepositDenom)
				suite.Require().Equal(balance, tc.deposit)

				// check deposit has increased
				deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr)
				suite.Require().Equal(deposit.Amount, tc.deposit.Amount)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerWithdrawToken() {
	params := suite.app.GameKeeper.GetParamSet(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc     string
		deposit  sdk.Coins
		withdraw sdk.Coin
		expPass  bool
	}{
		{
			desc:     "invalid denom withdraw",
			deposit:  sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 100000)},
			withdraw: sdk.NewInt64Coin("badtoken", 100000),
			expPass:  false,
		},
		{
			desc:     "not enough deposit",
			deposit:  sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 10000)},
			withdraw: sdk.NewInt64Coin(params.DepositDenom, 100000),
			expPass:  false,
		},
		{
			desc:     "successful withdraw",
			deposit:  sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 100000)},
			withdraw: sdk.NewInt64Coin(params.DepositDenom, 10000),
			expPass:  true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, tc.deposit)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, addr, tc.deposit)
			suite.Require().NoError(err)

			msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
			_, err = msgServer.DepositToken(sdk.WrapSDKContext(suite.ctx), &types.MsgDepositToken{
				Sender: addr.String(),
				Amount: tc.deposit[0],
			})
			suite.Require().NoError(err)

			_, err = msgServer.WithdrawToken(sdk.WrapSDKContext(suite.ctx), &types.MsgWithdrawToken{
				Sender: addr.String(),
				Amount: tc.withdraw,
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check balance has increased
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr, params.DepositDenom)
				suite.Require().Equal(balance.Amount, tc.withdraw.Amount)

				// check module balance has decreased
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
				balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, params.DepositDenom)
				suite.Require().Equal(balance.Amount, tc.deposit.Sub(sdk.Coins{tc.withdraw}).AmountOf(params.DepositDenom))

				// check deposit has decreased
				deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr)
				suite.Require().Equal(deposit.Amount, tc.deposit.Sub(sdk.Coins{tc.withdraw}).AmountOf(params.DepositDenom))
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerStakeInGameToken() {
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

	// claim staking rewards
	msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
	_, err := msgServer.StakeInGameToken(sdk.WrapSDKContext(suite.ctx), &types.MsgStakeInGameToken{
		Sender: addr1.String(),
		Amount: sdk.NewInt64Coin(params.DepositDenom, 500),
	})
	suite.Require().NoError(err)

	// check reward amount is correctly inreased on deposit object
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	increaseAmount := sdk.NewInt(int64(params.StakingInflation * 1000))
	suite.Require().Equal(deposit.Amount, sdk.NewInt(2000).Add(increaseAmount))

	// check staking amount is increased
	suite.Require().Equal(deposit.Staking, sdk.NewInt(1500))
}

func (suite *KeeperTestSuite) TestMsgServerBeginUnstakeInGameToken() {
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

	// check reward amount is correctly increased on deposit object
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	increaseAmount := sdk.NewInt(int64(params.StakingInflation * 1000))
	suite.Require().Equal(deposit.Amount, sdk.NewInt(2000).Add(increaseAmount))

	// check staking not changed
	suite.Require().Equal(deposit.Staking, sdk.NewInt(1000))

	// check unbonding value increased
	suite.Require().Equal(deposit.Unbonding, sdk.NewInt(500))

	// check last unbonding id increased
	lastUnbondingId := suite.app.GameKeeper.GetLastUnbondingId(suite.ctx)
	suite.Require().Equal(lastUnbondingId, uint64(1))

	// check unbonding object is correctly set
	unbondings := suite.app.GameKeeper.GetAllUnbondings(suite.ctx)
	suite.Require().Len(unbondings, 1)
	suite.Require().Equal(unbondings[0], types.Unbonding{
		Id:             1,
		StakerAddress:  addr1.String(),
		CreationHeight: suite.ctx.BlockHeight(),
		CompletionTime: suite.ctx.BlockTime().UTC().Add(params.UnstakingTime),
		Amount:         sdk.NewInt(500),
	})
}
