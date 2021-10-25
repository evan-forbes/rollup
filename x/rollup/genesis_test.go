package rollup_test

import (
	"testing"

	keepertest "github.com/evan-forbes/rollup/testutil/keeper"
	"github.com/evan-forbes/rollup/x/rollup"
	"github.com/evan-forbes/rollup/x/rollup/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RollupKeeper(t)
	rollup.InitGenesis(ctx, *k, genesisState)
	got := rollup.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
