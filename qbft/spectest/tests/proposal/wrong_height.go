package proposal

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// WrongHeight tests a proposal msg received with the wrong height
func WrongHeight() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()
	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessageWithHeight(ks.Shares[1], types.OperatorID(1), 2),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "wrong proposal height",
		Pre:           pre,
		PostRoot:      "16577c5ca1fcfd7e43549b13fc730edfd05ac8b0110e451e44d512f034ec6eb3",
		InputMessages: msgs,
		ExpectedError: "invalid signed message: wrong msg height",
	}
}
