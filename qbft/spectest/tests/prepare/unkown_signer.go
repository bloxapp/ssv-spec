package prepare

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// UnknownSigner tests a single prepare received with an unknown signer
func UnknownSigner() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	pre := testingutils.BaseInstance()
	pre.State.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessage(ks.Shares[1], types.OperatorID(1))

	msgs := []*qbft.SignedMessage{
		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(5)),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "prepare unknown signer",
		Pre:           pre,
		PostRoot:      "6f9b4793ad59ce5690725a1c9bd6c443dd7dcaa8e0c30060c4012a4c8a3bb36f",
		InputMessages: msgs,
		ExpectedError: "invalid signed message: msg signature invalid: unknown signer",
	}
}
