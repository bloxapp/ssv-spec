package commit

import (
	"crypto/rsa"

	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// DuplicateSigners tests a multi signer commit msg with duplicate signers
func DuplicateSigners() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()

	pre.State.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessage(ks.OperatorKeys[1], 1)
	commit := testingutils.TestingCommitMultiSignerMessage([]*rsa.PrivateKey{ks.OperatorKeys[1], ks.OperatorKeys[2]}, []types.OperatorID{1, 2})
	commit.OperatorIDs = []types.OperatorID{1, 1}

	return &tests.MsgProcessingSpecTest{
		Name:     "duplicate signers",
		Pre:      pre,
		PostRoot: "a711e11c0e08204861008b7f3b8affa498bfc777199e0692bd661e40b09d55d6",
		InputMessages: []*types.SignedSSVMessage{
			commit,
		},
		OutputMessages: []*types.SignedSSVMessage{},
		ExpectedError:  "invalid signed message: invalid SignedSSVMessage: non unique signer",
	}
}
