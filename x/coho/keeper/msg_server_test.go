package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmic-horizon/qwoyn/testutil/keeper"
	"github.com/cosmic-horizon/qwoyn/x/coho/keeper"
	"github.com/cosmic-horizon/qwoyn/x/coho/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CohoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
