package keeper

import (
	"github.com/cosmic-horizon/qwoyn/x/coho/types"
)

var _ types.QueryServer = Keeper{}
