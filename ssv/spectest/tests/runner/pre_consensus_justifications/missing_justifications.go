package pre_consensus_justifications

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// MissingQuorum tests missing quorum pre-consensus justification for relevant duties
func MissingQuorum() *tests.MultiMsgProcessingSpecTest {
	ks := testingutils.Testing4SharesSet()

	msgF := func(obj *types.ConsensusData, id []byte) *qbft.SignedMessage {
		fullData, _ := obj.Encode()
		root, _ := qbft.HashDataRoot(fullData)
		msg := &qbft.Message{
			MsgType:    qbft.ProposalMsgType,
			Height:     1,
			Round:      qbft.FirstRound,
			Identifier: id,
			Root:       root,
		}
		signed := testingutils.SignQBFTMsg(ks.Shares[1], 1, msg)
		signed.FullData = fullData

		return signed
	}

	expectedErr := "failed processing consensus message: invalid pre-consensus justification: no quorum"

	return &tests.MultiMsgProcessingSpecTest{
		Name: "pre consensus missing quorum",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:   "sync committee aggregator selection proof",
				Runner: testingutils.SyncCommitteeContributionRunner(ks),
				Duty:   &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgSyncCommitteeContribution(msgF(testingutils.TestSyncCommitteeContributionConsensusData, testingutils.SyncCommitteeContributionMsgID), nil),
				},
				PostDutyRunnerStateRoot: "29862cc6054edc8547efcb5ae753290971d664b9c39768503b4d66e1b52ecb06",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
				ExpectedError: expectedErr,
			},
			{
				Name:   "aggregator selection proof",
				Runner: testingutils.AggregatorRunner(ks),
				Duty:   &testingutils.TestingAggregatorDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgAggregator(msgF(testingutils.TestAggregatorConsensusData, testingutils.AggregatorMsgID), nil),
				},
				PostDutyRunnerStateRoot: "c54e71de23c3957b73abbb0e7b9e195b3f8f6370d62fbec256224faecf177fee",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
				ExpectedError: expectedErr,
			},
			{
				Name:   "randao",
				Runner: testingutils.ProposerRunner(ks),
				Duty:   &testingutils.TestingProposerDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgProposer(msgF(testingutils.TestProposerConsensusData, testingutils.ProposerMsgID), nil),
				},
				PostDutyRunnerStateRoot: "56eafcb33392ded888a0fefe30ba49e52aa00ab36841cb10c9dc1aa2935af347",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsg(ks.Shares[1], 1), // broadcasts when starting a new duty
				},
				ExpectedError: expectedErr,
			},
			{
				Name:   "randao (blinded block)",
				Runner: testingutils.ProposerBlindedBlockRunner(ks),
				Duty:   &testingutils.TestingProposerDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgProposer(msgF(testingutils.TestProposerBlindedBlockConsensusData, testingutils.ProposerMsgID), nil),
				},
				PostDutyRunnerStateRoot: "2ce3241658f324f352c77909f4043934eedf38e939ae638c5ce6acf28e965646",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsg(ks.Shares[1], 1), // broadcasts when starting a new duty
				},
				ExpectedError: expectedErr,
			},
			{

				Name:   "attester",
				Runner: testingutils.AttesterRunner(ks),
				Duty:   &testingutils.TestingAttesterDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgAttester(msgF(testingutils.TestAttesterConsensusData, testingutils.AttesterMsgID), nil),
				},
				PostDutyRunnerStateRoot: "757b432ae782aca13549cdaeb5df8b292691bd8eaab0057ae779c59ed222fd79",
				OutputMessages:          []*types.SignedPartialSignatureMessage{},
			},
			{
				Name:   "sync committee",
				Runner: testingutils.SyncCommitteeRunner(ks),
				Duty:   &testingutils.TestingSyncCommitteeDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgSyncCommittee(msgF(testingutils.TestSyncCommitteeConsensusData, testingutils.SyncCommitteeMsgID), nil),
				},
				PostDutyRunnerStateRoot: "044f5465db19e9e45ec6319dfd870fd8d63352585d6653a719090cba511b13ef",
				OutputMessages:          []*types.SignedPartialSignatureMessage{},
			},
		},
	}
}
