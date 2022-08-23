package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestWhitelistedContractGetSetDelete() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// check when not whitelisted
	whitelisted := suite.app.GameKeeper.IsWhitelistedContract(suite.ctx, addr1.String())
	suite.Require().Equal(whitelisted, false)

	// get all whitelisted contracts when not available
	allWhitelistedContracts := suite.app.GameKeeper.GetAllWhitelistedContracts(suite.ctx)
	suite.Require().Len(allWhitelistedContracts, 0)

	// set contracts
	contracts := []string{addr1.String(), addr2.String()}

	for _, contract := range contracts {
		suite.app.GameKeeper.SetWhitelistedContract(suite.ctx, contract)
	}

	for _, contract := range contracts {
		isWhitelisted := suite.app.GameKeeper.IsWhitelistedContract(suite.ctx, contract)
		suite.Require().Equal(isWhitelisted, true)
	}

	allWhitelistedContracts = suite.app.GameKeeper.GetAllWhitelistedContracts(suite.ctx)
	suite.Require().Len(allWhitelistedContracts, 2)

	for _, contract := range contracts {
		suite.app.GameKeeper.DeleteWhitelistedContract(suite.ctx, contract)
	}

	// get all whitelisted contracts after all deletion
	allWhitelistedContracts = suite.app.GameKeeper.GetAllWhitelistedContracts(suite.ctx)
	suite.Require().Len(allWhitelistedContracts, 0)
}
