package burner

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/vidulum/vidulum/x/burner/keeper"
	"github.com/vidulum/vidulum/x/burner/types"
)

func NewBurnerProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.BurnCoinsProposal:
			return k.HandleBurnCoinsProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized burner proposal content type: %T", c)
		}
	}
}
