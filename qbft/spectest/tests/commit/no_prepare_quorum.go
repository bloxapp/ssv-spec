package commit

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// NoPrepareQuorum tests a commit msg received without a previous prepare quorum
func NoPrepareQuorum() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()

	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], 1),

		testingutils.TestingPrepareMessage(ks.Shares[1], 1),
		testingutils.TestingPrepareMessage(ks.Shares[2], 2),
		// only 2 out of 4 prepare msgs

		testingutils.TestingCommitMessage(ks.Shares[1], 1),
		testingutils.TestingCommitMessage(ks.Shares[2], 2),
		testingutils.TestingCommitMessage(ks.Shares[3], 3),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "commit no prepare quorum",
		Pre:           pre,
		PostRoot:      "ec98a529e9a8736c72cf09c06ed3511728bdb13e27157b13ebd9c9cf279e0c15",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks.Shares[1], 1),
		},
	}
}
