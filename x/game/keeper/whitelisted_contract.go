package keeper

import (
	"github.com/cosmic-horizon/coho/x/game/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IsWhitelistedContract(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(append(types.PrefixWhitelistedContract, address...))
	return bz != nil
}

func (k Keeper) SetWhitelistedContract(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(types.PrefixWhitelistedContract, address...), address)
}

func (k Keeper) DeleteWhitelistedContract(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(types.PrefixWhitelistedContract, address...))
}

func (k Keeper) GetAllWhitelistedContracts(ctx sdk.Context) []sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, types.PrefixWhitelistedContract)
	defer it.Close()

	allContracts := []sdk.AccAddress{}
	for ; it.Valid(); it.Next() {
		allContracts = append(allContracts, sdk.AccAddress(it.Value()))
	}

	return allContracts
}
