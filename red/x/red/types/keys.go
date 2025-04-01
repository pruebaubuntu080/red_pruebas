package types

const (
	// ModuleName defines the module name
	ModuleName = "red"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_red"
)

var (
	ParamsKey = []byte("p_red")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
