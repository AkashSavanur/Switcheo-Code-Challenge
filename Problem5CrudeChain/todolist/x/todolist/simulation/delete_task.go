package simulation

import (
	"math/rand"

	"todolist/x/todolist/keeper"
	"todolist/x/todolist/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDeleteTask(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeleteTask{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DeleteTask simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "DeleteTask simulation not implemented"), nil, nil
	}
}
