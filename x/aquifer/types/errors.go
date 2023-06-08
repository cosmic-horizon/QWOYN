package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/aquifer module sentinel errors
var (
	ErrNotBaseAccount      = sdkerrors.Register(ModuleName, 1, "account type should be base account")
	ErrDepositTimeEnded    = sdkerrors.Register(ModuleName, 2, "deposit time ended")
	ErrDepositTimeNotEnded = sdkerrors.Register(ModuleName, 3, "deposit time not ended")
	ErrNotMaintainer       = sdkerrors.Register(ModuleName, 4, "not a maintainer")
	ErrConnectionNotFound  = sdkerrors.Register(ModuleName, 5, "connection not found")
)
