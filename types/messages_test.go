package types

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

var testingPubKey = []byte{0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3}

func TestMessageIDForValidatorPKAndRole(t *testing.T) {
	require.EqualValues(t, MessageID{0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0}, NewMsgID(testingPubKey, BNRoleAttester))
	require.EqualValues(t, MessageID{0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3, 0x2, 0x0, 0x0, 0x0}, NewMsgID(testingPubKey, BNRoleProposer))
	require.EqualValues(t, MessageID{0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3, 0x1, 0x0, 0x0, 0x0}, NewMsgID(testingPubKey, BNRoleAggregator))
}

func TestMessageID_GetRoleType(t *testing.T) {
	t.Run("attester", func(t *testing.T) {
		msgID := NewMsgID(testingPubKey, BNRoleAttester)
		require.EqualValues(t, BNRoleAttester, msgID.GetRoleType())
	})

	t.Run("proposer", func(t *testing.T) {
		msgID := NewMsgID(testingPubKey, BNRoleProposer)
		require.EqualValues(t, BNRoleProposer, msgID.GetRoleType())
	})

	t.Run("long pk", func(t *testing.T) {
		msgID := NewMsgID(testingPubKey, BNRoleProposer)
		require.EqualValues(t, BNRoleProposer, msgID.GetRoleType())
	})
}

func TestShare_Marshaling(t *testing.T) {
	expected, _ := hex.DecodeString("7b2264617461223a223232343135313439343434323431336433643232222c226964223a223031303230333034222c2274797065223a223330227d")

	t.Run("encode", func(t *testing.T) {
		msg := &SSVMessage{
			MsgID:   MessageID{1, 2, 3, 4},
			MsgType: SSVConsensusMsgType,
			Data:    []byte{1, 2, 3, 4},
		}

		byts, err := msg.Encode()
		require.NoError(t, err)
		require.EqualValues(t, expected, byts)
	})

	t.Run("decode", func(t *testing.T) {
		msg := &SSVMessage{}
		require.NoError(t, msg.Decode(expected))
		require.EqualValues(t, MessageID{1, 2, 3, 4}, msg.MsgID)
		require.EqualValues(t, SSVConsensusMsgType, msg.MsgType)
		require.EqualValues(t, []byte{1, 2, 3, 4}, msg.Data)
	})
}
