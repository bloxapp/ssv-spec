package pre_consensus_justifications

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/ssvlabs/ssv-spec/qbft"
	"github.com/ssvlabs/ssv-spec/ssv"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

var decideFirstHeight = func(r ssv.Runner) ssv.Runner {
	i := qbft.NewInstance(
		r.GetBaseRunner().QBFTController.GetConfig(),
		r.GetBaseRunner().QBFTController.CommitteeMember,
		r.GetBaseRunner().QBFTController.Identifier,
		qbft.FirstHeight)
	i.State.DecidedValue = []byte{1, 2, 3, 4}
	i.State.Decided = true
	r.GetBaseRunner().QBFTController.StoredInstances = append(r.GetBaseRunner().QBFTController.StoredInstances, i)
	return r
}

// Valid tests a valid pre-consensus justification
func Valid() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	msgF := func(obj *types.ValidatorConsensusData, id []byte) *types.SignedSSVMessage {
		fullData, _ := obj.Encode()
		root, _ := qbft.HashDataRoot(fullData)
		msg := &qbft.Message{
			MsgType:    qbft.ProposalMsgType,
			Height:     1,
			Round:      qbft.FirstRound,
			Identifier: id,
			Root:       root,
		}
		signed := testingutils.SignQBFTMsg(ks.OperatorKeys[1], 1, msg)
		signed.FullData = fullData

		return signed
	}

	return &tests.MultiMsgProcessingSpecTest{
		Name: "pre consensus valid quorum",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:   "sync committee aggregator selection proof",
				Runner: decideFirstHeight(testingutils.SyncCommitteeContributionRunner(ks)),
				Duty:   &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestContributionProofWithJustificationsConsensusData(ks), testingutils.SyncCommitteeContributionMsgID),
				},
				PostDutyRunnerStateRoot: "118bb07beaee54ffa3f4a676f752cdf391692b5981066b4827a8c29d100756ed",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
			},
			{
				Name:   "aggregator selection proof",
				Runner: decideFirstHeight(testingutils.AggregatorRunner(ks)),
				Duty:   &testingutils.TestingAggregatorDuty,
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestSelectionProofWithJustificationsConsensusData(ks), testingutils.AggregatorMsgID),
				},
				PostDutyRunnerStateRoot: "c19537feb74079e087bc1f3b22433b9a6c488257e1e1508c45852a63ce590e13",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
			},
			{
				Name:   "randao",
				Runner: decideFirstHeight(testingutils.ProposerRunner(ks)),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestProposerWithJustificationsConsensusDataV(ks, spec.DataVersionDeneb), testingutils.ProposerMsgID),
				},
				PostDutyRunnerStateRoot: "41fb0f6c6d7a6e4000ca874ad9d3d85c1f0a93dc67a85cfec89fa542b4d5a712",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionDeneb), // broadcasts when starting a new duty
				},
			},
			{
				Name:   "randao (blinded block)",
				Runner: decideFirstHeight(testingutils.ProposerBlindedBlockRunner(ks)),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestProposerBlindedWithJustificationsConsensusDataV(ks, spec.DataVersionDeneb), testingutils.ProposerMsgID),
				},
				PostDutyRunnerStateRoot: "0a40f16f4f05f444f7b4e589ad6edbe78de359302b10788c48a5e64e41047324",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionDeneb), // broadcasts when starting a new duty
				},
			},
		},
	}
}
