package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"todolist/x/task/types"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var task = types.Task{
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		Completed:   msg.Completed,
	}

	id := k.AppendTask(ctx, task)

	return &types.MsgCreateTaskResponse{
		Id: id,
	}, nil
}
