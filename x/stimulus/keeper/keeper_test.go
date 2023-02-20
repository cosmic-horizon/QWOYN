package keeper_test

import (
	"testing"

	simapp "github.com/cosmic-horizon/qwoyn/app"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	isCheckTx = false
)

type KeeperTestSuite struct {
	suite.Suite

	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	app         *simapp.App
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(isCheckTx)

	suite.legacyAmino = app.LegacyAmino()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
