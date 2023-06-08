package keeper_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmic-horizon/qwoyn/x/game/keeper"
	"github.com/cosmic-horizon/qwoyn/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func mustLoad(path string) []byte {
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}

var (
	shipNftContract = mustLoad("./cosmwasm/ship_nft.wasm")
)

func (suite *KeeperTestSuite) TestNftDepositWithdraw() {
	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	privKey1 := secp256k1.GenPrivKey()
	pubKey1 := privKey1.PubKey()
	addr1 := sdk.AccAddress(pubKey1.Address().Bytes())

	wasmParams := suite.app.WasmKeeper.GetParams(suite.ctx)
	wasmParams.CodeUploadAccess = wasmtypes.AllowEverybody
	wasmParams.InstantiateDefaultPermission = wasmtypes.AccessTypeEverybody
	suite.app.WasmKeeper.SetParams(suite.ctx, wasmParams)

	// store wasm code
	codeID, _, err := suite.app.GameKeeper.WasmKeeper.Create(suite.ctx, addr1, shipNftContract, &wasmtypes.AllowEverybody)
	suite.Require().NoError(err)

	moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)

	// instantiate contract
	initMsg := fmt.Sprintf(`{"name":"Ship NFT","symbol":"SHIP","minter":"%s","owner":"%s"}`, moduleAddr.String(), moduleAddr.String())
	contractAddr, _, err := suite.app.GameKeeper.WasmKeeper.Instantiate(suite.ctx, codeID, addr1, addr1, []byte(initMsg), "Ship-NFT", sdk.Coins{})
	suite.Require().NoError(err)

	// update game params to set owner to addr1
	gameParams := suite.app.GameKeeper.GetParamSet(suite.ctx)
	gameParams.Owner = addr1.String()
	suite.app.GameKeeper.SetParamSet(suite.ctx, gameParams)
	suite.app.AccountKeeper.SetAccount(suite.ctx, authtypes.NewBaseAccount(addr1, pubKey1, 0, 0))

	// whitelist nft contract
	msgServer := keeper.NewMsgServerImpl(suite.app.GameKeeper)
	_, err = msgServer.WhitelistNftContracts(sdk.WrapSDKContext(suite.ctx), types.NewMsgWhitelistNftContracts(addr1, []string{contractAddr.String()}))
	suite.Require().NoError(err)

	// check whitelist result
	contracts := suite.app.GameKeeper.GetAllWhitelistedContracts(suite.ctx)
	suite.Require().Len(contracts, 1)

	// mint and withdraw
	execMsg := fmt.Sprintf(`{"mint":{"token_id":"1","owner":"%s","extension":{"ship_type":12,"owner":"300"}}}`, moduleAddr.String())
	signerMsg := types.NewMsgSignerWithdrawUpdatedNft(addr1, contractAddr.String(), 1, execMsg)
	signBytes := signerMsg.GetSignBytes()
	signature, err := privKey1.Sign(signBytes)
	suite.Require().NoError(err)
	err = suite.app.GameKeeper.WithdrawUpdatedNft(suite.ctx, types.NewMsgWithdrawUpdatedNft(addr1, contractAddr.String(), 1, execMsg, signature))
	suite.Require().NoError(err)

	// check ownership by addr1
	resp, err := suite.app.WasmKeeper.QuerySmart(suite.ctx, contractAddr, []byte(`{"owner_of":{"token_id":"1"}}`))
	suite.Require().NoError(err)
	suite.Require().True(bytes.Contains(resp, []byte(addr1.String())))

	// deposit
	err = suite.app.GameKeeper.DepositNft(suite.ctx, types.NewMsgDepositNft(addr1, contractAddr.String(), 1))
	suite.Require().NoError(err)

	// check ownership transfer to module
	resp, err = suite.app.WasmKeeper.QuerySmart(suite.ctx, contractAddr, []byte(`{"owner_of":{"token_id":"1"}}`))
	suite.Require().NoError(err)
	suite.Require().True(bytes.Contains(resp, []byte(moduleAddr.String())))
}
