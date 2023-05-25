package preconsensus

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// InvalidSignedMessage tests SignedPartialSignatureMessage.Validate() != nil
func InvalidSignedMessage() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	invalidMsg := func(msg *types.SignedPartialSignatureMessage) *types.SignedPartialSignatureMessage {
		msg.Signer = 0
		return msg
	}

	return &tests.MultiMsgProcessingSpecTest{
		Name: "pre consensus invalid signed msg",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:   "sync committee aggregator selection proof",
				Runner: testingutils.SyncCommitteeContributionRunner(ks),
				Duty:   &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgSyncCommitteeContribution(nil, invalidMsg(testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1))),
				},
				PostDutyRunnerStateRoot: "eece7b3ec4c7e2c5576a8c28bba3e8ff816c8f84b769aac74dfd244de010e61a",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing sync committee selection proof message: invalid pre-consensus message: SignedPartialSignatureMessage invalid: signer ID 0 not allowed",
			},
			{
				Name:   "aggregator selection proof",
				Runner: testingutils.AggregatorRunner(ks),
				Duty:   &testingutils.TestingAggregatorDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgAggregator(nil, invalidMsg(testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1))),
				},
				PostDutyRunnerStateRoot: "be39d691ce8f6c800779bac909316a2bc869bd05cda24306a368bac9bf301678",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing selection proof message: invalid pre-consensus message: SignedPartialSignatureMessage invalid: signer ID 0 not allowed",
			},
			{
				Name:   "randao",
				Runner: testingutils.ProposerRunner(ks),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionBellatrix),
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgProposer(nil, invalidMsg(testingutils.PreConsensusRandaoDifferentSignerMsgV(ks.Shares[1], ks.Shares[1], 1, 1, spec.DataVersionBellatrix))),
				},
				PostDutyRunnerStateRoot: "75232add1f62f04503a93a162aaf353da9360399c3dc4131d94a4ea51b4ede88",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing randao message: invalid pre-consensus message: SignedPartialSignatureMessage invalid: signer ID 0 not allowed",
			},
			{
				Name:   "randao (blinded block)",
				Runner: testingutils.ProposerBlindedBlockRunner(ks),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionBellatrix),
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgProposer(nil, invalidMsg(testingutils.PreConsensusRandaoDifferentSignerMsgV(ks.Shares[1], ks.Shares[1], 1, 1, spec.DataVersionBellatrix))),
				},
				PostDutyRunnerStateRoot: "7a8cafd719c06d5ad08f96525b989c6f49661a0ef3bcf0b1ad7d644e258e8945",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionBellatrix), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing randao message: invalid pre-consensus message: SignedPartialSignatureMessage invalid: signer ID 0 not allowed",
			},
			{
				Name:   "validator registration",
				Runner: testingutils.ValidatorRegistrationRunner(ks),
				Duty:   &testingutils.TestingValidatorRegistrationDuty,
				Messages: []*types.SSVMessage{
					testingutils.SSVMsgValidatorRegistration(nil, invalidMsg(testingutils.PreConsensusValidatorRegistrationMsg(ks.Shares[1], 1))),
				},
				PostDutyRunnerStateRoot: "52c92a192a3ec5d7aae78adcc291ce50411d853acaad86dda679bbf02f0a59db",
				OutputMessages: []*types.SignedPartialSignatureMessage{
					testingutils.PreConsensusValidatorRegistrationMsg(ks.Shares[1], 1), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing validator registration message: invalid pre-consensus message: SignedPartialSignatureMessage invalid: signer ID 0 not allowed",
			},
		},
	}
}
