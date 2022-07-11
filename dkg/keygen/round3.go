package keygen

import (
	"bytes"
	"errors"
	"github.com/bloxapp/ssv-spec/dkg/dlog"
	"github.com/bloxapp/ssv-spec/dkg/vss"
	"github.com/herumi/bls-eth-go-binary/bls"
)

func (k *Keygen)calcSkI() *bls.SecretKey {
	skI := new(bls.SecretKey)
	skI.Deserialize(k.ownShare.Serialize())
	ind := int(k.PartyI - 1)
	for i, r3Msg := range k.Round3Msgs {
		if i == ind {
			continue
		}
		temp := new(bls.SecretKey)
		temp.Deserialize(r3Msg.Body.Round3.Shares[ind])
		skI.Add(temp)
	}
	return skI
}

func (k *Keygen) r3Proceed() error {
	if k.Round != 3 {
		return ErrInvalidRound
	}

	k.skI = k.calcSkI()
	knowledge := dlog.Knowledge{
		SecretKey:    k.skI,
		RandomNumber: k.DlogR,
	}
	proof := knowledge.Prove()
	msg := &Message{
		Sender: k.PartyI,
		Body: MessageBody{
			Round4: &Round4Msg{
				Commitment:        proof.Commitment.Serialize(),
				PubKey:            proof.PubKey.Serialize(),
				ChallengeResponse: proof.Response.Serialize(),
			},
		},
	}
	k.pushOutgoing(msg)
	k.Round = 4
	return nil
}

func (k *Keygen) r3CanProceed() error {
	var (
		ErrInvalidCoefficientCommitment = errors.New("invalid coefficient commitments")
		ErrInvalidSharesCount           = errors.New("invalid number of shares")
		ErrInvalidShare                 = errors.New("invalid share")
	)

	if k.Round != 3 {
		return ErrInvalidRound
	}
	for i, r3Msg := range k.Round3Msgs {
		if i == int(k.PartyI) -1 {
			continue
		}
		r2Msg := k.Round2Msgs[i]
		if r2Msg == nil || r2Msg.Body.Round2 == nil || r2Msg.Body.Round2.YI == nil || r3Msg == nil || r3Msg.Body.Round3 == nil {
			return errors.New("expected message not found")
		}
		if len(r3Msg.Body.Round3.Commitments) != len(k.Coefficients) {
			return ErrInvalidCoefficientCommitment
		}
		if len(r3Msg.Body.Round3.Shares) != int(k.PartyCount) {
			return ErrInvalidSharesCount
		}
		if bytes.Compare(r2Msg.Body.Round2.YI, r3Msg.Body.Round3.Commitments[0]) != 0 {
			return ErrInvalidCoefficientCommitment
		}
		shareBytes := r3Msg.Body.Round3.Shares[int(k.PartyI)-1]
		share := &vss.Share{
			Threshold: len(k.Coefficients) - 1,
			ID:        new(bls.Fr),
			Share:     new(bls.Fr),
		}
		share.ID.SetInt64(int64(k.PartyI))
		share.Share.Deserialize(shareBytes)
		if r3Msg.Sender == k.PartyI {
			share.Share = k.ownShare
		}
		commitments := make([]*bls.PublicKey, len(k.Coefficients))
		for j, commBytes := range r3Msg.Body.Round3.Commitments {
			// TODO: Improve conversion of multiple times
			commitments[j] = new(bls.PublicKey)
			commitments[j].Deserialize(commBytes)
		}
		if !share.Verify(len(k.Coefficients)-1, commitments) {
			return ErrInvalidShare
		}
	}

	return nil
}
