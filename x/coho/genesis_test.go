package coho_test

import (
	"testing"

	keepertest "github.com/cosmic-horizon/coho/testutil/keeper"
	"github.com/cosmic-horizon/coho/testutil/nullify"
	"github.com/cosmic-horizon/coho/x/coho"
	"github.com/cosmic-horizon/coho/x/coho/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CohoKeeper(t)
	coho.InitGenesis(ctx, *k, genesisState)
	got := coho.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
