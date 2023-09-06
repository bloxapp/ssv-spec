package commit

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// PostCutoff tests processing a commit msg when round >= cutoff
func PostCutoff() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	pre := testingutils.BaseInstance()
	pre.State.Round = 15

	msgs := []*qbft.SignedMessage{
		testingutils.TestingCommitMessageWithRound(ks.Shares[1], types.OperatorID(1), 15),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "round cutoff commit message",
		Pre:           pre,
		PostRoot:      "ed24ab0c1b9e4b8bf8ccfa74ae19525fed375a7aee2ee9dfb159d52358c09405",
		InputMessages: msgs,
		ExpectedError: "instance stopped processing messages",
	}
}
