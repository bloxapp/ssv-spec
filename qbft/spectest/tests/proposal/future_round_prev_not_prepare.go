package proposal

import (
	"github.com/ssvlabs/ssv-spec/qbft"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// FutureRoundPrevNotPrepared tests a proposal for future round, currently not prepared
func FutureRoundPrevNotPrepared() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = qbft.FirstRound
	ks := testingutils.Testing4SharesSet()

	rcMsgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 10),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[2], types.OperatorID(2), 10),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[3], types.OperatorID(3), 10),
	}

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingProposalMessageWithParams(ks.OperatorKeys[1], types.OperatorID(1), 10, qbft.FirstHeight,
			testingutils.TestingQBFTRootData,
			testingutils.MarshalJustifications(rcMsgs), nil,
		),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "proposal future round prev not prepared",
		Pre:           pre,
		PostRoot:      "dbf0f9f6c5e3aee3d5907028091520a9ee5bc68a2e987cb0bd3ddd431f6ae7ed",
		InputMessages: msgs,
		OutputMessages: []*types.SignedSSVMessage{
			testingutils.TestingPrepareMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 10),
		},
		ExpectedTimerState: &testingutils.TimerState{
			Timeouts: 1,
			Round:    qbft.Round(10),
		},
	}
}
