package proposal

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// PostDecided tests processing proposal msg after instance decided
func PostDecided() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks4 := testingutils.Testing4SharesSet()
	ks10 := testingutils.Testing10SharesSet() // TODO should be 4?

	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks4.Shares[1], types.OperatorID(1)),

		testingutils.TestingPrepareMessage(ks4.Shares[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks4.Shares[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks4.Shares[3], types.OperatorID(3)),

		testingutils.TestingCommitMessage(ks4.Shares[1], types.OperatorID(1)),
		testingutils.TestingCommitMessage(ks4.Shares[2], types.OperatorID(2)),
		testingutils.TestingCommitMessage(ks4.Shares[3], types.OperatorID(3)),

		testingutils.TestingProposalMessage(ks4.Shares[1], types.OperatorID(1)),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "proposal post decided",
		Pre:           pre,
		PostRoot:      "f0419eda7b9e11201353fcd800f9f177e020e93e44d9e164b6980f8d6fa545cf",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks10.Shares[1], types.OperatorID(1)),
			testingutils.TestingCommitMessage(ks10.Shares[1], types.OperatorID(1)),
		},
		ExpectedError: "invalid signed message: proposal is not valid with current state",
	}
}
