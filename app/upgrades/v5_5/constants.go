package v5_4

import (
	"github.com/cosmic-horizon/qwoyn/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	providertypes "github.com/cosmos/interchain-security/v3/x/ccv/provider/types"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v5.5.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{providertypes.StoreKey},
	},
}
