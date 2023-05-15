package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cosmic-horizon/qwoyn/osmosis/balancer"
	"github.com/cosmic-horizon/qwoyn/x/aquifer/keeper"
	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	minttypes "github.com/cosmic-horizon/qwoyn/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgServerBuyAllocationToken() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
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
			balance: sdk.Coins{sdk.NewInt64Coin(params.DepositToken, 10000)},
			deposit: sdk.NewInt64Coin(params.DepositToken, 100000),
			expPass: false,
		},
		{
			desc:    "successful deposit",
			balance: sdk.Coins{sdk.NewInt64Coin(params.DepositToken, 100000)},
			deposit: sdk.NewInt64Coin(params.DepositToken, 100000),
			expPass: true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			coins := sdk.Coins{sdk.NewInt64Coin(params.AllocationToken, 10000000)}
			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, coins)
			suite.Require().NoError(err)

			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.balance)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr, tc.balance)
			suite.Require().NoError(err)

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err = msgServer.BuyAllocationToken(sdk.WrapSDKContext(suite.ctx), &types.MsgBuyAllocationToken{
				Sender: addr.String(),
				Amount: tc.deposit,
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check balance has decreased
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr, params.DepositToken)
				suite.Require().Equal(balance.Amount, tc.balance.Sub(tc.deposit).AmountOf(params.DepositToken))

				// check module balance has increased
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
				balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, params.DepositToken)
				suite.Require().Equal(balance, tc.deposit)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerPutAllocationToken() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
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
			balance: sdk.Coins{sdk.NewInt64Coin(params.AllocationToken, 10000)},
			deposit: sdk.NewInt64Coin(params.AllocationToken, 100000),
			expPass: false,
		},
		{
			desc:    "successful deposit",
			balance: sdk.Coins{sdk.NewInt64Coin(params.AllocationToken, 100000)},
			deposit: sdk.NewInt64Coin(params.AllocationToken, 100000),
			expPass: true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.balance)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr, tc.balance)
			suite.Require().NoError(err)

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err = msgServer.PutAllocationToken(sdk.WrapSDKContext(suite.ctx), &types.MsgPutAllocationToken{
				Sender: addr.String(),
				Amount: tc.deposit,
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check balance has decreased
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr, params.AllocationToken)
				suite.Require().Equal(balance.Amount, tc.balance.Sub(tc.deposit).AmountOf(params.AllocationToken))

				// check module balance has increased
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
				balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, params.AllocationToken)
				suite.Require().Equal(balance, tc.deposit)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerTakeOutAllocationToken() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now()
	future := now.Add(time.Hour)

	for _, tc := range []struct {
		desc           string
		maintainer     string
		withdraw       sdk.Coin
		depositEndTime time.Time
		blockTime      time.Time
		expPass        bool
	}{
		{
			desc:           "invalid denom withdraw",
			maintainer:     addr.String(),
			withdraw:       sdk.NewInt64Coin("badtoken", 100000),
			depositEndTime: now,
			blockTime:      future,
			expPass:        false,
		},
		{
			desc:           "not enough balance",
			maintainer:     addr.String(),
			withdraw:       sdk.NewInt64Coin(params.AllocationToken, 100000000000),
			depositEndTime: now,
			blockTime:      future,
			expPass:        false,
		},
		{
			desc:           "not maintainer",
			maintainer:     "",
			withdraw:       sdk.NewInt64Coin(params.AllocationToken, 100000000000),
			depositEndTime: now,
			blockTime:      future,
			expPass:        false,
		},
		{
			desc:           "deposit end time not reached",
			maintainer:     "",
			withdraw:       sdk.NewInt64Coin(params.AllocationToken, 100000000000),
			depositEndTime: future,
			blockTime:      now,
			expPass:        false,
		},
		{
			desc:           "successful withdrawal",
			maintainer:     addr.String(),
			withdraw:       sdk.NewInt64Coin(params.AllocationToken, 100000),
			depositEndTime: now,
			blockTime:      future,
			expPass:        true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			params.Maintainer = tc.maintainer
			params.DepositEndTime = uint64(tc.depositEndTime.Unix())
			suite.app.AquiferKeeper.SetParams(suite.ctx, params)
			suite.ctx = suite.ctx.WithBlockTime(tc.blockTime)

			coins := sdk.Coins{sdk.NewInt64Coin(params.AllocationToken, 100000)}
			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, coins)
			suite.Require().NoError(err)

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err = msgServer.TakeOutAllocationToken(sdk.WrapSDKContext(suite.ctx), &types.MsgTakeOutAllocationToken{
				Sender: addr.String(),
				Amount: tc.withdraw,
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check balance has increased
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr, params.AllocationToken)
				suite.Require().Equal(balance.Amount, tc.withdraw.Amount)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerSetDepositEndTime() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now()
	future := now.Add(time.Hour)

	for _, tc := range []struct {
		desc           string
		maintainer     string
		depositEndTime time.Time
		expPass        bool
	}{
		{
			desc:           "not maintainer",
			maintainer:     "",
			depositEndTime: future,
			expPass:        false,
		},
		{
			desc:           "successful set",
			maintainer:     addr.String(),
			depositEndTime: future,
			expPass:        true,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			params.Maintainer = tc.maintainer
			suite.app.AquiferKeeper.SetParams(suite.ctx, params)

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err := msgServer.SetDepositEndTime(sdk.WrapSDKContext(suite.ctx), &types.MsgSetDepositEndTime{
				Sender:  addr.String(),
				EndTime: uint64(tc.depositEndTime.Unix()),
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check deposit end time changed
				params := suite.app.AquiferKeeper.GetParams(suite.ctx)
				suite.Require().Equal(params.DepositEndTime, uint64(tc.depositEndTime.Unix()))
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerInitICA() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc       string
		maintainer string
		expPass    bool
	}{
		{
			desc:       "not maintainer",
			maintainer: "",
			expPass:    false,
		},
		{
			desc:       "connection not found",
			maintainer: addr.String(),
			expPass:    false,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			params.Maintainer = tc.maintainer
			suite.app.AquiferKeeper.SetParams(suite.ctx, params)

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err := msgServer.InitICA(sdk.WrapSDKContext(suite.ctx), &types.MsgInitICA{
				Sender:       addr.String(),
				ConnectionId: "connection-0",
			})
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerExecTransfer() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc       string
		maintainer string
		expPass    bool
	}{
		{
			desc:       "not maintainer",
			maintainer: "",
			expPass:    false,
		},
		{
			desc:       "channel not found",
			maintainer: addr.String(),
			expPass:    false,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			params.Maintainer = tc.maintainer
			suite.app.AquiferKeeper.SetParams(suite.ctx, params)

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err := msgServer.ExecTransfer(sdk.WrapSDKContext(suite.ctx), &types.MsgExecTransfer{
				Sender:            addr.String(),
				TransferChannelId: "channel-1",
			})
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerExecAddLiquidity() {
	params := suite.app.AquiferKeeper.GetParams(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	for _, tc := range []struct {
		desc       string
		maintainer string
		expPass    bool
	}{
		{
			desc:       "not maintainer",
			maintainer: "",
			expPass:    false,
		},
		{
			desc:       "no active channel for this owner",
			maintainer: addr.String(),
			expPass:    false,
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			params.Maintainer = tc.maintainer
			suite.app.AquiferKeeper.SetParams(suite.ctx, params)

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err := msgServer.ExecAddLiquidity(sdk.WrapSDKContext(suite.ctx), &types.MsgExecAddLiquidity{
				Sender: addr.String(),
				Msg: balancer.MsgCreateBalancerPool{
					Sender: addr.String(),
				},
			})
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
