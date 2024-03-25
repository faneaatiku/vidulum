package burner

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vidulum/vidulum/x/burner/keeper"
	"github.com/vidulum/vidulum/x/burner/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, burnedCoins := range genState.BurnedCoinsList {
		k.SetBurnedCoins(ctx, burnedCoins)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.BurnedCoinsList = k.GetAllBurnedCoins(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
