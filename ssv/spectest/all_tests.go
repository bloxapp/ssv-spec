package spectest

import (
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/consensus"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/duties/newduty"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/postconsensus"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/runner/preconsensus"
	"testing"
)

type SpecTest interface {
	TestName() string
	Run(t *testing.T)
}

var AllTests = []SpecTest{
	// sanity tests - begin

	runner.FullHappyFlow(),
	newduty.NotDecided(),
	newduty.PostDecided(),

	preconsensus.NoRunningDuty(),
	preconsensus.DuplicateMsg(),
	preconsensus.PostFinish(),
	preconsensus.PostDecided(),
	preconsensus.PostQuorum(),
	preconsensus.Quorum(),
	preconsensus.Quorum7Operators(),
	preconsensus.UnknownSigner(),

	consensus.NoRunningDuty(),
	consensus.NoRunningConsensusInstance(),
	consensus.PostFinish(),
	consensus.PostDecided(),
	consensus.ValidDecided(),
	consensus.ValidDecided7Operators(),
	consensus.FutureDecided(),

	postconsensus.UnknownSigner(),
	postconsensus.InconsistentBeaconSigner(),
	postconsensus.DuplicateMsgDifferentRoots(),
	postconsensus.DuplicateMsg(),
	postconsensus.PreDecided(),
	postconsensus.PostQuorum(),
	postconsensus.Quorum(),
	postconsensus.Quorum7Operators(),

	// sanity tests - end

	//runner.FullHappyFlow(),
	//
	//postconsensus.TooManyRoots(),
	//postconsensus.TooFewRoots(),
	//postconsensus.UnorderedExpectedRoots(),
	//postconsensus.UnknownSigner(),
	//postconsensus.InconsistentBeaconSigner(),
	//postconsensus.PostFinish(),
	//postconsensus.NoRunningDuty(),
	//postconsensus.InvalidMessageSignature(),
	//postconsensus.InvalidBeaconSignature(),
	//postconsensus.DuplicateMsgDifferentRoots(),
	//postconsensus.DuplicateMsg(),
	//postconsensus.InvalidExpectedRoot(),
	//postconsensus.PreDecided(),
	//postconsensus.PostQuorum(),
	//postconsensus.InvalidMessage(),
	//postconsensus.InValidMessageSlot(),
	//postconsensus.ValidMessage(),
	//postconsensus.ValidMessage7Operators(),
	//postconsensus.ValidMessage10Operators(),
	//postconsensus.ValidMessage13Operators(),
	//postconsensus.Quorum(),
	//postconsensus.Quorum7Operators(),
	//postconsensus.Quorum10Operators(),
	//postconsensus.Quorum13Operators(),
	//postconsensus.InvalidDecidedValue(),
	//
	//newduty.ConsensusNotStarted(),
	//newduty.NotDecided(),
	//newduty.PostDecided(),
	//newduty.Finished(),
	//newduty.Valid(),
	//newduty.PostWrongDecided(),
	//newduty.PostInvalidDecided(),
	//newduty.PostFutureDecided(),
	//
	//consensus.FutureDecided(),
	//consensus.InvalidDecidedValue(),
	//consensus.NoRunningDuty(),
	//consensus.NoRunningConsensusInstance(),
	//consensus.PostFinish(),
	//consensus.PostDecided(),
	//consensus.ValidDecided(),
	//consensus.ValidDecided7Operators(),
	//consensus.ValidDecided10Operators(),
	//consensus.ValidDecided13Operators(),
	//consensus.ValidMessage(),
	//
	//synccommitteeaggregator.SomeAggregatorQuorum(),
	//synccommitteeaggregator.NoneAggregatorQuorum(),
	//synccommitteeaggregator.AllAggregatorQuorum(),
	//
	//proposer.ProposeBlindedBlockDecidedRegular(),
	//proposer.ProposeRegularBlockDecidedBlinded(),
	//
	//preconsensus.NoRunningDuty(),
	//preconsensus.TooFewRoots(),
	//preconsensus.TooManyRoots(),
	//preconsensus.UnorderedExpectedRoots(),
	//preconsensus.InvalidSignedMessage(),
	//preconsensus.InvalidExpectedRoot(),
	//preconsensus.DuplicateMsg(),
	//preconsensus.DuplicateMsgDifferentRoots(),
	//preconsensus.PostFinish(),
	//preconsensus.PostDecided(),
	//preconsensus.PostQuorum(),
	//preconsensus.Quorum(),
	//preconsensus.Quorum7Operators(),
	//preconsensus.Quorum10Operators(),
	//preconsensus.Quorum130Operators(),
	//preconsensus.ValidMessage(),
	//preconsensus.InValidMessageSlot(),
	//preconsensus.ValidMessage7Operators(),
	//preconsensus.ValidMessage10Operators(),
	//preconsensus.ValidMessage13Operators(),
	//preconsensus.InconsistentBeaconSigner(),
	//preconsensus.UnknownSigner(),
	//preconsensus.InvalidBeaconSignature(),
	//preconsensus.InvalidMessageSignature(),
	//
	//messages.EncodingAndRoot(),
	//messages.NoMsgs(),
	//messages.InvalidMsg(),
	//messages.ValidContributionProofMetaData(),
	//messages.SigValid(),
	//messages.SigTooShort(),
	//messages.SigTooLong(),
	//messages.PartialSigValid(),
	//messages.PartialSigTooShort(),
	//messages.PartialSigTooLong(),
	//messages.PartialRootValid(),
	//messages.MessageSigner0(),
	//messages.SignedMsgSigner0(),
	//
	//valcheckduty.WrongValidatorIndex(),
	//valcheckduty.WrongValidatorPK(),
	//valcheckduty.WrongDutyType(),
	//valcheckduty.FarFutureDutySlot(),
	//valcheckattestations.Slashable(),
	//valcheckattestations.SourceHigherThanTarget(),
	//valcheckattestations.FarFutureTarget(),
	//valcheckattestations.CommitteeIndexMismatch(),
	//valcheckattestations.SlotMismatch(),
	//valcheckattestations.AttestationDataNil(),
	//valcheckattestations.Valid(),
}
