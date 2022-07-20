package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "org/testutil/keeper"
	"org/x/org/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OrgKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
