package client

import (
	"github.com/vidulum/vidulum/x/burner/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var BurnCoinsProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitBurnCoinsProposal)
