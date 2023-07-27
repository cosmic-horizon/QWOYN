package v5_2

import (
	"github.com/cosmic-horizon/qwoyn/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v5.2.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{group.ModuleName},
	},
}
