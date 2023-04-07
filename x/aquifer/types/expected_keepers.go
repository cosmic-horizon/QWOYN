package types

import (
	gametypes "github.com/cosmic-horizon/qwoyn/x/game/types"
	minttypes "github.com/cosmic-horizon/qwoyn/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
	SetAccount(ctx sdk.Context, acc types.AccountI)
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// BankKeeper defines the expected game keeper
type GameKeeper interface {
	GetParamSet(ctx sdk.Context) gametypes.Params
	SwapFromModule(ctx sdk.Context, moduleName string, amount sdk.Coin) error
}

// MintKeeper defines the expected mint keeper
type MintKeeper interface {
	GetParams(ctx sdk.Context) minttypes.Params
}
