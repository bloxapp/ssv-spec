package proposer

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/herumi/bls-eth-go-binary/bls"

	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// ProposeBlindedBlockDecidedRegular tests proposing a blinded block but the decided block is a regular block. Full flow
func ProposeBlindedBlockDecidedRegular() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	return &tests.MsgProcessingSpecTest{
		Name:   "propose blinded decide regular",
		Runner: testingutils.ProposerBlindedBlockRunner(ks),
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
					testingutils.TestProposerConsensusDataBytsV(spec.DataVersionBellatrix),
				), nil),

			testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix)),
			testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[2], 2, spec.DataVersionBellatrix)),
			testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[3], 3, spec.DataVersionBellatrix)),
		},
		PostDutyRunnerStateRoot: "2d31e8b952f15adf6abf75e633ff191c032faad87df4ea97d638514ba054a2ae",
		OutputMessages: []*types.SignedPartialSignatureMessage{
			testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix),
			testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix),
		},
		BeaconBroadcastedRoots: []string{
			testingutils.GetSSZRootNoError(testingutils.TestingSignedBeaconBlockV(ks, spec.DataVersionBellatrix)),
		},
	}
}
