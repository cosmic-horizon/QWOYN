package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/qwoyn/x/game/keeper"
	"github.com/cosmic-horizon/qwoyn/x/game/types"
	minttypes "github.com/cosmic-horizon/qwoyn/x/mint/types"
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

func (suite *KeeperTestSuite) TestMsgServerAddLiquidity() {
	moduleOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc     string
		executor sdk.AccAddress
		balance  sdk.Coins
		deposit  sdk.Coins
		expPass  bool
	}{
		{
			desc:     "ensure owner to add liquidity",
			executor: addr1,
			balance:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			deposit:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			expPass:  false,
		},
		{
			desc:     "not enough balance",
			executor: moduleOwner,
			balance:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 100), sdk.NewInt64Coin("ucoho", 100)},
			deposit:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			expPass:  false,
		},
		{
			desc:     "1 coin deposit",
			executor: moduleOwner,
			balance:  sdk.Coins{sdk.NewInt64Coin("vvv", 1000)},
			deposit:  sdk.Coins{sdk.NewInt64Coin("vvv", 1000)},
			expPass:  false,
		},
		{
			desc:     "3 coins deposit",
			executor: moduleOwner,
			balance:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000), sdk.NewInt64Coin("zzz", 1000)},
			deposit:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000), sdk.NewInt64Coin("zzz", 1000)},
			expPass:  false,
		},
		{
			desc:     "successful deposit",
			executor: moduleOwner,
			balance:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			deposit:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			expPass:  true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			now := time.Now().UTC()
			suite.ctx = suite.ctx.WithBlockTime(now)

			params := suite.app.GameKeeper.GetParamSet(suite.ctx)
			params.Owner = moduleOwner.String()
			suite.app.GameKeeper.SetParamSet(suite.ctx, params)

			// allocate coins to executor
			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.balance)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, tc.balance)
			suite.Require().NoError(err)

			// execute liquidity addition
			msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
			_, err = msgServer.AddLiquidity(sdk.WrapSDKContext(suite.ctx), &types.MsgAddLiquidity{
				Sender:  tc.executor.String(),
				Amounts: tc.deposit,
			})

			if tc.expPass {
				suite.Require().NoError(err)

				// check liquidity is increased by correct amount
				liq := suite.app.GameKeeper.GetLiquidity(suite.ctx)
				suite.Require().Equal(sdk.Coins(liq.Amounts).String(), tc.deposit.String())

				// check coin movement by correct amount
				bal := suite.app.BankKeeper.GetAllBalances(suite.ctx, tc.executor)
				suite.Require().Equal(bal.String(), tc.balance.Sub(tc.deposit).String())

				// module address
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
				bal = suite.app.BankKeeper.GetAllBalances(suite.ctx, moduleAddr)
				suite.Require().Equal(bal.String(), tc.deposit.String())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerRemoveLiquidity() {
	moduleOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc      string
		executor  sdk.AccAddress
		liquidity sdk.Coins
		withdraw  sdk.Coins
		expPass   bool
	}{
		{
			desc:      "ensure owner to remove liquidity",
			executor:  addr1,
			liquidity: sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			withdraw:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			expPass:   false,
		},
		{
			desc:      "withdrawing more than existing liquidity",
			executor:  moduleOwner,
			liquidity: sdk.Coins{sdk.NewInt64Coin("qwoyn", 100), sdk.NewInt64Coin("ucoho", 100)},
			withdraw:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			expPass:   false,
		},
		{
			desc:      "successful withdraw",
			executor:  moduleOwner,
			liquidity: sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			withdraw:  sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			expPass:   true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			now := time.Now().UTC()
			suite.ctx = suite.ctx.WithBlockTime(now)

			params := suite.app.GameKeeper.GetParamSet(suite.ctx)
			params.Owner = moduleOwner.String()
			suite.app.GameKeeper.SetParamSet(suite.ctx, params)

			// allocate coins to module owner
			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.liquidity)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, moduleOwner, tc.liquidity)
			suite.Require().NoError(err)

			// execute liquidity addition
			msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
			_, err = msgServer.AddLiquidity(sdk.WrapSDKContext(suite.ctx), &types.MsgAddLiquidity{
				Sender:  moduleOwner.String(),
				Amounts: tc.liquidity,
			})
			suite.Require().NoError(err)

			// execute liquidity withdraw
			_, err = msgServer.RemoveLiquidity(sdk.WrapSDKContext(suite.ctx), &types.MsgRemoveLiquidity{
				Sender:  tc.executor.String(),
				Amounts: tc.withdraw,
			})

			if tc.expPass {
				suite.Require().NoError(err)

				// check liquidity is decreased by correct amount
				liq := suite.app.GameKeeper.GetLiquidity(suite.ctx)
				suite.Require().Equal(sdk.Coins(liq.Amounts).String(), tc.liquidity.Sub(tc.withdraw).String())

				// check coin movement by correct amount
				bal := suite.app.BankKeeper.GetAllBalances(suite.ctx, tc.executor)
				suite.Require().Equal(bal.String(), tc.withdraw.String())

				// module address
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
				bal = suite.app.BankKeeper.GetAllBalances(suite.ctx, moduleAddr)
				suite.Require().Equal(bal.String(), tc.liquidity.Sub(tc.withdraw).String())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerSwap() {
	moduleOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc         string
		executor     sdk.AccAddress
		liquidity    sdk.Coins
		inAmount     sdk.Coin
		expPass      bool
		expOutAmount sdk.Coin
	}{
		{
			desc:         "no liquidity available case",
			executor:     addr1,
			liquidity:    sdk.Coins{},
			inAmount:     sdk.NewInt64Coin("qwoyn", 1000),
			expPass:      false,
			expOutAmount: sdk.NewInt64Coin("ucoho", 500),
		},
		{
			desc:         "successful swap from qwoyn -> ucoho",
			executor:     addr1,
			liquidity:    sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			inAmount:     sdk.NewInt64Coin("qwoyn", 1000),
			expPass:      true,
			expOutAmount: sdk.NewInt64Coin("ucoho", 500),
		},
		{
			desc:         "successful swap from ucoho -> qwoyn",
			executor:     addr1,
			liquidity:    sdk.Coins{sdk.NewInt64Coin("qwoyn", 1000), sdk.NewInt64Coin("ucoho", 1000)},
			inAmount:     sdk.NewInt64Coin("ucoho", 1000),
			expPass:      true,
			expOutAmount: sdk.NewInt64Coin("qwoyn", 500),
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			now := time.Now().UTC()
			suite.ctx = suite.ctx.WithBlockTime(now)

			params := suite.app.GameKeeper.GetParamSet(suite.ctx)
			params.Owner = moduleOwner.String()
			suite.app.GameKeeper.SetParamSet(suite.ctx, params)

			// allocate coins to module owner
			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.liquidity)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, moduleOwner, tc.liquidity)
			suite.Require().NoError(err)

			// execute liquidity addition
			msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
			if tc.liquidity.String() != "" {
				_, err = msgServer.AddLiquidity(sdk.WrapSDKContext(suite.ctx), &types.MsgAddLiquidity{
					Sender:  moduleOwner.String(),
					Amounts: tc.liquidity,
				})
				suite.Require().NoError(err)
			}

			// execute swap operation
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.inAmount})
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{tc.inAmount})
			suite.Require().NoError(err)
			_, err = msgServer.Swap(sdk.WrapSDKContext(suite.ctx), &types.MsgSwap{
				Sender: tc.executor.String(),
				Amount: tc.inAmount,
			})

			if tc.expPass {
				suite.Require().NoError(err)

				// check liquidity changes bi-directional by correct amount
				liq := suite.app.GameKeeper.GetLiquidity(suite.ctx)
				expLiqRemaining := tc.liquidity.Add(tc.inAmount).Sub(sdk.Coins{tc.expOutAmount})
				suite.Require().Equal(sdk.Coins(liq.Amounts).String(), expLiqRemaining.String())

				// check coin movement bi-directional by correct amount
				bal := suite.app.BankKeeper.GetAllBalances(suite.ctx, tc.executor)
				suite.Require().Equal(bal.String(), tc.expOutAmount.String())

				// module address
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
				bal = suite.app.BankKeeper.GetAllBalances(suite.ctx, moduleAddr)
				suite.Require().Equal(bal.String(), expLiqRemaining.String())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
