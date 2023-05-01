package runner

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/bloxapp/ssv-spec/qbft"
	qbftcomparable "github.com/bloxapp/ssv-spec/qbft/spectest/comparable"
	"github.com/bloxapp/ssv-spec/ssv"
	ssvcomparable "github.com/bloxapp/ssv-spec/ssv/spectest/comparable"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	comparable2 "github.com/bloxapp/ssv-spec/types/testingutils/comparable"
)

// fullHappyFlowStateComparison returns state comparison object for the FullHappyFlow spec test
func fullHappyFlowStateComparison() *ssvcomparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	return &ssvcomparable.StateComparison{
		SyncCommitteeContribution: func() ssv.Runner {
			ret := testingutils.SyncCommitteeContributionRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				Finished:     true,
				DecidedValue: comparable2.FixIssue178(testingutils.TestSyncCommitteeContributionConsensusData, spec.DataVersionPhase0),
				StartingDuty: &testingutils.TestSyncCommitteeContributionConsensusData.Duty,
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					testingutils.SSVDecidingMsgs(testingutils.TestSyncCommitteeContributionConsensusData, ks, types.BNRoleSyncCommitteeContribution)[:3]),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{
						testingutils.SSVMsgSyncCommitteeContribution(nil, testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[1], 1, ks)),
						testingutils.SSVMsgSyncCommitteeContribution(nil, testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[2], 2, ks)),
						testingutils.SSVMsgSyncCommitteeContribution(nil, testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[3], 3, ks)),
					}),
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				StartValue: comparable2.NoErrorEncoding(comparable2.FixIssue178(testingutils.TestSyncCommitteeContributionConsensusData, spec.DataVersionBellatrix)),
				State: &qbft.State{
					Share:  testingutils.TestingShare(testingutils.Testing4SharesSet()),
					ID:     ret.GetBaseRunner().QBFTController.Identifier,
					Round:  qbft.FirstRound,
					Height: qbft.FirstHeight,
					ProposalAcceptedForCurrentRound: testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.Shares[1], types.OperatorID(1),
						ret.GetBaseRunner().QBFTController.Identifier,
						testingutils.TestSyncCommitteeContributionConsensusDataByts,
					),
					LastPreparedRound: 1,
					LastPreparedValue: testingutils.TestSyncCommitteeContributionConsensusDataByts,
					Decided:           true,
					DecidedValue:      testingutils.TestSyncCommitteeContributionConsensusDataByts,
				},
			}
			qbftcomparable.SetMessages(
				ret.GetBaseRunner().State.RunningInstance,
				testingutils.SSVDecidingMsgs(testingutils.TestSyncCommitteeContributionConsensusData, ks, types.BNRoleSyncCommitteeContribution)[3:10],
			)
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)
			return ret
		}(),
		SyncCommittee: func() ssv.Runner {
			ret := testingutils.SyncCommitteeRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				Finished:     true,
				DecidedValue: comparable2.FixIssue178(testingutils.TestSyncCommitteeConsensusData, spec.DataVersionPhase0),
				StartingDuty: &testingutils.TestSyncCommitteeConsensusData.Duty,
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{}),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[1], 1)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[2], 2)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[3], 3)),
					}),
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				StartValue: comparable2.NoErrorEncoding(testingutils.TestSyncCommitteeConsensusData),
				State: &qbft.State{
					Share:  testingutils.TestingShare(testingutils.Testing4SharesSet()),
					ID:     ret.GetBaseRunner().QBFTController.Identifier,
					Round:  qbft.FirstRound,
					Height: qbft.FirstHeight,
					ProposalAcceptedForCurrentRound: testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.Shares[1], types.OperatorID(1),
						ret.GetBaseRunner().QBFTController.Identifier,
						testingutils.TestSyncCommitteeConsensusDataByts,
					),
					LastPreparedRound: 1,
					LastPreparedValue: testingutils.TestSyncCommitteeConsensusDataByts,
					Decided:           true,
					DecidedValue:      testingutils.TestSyncCommitteeConsensusDataByts,
				},
			}
			qbftcomparable.SetMessages(
				ret.GetBaseRunner().State.RunningInstance,
				testingutils.SSVDecidingMsgs(testingutils.TestSyncCommitteeConsensusData, ks, types.BNRoleSyncCommittee)[0:7],
			)
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)
			return ret
		}(),
		Aggregator: func() ssv.Runner {
			ret := testingutils.AggregatorRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				Finished:     true,
				DecidedValue: comparable2.FixIssue178(testingutils.TestAggregatorConsensusData, spec.DataVersionPhase0),
				StartingDuty: &testingutils.TestAggregatorConsensusData.Duty,
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					testingutils.SSVDecidingMsgs(testingutils.TestAggregatorConsensusData, ks, types.BNRoleAggregator)[0:3]),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusAggregatorMsg(ks.Shares[1], 1)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusAggregatorMsg(ks.Shares[2], 2)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusAggregatorMsg(ks.Shares[3], 3)),
					}),
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				StartValue: comparable2.NoErrorEncoding(testingutils.TestAggregatorConsensusData),
				State: &qbft.State{
					Share:  testingutils.TestingShare(testingutils.Testing4SharesSet()),
					ID:     ret.GetBaseRunner().QBFTController.Identifier,
					Round:  qbft.FirstRound,
					Height: qbft.FirstHeight,
					ProposalAcceptedForCurrentRound: testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.Shares[1], types.OperatorID(1),
						ret.GetBaseRunner().QBFTController.Identifier,
						testingutils.TestAggregatorConsensusDataByts,
					),
					LastPreparedRound: 1,
					LastPreparedValue: testingutils.TestAggregatorConsensusDataByts,
					Decided:           true,
					DecidedValue:      testingutils.TestAggregatorConsensusDataByts,
				},
			}
			qbftcomparable.SetMessages(
				ret.GetBaseRunner().State.RunningInstance,
				testingutils.SSVDecidingMsgs(testingutils.TestAggregatorConsensusData, ks, types.BNRoleAggregator)[3:10],
			)
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)
			return ret
		}(),
		Proposer: func() ssv.Runner {
			ret := testingutils.ProposerRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				Finished:     true,
				DecidedValue: comparable2.FixIssue178(testingutils.TestProposerConsensusData, spec.DataVersionCapella),
				StartingDuty: &testingutils.TestProposerConsensusData.Duty,
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					testingutils.SSVDecidingMsgs(testingutils.TestProposerConsensusData, ks, types.BNRoleProposer)[0:3]),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusProposerMsg(ks.Shares[1], 1)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusProposerMsg(ks.Shares[2], 2)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusProposerMsg(ks.Shares[3], 3)),
					}),
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				StartValue: comparable2.NoErrorEncoding(testingutils.TestProposerConsensusData),
				State: &qbft.State{
					Share:  testingutils.TestingShare(testingutils.Testing4SharesSet()),
					ID:     ret.GetBaseRunner().QBFTController.Identifier,
					Round:  qbft.FirstRound,
					Height: qbft.FirstHeight,
					ProposalAcceptedForCurrentRound: testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.Shares[1], types.OperatorID(1),
						ret.GetBaseRunner().QBFTController.Identifier,
						testingutils.TestProposerConsensusDataByts,
					),
					LastPreparedRound: 1,
					LastPreparedValue: testingutils.TestProposerConsensusDataByts,
					Decided:           true,
					DecidedValue:      testingutils.TestProposerConsensusDataByts,
				},
			}
			qbftcomparable.SetMessages(
				ret.GetBaseRunner().State.RunningInstance,
				testingutils.SSVDecidingMsgs(testingutils.TestProposerConsensusData, ks, types.BNRoleProposer)[3:10],
			)
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)
			return ret
		}(),
		BlindedProposer: func() ssv.Runner {
			ret := testingutils.ProposerBlindedBlockRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				Finished:     true,
				DecidedValue: comparable2.FixIssue178(testingutils.TestProposerBlindedBlockConsensusData, spec.DataVersionCapella),
				StartingDuty: &testingutils.TestProposerConsensusData.Duty,
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					testingutils.SSVDecidingMsgs(testingutils.TestProposerBlindedBlockConsensusData, ks, types.BNRoleProposer)[0:3]),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusProposerMsg(ks.Shares[1], 1)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusProposerMsg(ks.Shares[2], 2)),
						testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusProposerMsg(ks.Shares[3], 3)),
					}),
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				StartValue: comparable2.NoErrorEncoding(testingutils.TestProposerBlindedBlockConsensusData),
				State: &qbft.State{
					Share:  testingutils.TestingShare(testingutils.Testing4SharesSet()),
					ID:     ret.GetBaseRunner().QBFTController.Identifier,
					Round:  qbft.FirstRound,
					Height: qbft.FirstHeight,
					ProposalAcceptedForCurrentRound: testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.Shares[1], types.OperatorID(1),
						ret.GetBaseRunner().QBFTController.Identifier,
						testingutils.TestProposerBlindedBlockConsensusDataByts,
					),
					LastPreparedRound: 1,
					LastPreparedValue: testingutils.TestProposerBlindedBlockConsensusDataByts,
					Decided:           true,
					DecidedValue:      testingutils.TestProposerBlindedBlockConsensusDataByts,
				},
			}
			qbftcomparable.SetMessages(
				ret.GetBaseRunner().State.RunningInstance,
				testingutils.SSVDecidingMsgs(testingutils.TestProposerBlindedBlockConsensusData, ks, types.BNRoleProposer)[3:10],
			)
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)
			return ret
		}(),
		Attester: func() ssv.Runner {
			ret := testingutils.AttesterRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				Finished:     true,
				DecidedValue: comparable2.FixIssue178(testingutils.TestAttesterConsensusData, spec.DataVersionPhase0),
				StartingDuty: &testingutils.TestAttesterConsensusData.Duty,
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{}),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{
						testingutils.SSVMsgAttester(nil, testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight)),
						testingutils.SSVMsgAttester(nil, testingutils.PostConsensusAttestationMsg(ks.Shares[2], 2, qbft.FirstHeight)),
						testingutils.SSVMsgAttester(nil, testingutils.PostConsensusAttestationMsg(ks.Shares[3], 3, qbft.FirstHeight)),
					}),
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				StartValue: comparable2.NoErrorEncoding(testingutils.TestAttesterConsensusData),
				State: &qbft.State{
					Share:  testingutils.TestingShare(testingutils.Testing4SharesSet()),
					ID:     ret.GetBaseRunner().QBFTController.Identifier,
					Round:  qbft.FirstRound,
					Height: qbft.FirstHeight,
					ProposalAcceptedForCurrentRound: testingutils.TestingProposalMessageWithIdentifierAndFullData(
						ks.Shares[1], types.OperatorID(1),
						ret.GetBaseRunner().QBFTController.Identifier,
						testingutils.TestAttesterConsensusDataByts,
					),
					LastPreparedRound: 1,
					LastPreparedValue: testingutils.TestAttesterConsensusDataByts,
					Decided:           true,
					DecidedValue:      testingutils.TestAttesterConsensusDataByts,
				},
			}
			qbftcomparable.SetMessages(
				ret.GetBaseRunner().State.RunningInstance,
				testingutils.SSVDecidingMsgs(testingutils.TestAttesterConsensusData, ks, types.BNRoleAttester)[0:7],
			)
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)
			return ret
		}(),
		ValidatorRegistration: func() ssv.Runner {
			ret := testingutils.ValidatorRegistrationRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				Finished:     true,
				StartingDuty: &testingutils.TestingValidatorRegistrationDuty,
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{
						testingutils.SSVMsgValidatorRegistration(nil, testingutils.PreConsensusValidatorRegistrationMsg(ks.Shares[1], 1)),
						testingutils.SSVMsgValidatorRegistration(nil, testingutils.PreConsensusValidatorRegistrationMsg(ks.Shares[2], 2)),
						testingutils.SSVMsgValidatorRegistration(nil, testingutils.PreConsensusValidatorRegistrationMsg(ks.Shares[3], 3)),
					}),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SSVMessage{}),
			}
			return ret
		}(),
	}
}
