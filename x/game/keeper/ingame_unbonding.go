package keeper

import (
	"time"

	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLastUnbondingId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastUnbondingIndex, sdk.Uint64ToBigEndian(id))
}

func (k Keeper) GetLastUnbondingId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLastUnbondingIndex)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) GetAllUnbondings(ctx sdk.Context) []types.Unbonding {
	store := ctx.KVStore(k.storeKey)

	unbondings := []types.Unbonding{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixUnbondingKey))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		unbonding := types.Unbonding{}
		k.cdc.MustUnmarshal(it.Value(), &unbonding)

		unbondings = append(unbondings, unbonding)
	}
	return unbondings
}

func (k Keeper) GetCompletedUnbondingsAt(ctx sdk.Context, endTime time.Time) []types.Unbonding {
	store := ctx.KVStore(k.storeKey)

	unbondings := []types.Unbonding{}
	it := store.Iterator(types.PrefixInGameUnbondingTimeKey, sdk.InclusiveEndBytes(types.InGameUnbondingTimePrefixKey(endTime)))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		id := sdk.BigEndianToUint64(it.Value())
		unbonding := k.GetUnbonding(ctx, id)
		unbondings = append(unbondings, unbonding)
	}
	return unbondings
}

func (k Keeper) GetUserUnbondings(ctx sdk.Context, addr sdk.AccAddress) []types.Unbonding {
	store := ctx.KVStore(k.storeKey)

	unbondings := []types.Unbonding{}
	it := store.Iterator(types.PrefixInGameUnbondingTimeKey, sdk.InclusiveEndBytes(types.InGameUnbondingUserPrefixKey(addr)))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		id := sdk.BigEndianToUint64(it.Value())
		unbonding := k.GetUnbonding(ctx, id)
		unbondings = append(unbondings, unbonding)
	}
	return unbondings
}

func (k Keeper) SetUnbonding(ctx sdk.Context, unbonding types.Unbonding) {
	bz := k.cdc.MustMarshal(&unbonding)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UnbondingKey(unbonding.Id), bz)

	idBytes := sdk.Uint64ToBigEndian(unbonding.Id)
	stakerAddr, err := sdk.AccAddressFromBech32(unbonding.StakerAddress)
	if err != nil {
		panic(err)
	}
	store.Set(types.InGameUnbondingUserKey(stakerAddr, unbonding.Id), idBytes)
	store.Set(types.InGameUnbondingTimeKey(unbonding.CompletionTime, unbonding.Id), idBytes)
}

func (k Keeper) DeleteUnbonding(ctx sdk.Context, unbonding types.Unbonding) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UnbondingKey(unbonding.Id))
	stakerAddr, err := sdk.AccAddressFromBech32(unbonding.StakerAddress)
	if err != nil {
		panic(err)
	}
	store.Delete(types.InGameUnbondingUserKey(stakerAddr, unbonding.Id))
	store.Delete(types.InGameUnbondingTimeKey(unbonding.CompletionTime, unbonding.Id))
}

func (k Keeper) GetUnbonding(ctx sdk.Context, id uint64) types.Unbonding {
	unbonding := types.Unbonding{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UnbondingKey(id))
	if bz == nil {
		return unbonding
	}
	k.cdc.MustUnmarshal(bz, &unbonding)
	return unbonding
}
