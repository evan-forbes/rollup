package keeper

import (
	"github.com/evan-forbes/rollup/x/rollup/types"
)

var _ types.QueryServer = Keeper{}
