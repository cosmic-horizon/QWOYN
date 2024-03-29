package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmic-horizon/qwoyn/x/aquifer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
)

type Keeper struct {
	cdc                 codec.BinaryCodec
	storeKey            storetypes.StoreKey
	paramstore          paramtypes.Subspace
	ak                  types.AccountKeeper
	bk                  types.BankKeeper
	gk                  types.GameKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
	IBCKeeper           ibckeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper

	ScopedKeeper     capabilitykeeper.ScopedKeeper
	IBCScopperKeeper capabilitykeeper.ScopedKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	gk types.GameKeeper,
	icacontrollerKeeper icacontrollerkeeper.Keeper,
	ibcKeeper ibckeeper.Keeper,
	TransferKeeper ibctransferkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	IBCScopperKeeper capabilitykeeper.ScopedKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                 cdc,
		storeKey:            storeKey,
		paramstore:          ps,
		ak:                  ak,
		bk:                  bk,
		gk:                  gk,
		icaControllerKeeper: icacontrollerKeeper,
		IBCKeeper:           ibcKeeper,
		TransferKeeper:      TransferKeeper,
		ScopedKeeper:        scopedKeeper,
		IBCScopperKeeper:    IBCScopperKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.ScopedKeeper.ClaimCapability(ctx, cap, name)
}
