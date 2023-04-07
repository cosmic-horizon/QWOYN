package keeper_test

import (
	"github.com/cosmic-horizon/qwoyn/x/aquifer/keeper"
	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	minttypes "github.com/cosmic-horizon/qwoyn/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestMsgServerBuyAllocationToken() {
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

			msgServer := keeper.NewMsgServerImpl(suite.app.AquiferKeeper)
			_, err = msgServer.BuyAllocationToken(sdk.WrapSDKContext(suite.ctx), &types.MsgBuyAllocationToken{
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
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
