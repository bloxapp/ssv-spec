package ssv

import (
	"bytes"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
)

func (b *BaseRunner) validatePreConsensusMsg(runner Runner, signedMsg *SignedPartialSignatureMessage) error {
	if !b.hasRunningDuty() {
		return errors.New("no running duty")
	}

	if err := b.validatePartialSigMsg(signedMsg, b.State.StartingDuty.Slot); err != nil {
		return err
	}

	roots, domain, err := runner.expectedPreConsensusRootsAndDomain()
	if err != nil {
		return err
	}

	return b.verifyExpectedRoot(runner, signedMsg, roots, domain)
}

func (b *BaseRunner) validateConsensusMsg(msg *qbft.SignedMessage) error {
	if !b.hasRunningDuty() {
		return errors.New("no running duty")
	}
	if b.State.RunningInstance == nil {
		return errors.New("no running consensus instance")
	}
	return nil
}

func (b *BaseRunner) validatePostConsensusMsg(runner Runner, signedMsg *SignedPartialSignatureMessage) error {
	if !b.hasRunningDuty() {
		return errors.New("no running duty")
	}

	if b.State.RunningInstance == nil {
		return errors.New("no running consensus instance")
	}
	if decided, _ := b.State.RunningInstance.IsDecided(); !decided {
		return errors.New("consensus instance not decided")
	}

	if err := b.validatePartialSigMsg(signedMsg, b.State.StartingDuty.Slot); err != nil {
		return err
	}

	roots, domain, err := runner.expectedPostConsensusRootsAndDomain()
	if err != nil {
		return err
	}

	return b.verifyExpectedRoot(runner, signedMsg, roots, domain)
}

func (b *BaseRunner) validateDecidedConsensusData(runner Runner, val *types.ConsensusData) error {
	byts, err := val.Encode()
	if err != nil {
		return errors.Wrap(err, "could not encode decided value")
	}
	if err := runner.GetValCheckF()(byts); err != nil {
		return errors.Wrap(err, "decided value is invalid")
	}

	return nil
}

func (b *BaseRunner) verifyExpectedRoot(runner Runner, signedMsg *SignedPartialSignatureMessage, expectedRootObjs []ssz.HashRoot, domain spec.DomainType) error {
	if len(expectedRootObjs) != len(signedMsg.Message.Messages) {
		return errors.New("wrong expected roots count")
	}
	for i, msg := range signedMsg.Message.Messages {
		epoch := b.BeaconNetwork.EstimatedEpochAtSlot(b.State.StartingDuty.Slot)
		d, err := runner.GetBeaconNode().DomainData(epoch, domain)
		if err != nil {
			return errors.Wrap(err, "could not get pre consensus root domain")
		}

		r, err := types.ComputeETHSigningRoot(expectedRootObjs[i], d)
		if err != nil {
			return errors.Wrap(err, "could not compute ETH signing root")
		}
		if !bytes.Equal(r[:], msg.SigningRoot) {
			return errors.New("wrong signing root")
		}
	}
	return nil
}
