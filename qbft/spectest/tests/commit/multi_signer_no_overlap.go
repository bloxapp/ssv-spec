package commit

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/herumi/bls-eth-go-binary/bls"
)

// MultiSignerNoOverlap tests a multi signer commit msg which doesn't overlap previous valid commits
func MultiSignerNoOverlap() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()

	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], 1),

		testingutils.TestingPrepareMessage(ks.Shares[1], 1),
		testingutils.TestingPrepareMessage(ks.Shares[2], 2),
		testingutils.TestingPrepareMessage(ks.Shares[3], 3),

		testingutils.TestingCommitMessage(ks.Shares[1], 1),
		testingutils.TestingCommitMultiSignerMessage([]*bls.SecretKey{ks.Shares[2], ks.Shares[3]}, []types.OperatorID{2, 3}),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "multi signer, no overlap",
		Pre:           pre,
		PostRoot:      "b72238ccab63b7aff7c8842f4afe17c622f52da71f101b20a042568edba4edb0",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks.Shares[1], 1),
			testingutils.TestingCommitMessage(ks.Shares[1], 1),
		},
		ExpectedError: "invalid signed message: msg allows 1 signer",
	}
}
