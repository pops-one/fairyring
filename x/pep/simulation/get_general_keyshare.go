package simulation

import (
	"math/rand"

	"github.com/Fairblock/fairyring/x/pep/keeper"
	"github.com/Fairblock/fairyring/x/pep/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgGetGeneralKeyshare(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgGetGeneralKeyshare{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the GetGeneralKeyshare simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "GetGeneralKeyshare simulation not implemented"), nil, nil
	}
}
