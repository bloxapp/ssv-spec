package prepare

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// WrongSignature tests a single prepare received with a wrong signature
func WrongSignature() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	pre := testingutils.BaseInstance()
	pre.State.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessage(ks.Shares[1], types.OperatorID(1))

	msgs := []*qbft.SignedMessage{
		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(2)),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "prepare wrong sig",
		Pre:           pre,
		InputMessages: msgs,
		ExpectedError: "invalid signed message: msg signature invalid: failed to verify signature",
	}
}
