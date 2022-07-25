package dkg

import (
	"bytes"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	dkgtypes "github.com/bloxapp/ssv-spec/dkg/types"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
)

// Runner manages the execution of a DKG, start to finish.
type Runner struct {
	Operator *dkgtypes.Operator
	// InitMsg holds the init method which started this runner
	InitMsg *dkgtypes.Init
	// Identifier unique for DKG session
	Identifier dkgtypes.RequestID
	// DepositDataRoot is the signing root for the deposit data
	DepositDataRoot []byte
	// PartialSignatures holds partial sigs on deposit data
	PartialSignatures map[types.OperatorID]*dkgtypes.ParsedPartialSigMessage
	// OutputMsgs holds all output messages received
	OutputMsgs map[types.OperatorID]*dkgtypes.ParsedSignedDepositDataMessage

	KeygenSubProtocol dkgtypes.Protocol
	keygenOutput      *dkgtypes.LocalKeyShare
	signOutput        *dkgtypes.SignedDepositDataMsgBody
	Config            *dkgtypes.Config
}

func (r *Runner) Start() error {
	data, err := r.InitMsg.Encode()
	if err != nil {
		return err
	}
	outgoing, err := r.KeygenSubProtocol.ProcessMsg(&dkgtypes.Message{
		Header: &dkgtypes.MessageHeader{
			SessionId: r.Identifier[:],
			MsgType:   int32(dkgtypes.InitMsgType),
			Sender:    0,
			Receiver:  0,
		},
		Data: data,
	})
	if err != nil {
		return err
	}
	for _, message := range outgoing {
		err = r.signAndBroadcast(&message)
		if err != nil {
			return err
		}
	}
	return nil
}

// ProcessMsg processes a DKG signed message and returns true and signed output if finished
func (r *Runner) ProcessMsg(msg *dkgtypes.Message) (bool, map[types.OperatorID]*dkgtypes.ParsedSignedDepositDataMessage, error) {
	// TODO - validate message

	switch msg.Header.MsgType {
	case int32(dkgtypes.ProtocolMsgType):
		outgoing, err := r.KeygenSubProtocol.ProcessMsg(msg)
		if err != nil {
			return false, nil, errors.Wrap(err, "failed to process dkg msg")
		}

		err = r.broadcastMessages(outgoing, dkgtypes.ProtocolMsgType)
		if err != nil {
			return false, nil, err
		}

		output, err := r.KeygenSubProtocol.Output()
		if err != nil {
			return false, nil, err
		}

		if output != nil {
			r.keygenOutput = &dkgtypes.LocalKeyShare{}
			if err = r.keygenOutput.Decode(output); err != nil {
				return false, nil, err
			}
			if err = r.startSigning(); err != nil {
				return false, nil, err
			}
		}
	case int32(dkgtypes.DepositDataMsgType):

		if msg.Header.MsgType != int32(dkgtypes.DepositDataMsgType) {
			return false, nil, errors.New("invalid message type")
		}

		pMsg := &dkgtypes.ParsedPartialSigMessage{}
		err := pMsg.FromBase(msg)
		if err != nil {
			return false, nil, err
		}
		err = r.handlePartialSigMessage(pMsg)
		if err != nil {
			return false, nil, err
		}

		if err != nil {
			return false, nil, errors.Wrap(err, "failed to partial sig msg")
		}

		if r.signOutput != nil {
			sig, err := r.Config.Signer.SignDKGOutput(r.signOutput, r.Operator.ETHAddress)
			if err != nil {
				return false, nil, err
			}
			r.signOutput.OperatorSignature = sig
			out := &dkgtypes.ParsedSignedDepositDataMessage{
				Header: &dkgtypes.MessageHeader{
					SessionId: r.Identifier[:],
					MsgType:   int32(dkgtypes.OutputMsgType),
					Sender:    uint64(r.Operator.OperatorID),
					Receiver:  0,
				},
				Body:      r.signOutput,
				Signature: nil,
			}
			base, err := out.ToBase()
			if err != nil {
				return false, nil, err
			}
			r.OutputMsgs[r.Operator.OperatorID] = out
			r.broadcastMessages([]dkgtypes.Message{*base}, dkgtypes.OutputMsgType)
			return false, nil, nil
		}
		//if hasOutput(outgoing, dkgtypes.PartialOutputMsgType) {
		//	return true, nil, err
		//}

		/*
				// TODO: Do we need to aggregate the signed outputs.
			case DepositDataMsgType:
				depSig := &PartialSigMsgBody{}
				if err := depSig.Decode(msg.Message.Data); err != nil {
					return false, nil, errors.Wrap(err, "could not decode PartialSigMsgBody")
				}

				if err := r.validateDepositDataSig(depSig); err != nil {
					return false, nil, errors.Wrap(err, "PartialSigMsgBody invalid")
				}

				r.DepositDataSignatures[msg.Signer] = depSig
				if len(r.DepositDataSignatures) == int(r.InitMsg.Threshold) {
					// reconstruct deposit data sig
					depositSig, err := r.reconstructDepositDataSignature()
					if err != nil {
						return false, nil, errors.Wrap(err, "could not reconstruct deposit data sig")
					}

					// encrypt Operator's share
					encryptedShare, err := r.Config.Signer.Encrypt(r.Operator.EncryptionPubKey, r.KeyGenOutput.Share.Serialize())
					if err != nil {
						return false, nil, errors.Wrap(err, "could not encrypt share")
					}

					ret, err := r.generateSignedOutput(&Output{
						RequestID:            r.Identifier,
						EncryptedShare:       encryptedShare,
						SharePubKey:          r.KeyGenOutput.Share.GetPublicKey().Serialize(),
						ValidatorPubKey:      r.KeyGenOutput.ValidatorPK,
						DepositDataSignature: depositSig,
					})
					if err != nil {
						return false, nil, errors.Wrap(err, "could not generate dkg SignedOutput")
					}

					if err := r.signAndBroadcastMsg(ret, OutputMsgType); err != nil {
						return false, nil, errors.Wrap(err, "could not broadcast SignedOutput")
					}
					return false, nil, nil
				} */
	case int32(dkgtypes.OutputMsgType):
		output := &dkgtypes.ParsedSignedDepositDataMessage{}
		if err := output.FromBase(msg); err != nil {
			return false, nil, errors.Wrap(err, "could not decode SignedOutput")
		}
		if output.Header.RequestID() != r.Identifier {
			return false, nil, errors.New("request id mismatch")
		}
		r.OutputMsgs[types.OperatorID(msg.Header.Sender)] = output
		if len(r.OutputMsgs) == len(r.InitMsg.OperatorIDs) {
			for _, message := range r.OutputMsgs {
				if message.Header.RequestID() != r.Identifier {
					return true, r.OutputMsgs, errors.New("one of more messages have mismatched request id")
				}
			}
			return true, r.OutputMsgs, nil
		}
		return false, nil, nil
	default:
		return false, nil, errors.New("msg type invalid")
	}

	return false, nil, nil
}

func (r *Runner) startSigning() error {

	pSig, err := r.partialSign()
	if err != nil {
		return err
	}
	partialSigMsg := dkgtypes.ParsedPartialSigMessage{
		Header: &dkgtypes.MessageHeader{
			SessionId: r.Identifier[:],
			MsgType:   int32(dkgtypes.DepositDataMsgType),
			Sender:    uint64(r.Operator.OperatorID),
			Receiver:  0,
		},
		Body: pSig,
	}
	r.PartialSignatures[r.Operator.OperatorID] = &partialSigMsg
	base, err := partialSigMsg.ToBase()
	if err != nil {
		return err
	}
	err = r.broadcastMessages([]dkgtypes.Message{*base}, dkgtypes.DepositDataMsgType)
	return nil
}

func (r *Runner) handlePartialSigMessage(msg *dkgtypes.ParsedPartialSigMessage) error {

	if found := r.PartialSignatures[types.OperatorID(msg.Header.Sender)]; found == nil {
		r.PartialSignatures[types.OperatorID(msg.Header.Sender)] = msg
	} else if bytes.Compare(found.Body.Signature, msg.Body.Signature) != 0 {
		return errors.New("inconsistent partial signature received")
	}

	if len(r.PartialSignatures) > int(r.InitMsg.Threshold) {
		sigBytes := map[types.OperatorID][]byte{}
		for id, pSig := range r.PartialSignatures {
			if err := r.validateDepositDataSig(pSig.Body); err != nil {
				return errors.Wrap(err, "PartialSigMsgBody invalid")
			}
			sigBytes[id] = pSig.Body.Signature
		}

		sig, err := types.ReconstructSignatures(sigBytes)
		if err != nil {
			return err
		}

		// encrypt Operator's share
		encryptedShare, err := r.Config.Signer.Encrypt(r.Operator.EncryptionPubKey, r.keygenOutput.SecretShare)
		if err != nil {
			return errors.Wrap(err, "could not encrypt share")
		}

		r.signOutput = &dkgtypes.SignedDepositDataMsgBody{
			RequestID:             r.Identifier[:],
			OperatorID:            uint64(r.Operator.OperatorID),
			EncryptedShare:        encryptedShare,
			Committee:             r.keygenOutput.Committee,
			Threshold:             r.InitMsg.Threshold,
			ValidatorPublicKey:    r.keygenOutput.PublicKey,
			WithdrawalCredentials: r.InitMsg.WithdrawalCredentials,
			DepositDataSignature:  sig.Serialize(),
		}
		return nil
	}
	return nil
}

func (r *Runner) partialSign() (*dkgtypes.PartialSigMsgBody, error) {
	share := bls.SecretKey{}
	err := share.Deserialize(r.keygenOutput.SecretShare)
	if err != nil {
		return nil, err
	}

	fork := spec.Version{}
	copy(fork[:], r.InitMsg.Fork)
	root, depData, err := types.GenerateETHDepositData(
		r.keygenOutput.PublicKey,
		r.InitMsg.WithdrawalCredentials,
		fork,
		types.DomainDeposit,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not generate deposit data")
	}
	r.DepositDataRoot = make([]byte, len(root))
	copy(r.DepositDataRoot[:], root)

	//root := make([]byte, len(depData.DepositDataRoot))
	//copy(root, depData.DepositDataRoot[:])
	rawSig := share.SignByte(root[:])
	sigBytes := rawSig.Serialize()
	var sig spec.BLSSignature
	copy(sig[:], sigBytes)

	copy(depData.DepositData.Signature[:], sigBytes)

	return &dkgtypes.PartialSigMsgBody{
		Signer:    uint64(r.Operator.OperatorID),
		Root:      root,
		Signature: sig[:],
	}, nil
}

func (r *Runner) validateDepositDataSig(msg *dkgtypes.PartialSigMsgBody) error {
	if !bytes.Equal(r.DepositDataRoot[:], msg.Root) {
		return errors.New("deposit data roots not equal")
	}

	index := -1
	for i, d := range r.InitMsg.OperatorIDs {
		if d == msg.Signer {
			index = i
		}
	}

	if index == -1 {
		return errors.New("signer not part of committee")
	}

	// find operator and verify msg
	sharePkBytes := r.keygenOutput.SharePublicKeys[index]
	sharePk := &bls.PublicKey{} // TODO: cache this PubKey
	if err := sharePk.Deserialize(sharePkBytes); err != nil {
		return errors.Wrap(err, "could not deserialize public key")
	}

	sig := &bls.Sign{}
	if err := sig.Deserialize(msg.Signature); err != nil {
		return errors.Wrap(err, "could not deserialize partial sig")
	}

	root := make([]byte, 32)
	copy(root, r.DepositDataRoot[:])
	if !sig.VerifyByte(sharePk, root) {
		return errors.New("partial deposit data sig invalid")
	}

	return nil
}

//func (r *Runner) generateSignedOutput(o *dkgtypes.Output) (*dkgtypes.SignedOutput, error) {
//	sig, err := r.Config.Signer.SignDKGOutput(o, r.Operator.ETHAddress)
//	if err != nil {
//		return nil, errors.Wrap(err, "could not sign output")
//	}
//
//	return &dkgtypes.SignedOutput{
//		Data:      o,
//		Signer:    r.Operator.OperatorID,
//		Signature: sig,
//	}, nil
//}

func (r *Runner) broadcastMessages(msgs []dkgtypes.Message, msgType dkgtypes.MsgType) error {
	for _, message := range msgs {
		if message.Header.MsgType == int32(msgType) {
			err := r.signAndBroadcast(&message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Runner) signAndBroadcast(msg *dkgtypes.Message) error {
	sig, err := r.Config.Signer.SignDKGOutput(msg, r.Operator.ETHAddress)
	if err != nil {
		return err
	}
	err = msg.SetSignature(sig)
	if err != nil {
		return err
	}
	r.Config.Network.Broadcast(msg)
	return nil
}

func hasOutput(msgs []dkgtypes.Message, msgType dkgtypes.MsgType) bool {
	return msgs != nil && len(msgs) > 0 && msgs[len(msgs)-1].Header.MsgType == int32(msgType)
}

func (r *Runner) validateSignedOutput(msg *dkgtypes.ParsedSignedDepositDataMessage) error {
	if msg == nil {
		return errors.New("msg is nil")
	}
	if !r.signOutput.SameDepositData(msg.Body) {
		return errors.New("deposit data doesn't match")
	}
	// TODO: Verify OperatorSignature
	return nil
}
