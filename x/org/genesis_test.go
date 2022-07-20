package org_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "org/testutil/keeper"
	"org/testutil/nullify"
	"org/x/org"
	"org/x/org/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OrgKeeper(t)
	org.InitGenesis(ctx, *k, genesisState)
	got := org.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
