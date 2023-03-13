package consensus

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// NoRunningConsensusInstance tests a valid proposal msg before consensus instance starts
func NoRunningConsensusInstance() *tests.MultiMsgProcessingSpecTest {
	ks := testingutils.Testing4SharesSet()
	return &tests.MultiMsgProcessingSpecTest{
		Name: "consensus no running consensus instance",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:   "sync committee contribution",
				Runner: testingutils.SyncCommitteeContributionRunner(ks),
				Duty:   &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgSyncCommitteeContribution(
						testingutils.TestingPrepareMessageWithParams(
							ks.Shares[1],
							1,
							qbft.FirstRound,
							qbft.FirstHeight,
							testingutils.SyncCommitteeContributionMsgID,
							testingutils.TestingQBFTRootData,
						),
						nil),
				},
				PostDutyRunnerStateRoot: "b2e883cc81caaed04f3e40e8561ae55aa1f6abcdb3168e5cc5c834b1d327026e",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
				},
				ExpectedError: "", // with pre-consensus justifications there will always be an instance
			},
			{
				Name:   "sync committee",
				Runner: testingutils.SyncCommitteeRunner(ks),
				Duty:   &testingutils.TestingSyncCommitteeDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgSyncCommittee(
						testingutils.TestingProposalMessageWithIdentifierAndFullData(
							ks.Shares[1], types.OperatorID(1), testingutils.SyncCommitteeMsgID,
							testingutils.TestSyncCommitteeConsensusDataByts,
						),
						nil),
				},
				PostDutyRunnerStateRoot: "941338be935cec87018b8c076d81ebd28a54e46097352fd4585a89a2bf59b3a2",
				OutputMessages:          []*types.SignedPartialSignatureMessage{},
			},
			{
				Name:   "aggregator",
				Runner: testingutils.AggregatorRunner(ks),
				Duty:   &testingutils.TestingAggregatorDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgAggregator(
						testingutils.TestingPrepareMessageWithParams(
							ks.Shares[1],
							1,
							qbft.FirstRound,
							qbft.FirstHeight,
							testingutils.AggregatorMsgID,
							testingutils.TestingQBFTRootData,
						),
						nil),
				},
				PostDutyRunnerStateRoot: "6e1095601c6fbbd6ba5912dfe296b50db2ae67d4115bce7aa2ad0b091c693ea5",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
				},
				ExpectedError: "", // with pre-consensus justifications there will always be an instance
			},
			{
				Name:   "proposer",
				Runner: testingutils.ProposerRunner(ks),
				Duty:   &testingutils.TestingProposerDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgProposer(
						testingutils.TestingPrepareMessageWithParams(
							ks.Shares[1],
							1,
							qbft.FirstRound,
							qbft.FirstHeight,
							testingutils.ProposerMsgID,
							testingutils.TestingQBFTRootData,
						),
						nil),
				},
				PostDutyRunnerStateRoot: "a70427708a0ab6995225538b39e7de5cb622af9651fb02a162c6bfbdf5d0966d",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsg(ks.Shares[1], 1),
				},
				ExpectedError: "", // with pre-consensus justifications there will always be an instance
			},
			{
				Name:   "proposer (blinded block)",
				Runner: testingutils.ProposerBlindedBlockRunner(ks),
				Duty:   &testingutils.TestingProposerDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgProposer(
						testingutils.TestingPrepareMessageWithParams(
							ks.Shares[1],
							1,
							qbft.FirstRound,
							qbft.FirstHeight,
							testingutils.ProposerMsgID,
							testingutils.TestingQBFTRootData,
						),
						nil),
				},
				PostDutyRunnerStateRoot: "dc9ee0b1b1d1562763855898c9962957bc5d4f3090890419c22e0162705e9ca0",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsg(ks.Shares[1], 1),
				},
				ExpectedError: "", // with pre-consensus justifications there will always be an instance
			},
			{
				Name:   "attester",
				Runner: testingutils.AttesterRunner(ks),
				Duty:   &testingutils.TestingAttesterDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgAttester(
						testingutils.TestingProposalMessageWithIdentifierAndFullData(
							ks.Shares[1], types.OperatorID(1), testingutils.AttesterMsgID,
							testingutils.TestAttesterConsensusDataByts,
						),
						nil),
				},
				PostDutyRunnerStateRoot: "0d5b671f94eeddcb00025dd70fa52d259cafaa5f284645db4fd20e943e2e900d",
				OutputMessages:          []*types.SignedPartialSignatureMessage{},
			},
		},
	}
}
