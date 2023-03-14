package keeper_test

import (
	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	gamekeeper "github.com/cosmic-horizon/qwoyn/x/game/keeper"
	gametypes "github.com/cosmic-horizon/qwoyn/x/game/types"
	minttypes "github.com/cosmic-horizon/qwoyn/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestBeginBlocker() {
	params := suite.app.GameKeeper.GetParamSet(suite.ctx)
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params.Owner = addr.String()
	mintDenom := suite.app.MintKeeper.GetParams(suite.ctx).MintDenom

	for _, tc := range []struct {
		desc       string
		liquidity  sdk.Coins
		balance    sdk.Coin
		expBalance sdk.Coins
	}{
		{
			desc:       "no liquidity to swap",
			liquidity:  sdk.Coins{},
			balance:    sdk.NewInt64Coin(mintDenom, 1000),
			expBalance: sdk.Coins{sdk.NewInt64Coin(mintDenom, 1000)},
		},
		{
			desc:       "no balance to swap",
			liquidity:  sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 100000), sdk.NewInt64Coin(mintDenom, 100000)},
			balance:    sdk.NewInt64Coin(params.DepositDenom, 0),
			expBalance: sdk.Coins{},
		},
		{
			desc:       "successful swap",
			liquidity:  sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 100000), sdk.NewInt64Coin(mintDenom, 100000)},
			balance:    sdk.NewInt64Coin(mintDenom, 1000),
			expBalance: sdk.Coins{sdk.NewInt64Coin(params.DepositDenom, 991)},
		},
	} {
		suite.Run(tc.desc, func() {
			suite.SetupTest()
			suite.app.GameKeeper.SetParamSet(suite.ctx, params)

			if !tc.liquidity.IsZero() {
				err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.liquidity)
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr, tc.liquidity)
				suite.Require().NoError(err)

				gameMsgServer := gamekeeper.NewMsgServerImpl(suite.app.GameKeeper)
				_, err = gameMsgServer.AddLiquidity(sdk.WrapSDKContext(suite.ctx), &gametypes.MsgAddLiquidity{
					Sender:  addr.String(),
					Amounts: tc.liquidity,
				})
			}

			if tc.balance.IsPositive() {
				err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.balance})
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.OutpostFundingPoolName, sdk.Coins{tc.balance})
				suite.Require().NoError(err)
			}

			suite.app.aquiferKeeper.BeginBlocker(suite.ctx)

			// check balance has increased
			moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.OutpostFundingPoolName)
			balance := suite.app.BankKeeper.GetAllBalances(suite.ctx, moduleAddr)
			suite.Require().Equal(balance.String(), tc.expBalance.String())
		})
	}
}
