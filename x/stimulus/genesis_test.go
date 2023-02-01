package stimulus_test

import (
	"testing"

	keepertest "github.com/cosmic-horizon/qwoyn/testutil/keeper"
	"github.com/cosmic-horizon/qwoyn/testutil/nullify"
	"github.com/cosmic-horizon/qwoyn/x/stimulus"
	"github.com/cosmic-horizon/qwoyn/x/stimulus/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.CohoKeeper(t)
	stimulus.InitGenesis(ctx, *k, genesisState)
	got := stimulus.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
