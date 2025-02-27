package keeper

import (
	"todolist/x/todolist/types"
)

var _ types.QueryServer = Keeper{}
