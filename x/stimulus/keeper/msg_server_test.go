package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	minttypes "github.com/cosmic-horizon/qwoyn/x/mint/types"
	"github.com/cosmic-horizon/qwoyn/x/stimulus/keeper"
	"github.com/cosmic-horizon/qwoyn/x/stimulus/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgServerDepositIntoOutpostFunding() {
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
			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.balance)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr, tc.balance)
			suite.Require().NoError(err)

			msgServer := keeper.NewMsgServerImpl(suite.app.StimulusKeeper)
			_, err = msgServer.DepositIntoOutpostFunding(sdk.WrapSDKContext(suite.ctx), &types.MsgDepositIntoOutpostFunding{
				Sender: addr.String(),
				Amount: tc.deposit,
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check balance has decreased
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr, params.DepositDenom)
				suite.Require().Equal(balance.Amount, tc.balance.Sub(tc.deposit).AmountOf(params.DepositDenom))

				// check module balance has increased
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.OutpostFundingPoolName)
				balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, params.DepositDenom)
				suite.Require().Equal(balance, tc.deposit)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServerWithdrawFromOutpostFunding() {
	params := suite.app.GameKeeper.GetParamSet(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params.Owner = addr.String()

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
			suite.app.GameKeeper.SetParamSet(suite.ctx, params)

			err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.deposit)
			suite.Require().NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr, tc.deposit)
			suite.Require().NoError(err)

			msgServer := keeper.NewMsgServerImpl(suite.app.StimulusKeeper)
			_, err = msgServer.DepositIntoOutpostFunding(sdk.WrapSDKContext(suite.ctx), &types.MsgDepositIntoOutpostFunding{
				Sender: addr.String(),
				Amount: tc.deposit[0],
			})
			suite.Require().NoError(err)

			_, err = msgServer.WithdrawFromOutpostFunding(sdk.WrapSDKContext(suite.ctx), &types.MsgWithdrawFromOutpostFunding{
				Sender: addr.String(),
				Amount: tc.withdraw,
			})
			if tc.expPass {
				suite.Require().NoError(err)

				// check balance has increased
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr, params.DepositDenom)
				suite.Require().Equal(balance.Amount, tc.withdraw.Amount)

				// check module balance has decreased
				moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.OutpostFundingPoolName)
				balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, params.DepositDenom)
				suite.Require().Equal(balance.Amount, tc.deposit.Sub(tc.withdraw).AmountOf(params.DepositDenom))
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
