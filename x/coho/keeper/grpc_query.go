package keeper

import (
	"github.com/cosmic-horizon/coho/x/coho/types"
)

var _ types.QueryServer = Keeper{}
