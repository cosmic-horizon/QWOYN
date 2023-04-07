package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/aquifer module sentinel errors
var (
	ErrNotBaseAccount = sdkerrors.Register(ModuleName, 1, "account type should be base account")
)
