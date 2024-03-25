package keeper

import (
	"github.com/vidulum/vidulum/x/burner/types"
)

var _ types.QueryServer = Keeper{}
