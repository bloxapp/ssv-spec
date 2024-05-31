package commit

import (
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// NoPrevAcceptedProposal tests a commit msg received without a previous accepted proposal
func NoPrevAcceptedProposal() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()

	pre.State.ProposalAcceptedForCurrentRound = nil
	msgs := []*types.SignedSSVMessage{
		testingutils.TestingCommitMessage(ks.OperatorKeys[1], 1),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "no previous accepted proposal",
		Pre:           pre,
		PostRoot:      "57e323705826bc5d475ead7f48015a785306be56b33b9fed163fdedf03743754",
		InputMessages: msgs,
		ExpectedError: "invalid signed message: did not receive proposal for this round",
	}
}
