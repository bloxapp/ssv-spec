package proposal

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// DuplicateMsg tests a duplicate proposal msg processing
func DuplicateMsg() *tests.MsgProcessingSpecTest {
	pre := testingutils.BaseInstance()
	signMsgEncoded, _ := testingutils.SignQBFTMsg(testingutils.Testing4SharesSet().Shares[1], types.OperatorID(1), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  qbft.FirstRound,
		Input:  []byte{1, 2, 3, 4},
	}).Encode()

	msgs := []*types.Message{
		{
			ID:   types.PopulateMsgType(pre.State.ID, types.ConsensusProposeMsgType),
			Data: signMsgEncoded,
		},
		{
			ID:   types.PopulateMsgType(pre.State.ID, types.ConsensusProposeMsgType),
			Data: signMsgEncoded,
		},
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "proposal duplicate message",
		Pre:           pre,
		PostRoot:      "93f809fcad476369d03647b6a6661ba3313af3e7b3446d9f639513c2c25f53a9",
		InputMessages: msgs,
		OutputMessages: []*types.Message{
			{
				ID:   types.PopulateMsgType(pre.State.ID, types.ConsensusPrepareMsgType),
				Data: signMsgEncoded,
			},
		},
		ExpectedError: "proposal invalid: proposal is not valid with current state",
	}
}
