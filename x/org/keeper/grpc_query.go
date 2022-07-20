package keeper

import (
	"org/x/org/types"
)

var _ types.QueryServer = Keeper{}
