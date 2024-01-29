package spectest

import (
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/consensus"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/duties/newduty"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/duties/proposer"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/duties/synccommitteeaggregator"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/preconsensus"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck/valcheckattestations"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck/valcheckduty"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck/valcheckproposer"
)

var AllTests = []tests.TestF{
	runner.FullHappyFlow,

	newduty.ConsensusNotStarted,
	newduty.NotDecided,
	newduty.PostDecided,
	newduty.Finished,
	newduty.Valid,
	newduty.PostWrongDecided,
	newduty.PostInvalidDecided,
	newduty.PostFutureDecided,
	newduty.DuplicateDutyFinished,
	newduty.DuplicateDutyNotFinished,
	newduty.FirstHeight,

	consensus.FutureDecided,
	consensus.InvalidDecidedValue,
	consensus.FutureMessage,
	consensus.PastMessage,
	consensus.PostFinish,
	consensus.PostDecided,
	consensus.ValidDecided,
	consensus.ValidDecided7Operators,
	consensus.ValidDecided10Operators,
	consensus.ValidDecided13Operators,
	consensus.ValidMessage,

	synccommitteeaggregator.SomeAggregatorQuorum,
	synccommitteeaggregator.NoneAggregatorQuorum,
	synccommitteeaggregator.AllAggregatorQuorum,

	proposer.ProposeBlindedBlockDecidedRegular,
	proposer.ProposeRegularBlockDecidedBlinded,
	proposer.BlindedRunnerAcceptsNormalBlock,
	proposer.NormalProposerAcceptsBlindedBlock,

	// pre_consensus_justifications.PastSlot,
	// pre_consensus_justifications.InvalidData,
	// pre_consensus_justifications.FutureHeight,
	// pre_consensus_justifications.PastHeight,
	// pre_consensus_justifications.InvalidMsgType,
	// pre_consensus_justifications.WrongBeaconRole,
	// pre_consensus_justifications.InvalidConsensusData,
	// pre_consensus_justifications.InvalidSlot,
	// pre_consensus_justifications.UnknownSigner,
	// pre_consensus_justifications.InvalidJustificationSignature,
	// pre_consensus_justifications.DuplicateJustificationSigner,
	// pre_consensus_justifications.DuplicateRoots,
	// pre_consensus_justifications.InconsistentRootCount,
	// pre_consensus_justifications.InconsistentRoots,
	// pre_consensus_justifications.InvalidJustification,
	// pre_consensus_justifications.MissingQuorum,
	// pre_consensus_justifications.DecidedInstance,
	// pre_consensus_justifications.ExistingValidPreConsensus,
	// pre_consensus_justifications.Valid,
	// pre_consensus_justifications.Valid7Operators,
	// pre_consensus_justifications.Valid10Operators,
	// pre_consensus_justifications.Valid13Operators,
	// pre_consensus_justifications.ValidFirstHeight,
	// pre_consensus_justifications.ValidNoRunningDuty,
	// pre_consensus_justifications.ValidRoundChangeMsg,
	// pre_consensus_justifications.HappyFlow,

	preconsensus.NoRunningDuty,
	preconsensus.TooFewRoots,
	preconsensus.TooManyRoots,
	preconsensus.UnorderedExpectedRoots,
	preconsensus.InvalidSignedMessage,
	preconsensus.InvalidExpectedRoot,
	preconsensus.DuplicateMsg,
	preconsensus.DuplicateMsgDifferentRoots,
	preconsensus.PostFinish,
	preconsensus.PostDecided,
	preconsensus.PostQuorum,
	preconsensus.Quorum,
	preconsensus.Quorum7Operators,
	preconsensus.Quorum10Operators,
	preconsensus.Quorum13Operators,
	preconsensus.ValidMessage,
	preconsensus.InvalidMessageSlot,
	preconsensus.ValidMessage7Operators,
	preconsensus.ValidMessage10Operators,
	preconsensus.ValidMessage13Operators,
	preconsensus.InconsistentBeaconSigner,
	preconsensus.UnknownSigner,
	preconsensus.InvalidBeaconSignature,
	preconsensus.InvalidMessageSignature,

	valcheckduty.WrongValidatorIndex,
	valcheckduty.WrongValidatorPK,
	valcheckduty.WrongDutyType,
	valcheckduty.FarFutureDutySlot,
	valcheckattestations.Slashable,
	valcheckattestations.SourceHigherThanTarget,
	valcheckattestations.FarFutureTarget,
	valcheckattestations.CommitteeIndexMismatch,
	valcheckattestations.SlotMismatch,
	valcheckattestations.ConsensusDataNil,
	valcheckattestations.Valid,
	valcheckproposer.BlindedBlock,
}
