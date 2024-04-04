package commit

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/herumi/bls-eth-go-binary/bls"
)

// MultiSignerWithOverlap tests a multi signer commit msg which does overlap previous valid commit signers
func MultiSignerWithOverlap() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()

	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], 1),

		testingutils.TestingPrepareMessage(ks.Shares[1], 1),
		testingutils.TestingPrepareMessage(ks.Shares[2], 2),
		testingutils.TestingPrepareMessage(ks.Shares[3], 3),

		testingutils.TestingCommitMessage(ks.Shares[1], 1),
		testingutils.TestingCommitMultiSignerMessage([]*bls.SecretKey{ks.Shares[2], ks.Shares[3]}, []types.OperatorID{2, 3}),
		testingutils.TestingCommitMessage(ks.Shares[3], 3),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "multi signer, with overlap",
		Pre:           pre,
		PostRoot:      "c47275774ac12fab96625660db46e936d2834eaa60cf83a23d07c94556604358",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks.Shares[1], 1),
			testingutils.TestingCommitMessage(ks.Shares[1], 1),
		},
		ExpectedError: "invalid signed message: msg allows 1 signer",
	}
}
