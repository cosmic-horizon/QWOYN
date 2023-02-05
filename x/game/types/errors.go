package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/game module sentinel errors
var (
	ErrNotModuleOwner                   = sdkerrors.Register(ModuleName, 1, "not module owner")
	ErrNotWhitelistedContract           = sdkerrors.Register(ModuleName, 2, "not whitelisted contract")
	ErrMinterIsNotModuleAddress         = sdkerrors.Register(ModuleName, 3, "not the minter of the nft contract")
	ErrOwnerIsNotModuleAddress          = sdkerrors.Register(ModuleName, 4, "not the owner of the nft contract")
	ErrSignerAccountNotRegistered       = sdkerrors.Register(ModuleName, 5, "signer account is not registered")
	ErrSignerAccountPubKeyNotRegistered = sdkerrors.Register(ModuleName, 6, "signer account public key is not registered")
	ErrInvalidDepositDenom              = sdkerrors.Register(ModuleName, 7, "invalid deposit denom")
	ErrInvalidWithdrawDenom             = sdkerrors.Register(ModuleName, 8, "invalid withdraw denom")
	ErrInsufficientDepositAmount        = sdkerrors.Register(ModuleName, 9, "insufficient deposit amount")
	ErrInsufficientStakingAmount        = sdkerrors.Register(ModuleName, 10, "insufficient staking amount")
	ErrInsufficientUnbondingAmount      = sdkerrors.Register(ModuleName, 11, "insufficient unbonding amount")
	ErrInsufficientLiquidityAmount      = sdkerrors.Register(ModuleName, 12, "insufficient liquidity amount")
	ErrLiquidityShouldHoldTwoTokens     = sdkerrors.Register(ModuleName, 13, "liquidity should hold two tokens")
)
