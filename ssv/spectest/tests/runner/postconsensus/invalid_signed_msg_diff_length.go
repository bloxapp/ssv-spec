package postconsensus

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// InvalidSignedMessageDifferentLength tests an invalid SignedSSVMessage with a different number of signers and signatures
func InvalidSignedMessageDifferentLength() tests.SpecTest {

	ks := testingutils.Testing4SharesSet()

	differentLength := func(msg *types.SignedSSVMessage) *types.SignedSSVMessage {
		msg.Signatures = [][]byte{{1, 2, 3, 4}, {2, 3, 4, 5}}
		return msg
	}

	expectedError := "invalid SignedSSVMessage: number of signatures is different than number of signers"

	return &tests.MultiMsgProcessingSpecTest{
		Name: "post consensus invalid signed message different length",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name: "attester",
				Runner: decideCommitteeRunner(
					testingutils.CommitteeRunner(ks),
					testingutils.TestingAttesterDuty,
					&testingutils.TestBeaconVote,
				),
				Duty: testingutils.TestingAttesterDuty,
				Messages: []*types.SignedSSVMessage{
					differentLength(testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationMsg(ks.Shares[1], 0, testingutils.TestingDutySlot)))),
				},
				OutputMessages:         []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots: []string{},
				DontStartDuty:          true,
				ExpectedError:          expectedError,
			},
			{
				Name: "sync committee",
				Runner: decideCommitteeRunner(
					testingutils.CommitteeRunner(ks),
					testingutils.TestingSyncCommitteeDuty,
					&testingutils.TestBeaconVote,
				),
				Duty: testingutils.TestingSyncCommitteeDuty,
				Messages: []*types.SignedSSVMessage{
					differentLength(testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[1], 0)))),
				},
				OutputMessages:         []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots: []string{},
				DontStartDuty:          true,
				ExpectedError:          expectedError,
			},
			{
				Name: "attester and sync committee",
				Runner: decideCommitteeRunner(
					testingutils.CommitteeRunner(ks),
					testingutils.TestingAttesterAndSyncCommitteeDuties,
					&testingutils.TestBeaconVote,
				),
				Duty: testingutils.TestingAttesterAndSyncCommitteeDuties,
				Messages: []*types.SignedSSVMessage{
					differentLength(testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationAndSyncCommitteeMsg(ks.Shares[1], 0, testingutils.TestingDutySlot)))),
				},
				OutputMessages:         []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots: []string{},
				DontStartDuty:          true,
				ExpectedError:          expectedError,
			},
			{
				Name: "sync committee contribution",
				Runner: decideRunner(
					testingutils.SyncCommitteeContributionRunner(ks),
					&testingutils.TestingSyncCommitteeContributionDuty,
					testingutils.TestSyncCommitteeContributionConsensusData,
				),
				Duty: &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SignedSSVMessage{
					differentLength(testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], testingutils.SSVMsgSyncCommitteeContribution(nil, testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[1], 0, ks)))),
				},
				PostDutyRunnerStateRoot: "f58387d4d4051a2de786e4cbf9dc370a8b19a544f52af04f71195feb3863fc5c",
				OutputMessages:          []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots:  []string{},
				DontStartDuty:           true,
				ExpectedError:           expectedError,
			},
			{
				Name: "proposer",
				Runner: decideRunner(
					testingutils.ProposerRunner(ks),
					testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
					testingutils.TestProposerConsensusDataV(spec.DataVersionDeneb),
				),
				Duty: testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					differentLength(testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[1], 0, spec.DataVersionDeneb)))),
				},
				PostDutyRunnerStateRoot: "ff213af6f0bf2350bb37f48021c137dd5552b1c25cb5c6ebd0c1d27debf6080e",
				OutputMessages:          []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots:  []string{},
				DontStartDuty:           true,
				ExpectedError:           expectedError,
			},
			{
				Name: "proposer (blinded block)",
				Runner: decideRunner(
					testingutils.ProposerBlindedBlockRunner(ks),
					testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
					testingutils.TestProposerBlindedBlockConsensusDataV(spec.DataVersionDeneb),
				),
				Duty: testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					differentLength(testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[1], 0, spec.DataVersionDeneb)))),
				},
				PostDutyRunnerStateRoot: "9b4524d5100835df4d71d0a1e559acdc33d541c44a746ebda115c5e7f3eaa85a",
				OutputMessages:          []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots:  []string{},
				DontStartDuty:           true,
				ExpectedError:           expectedError,
			},
			{
				Name: "aggregator",
				Runner: decideRunner(
					testingutils.AggregatorRunner(ks),
					&testingutils.TestingAggregatorDuty,
					testingutils.TestAggregatorConsensusData,
				),
				Duty: &testingutils.TestingAggregatorDuty,
				Messages: []*types.SignedSSVMessage{
					differentLength(testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], testingutils.SSVMsgAggregator(nil, testingutils.PostConsensusAggregatorMsg(ks.Shares[1], 0)))),
				},
				PostDutyRunnerStateRoot: "1fb182fb19e446d61873abebc0ac85a3a9637b51d139cdbd7d8cb70cf7ffec82",
				OutputMessages:          []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots:  []string{},
				DontStartDuty:           true,
				ExpectedError:           expectedError,
			},
		},
	}
}
