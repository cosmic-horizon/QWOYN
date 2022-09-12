package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "game"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_game"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	PrefixWhitelistedContract    = []byte{0x01}
	PrefixAccountDeposit         = []byte{0x02}
	PrefixUnbondingKey           = []byte{0x03}
	PrefixInGameUnbondingUserKey = []byte{0x04}
	PrefixInGameUnbondingTimeKey = []byte{0x05}
	KeyLastUnbondingIndex        = []byte{0x06}
	KeyLiquidity                 = []byte{0x07}
)

func AccountDepositKey(addr sdk.AccAddress) []byte {
	return append(PrefixAccountDeposit, addr...)
}

func UnbondingKey(id uint64) []byte {
	return append(PrefixUnbondingKey, sdk.Uint64ToBigEndian(id)...)
}

func InGameUnbondingUserKey(addr sdk.AccAddress, id uint64) []byte {
	return append(InGameUnbondingUserPrefixKey(addr), sdk.Uint64ToBigEndian(id)...)
}

func InGameUnbondingUserPrefixKey(addr sdk.AccAddress) []byte {
	return append(PrefixInGameUnbondingUserKey, address.MustLengthPrefix(addr)...)
}

func InGameUnbondingTimeKey(time time.Time, id uint64) []byte {
	return append(InGameUnbondingTimePrefixKey(time), sdk.Uint64ToBigEndian(id)...)
}

func InGameUnbondingTimePrefixKey(time time.Time) []byte {
	return append(PrefixInGameUnbondingTimeKey, sdk.FormatTimeBytes(time)...)
}
