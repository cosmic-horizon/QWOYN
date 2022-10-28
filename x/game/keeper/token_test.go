package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestDepositGetSetDelete() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// get not available deposit
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit.Address, "")

	// get all deposits when not available
	allDeposits := suite.app.GameKeeper.GetAllDeposits(suite.ctx)
	suite.Require().Len(allDeposits, 0)

	// set deposits
	deposits := []types.Deposit{
		{
			Address:         addr1.String(),
			Amount:          sdk.NewInt(10000),
			Staking:         sdk.NewInt(0),
			Unbonding:       sdk.NewInt(0),
			RewardClaimTime: now,
		},
		{
			Address:         addr2.String(),
			Amount:          sdk.NewInt(10000),
			Staking:         sdk.NewInt(1000),
			Unbonding:       sdk.NewInt(500),
			RewardClaimTime: now,
		},
	}

	for _, deposit := range deposits {
		suite.app.GameKeeper.SetDeposit(suite.ctx, deposit)
	}

	for _, deposit := range deposits {
		addr, _ := sdk.AccAddressFromBech32(deposit.Address)
		d := suite.app.GameKeeper.GetDeposit(suite.ctx, addr)
		suite.Require().Equal(deposit, d)
	}

	allDeposits = suite.app.GameKeeper.GetAllDeposits(suite.ctx)
	suite.Require().Len(allDeposits, 2)

	for _, deposit := range deposits {
		addr, _ := sdk.AccAddressFromBech32(deposit.Address)
		suite.app.GameKeeper.DeleteDeposit(suite.ctx, addr)
	}

	// get all deposits after all deletion
	allDeposits = suite.app.GameKeeper.GetAllDeposits(suite.ctx)
	suite.Require().Len(allDeposits, 0)
}

func (suite *KeeperTestSuite) TestIncreaseDeposit() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// get not available deposit
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit.Address, "")

	// get all deposits when not available
	allDeposits := suite.app.GameKeeper.GetAllDeposits(suite.ctx)
	suite.Require().Len(allDeposits, 0)

	// deposit once
	suite.app.GameKeeper.IncreaseDeposit(suite.ctx, addr1, sdk.NewInt(1000))
	deposit = suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.ZeroInt(),
		Unbonding: sdk.ZeroInt(),
	})

	allDeposits = suite.app.GameKeeper.GetAllDeposits(suite.ctx)
	suite.Require().Len(allDeposits, 1)

	// deposit once more
	suite.app.GameKeeper.IncreaseDeposit(suite.ctx, addr1, sdk.NewInt(1000))
	deposit = suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(2000),
		Staking:   sdk.ZeroInt(),
		Unbonding: sdk.ZeroInt(),
	})
}

func (suite *KeeperTestSuite) TestDecreaseDeposit() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// decrease deposit when not available
	err := suite.app.GameKeeper.DecreaseDeposit(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.ZeroInt(),
		Unbonding: sdk.ZeroInt(),
	})

	// try decreasing more than deposit
	err = suite.app.GameKeeper.DecreaseDeposit(suite.ctx, addr1, sdk.NewInt(2000))
	suite.Require().Error(err)

	// decrease max
	err = suite.app.GameKeeper.DecreaseDeposit(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().NoError(err)
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.ZeroInt(),
		Staking:   sdk.ZeroInt(),
		Unbonding: sdk.ZeroInt(),
	})

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(500),
		Unbonding: sdk.ZeroInt(),
	})

	// try decreasing full after staking
	err = suite.app.GameKeeper.DecreaseDeposit(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// try decreasing considering staking
	err = suite.app.GameKeeper.DecreaseDeposit(suite.ctx, addr1, sdk.NewInt(500))
	suite.Require().NoError(err)
	deposit = suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(500),
		Staking:   sdk.NewInt(500),
		Unbonding: sdk.ZeroInt(),
	})
}

func (suite *KeeperTestSuite) TestIncreaseStaking() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// increase staking when deposit not available
	err := suite.app.GameKeeper.IncreaseStaking(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.ZeroInt(),
		Unbonding: sdk.ZeroInt(),
	})

	// try increase staking more than deposit
	err = suite.app.GameKeeper.IncreaseStaking(suite.ctx, addr1, sdk.NewInt(2000))
	suite.Require().Error(err)

	// stake max
	err = suite.app.GameKeeper.IncreaseStaking(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().NoError(err)
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.ZeroInt(),
	})

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(500),
		Unbonding: sdk.ZeroInt(),
	})

	// try staking full after staking partial
	err = suite.app.GameKeeper.IncreaseStaking(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// try staking considering previous staking
	err = suite.app.GameKeeper.IncreaseStaking(suite.ctx, addr1, sdk.NewInt(500))
	suite.Require().NoError(err)
	deposit = suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.ZeroInt(),
	})
}

func (suite *KeeperTestSuite) TestDecreaseStaking() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// decrease staking when not available
	err := suite.app.GameKeeper.DecreaseStaking(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.ZeroInt(),
	})

	// try decreasing more than total staking
	err = suite.app.GameKeeper.DecreaseStaking(suite.ctx, addr1, sdk.NewInt(2000))
	suite.Require().Error(err)

	// decrease max
	err = suite.app.GameKeeper.DecreaseStaking(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().NoError(err)
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(0),
		Unbonding: sdk.ZeroInt(),
	})

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.NewInt(500),
	})

	// try decreasing full after unbonding
	err = suite.app.GameKeeper.DecreaseStaking(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// try decreasing considering unbonding
	err = suite.app.GameKeeper.DecreaseStaking(suite.ctx, addr1, sdk.NewInt(500))
	suite.Require().NoError(err)
	deposit = suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(500),
		Unbonding: sdk.NewInt(500),
	})
}

func (suite *KeeperTestSuite) TestIncreaseUnbonding() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// increase staking when deposit not available
	err := suite.app.GameKeeper.IncreaseUnbonding(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.ZeroInt(),
	})

	// try increase staking more than deposit
	err = suite.app.GameKeeper.IncreaseUnbonding(suite.ctx, addr1, sdk.NewInt(2000))
	suite.Require().Error(err)

	// stake max
	err = suite.app.GameKeeper.IncreaseUnbonding(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().NoError(err)
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.NewInt(1000),
	})

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.NewInt(500),
	})

	// try unbonding full after unbonding partial
	err = suite.app.GameKeeper.IncreaseUnbonding(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// try unbonding considering previous unbonding
	err = suite.app.GameKeeper.IncreaseUnbonding(suite.ctx, addr1, sdk.NewInt(500))
	suite.Require().NoError(err)
	deposit = suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.NewInt(1000),
	})
}

func (suite *KeeperTestSuite) TestDecreaseUnbonding() {

	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// decrease unbonding when not available
	err := suite.app.GameKeeper.DecreaseUnbonding(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.NewInt(1000),
	})

	// try decreasing more than total staking
	err = suite.app.GameKeeper.DecreaseUnbonding(suite.ctx, addr1, sdk.NewInt(2000))
	suite.Require().Error(err)

	// decrease max
	err = suite.app.GameKeeper.DecreaseUnbonding(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().NoError(err)
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.ZeroInt(),
	})

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.NewInt(500),
	})

	// try decreasing more than unbonding
	err = suite.app.GameKeeper.DecreaseUnbonding(suite.ctx, addr1, sdk.NewInt(1000))
	suite.Require().Error(err)

	// try decreasing considering unbonding value
	err = suite.app.GameKeeper.DecreaseUnbonding(suite.ctx, addr1, sdk.NewInt(500))
	suite.Require().NoError(err)
	deposit = suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	suite.Require().Equal(deposit, types.Deposit{
		Address:   addr1.String(),
		Amount:    sdk.NewInt(1000),
		Staking:   sdk.NewInt(1000),
		Unbonding: sdk.NewInt(0),
	})
}

func (suite *KeeperTestSuite) TestClaimInGameStakingReward() {

	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// set deposit
	suite.app.GameKeeper.SetDeposit(suite.ctx, types.Deposit{
		Address:         addr1.String(),
		Amount:          sdk.NewInt(1000),
		Staking:         sdk.NewInt(1000),
		Unbonding:       sdk.NewInt(0),
		RewardClaimTime: now,
	})

	future := now.Add(365 * 24 * time.Hour)
	suite.ctx = suite.ctx.WithBlockTime(future)

	// claim staking rewards
	err := suite.app.GameKeeper.ClaimInGameStakingReward(suite.ctx, addr1)
	suite.Require().NoError(err)

	// check reward amount is correctly inreased on deposit object
	params := suite.app.GameKeeper.GetParamSet(suite.ctx)
	deposit := suite.app.GameKeeper.GetDeposit(suite.ctx, addr1)
	increaseAmount := sdk.NewInt(int64(params.StakingInflation * 1000))
	suite.Require().Equal(deposit.Amount, sdk.NewInt(1000).Add(increaseAmount))

	// check reward claim time is set correctly
	suite.Require().Equal(deposit.RewardClaimTime, future)

	// check tokens are minted correctly into game module
	moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
	balance := suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, params.DepositDenom)
	suite.Require().Equal(increaseAmount, balance.Amount)
}
