package ssv

import (
	"github.com/bloxapp/ssv-spec/types"
	"github.com/pkg/errors"
)

// Validator represents an SSV ETH consensus validator Share assigned, coordinates duty execution and more.
// Every validator has a validatorID which is validator's public key.
// Each validator has multiple DutyRunners, for each duty type.
type Validator struct {
	DutyRunners DutyRunners
	Network     Network
	Beacon      BeaconNode
	Share       *types.Share
	Signer      types.KeyManager
}

func NewValidator(
	network Network,
	beacon BeaconNode,
	share *types.Share,
	signer types.KeyManager,
	runners map[types.BeaconRole]Runner,
) *Validator {
	return &Validator{
		DutyRunners: runners,
		Network:     network,
		Beacon:      beacon,
		Share:       share,
		Signer:      signer,
	}
}

// StartDuty starts a duty for the validator
func (v *Validator) StartDuty(duty *types.Duty) error {
	dutyRunner := v.DutyRunners[duty.Type]
	if dutyRunner == nil {
		return errors.Errorf("duty type %s not supported", duty.Type.String())
	}
	return dutyRunner.StartNewDuty(duty)
}

// ProcessMessage processes Network Message of all types
func (v *Validator) ProcessMessage(signedSSVMsg *types.SignedSSVMessage) error {

	msg := signedSSVMsg.SSVMessage

	dutyRunner := v.DutyRunners.DutyRunnerForMsgID(msg.GetID())
	if dutyRunner == nil {
		return errors.Errorf("could not get duty runner for msg ID")
	}

	if err := v.validateMessage(dutyRunner, msg); err != nil {
		return errors.Wrap(err, "Message invalid")
	}

	switch msg.GetType() {
	case types.SSVConsensusMsgType:
		return dutyRunner.ProcessConsensus(signedSSVMsg)
	case types.SSVPartialSignatureMsgType:
		partialSignatureMessages := &types.PartialSignatureMessages{}
		if err := partialSignatureMessages.Decode(msg.GetData()); err != nil {
			return errors.Wrap(err, "could not get post consensus Message from network Message")
		}

		if partialSignatureMessages.Type == types.PostConsensusPartialSig {
			return dutyRunner.ProcessPostConsensus(partialSignatureMessages)
		}
		return dutyRunner.ProcessPreConsensus(partialSignatureMessages)
	default:
		return errors.New("unknown msg")
	}
}

func (v *Validator) validateMessage(runner Runner, msg *types.SSVMessage) error {
	if !v.Share.ValidatorPubKey.MessageIDBelongs(msg.GetID()) {
		return errors.New("msg ID doesn't match validator ID")
	}

	if len(msg.GetData()) == 0 {
		return errors.New("msg data is invalid")
	}

	return nil
}
