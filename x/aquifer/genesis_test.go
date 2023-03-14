package aquifer_test

import (
	"testing"

	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
)

func TestGenesis(t *testing.T) {
	_ = types.GenesisState{
		Params: types.DefaultParams(),
	}
}
