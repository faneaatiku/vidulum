package burner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vidulum/vidulum/testutil/keeper"
	"github.com/vidulum/vidulum/testutil/nullify"
	"github.com/vidulum/vidulum/x/burner"
	"github.com/vidulum/vidulum/x/burner/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BurnerKeeper(t)
	burner.InitGenesis(ctx, *k, genesisState)
	got := burner.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
