package consensus

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/ssvlabs/ssv-spec/qbft"
	"github.com/ssvlabs/ssv-spec/ssv"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// PastMessage tests a valid proposal past msg
func PastMessage() tests.SpecTest {

	ks := testingutils.Testing4SharesSet()

	bumpHeight := func(r ssv.Runner) ssv.Runner {
		r.GetBaseRunner().QBFTController.StoredInstances = append(r.GetBaseRunner().QBFTController.StoredInstances, qbft.NewInstance(
			r.GetBaseRunner().QBFTController.GetConfig(),
			r.GetBaseRunner().QBFTController.Share,
			r.GetBaseRunner().QBFTController.Identifier,
			qbft.FirstHeight))

		r.GetBaseRunner().QBFTController.Height = 10
		return r
	}

	pastMsgF := func(cd *types.ConsensusData, beaconVote *types.BeaconVote, id []byte) *types.SignedSSVMessage {
		var fullData []byte
		if cd != nil {
			fullData, _ = cd.Encode()
		} else if beaconVote != nil {
			fullData, _ = beaconVote.Encode()
		} else {
			panic("no consensus data or beacon vote")
		}
		root, _ := qbft.HashDataRoot(fullData)
		msg := &qbft.Message{
			MsgType:    qbft.ProposalMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: id,
			Root:       root,
		}
		signed := testingutils.SignQBFTMsg(ks.OperatorKeys[1], 1, msg)
		signed.FullData = fullData

		return signed
	}

	return &tests.MultiMsgProcessingSpecTest{
		Name: "consensus past message",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:   "attester",
				Runner: bumpHeight(testingutils.CommitteeRunner(ks)),
				Duty:   testingutils.TestingAttesterDuty,
				Messages: []*types.SignedSSVMessage{
					pastMsgF(nil, &testingutils.TestBeaconVote, testingutils.CommitteeMsgID),
				},
				OutputMessages: []*types.PartialSignatureMessages{},
				DontStartDuty:  true,
			},
			{
				Name:   "sync committee",
				Runner: bumpHeight(testingutils.CommitteeRunner(ks)),
				Duty:   testingutils.TestingSyncCommitteeDuty,
				Messages: []*types.SignedSSVMessage{
					pastMsgF(nil, &testingutils.TestBeaconVote, testingutils.CommitteeMsgID),
				},
				OutputMessages: []*types.PartialSignatureMessages{},
				DontStartDuty:  true,
			},
			{
				Name:   "attester and sync committee",
				Runner: bumpHeight(testingutils.CommitteeRunner(ks)),
				Duty:   testingutils.TestingAttesterAndSyncCommitteeDuties,
				Messages: []*types.SignedSSVMessage{
					pastMsgF(nil, &testingutils.TestBeaconVote, testingutils.CommitteeMsgID),
				},
				OutputMessages: []*types.PartialSignatureMessages{},
				DontStartDuty:  true,
			},
			{
				Name:   "sync committee contribution",
				Runner: bumpHeight(testingutils.SyncCommitteeContributionRunner(ks)),
				Duty:   &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SignedSSVMessage{
					pastMsgF(testingutils.TestContributionProofWithJustificationsConsensusData(ks), nil, testingutils.SyncCommitteeContributionMsgID),
				},
				PostDutyRunnerStateRoot: "d1ba71cab348c80ebb7b4533c9c482eaba407f6a73864ee742aab93e73b94dab",
				OutputMessages:          []*types.PartialSignatureMessages{},
				DontStartDuty:           true,
			},
			{
				Name:   "aggregator",
				Runner: bumpHeight(testingutils.AggregatorRunner(ks)),
				Duty:   &testingutils.TestingAggregatorDuty,
				Messages: []*types.SignedSSVMessage{
					pastMsgF(testingutils.TestSelectionProofWithJustificationsConsensusData(ks), nil, testingutils.AggregatorMsgID),
				},
				PostDutyRunnerStateRoot: "5a1a9b9fb21682ea854f919be531a692fe5c3a6c5302214dbf3421faed57cff8",
				OutputMessages:          []*types.PartialSignatureMessages{},
				DontStartDuty:           true,
			},
			{
				Name:   "proposer",
				Runner: bumpHeight(testingutils.ProposerRunner(ks)),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					pastMsgF(testingutils.TestProposerWithJustificationsConsensusDataV(ks, spec.DataVersionDeneb), nil, testingutils.ProposerMsgID),
				},
				PostDutyRunnerStateRoot: "1c939726a237c02013fab61901e819e34ec99e2ef62dadb6c847e5ad118fc4e7",
				OutputMessages:          []*types.PartialSignatureMessages{},
				DontStartDuty:           true,
			},
			{
				Name:   "proposer (blinded block)",
				Runner: bumpHeight(testingutils.ProposerBlindedBlockRunner(ks)),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					pastMsgF(testingutils.TestProposerBlindedWithJustificationsConsensusDataV(ks, spec.DataVersionDeneb), nil, testingutils.ProposerMsgID),
				},
				PostDutyRunnerStateRoot: "49edaab0d759ba8a35a37ab26416ae04962d77ec088b87c4f1e65f781c1ed96f",
				OutputMessages:          []*types.PartialSignatureMessages{},
				DontStartDuty:           true,
			},
			{
				Name:   "validator registration",
				Runner: testingutils.ValidatorRegistrationRunner(ks),
				Duty:   &testingutils.TestingValidatorRegistrationDuty,
				Messages: []*types.SignedSSVMessage{
					testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.OperatorKeys[1], types.OperatorID(1), testingutils.ValidatorRegistrationMsgID,
						testingutils.TestAttesterConsensusDataByts, qbft.Height(testingutils.TestingDutySlot)),
				},
				PostDutyRunnerStateRoot: "2ac409163b617c79a2a11d3919d6834d24c5c32f06113237a12afcf43e7757a0",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusValidatorRegistrationMsg(ks.Shares[1], 1), // broadcasts when starting a new duty
				},
				ExpectedError: "no consensus phase for validator registration",
			},
			{
				Name:   "voluntary exit",
				Runner: testingutils.VoluntaryExitRunner(ks),
				Duty:   &testingutils.TestingVoluntaryExitDuty,
				Messages: []*types.SignedSSVMessage{
					testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.OperatorKeys[1], types.OperatorID(1), testingutils.VoluntaryExitMsgID,
						testingutils.TestAttesterConsensusDataByts, qbft.Height(testingutils.TestingDutySlot)),
				},
				PostDutyRunnerStateRoot: "2ac409163b617c79a2a11d3919d6834d24c5c32f06113237a12afcf43e7757a0",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusVoluntaryExitMsg(ks.Shares[1], 1), // broadcasts when starting a new duty
				},
				ExpectedError: "no consensus phase for voluntary exit",
			},
		},
	}
}
