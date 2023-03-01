package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/capability/types"
)

//go:generate mockgen -source=expected_keepers.go -package mocks -destination mocks/expected_keepers.go

type CapabilityKeeper interface {
	ClaimCapability(ctx sdk.Context, cap *types.Capability, name string) error
	GetCapability(ctx sdk.Context, name string) (*types.Capability, bool)
}
