package validation

import (
	"context"

	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/ssv"
	"github.com/bloxapp/ssv-spec/types"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
)

// MsgValidatorFunc represents a message validator
type MsgValidatorFunc = func(ctx context.Context, p peer.ID, msg *pubsub.Message) pubsub.ValidationResult

func MsgValidation(runner ssv.Runner) MsgValidatorFunc {
	return func(ctx context.Context, p peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
		signedSSVMsg, err := DecodePubsubMsg(msg)
		if err != nil {
			return pubsub.ValidationReject
		}

		// Validate SignedSSVMessage
		if validateSignedSSVMessage(runner, signedSSVMsg) != nil {
			return pubsub.ValidationReject
		}

		switch signedSSVMsg.SSVMessage.GetType() {
		case types.SSVConsensusMsgType:
			if validateConsensusMsg(runner, signedSSVMsg) != nil {
				return pubsub.ValidationReject
			}
		case types.SSVPartialSignatureMsgType:
			if validatePartialSigMsg(runner, signedSSVMsg.SSVMessage.Data) != nil {
				return pubsub.ValidationReject
			}
		default:
			return pubsub.ValidationReject
		}

		return pubsub.ValidationAccept
	}
}

func DecodePubsubMsg(msg *pubsub.Message) (*types.SignedSSVMessage, error) {
	byts := msg.GetData()
	ret := &types.SignedSSVMessage{}
	if err := ret.Decode(byts); err != nil {
		return nil, err
	}
	return ret, nil
}

func validateSignedSSVMessage(runner ssv.Runner, msg *types.SignedSSVMessage) error {

	if err := msg.Validate(); err != nil {
		return err
	}

	return validateSSVMessage(runner, msg.SSVMessage)
}

func validateSSVMessage(runner ssv.Runner, msg *types.SSVMessage) error {
	if !runner.GetBaseRunner().Share.ValidatorPubKey.MessageIDBelongs(msg.GetID()) {
		return errors.New("msg ID doesn't match validator ID")
	}

	if len(msg.GetData()) == 0 {
		return errors.New("msg data is invalid")
	}

	return nil
}

func validateConsensusMsg(runner ssv.Runner, signedSSVMessage *types.SignedSSVMessage) error {

	contr := runner.GetBaseRunner().QBFTController

	// Decode
	message := &qbft.Message{}
	if err := message.Decode(signedSSVMessage.SSVMessage.Data); err != nil {
		return errors.Wrap(err, "could not decode Message")
	}

	if err := contr.BaseMsgValidation(message); err != nil {
		return err
	}

	/**
	Main controller processing flow
	_______________________________
	All decided msgs are processed the same, out of instance
	All valid future msgs are saved in a container and can trigger highest decided futuremsg
	All other msgs (not future or decided) are processed normally by an existing instance (if found)
	*/
	if qbft.IsDecidedMsg(contr.Share, signedSSVMessage) {
		return qbft.ValidateDecided(contr.GetConfig(), signedSSVMessage, contr.Share)
	} else if message.Height > contr.Height {
		return validateFutureMsg(contr.GetConfig(), signedSSVMessage, contr.Share.Committee)
	} else {
		if inst := contr.StoredInstances.FindInstance(message.Height); inst != nil {
			return inst.BaseMsgValidation(signedSSVMessage)
		}
		return errors.New("unknown instance")
	}
}

func validatePartialSigMsg(runner ssv.Runner, data []byte) error {
	partialSignatureMessages := &types.PartialSignatureMessages{}
	if err := partialSignatureMessages.Decode(data); err != nil {
		return err
	}

	if partialSignatureMessages.Type == types.PostConsensusPartialSig {
		return runner.GetBaseRunner().ValidatePostConsensusMsg(runner, partialSignatureMessages)
	}
	return runner.GetBaseRunner().ValidatePreConsensusMsg(runner, partialSignatureMessages)
}

func validateFutureMsg(
	config qbft.IConfig,
	msg *types.SignedSSVMessage,
	operators []*types.Operator,
) error {
	if err := msg.Validate(); err != nil {
		return errors.Wrap(err, "invalid decided msg")
	}

	if len(msg.GetOperatorIDs()) != 1 {
		return errors.New("allows 1 signer")
	}

	// verify signature
	if err := types.VerifySignedSSVMessage(msg, operators); err != nil {
		return errors.Wrap(err, "msg signature invalid")
	}

	return nil
}
