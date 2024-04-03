package proposer

import (
	"fmt"

	"github.com/attestantio/go-eth2-client/spec"

	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// NormalProposerAcceptsBlindedBlock tests a full happy flow for a proposer runner that produces normal blocks but receives as proposal a blinded block
func NormalProposerAcceptsBlindedBlock() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	multiSpecTest := &tests.MultiMsgProcessingSpecTest{
		Name:  "normal proposer accepts blinded block proposal",
		Tests: []*tests.MsgProcessingSpecTest{},
	}

	// proposerReceivingBlindedBlockV creates a test specification for versioned normal proposer receiving blinded block as proposal.
	proposerReceivingBlindedBlockV := func(version spec.DataVersion) *tests.MsgProcessingSpecTest {
		return &tests.MsgProcessingSpecTest{
			Name:   fmt.Sprintf("normal proposer accepts blinded block proposal (%s)", version.String()),
			Runner: testingutils.ProposerRunner(ks),
			Duty:   testingutils.TestingProposerDutyV(version),
			Messages: append(
				testingutils.SSVDecidingMsgsV(testingutils.TestProposerBlindedBlockConsensusDataV(version), ks, types.BNRoleProposer), // consensus
				[]*types.SignedSSVMessage{ // post consensus
					testingutils.SSVMsgProposer(1, ks.NetworkKeys[1], nil, testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, version)),
					testingutils.SSVMsgProposer(2, ks.NetworkKeys[2], nil, testingutils.PostConsensusProposerMsgV(ks.Shares[2], 2, version)),
					testingutils.SSVMsgProposer(3, ks.NetworkKeys[3], nil, testingutils.PostConsensusProposerMsgV(ks.Shares[3], 3, version)),
				}...,
			),
			PostDutyRunnerStateRoot: fullHappyFlowProposerReceivingBlindedBlockSC(version).Root(),
			PostDutyRunnerState:     fullHappyFlowProposerReceivingBlindedBlockSC(version).ExpectedState,
			OutputMessages: []*types.PartialSignatureMessages{
				testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, version),
				testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, version),
			},
			BeaconBroadcastedRoots: []string{
				testingutils.GetSSZRootNoError(testingutils.TestingSignedBlindedBeaconBlockV(ks, version)),
			},
		}
	}

	for _, v := range testingutils.SupportedBlockVersions {
		multiSpecTest.Tests = append(multiSpecTest.Tests, []*tests.MsgProcessingSpecTest{proposerReceivingBlindedBlockV(v)}...)
	}

	return multiSpecTest
}
