package roundchange

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// PeerPreparedDifferentHeights tests a round change quorum where peers prepared on different heights
func PeerPreparedDifferentHeights() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 3

	ks := testingutils.Testing4SharesSet()
	prepareMsgs1 := []*types.SignedSSVMessage{
		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[3], types.OperatorID(3)),
	}
	prepareMsgs2 := []*types.SignedSSVMessage{
		testingutils.TestingPrepareMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 2),
		testingutils.TestingPrepareMessageWithRound(ks.OperatorKeys[2], types.OperatorID(2), 2),
		testingutils.TestingPrepareMessageWithRound(ks.OperatorKeys[3], types.OperatorID(3), 2),
	}
	msgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 3),
		testingutils.TestingRoundChangeMessageWithParamsAndFullData(
			ks.OperatorKeys[2], types.OperatorID(2), 3, qbft.FirstHeight, testingutils.TestingQBFTRootData,
			qbft.FirstRound, testingutils.MarshalJustifications(prepareMsgs1), testingutils.TestingQBFTFullData),
		testingutils.TestingRoundChangeMessageWithParamsAndFullData(ks.OperatorKeys[3], types.OperatorID(3), 3, qbft.FirstHeight,
			testingutils.TestingQBFTRootData, 2, testingutils.MarshalJustifications(prepareMsgs2), testingutils.TestingQBFTFullData),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "round change peer prepared different heights",
		Pre:           pre,
		PostRoot:      "2438df561da0726f0811116d18dafa3c9ef0b7ab2ad0bbd101345cff5dd14b90",
		InputMessages: msgs,
		OutputMessages: []*types.SignedSSVMessage{
			testingutils.TestingProposalMessageWithParams(ks.OperatorKeys[1], types.OperatorID(1), 3, qbft.FirstHeight,
				testingutils.TestingQBFTRootData,
				testingutils.MarshalJustifications(msgs), testingutils.MarshalJustifications(prepareMsgs2)),
		},
	}
}
