package proposal

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// WrongHeight tests a proposal msg received with the wrong height
func WrongHeight() *tests.MsgProcessingSpecTest {
	pre := testingutils.BaseInstance()
	signMsgEncoded, _ := testingutils.SignQBFTMsg(testingutils.Testing4SharesSet().Shares[1], types.OperatorID(1), &qbft.Message{
		Height: 2,
		Round:  qbft.FirstRound,
		Input:  []byte{1, 2, 3, 4},
	}).Encode()

	msgs := []*types.Message{
		{
			ID:   types.PopulateMsgType(pre.State.ID, types.ConsensusProposeMsgType),
			Data: signMsgEncoded,
		},
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "wrong proposal height",
		Pre:           pre,
		PostRoot:      "3e721f04a2a64737ec96192d59e90dfdc93f166ec9a21b88cc33ee0c43f2b26a",
		InputMessages: msgs,
		ExpectedError: "proposal invalid: proposal Height is wrong",
	}
}
