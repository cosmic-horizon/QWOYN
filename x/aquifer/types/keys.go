package types

const (
	// ModuleName defines the module name
	ModuleName = "aquifer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_aquifer"

	// OutpostFundingPoolName the root string for the outpost funding pool account address
	OutpostFundingPoolName = "outpost_funding_pool"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
