package keeper_test

import (
	"time"

	"github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestLastUnbondingId() {
	// check initial last unbonding id
	lastUnbondingId := suite.app.GameKeeper.GetLastUnbondingId(suite.ctx)
	suite.Require().Equal(lastUnbondingId, uint64(0))

	// set new last unbonding id
	suite.app.GameKeeper.SetLastUnbondingId(suite.ctx, 10)

	// check updated last unbonding id
	lastUnbondingId = suite.app.GameKeeper.GetLastUnbondingId(suite.ctx)
	suite.Require().Equal(lastUnbondingId, uint64(10))
}

func (suite *KeeperTestSuite) TesetUnbondingsGetSetDelete() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	future := time.Now().UTC().Add(time.Second)
	suite.ctx = suite.ctx.WithBlockTime(now)

	// get not available unbonding
	unbonding := suite.app.GameKeeper.GetUnbonding(suite.ctx, 0)
	suite.Require().Equal(unbonding, types.Unbonding{})

	// get all unbondings when not available
	allUnbondings := suite.app.GameKeeper.GetAllUnbondings(suite.ctx)
	suite.Require().Len(allUnbondings, 0)

	// get ending queue unbondings when not available
	maturedUnbondings := suite.app.GameKeeper.GetCompletedUnbondingsAt(suite.ctx, now)
	suite.Require().Len(maturedUnbondings, 0)

	// get user unbondings
	userUnbondings := suite.app.GameKeeper.GetUserUnbondings(suite.ctx, addr1)
	suite.Require().Len(userUnbondings, 0)

	// set unbondings
	unbondings := []types.Unbonding{
		{
			Id:             1,
			StakerAddress:  addr1.String(),
			CreationHeight: 1,
			CompletionTime: now,
			Amount:         sdk.NewInt(100),
		},
		{
			Id:             2,
			StakerAddress:  addr1.String(),
			CreationHeight: 2,
			CompletionTime: future,
			Amount:         sdk.NewInt(100),
		},
		{
			Id:             3,
			StakerAddress:  addr2.String(),
			CreationHeight: 1,
			CompletionTime: now,
			Amount:         sdk.NewInt(100),
		},
	}

	for _, unbonding := range unbondings {
		suite.app.GameKeeper.SetUnbonding(suite.ctx, unbonding)
	}

	for _, unbonding := range unbondings {
		u := suite.app.GameKeeper.GetUnbonding(suite.ctx, unbonding.Id)
		suite.Require().Equal(unbonding, u)
	}

	allUnbondings = suite.app.GameKeeper.GetAllUnbondings(suite.ctx)
	suite.Require().Len(allUnbondings, 4)
	suite.Require().Equal(unbondings, allUnbondings)

	maturedUnbondings = suite.app.GameKeeper.GetCompletedUnbondingsAt(suite.ctx, now)
	suite.Require().Len(maturedUnbondings, 2)

	// get user unbondings
	userUnbondings = suite.app.GameKeeper.GetUserUnbondings(suite.ctx, addr1)
	suite.Require().Len(userUnbondings, 2)

	for _, unbonding := range unbondings {
		suite.app.GameKeeper.DeleteUnbonding(suite.ctx, unbonding)
	}

	// get all unbondings after all deletion
	allUnbondings = suite.app.GameKeeper.GetAllUnbondings(suite.ctx)
	suite.Require().Len(allUnbondings, 0)

	// get ending queue unbondings after all deletion
	maturedUnbondings = suite.app.GameKeeper.GetCompletedUnbondingsAt(suite.ctx, now)
	suite.Require().Len(maturedUnbondings, 0)

	// get user unbondings
	userUnbondings = suite.app.GameKeeper.GetUserUnbondings(suite.ctx, addr1)
	suite.Require().Len(userUnbondings, 0)
}
