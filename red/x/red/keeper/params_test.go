package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "red/testutil/keeper"
	"red/x/red/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RedKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
