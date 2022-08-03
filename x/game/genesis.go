package game

import (
	"github.com/cosmic-horizon/coho/x/game/keeper"
	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParamSet(ctx, genState.Params)

	for _, deposit := range genState.Deposits {
		addr, err := sdk.AccAddressFromBech32(deposit.Address)
		if err != nil {
			panic(err)
		}
		k.SetDeposit(ctx, addr, deposit.Amount)
	}

	for _, contract := range genState.WhitelistedContracts {
		k.SetWhitelistedContract(ctx, contract)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParamSet(ctx)
	genesis.Deposits = k.GetAllDeposits(ctx)
	genesis.WhitelistedContracts = k.GetAllWhitelistedContracts(ctx)

	return genesis
}
