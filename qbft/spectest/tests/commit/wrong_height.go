package commit

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// WrongHeight tests a commit msg received with the wrong height
func WrongHeight() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()

	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], 1),

		testingutils.TestingPrepareMessage(ks.Shares[1], 1),
		testingutils.TestingPrepareMessage(ks.Shares[2], 2),
		testingutils.TestingPrepareMessage(ks.Shares[3], 3),

		testingutils.TestingCommitMessageWrongHeight(ks.Shares[1], 1),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "wrong commit height",
		Pre:           pre,
		PostRoot:      "f92657c2f9b7225573bcd5b03a0f59dfda3f04c361e022e37357bf22c7a10ab9",
		InputMessages: msgs,
		ExpectedError: "invalid signed message: wrong msg height",
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks.Shares[1], 1),
			testingutils.TestingCommitMessage(ks.Shares[1], 1),
		},
	}
}
