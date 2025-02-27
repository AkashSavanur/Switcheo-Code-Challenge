package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "todolist/testutil/keeper"
	"todolist/x/todolist/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TodolistKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
