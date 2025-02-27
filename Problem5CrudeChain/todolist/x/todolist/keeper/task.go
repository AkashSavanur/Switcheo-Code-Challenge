package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"todolist/x/task/types"
)

// Append a new task to the blockchain
func (k Keeper) AppendTask(ctx context.Context, task types.Task) uint64 {
	count := k.GetTaskCount(ctx)
	task.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKey))
	appendedValue := k.cdc.MustMarshal(&task)
	store.Set(GetTaskIDBytes(task.Id), appendedValue)

	k.SetTaskCount(ctx, count+1)
	return count
}

// Retrieve total number of tasks
func (k Keeper) GetTaskCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TaskCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// Generate unique task ID
func GetTaskIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// Set task count
func (k Keeper) SetTaskCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TaskCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// Retrieve a task by ID
func (k Keeper) GetTask(ctx context.Context, id uint64) (val types.Task, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKey))
	b := store.Get(GetTaskIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetTask(ctx context.Context, task types.Task) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKey))
	b := k.cdc.MustMarshal(&task)
	store.Set(GetTaskIDBytes(task.Id), b)
}

func (k Keeper) RemoveTask(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKey))
	store.Delete(GetTaskIDBytes(id))
}
