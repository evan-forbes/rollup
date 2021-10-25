package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/evan-forbes/rollup/testutil/keeper"
	"github.com/evan-forbes/rollup/x/rollup/keeper"
	"github.com/evan-forbes/rollup/x/rollup/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RollupKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
