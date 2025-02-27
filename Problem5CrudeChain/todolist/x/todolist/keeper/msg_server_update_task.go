package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"todolist/x/task/types"
)

func (k msgServer) UpdateTask(goCtx context.Context, msg *types.MsgUpdateTask) (*types.MsgUpdateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("task %d doesn't exist", msg.Id))
	}

	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(errors.ErrUnauthorized, "only the creator can update this task")
	}

	val.Title = msg.Title
	val.Description = msg.Description
	val.Completed = msg.Completed

	k.SetTask(ctx, val)

	return &types.MsgUpdateTaskResponse{}, nil
}
