package proposal

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// PostPrepared tests processing proposal msg after instance prepared
func PostPrepared() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()
	ks10 := testingutils.Testing10SharesSet() // TODO: should be 4?
	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], types.OperatorID(1)),

		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.Shares[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.Shares[3], types.OperatorID(3)),

		testingutils.TestingProposalMessage(ks.Shares[1], types.OperatorID(1)),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "proposal post prepare",
		Pre:           pre,
		PostRoot:      "fd088c77032874833a90280172966c4a24cc0e661bb132ab028ff3b2b36a99f2",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks10.Shares[1], types.OperatorID(1)),
			testingutils.TestingCommitMessage(ks10.Shares[1], types.OperatorID(1)),
		},
		ExpectedError: "invalid signed message: proposal is not valid with current state",
	}
}
