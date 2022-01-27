package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmic-horizon/coho/testutil/keeper"
	"github.com/cosmic-horizon/coho/x/coho/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CohoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
