package v5_3

import (
	"github.com/cosmic-horizon/qwoyn/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v5.3.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{group.ModuleName},
	},
}
