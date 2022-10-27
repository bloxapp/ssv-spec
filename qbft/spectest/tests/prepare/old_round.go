package prepare

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// OldRound tests prepare for signedProposal.Message.Round < state.Round
func OldRound() *tests.MsgProcessingSpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 10

	rcMsg := testingutils.SignQBFTMsg(testingutils.Testing4SharesSet().Shares[1], types.OperatorID(1), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  10,
	})
	rcMsg2 := testingutils.SignQBFTMsg(testingutils.Testing4SharesSet().Shares[2], types.OperatorID(2), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  10,
	})
	rcMsg3 := testingutils.SignQBFTMsg(testingutils.Testing4SharesSet().Shares[3], types.OperatorID(3), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  10,
	})
	prepareMsg := testingutils.SignQBFTMsg(testingutils.Testing4SharesSet().Shares[1], types.OperatorID(1), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  9,
		Input:  &qbft.Data{Root: [32]byte{}, Source: []byte{1, 2, 3, 4}},
	})
	proposalMsg := testingutils.SignQBFTMsg(testingutils.Testing4SharesSet().Shares[1], types.OperatorID(1), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  10,
		Input:  &qbft.Data{Root: [32]byte{}, Source: []byte{1, 2, 3, 4}},
	})

	rcMsgHeader, _ := rcMsg.ToSignedMessage()
	rcMsgHeader2, _ := rcMsg2.ToSignedMessage()
	rcMsgHeader3, _ := rcMsg3.ToSignedMessage()

	rcJustifications := []*qbft.SignedMessage{
		rcMsgHeader,
		rcMsgHeader2,
		rcMsgHeader3,
	}
	proposalMsg.RoundChangeJustifications = rcJustifications
	pre.State.ProposalAcceptedForCurrentRound = proposalMsg

	prepareMsgEncoded, _ := prepareMsg.Encode()

	msgs := []*types.Message{
		{
			ID:   types.PopulateMsgType(pre.State.ID, types.ConsensusPrepareMsgType),
			Data: prepareMsgEncoded,
		},
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "prepare prev round",
		Pre:           pre,
		PostRoot:      "fe4b14e78ddeda2c5440fb74a6c2d2bf8fdb6b81429c1b5b442018472756ed05",
		InputMessages: msgs,
		ExpectedError: "invalid prepare msg: msg round wrong",
	}
}
