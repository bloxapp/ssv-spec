package proposer

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/herumi/bls-eth-go-binary/bls"

	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// ProposeRegularBlockDecidedBlinded tests proposing a regular block but the decided block is a blinded block. Full flow
func ProposeRegularBlockDecidedBlinded() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	return &tests.MsgProcessingSpecTest{
		Name:   "propose regular decide blinded",
		Runner: testingutils.ProposerRunner(ks),
		Duty:   testingutils.TestingProposerDutyV(spec.DataVersionBellatrix),
		Messages: []*types.SSVMessage{
			testingutils.SSVMsgProposer(nil, testingutils.PreConsensusRandaoDifferentSignerMsgV(ks.Shares[1], ks.Shares[1], 1, 1, spec.DataVersionBellatrix)),
			testingutils.SSVMsgProposer(nil, testingutils.PreConsensusRandaoDifferentSignerMsgV(ks.Shares[2], ks.Shares[2], 2, 2, spec.DataVersionBellatrix)),
			testingutils.SSVMsgProposer(nil, testingutils.PreConsensusRandaoDifferentSignerMsgV(ks.Shares[3], ks.Shares[3], 3, 3, spec.DataVersionBellatrix)),

			testingutils.SSVMsgProposer(
				testingutils.TestingCommitMultiSignerMessageWithIdentifierAndFullData(
					[]*bls.SecretKey{ks.Shares[1], ks.Shares[2], ks.Shares[3]},
					[]types.OperatorID{1, 2, 3},
					testingutils.ProposerMsgID,
					testingutils.TestProposerBlindedBlockConsensusDataBytsV(spec.DataVersionBellatrix),
				), nil),

			testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix)),
			testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[2], 2, spec.DataVersionBellatrix)),
			testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[3], 3, spec.DataVersionBellatrix)),
		},
		PostDutyRunnerStateRoot: "c676d4f845d22b7a4bf7ec323b7f583b1c2d4dabfc4be579b2a5e3b38251e63e",
		OutputMessages: []*types.SignedPartialSignatureMessage{
			testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix),
			testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix),
		},
		BeaconBroadcastedRoots: []string{
			testingutils.GetSSZRootNoError(testingutils.TestingSignedBeaconBlockV(ks, spec.DataVersionBellatrix)),
		},
	}
}
