package roundchange

import (
	"github.com/bloxapp/ssv-spec/qbft"
	qbftcomparable "github.com/bloxapp/ssv-spec/qbft/spectest/comparable"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// DuplicateMsgQuorumPreparedRCFirst tests a duplicate rc msg (the prev prepared one first)
func DuplicateMsgQuorumPreparedRCFirst() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 2
	ks := testingutils.Testing4SharesSet()
	sc := duplicateMsgQuorumPreparedRCFirstStateComparison()

	prepareMsgs := []*qbft.SignedMessage{
		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.Shares[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.Shares[3], types.OperatorID(3)),
	}
	msgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRoundAndRC(ks.Shares[1], types.OperatorID(1), 2,
			testingutils.MarshalJustifications(prepareMsgs)),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[1], types.OperatorID(1), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[2], types.OperatorID(2), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[3], types.OperatorID(3), 2),
	}

	rcMsgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRoundAndRC(ks.Shares[1], types.OperatorID(1), 2,
			testingutils.MarshalJustifications(prepareMsgs)),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[2], types.OperatorID(2), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[3], types.OperatorID(3), 2),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "round change duplicate msg quorum (prev prepared rc first)",
		Pre:           pre,
		PostRoot:      sc.Root(),
		PostState:     sc.ExpectedState,
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingProposalMessageWithParams(ks.Shares[1], types.OperatorID(1), 2, qbft.FirstHeight,
				testingutils.TestingQBFTRootData,
				testingutils.MarshalJustifications(rcMsgs), testingutils.MarshalJustifications(prepareMsgs)),
		},
	}
}

func duplicateMsgQuorumPreparedRCFirstStateComparison() *qbftcomparable.StateComparison {
	ks := testingutils.Testing4SharesSet()

	prepareMsgs := []*qbft.SignedMessage{
		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.Shares[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.Shares[3], types.OperatorID(3)),
	}

	msgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRoundAndRC(ks.Shares[1], types.OperatorID(1), 2,
			testingutils.MarshalJustifications(prepareMsgs)),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[2], types.OperatorID(2), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[3], types.OperatorID(3), 2),
	}

	instance := &qbft.Instance{
		State: &qbft.State{
			Share: testingutils.TestingShare(testingutils.Testing4SharesSet()),
			ID:    testingutils.TestingIdentifier,
			Round: 2,
		},
	}
	qbftcomparable.SetSignedMessages(instance, msgs)
	return &qbftcomparable.StateComparison{ExpectedState: instance.State}
}
