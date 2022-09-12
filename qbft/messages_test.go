package qbft_test

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSignedMessage_MatchedSigners(t *testing.T) {
	t.Run("all unique", func(t *testing.T) {
		m1 := &qbft.SignedMessage{Signers: []types.OperatorID{1}}
		require.False(t, m1.MatchedSigners([]types.OperatorID{2}))
	})

	t.Run("unique multi signer", func(t *testing.T) {
		m1 := &qbft.SignedMessage{Signers: []types.OperatorID{1, 2}}
		require.False(t, m1.MatchedSigners([]types.OperatorID{3, 4}))
	})

	t.Run("common multi signer", func(t *testing.T) {
		m1 := &qbft.SignedMessage{Signers: []types.OperatorID{1}}
		require.False(t, m1.MatchedSigners([]types.OperatorID{1, 2}))
	})

	t.Run("common multi signer", func(t *testing.T) {
		m1 := &qbft.SignedMessage{Signers: []types.OperatorID{1, 2}}
		require.False(t, m1.MatchedSigners([]types.OperatorID{1}))
	})

	t.Run("common multi signer", func(t *testing.T) {
		m1 := &qbft.SignedMessage{Signers: []types.OperatorID{1, 2}}
		require.False(t, m1.MatchedSigners([]types.OperatorID{1, 3}))
	})
}

//func TestSignedMessage_Aggregate(t *testing.T) {
//	t.Run("roots don't match", func(t *testing.T) {
//		m1 := testingSignedMsg.DeepCopy()
//		m2 := testingSignedMsg.DeepCopy()
//
//		m1.Signers = []types.OperatorID{11}
//		m1.Message.Height = 1222 // changing root
//
//		require.EqualError(t, m1.Aggregate(m2), "can't aggregate, roots not equal")
//	})
//
//	t.Run("common signers", func(t *testing.T) {
//		m1 := testingSignedMsg.DeepCopy()
//		m2 := testingSignedMsg.DeepCopy()
//
//		require.EqualError(t, m1.Aggregate(m2), "can't aggregate 2 signed messages with mutual signers")
//	})
//}

func TestMessage_Validate(t *testing.T) {
	t.Run("valid proposal", func(t *testing.T) {
		m := &qbft.Message{
			MsgType:    qbft.ProposalMsgType,
			Identifier: []byte{1, 2, 3, 4},
			Data:       []byte{1, 2, 3, 4},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("valid prepare", func(t *testing.T) {
		m := &qbft.Message{
			MsgType:    qbft.PrepareMsgType,
			Identifier: []byte{1, 2, 3, 4},
			Data:       []byte{1, 2, 3, 4},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("valid commit", func(t *testing.T) {
		m := &qbft.Message{
			MsgType:    qbft.CommitMsgType,
			Identifier: []byte{1, 2, 3, 4},
			Data:       []byte{1, 2, 3, 4},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("valid round change", func(t *testing.T) {
		m := &qbft.Message{
			MsgType:    qbft.RoundChangeMsgType,
			Identifier: []byte{1, 2, 3, 4},
			Data:       []byte{1, 2, 3, 4},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("invalid msg type", func(t *testing.T) {
		m := &qbft.Message{
			MsgType:    6,
			Identifier: []byte{1, 2, 3, 4},
			Data:       []byte{1, 2, 3, 4},
		}
		require.EqualError(t, m.Validate(), "message type is invalid")
	})
}

func TestSignedMessage_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := &qbft.SignedMessage{
			Signature: make([]byte, 96),
			Signers:   []types.OperatorID{1},
			Message: &qbft.Message{
				MsgType:    qbft.ProposalMsgType,
				Identifier: []byte{1, 2, 3, 4},
				Data:       []byte{1, 2, 3, 4},
			},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("invalid signature", func(t *testing.T) {
		m := &qbft.SignedMessage{
			Signature: []byte{1, 2, 3, 4},
			Signers:   []types.OperatorID{1},
			Message: &qbft.Message{
				MsgType:    qbft.ProposalMsgType,
				Identifier: []byte{1, 2, 3, 4},
				Data:       []byte{1, 2, 3, 4},
			},
		}
		require.EqualError(t, m.Validate(), "message signature is invalid")
	})
	t.Run("invalid signers", func(t *testing.T) {
		m := &qbft.SignedMessage{
			Signature: make([]byte, 96),
			Signers:   []types.OperatorID{},
			Message: &qbft.Message{
				MsgType:    qbft.ProposalMsgType,
				Identifier: []byte{1, 2, 3, 4},
				Data:       []byte{1, 2, 3, 4},
			},
		}
		require.EqualError(t, m.Validate(), "message signers is empty")
	})
	t.Run("invalid msg", func(t *testing.T) {
		m := &qbft.SignedMessage{
			Signature: make([]byte, 96),
			Signers:   []types.OperatorID{1},
			Message: &qbft.Message{
				MsgType:    100,
				Identifier: []byte{1, 2, 3, 4},
				Data:       []byte{1, 2, 3, 4},
			},
		}
		require.EqualError(t, m.Validate(), "message type is invalid")
	})
}
func TestProposalData_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := &qbft.ProposalData{
			Data: []byte{1, 2, 3, 4},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("invalid data", func(t *testing.T) {
		m := &qbft.ProposalData{
			Data: []byte{},
		}
		require.EqualError(t, m.Validate(), "ProposalData data is invalid")
	})
	t.Run("invalid data", func(t *testing.T) {
		m := &qbft.ProposalData{}
		require.EqualError(t, m.Validate(), "ProposalData data is invalid")
	})
}

func TestPrepareData_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := &qbft.PrepareData{
			Data: []byte{1, 2, 3, 4},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("invalid data", func(t *testing.T) {
		m := &qbft.PrepareData{
			Data: []byte{},
		}
		require.EqualError(t, m.Validate(), "PrepareData data is invalid")
	})
	t.Run("invalid data", func(t *testing.T) {
		m := &qbft.PrepareData{}
		require.EqualError(t, m.Validate(), "PrepareData data is invalid")
	})
}

func TestCommitData_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := &qbft.CommitData{
			Data: []byte{1, 2, 3, 4},
		}
		require.NoError(t, m.Validate())
	})
	t.Run("invalid data", func(t *testing.T) {
		m := &qbft.CommitData{
			Data: []byte{},
		}
		require.EqualError(t, m.Validate(), "CommitData data is invalid")
	})
	t.Run("invalid data", func(t *testing.T) {
		m := &qbft.CommitData{}
		require.EqualError(t, m.Validate(), "CommitData data is invalid")
	})
}
