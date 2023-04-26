package keeper

import (
	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
)

var _ types.QueryServer = Keeper{}
