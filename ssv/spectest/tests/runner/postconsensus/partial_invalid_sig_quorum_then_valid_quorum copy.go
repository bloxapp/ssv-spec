package postconsensus

import (
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// PartialInvalidSigQuorumThenValidQuorum tests a runner receiving a partially invalid message (due to wrong sigs) forming an invalid quorum, then receiving a valid message forming a valid quorum, terminating successfully
func PartialInvalidSigQuorumThenValidQuorum() tests.SpecTest {

	ks := testingutils.Testing4SharesSet()
	expectedError := "got post-consensus quorum but it has invalid signatures: could not reconstruct beacon sig: failed to verify reconstruct signature: could not reconstruct a valid signature"

	numValidators := 30
	validatorsIndexList := testingutils.ValidatorIndexList(numValidators)
	ksMap := testingutils.KeySetMapForValidators(numValidators)
	shareMap := testingutils.ShareMapFromKeySetMap(ksMap)

	multiSpecTest := &tests.MultiMsgProcessingSpecTest{
		Name: "post consensus partial invalid sig quorum then valid quorum",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name: "attester",
				Runner: decideCommitteeRunner(
					testingutils.CommitteeRunnerWithShareMap(shareMap),
					testingutils.TestingCommitteeAttesterDuty(testingutils.TestingDutySlot, validatorsIndexList),
					&testingutils.TestBeaconVote,
				),
				Duty: testingutils.TestingCommitteeAttesterDuty(testingutils.TestingDutySlot, validatorsIndexList),
				Messages: []*types.SignedSSVMessage{
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusPartiallyWrongBeaconSigAttestationMsgForKeySet(ksMap, 1, testingutils.TestingDutySlot))),

					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationMsgForKeySet(ksMap, 2, testingutils.TestingDutySlot))),
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationMsgForKeySet(ksMap, 3, testingutils.TestingDutySlot))),
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationMsgForKeySet(ksMap, 4, testingutils.TestingDutySlot))),
				},
				OutputMessages:         []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots: testingutils.TestingSignedAttestationSSZRootForKeyMap(ksMap),
				DontStartDuty:          true,
				ExpectedError:          expectedError,
			},
			{
				Name: "sync committee",
				Runner: decideCommitteeRunner(
					testingutils.CommitteeRunnerWithShareMap(shareMap),
					testingutils.TestingCommitteeSyncCommitteeDuty(testingutils.TestingDutySlot, validatorsIndexList),
					&testingutils.TestBeaconVote,
				),
				Duty: testingutils.TestingCommitteeSyncCommitteeDuty(testingutils.TestingDutySlot, validatorsIndexList),
				Messages: []*types.SignedSSVMessage{
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusPartiallyWrongBeaconSigSyncCommitteeMsgForKeySet(ksMap, 1))),

					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusSyncCommitteeMsgForKeySet(ksMap, 2))),
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusSyncCommitteeMsgForKeySet(ksMap, 3))),
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusSyncCommitteeMsgForKeySet(ksMap, 4))),
				},
				OutputMessages:         []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots: testingutils.TestingSignedSyncCommitteeBlockRootSSZRootForKeyMap(ksMap),
				DontStartDuty:          true,
				ExpectedError:          expectedError,
			},
			{
				Name: "attester and sync committee",
				Runner: decideCommitteeRunner(
					testingutils.CommitteeRunnerWithShareMap(shareMap),
					testingutils.TestingCommitteeDuty(testingutils.TestingDutySlot, validatorsIndexList, validatorsIndexList),
					&testingutils.TestBeaconVote,
				),
				Duty: testingutils.TestingCommitteeDuty(testingutils.TestingDutySlot, validatorsIndexList, validatorsIndexList),
				Messages: []*types.SignedSSVMessage{
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusPartiallyWrongBeaconSigAttestationAndSyncCommitteeMsgForKeySet(ksMap, 1, testingutils.TestingDutySlot))),

					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationAndSyncCommitteeMsgForKeySet(ksMap, 2, testingutils.TestingDutySlot))),
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationAndSyncCommitteeMsgForKeySet(ksMap, 3, testingutils.TestingDutySlot))),
					testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgCommittee(ks, nil, testingutils.PostConsensusAttestationAndSyncCommitteeMsgForKeySet(ksMap, 4, testingutils.TestingDutySlot))),
				},
				OutputMessages: []*types.PartialSignatureMessages{},
				BeaconBroadcastedRoots: append(
					testingutils.TestingSignedAttestationSSZRootForKeyMap(ksMap),
					testingutils.TestingSignedSyncCommitteeBlockRootSSZRootForKeyMap(ksMap)...),
				DontStartDuty: true,
				ExpectedError: expectedError,
			},
		},
	}

	return multiSpecTest
}