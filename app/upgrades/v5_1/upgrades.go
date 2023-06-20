package v5_1

import (
	"github.com/cosmic-horizon/qwoyn/app/keepers"
	minttypes "github.com/cosmic-horizon/qwoyn/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("start to run module migrations...")

		gameManager := sdk.MustAccAddressFromBech32(keepers.GameKeeper.GetParamSet(ctx).Owner)
		cohoAmount := sdk.NewInt(500_000_000).Mul(sdk.NewInt(1000_000)) // 500M COHO
		err := keepers.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin("ucoho", cohoAmount)))
		if err != nil {
			return vm, err
		}

		err = keepers.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, gameManager, sdk.NewCoins(sdk.NewCoin("ucoho", cohoAmount)))
		if err != nil {
			return vm, err
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
