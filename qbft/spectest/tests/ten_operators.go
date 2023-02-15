package tests

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// TenOperators tests a simple full happy flow until decided
func TenOperators() *MsgProcessingSpecTest {
	pre := testingutils.TenOperatorsInstance()
	ks := testingutils.Testing10SharesSet()

	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], types.OperatorID(1)),

		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.Shares[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.Shares[3], types.OperatorID(3)),
		testingutils.TestingPrepareMessage(ks.Shares[4], types.OperatorID(4)),
		testingutils.TestingPrepareMessage(ks.Shares[5], types.OperatorID(5)),
		testingutils.TestingPrepareMessage(ks.Shares[6], types.OperatorID(6)),
		testingutils.TestingPrepareMessage(ks.Shares[7], types.OperatorID(7)),

		testingutils.TestingCommitMessage(ks.Shares[1], types.OperatorID(1)),
		testingutils.TestingCommitMessage(ks.Shares[2], types.OperatorID(2)),
		testingutils.TestingCommitMessage(ks.Shares[3], types.OperatorID(3)),
		testingutils.TestingCommitMessage(ks.Shares[4], types.OperatorID(4)),
		testingutils.TestingCommitMessage(ks.Shares[5], types.OperatorID(5)),
		testingutils.TestingCommitMessage(ks.Shares[6], types.OperatorID(6)),
		testingutils.TestingCommitMessage(ks.Shares[7], types.OperatorID(7)),
	}
	return &MsgProcessingSpecTest{
		Name:          "happy flow ten operators",
		Pre:           pre,
		PostRoot:      "9eb7b796200c6aad6b5014a241f08802f35a37c949d9d3a8489a2b61e91fd24b",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(1)),
			testingutils.TestingCommitMessage(ks.Shares[1], types.OperatorID(1)),
		},
	}
}
