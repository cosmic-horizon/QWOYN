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
	k.SetLastUnbondingId(ctx, genState.LastUnbondingId)

	for _, deposit := range genState.Deposits {
		k.SetDeposit(ctx, deposit)
	}

	for _, contract := range genState.WhitelistedContracts {
		k.SetWhitelistedContract(ctx, contract)
	}

	for _, unbonding := range genState.Unbondings {
		k.SetUnbonding(ctx, unbonding)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParamSet(ctx)
	genesis.Deposits = k.GetAllDeposits(ctx)
	genesis.WhitelistedContracts = k.GetAllWhitelistedContracts(ctx)
	genesis.Unbondings = k.GetAllUnbondings(ctx)
	genesis.LastUnbondingId = k.GetLastUnbondingId(ctx)

	return genesis
}
