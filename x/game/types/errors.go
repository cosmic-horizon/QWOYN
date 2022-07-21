package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/game module sentinel errors
var (
	ErrNotModuleOwner         = sdkerrors.Register(ModuleName, 1, "not module owner")
	ErrNotWhitelistedContract = sdkerrors.Register(ModuleName, 2, "not whitelisted contract")
)
