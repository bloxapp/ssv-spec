package roundchange

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// DuplicateMsg tests a duplicate rc msg (first one inserted, second not)
func DuplicateMsg() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 2
	ks := testingutils.Testing4SharesSet()

	prepareMsgs := []*qbft.SignedMessage{
		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.Shares[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.Shares[3], types.OperatorID(3)),
	}
	msgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[1], types.OperatorID(1), 5),
		testingutils.TestingRoundChangeMessageWithParams(ks.Shares[1], types.OperatorID(1), 5, qbft.FirstHeight,
			testingutils.TestingQBFTRootData, qbft.FirstRound, testingutils.MarshalJustifications(prepareMsgs)),
	}

	return &tests.MsgProcessingSpecTest{
		Name:           "round change duplicate msg",
		Pre:            pre,
		PostRoot:       "49cbc65f38d8432a0621d6be406118248c342b41c2544091a6dd56de494861ed",
		InputMessages:  msgs,
		OutputMessages: []*qbft.SignedMessage{},
	}
}
