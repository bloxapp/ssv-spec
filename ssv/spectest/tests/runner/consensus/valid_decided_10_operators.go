package consensus

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// ValidDecided10Operators tests a valid decided value (10 operators)
func ValidDecided10Operators() *tests.MultiMsgProcessingSpecTest {
	ks := testingutils.Testing10SharesSet()
	return &tests.MultiMsgProcessingSpecTest{
		Name: "consensus valid decided 10 operators",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:                    "sync committee contribution",
				Runner:                  testingutils.SyncCommitteeContributionRunner(ks),
				Duty:                    &testingutils.TestingSyncCommitteeContributionDuty,
				Messages:                testingutils.SSVDecidingMsgs(testingutils.TestSyncCommitteeContributionConsensusData, ks, types.BNRoleSyncCommitteeContribution),
				PostDutyRunnerStateRoot: "275f20b2a276fe81528ac1e723972a79d912e5b0be976c267ac150ea89ac9398",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
					testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[1], 1, ks),
				},
			},
			{
				Name:                    "sync committee",
				Runner:                  testingutils.SyncCommitteeRunner(ks),
				Duty:                    &testingutils.TestingSyncCommitteeDuty,
				Messages:                testingutils.SSVDecidingMsgs(testingutils.TestSyncCommitteeConsensusData, ks, types.BNRoleSyncCommittee),
				PostDutyRunnerStateRoot: "e0d46701cedd550808918a83ff8791127a5a4c81f3a88cceeb657695bf68beb0",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[1], 1),
				},
			},
			{
				Name:                    "aggregator",
				Runner:                  testingutils.AggregatorRunner(ks),
				Duty:                    &testingutils.TestingAggregatorDuty,
				Messages:                testingutils.SSVDecidingMsgs(testingutils.TestAggregatorConsensusData, ks, types.BNRoleAggregator),
				PostDutyRunnerStateRoot: "6c2332c81220f5f10a233c6224e8d457fe71c316886d64b79420e8200a5c51cf",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
					testingutils.PostConsensusAggregatorMsg(ks.Shares[1], 1),
				},
			},
			{
				Name:                    "proposer",
				Runner:                  testingutils.ProposerRunner(ks),
				Duty:                    &testingutils.TestingProposerDuty,
				Messages:                testingutils.SSVDecidingMsgs(testingutils.TestProposerConsensusData(spec.DataVersionBellatrix), ks, types.BNRoleProposer),
				PostDutyRunnerStateRoot: "c6e565152219704d6a445d539b3b569b05696532d3854532d82c44880b15ca97",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsg(ks.Shares[1], 1),
					testingutils.PostConsensusProposerMsg(ks.Shares[1], 1, spec.DataVersionBellatrix),
				},
			},
			{
				Name:                    "proposer (blinded block)",
				Runner:                  testingutils.ProposerBlindedBlockRunner(ks),
				Duty:                    &testingutils.TestingProposerDuty,
				Messages:                testingutils.SSVDecidingMsgs(testingutils.TestProposerBlindedBlockConsensusData(spec.DataVersionBellatrix), ks, types.BNRoleProposer),
				PostDutyRunnerStateRoot: "edbe77ae3456264f3aecd1259ada249bb5ddaca5f31f76b8e72bcacf9df3abc8",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsg(ks.Shares[1], 1),
					testingutils.PostConsensusProposerMsg(ks.Shares[1], 1, spec.DataVersionBellatrix),
				},
			},
			{
				Name:                    "attester",
				Runner:                  testingutils.AttesterRunner(ks),
				Duty:                    &testingutils.TestingAttesterDuty,
				Messages:                testingutils.SSVDecidingMsgs(testingutils.TestAttesterConsensusData, ks, types.BNRoleAttester),
				PostDutyRunnerStateRoot: "0535cd4361ffbe961f71420c8b85dd784c89566f47b2ef71da00b79c836247c1",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight),
				},
			},
		},
	}
}
