package decided

import (
	"crypto/rsa"

	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// WrongMsgType tests a non commit msg with 2f+1 signers
func WrongMsgType() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	return &tests.ControllerSpecTest{
		Name: "decide wrong msg type",
		RunInstanceData: []*tests.RunInstanceData{
			{
				InputValue: []byte{1, 2, 3, 4},
				InputMessages: []*types.SignedSSVMessage{
					testingutils.TestingMultiSignerProposalMessageWithHeight(
						[]*rsa.PrivateKey{ks.NetworkKeys[1], ks.NetworkKeys[2], ks.NetworkKeys[3]},
						[]types.OperatorID{1, 2, 3},
						qbft.FirstHeight,
					),
				},
				ControllerPostRoot: "47713c38fe74ce55959980781287886c603c2117a14dc8abce24dcb9be0093af",
			},
		},
		ExpectedError: "could not process msg: invalid signed message: msg allows 1 signer",
	}
}
