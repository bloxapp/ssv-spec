package preconsensus

import (
	"github.com/bloxapp/ssv-spec/ssv"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// InvalidSignedMessage tests SignedPartialSignature.Validate() != nil
func InvalidSignedMessage() *tests.MultiMsgProcessingSpecTest {
	ks := testingutils.Testing4SharesSet()

	invaldiateMsg := func(msg *ssv.SignedPartialSignature) *ssv.SignedPartialSignature {
		msg.Signature = nil
		return msg
	}

	return &tests.MultiMsgProcessingSpecTest{
		Name: "pre consensus invalid signed msg",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:   "sync committee aggregator selection proof",
				Runner: testingutils.SyncCommitteeContributionRunner(ks),
				Duty:   testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.Message{
					testingutils.SSVMsgSyncCommitteeContribution(nil, invaldiateMsg(testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1)), types.PartialContributionProofSignatureMsgType),
				},
				PostDutyRunnerStateRoot: "7ec85b6d9e1f568c778ba3fdbb26df3d5e32c56628775d7facd8a1d0038dc59e",
				OutputMessages: []*ssv.SignedPartialSignature{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing sync committee selection proof message: invalid pre-consensus message: SignedPartialSignature invalid: SignedPartialSignature sig invalid",
			},
			{
				Name:   "aggregator selection proof",
				Runner: testingutils.AggregatorRunner(ks),
				Duty:   testingutils.TestingAggregatorDuty,
				Messages: []*types.Message{
					testingutils.SSVMsgAggregator(nil, invaldiateMsg(testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1)), types.PartialSelectionProofSignatureMsgType),
				},
				PostDutyRunnerStateRoot: "ac64bab0a5194cd5b5a426b369839dbaeab4ee98f1dc682a558137d4b0ce7063",
				OutputMessages: []*ssv.SignedPartialSignature{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing selection proof message: invalid pre-consensus message: SignedPartialSignature invalid: SignedPartialSignature sig invalid",
			},
			{
				Name:   "randao",
				Runner: testingutils.ProposerRunner(ks),
				Duty:   testingutils.TestingProposerDuty,
				Messages: []*types.Message{
					testingutils.SSVMsgProposer(nil, invaldiateMsg(testingutils.PreConsensusRandaoDifferentSignerMsg(ks.Shares[1], ks.Shares[1], 1, 1)), types.PartialRandaoSignatureMsgType),
				},
				PostDutyRunnerStateRoot: "4011c842c4b65f678c065595a288696dbf88869dbf74b3d6bd980dfcaba07843",
				OutputMessages: []*ssv.SignedPartialSignature{
					testingutils.PreConsensusRandaoMsg(ks.Shares[1], 1), // broadcasts when starting a new duty
				},
				ExpectedError: "failed processing randao message: invalid pre-consensus message: SignedPartialSignature invalid: SignedPartialSignature sig invalid",
			},
		},
	}
}
