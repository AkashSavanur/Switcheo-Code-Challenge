package types

const (
	// ModuleName defines the module name
	ModuleName = "todolist"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_todolist"

	TaskKey = "Task/value/"

	TaskCountKey = "Task/count/"
)

var (
	ParamsKey = []byte("p_todolist")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
