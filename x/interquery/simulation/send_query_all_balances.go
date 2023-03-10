package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/junkai121/interchain-swap/x/interquery/keeper"
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

func SimulateMsgSendQueryOsmosisPrice(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSendQueryOsmosisPrice{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SendQueryOsmosisPrice simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SendQueryOsmosisPrice simulation not implemented"), nil, nil
	}
}
