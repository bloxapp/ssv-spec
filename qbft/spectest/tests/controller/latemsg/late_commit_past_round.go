package latemsg

import (
	"crypto/rsa"

	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// LateCommitPastRound tests process late commit msg for an instance which just decided for a round < decided round
func LateCommitPastRound() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	rcMsgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.NetworkKeys[1], types.OperatorID(1), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.NetworkKeys[2], types.OperatorID(2), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.NetworkKeys[3], types.OperatorID(3), 2),
	}

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingProposalMessage(ks.NetworkKeys[1], types.OperatorID(1)),

		testingutils.TestingPrepareMessage(ks.NetworkKeys[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.NetworkKeys[2], types.OperatorID(2)),
	}
	msgs = append(msgs, rcMsgs...)
	msgs = append(msgs, []*types.SignedSSVMessage{
		testingutils.TestingProposalMessageWithRoundAndRC(ks.NetworkKeys[1], types.OperatorID(1), 2,
			testingutils.MarshalJustifications(rcMsgs)),

		testingutils.TestingPrepareMessageWithRound(ks.NetworkKeys[1], types.OperatorID(1), 2),
		testingutils.TestingPrepareMessageWithRound(ks.NetworkKeys[2], types.OperatorID(2), 2),
		testingutils.TestingPrepareMessageWithRound(ks.NetworkKeys[3], types.OperatorID(3), 2),

		testingutils.TestingCommitMessageWithRound(ks.NetworkKeys[1], types.OperatorID(1), 2),
		testingutils.TestingCommitMessageWithRound(ks.NetworkKeys[2], types.OperatorID(2), 2),
		testingutils.TestingCommitMessageWithRound(ks.NetworkKeys[3], types.OperatorID(3), 2),

		testingutils.TestingCommitMessage(ks.NetworkKeys[4], 4),
	}...)

	return &tests.ControllerSpecTest{
		Name: "late commit past round",
		RunInstanceData: []*tests.RunInstanceData{
			{
				InputValue:    []byte{1, 2, 3, 4},
				InputMessages: msgs,
				ExpectedDecidedState: tests.DecidedState{
					DecidedVal: testingutils.TestingQBFTFullData,
					DecidedCnt: 1,
					BroadcastedDecided: testingutils.TestingCommitMultiSignerMessageWithRound(
						[]*rsa.PrivateKey{ks.NetworkKeys[1], ks.NetworkKeys[2], ks.NetworkKeys[3]},
						[]types.OperatorID{1, 2, 3},
						2,
					),
				},

				ControllerPostRoot: "04eba8412e564e9b000f61035cfbc663216671d98d39d74119612614916542ad",
			},
		},
		ExpectedError: "could not process msg: invalid signed message: past round",
	}
}
