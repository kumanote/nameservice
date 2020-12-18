package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kumanote/nameservice/x/nameservice/types"
	"github.com/kumanote/nameservice/x/nameservice/keeper"
)

func handleMsgCreateWhois(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateWhois) (*sdk.Result, error) {
	k.CreateWhois(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
