package keeper

import (
	"github.com/cosmic-horizon/coho/x/game/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IsWhitelistedContract(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(append(types.PrefixWhitelistedContract, address...))
	return bz != nil
}

func (k Keeper) SetWhitelistedContract(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(types.PrefixWhitelistedContract, address...), []byte(address))
}

func (k Keeper) DeleteWhitelistedContract(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(types.PrefixWhitelistedContract, address...))
}

func (k Keeper) GetAllWhitelistedContracts(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, types.PrefixWhitelistedContract)
	defer it.Close()

	allContracts := []string{}
	for ; it.Valid(); it.Next() {
		allContracts = append(allContracts, string(it.Value()))
	}

	return allContracts
}
