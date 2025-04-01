package red_test

import (
	"testing"

	keepertest "red/testutil/keeper"
	"red/testutil/nullify"
	red "red/x/red/module"
	"red/x/red/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RedKeeper(t)
	red.InitGenesis(ctx, k, genesisState)
	got := red.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
