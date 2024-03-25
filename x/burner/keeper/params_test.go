package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/vidulum/vidulum/testutil/keeper"
	"github.com/vidulum/vidulum/x/burner/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BurnerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
