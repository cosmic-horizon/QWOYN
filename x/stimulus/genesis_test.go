package stimulus_test

import (
	"testing"

	"github.com/cosmic-horizon/qwoyn/x/stimulus/types"
)

func TestGenesis(t *testing.T) {
	_ = types.GenesisState{
		Params: types.DefaultParams(),
	}
}
