package tests

import (
	"encoding/hex"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	CreateProposal    = "createProposal"
	CreatePrepare     = "CreatePrepare"
	CreateCommit      = "CreateCommit"
	CreateRoundChange = "CreateRoundChange"
)

type CreateMsgSpecTest struct {
	Name                                             string
	Value                                            []byte
	ValueRoot                                        [32]byte
	Round                                            qbft.Round
	RoundChangeJustifications, PrepareJustifications []*qbft.SignedMessage
	CreateType                                       string
	ExpectedRoot                                     string
	ExpectedError                                    string
}

func (test *CreateMsgSpecTest) Run(t *testing.T) {
	var msg *qbft.SignedMessage
	var r []byte
	var lastErr error
	switch test.CreateType {
	case CreateProposal:
		msg, _ = test.createProposal()
	case CreatePrepare:
		msg, _ = test.createPrepare()
	case CreateCommit:
		msg, _ = test.createCommit()
	case CreateRoundChange:
		msg, _ = test.createRoundChange()
	default:
		t.Fail()
	}

	r, err := msg.GetRoot()
	if err != nil {
		lastErr = err
	}

	if len(test.ExpectedError) != 0 {
		require.EqualError(t, lastErr, test.ExpectedError)
	} else {
		require.NoError(t, lastErr)
	}

	require.EqualValues(t, test.ExpectedRoot, hex.EncodeToString(r))
}

func (test *CreateMsgSpecTest) createCommit() (*qbft.SignedMessage, error) {
	ks := testingutils.Testing4SharesSet()
	msgId := types.NewBaseMsgID([]byte{1, 2, 3, 4}, types.BNRoleAttester)
	state := &qbft.State{
		Share: testingutils.TestingShare(ks),
		ID:    msgId,
	}
	config := testingutils.TestingConfig(ks)

	return qbft.CreateCommit(state, config, test.ValueRoot)
}

func (test *CreateMsgSpecTest) createPrepare() (*qbft.SignedMessage, error) {
	ks := testingutils.Testing4SharesSet()
	msgId := types.NewBaseMsgID([]byte{1, 2, 3, 4}, types.BNRoleAttester)
	state := &qbft.State{
		Share: testingutils.TestingShare(ks),
		ID:    msgId,
	}
	config := testingutils.TestingConfig(ks)

	return qbft.CreatePrepare(state, config, test.Round, test.ValueRoot)
}

func (test *CreateMsgSpecTest) createProposal() (*qbft.SignedMessage, error) {
	ks := testingutils.Testing4SharesSet()
	msgId := types.NewBaseMsgID([]byte{1, 2, 3, 4}, types.BNRoleAttester)
	state := &qbft.State{
		Share: testingutils.TestingShare(ks),
		ID:    msgId,
	}
	config := testingutils.TestingConfig(ks)

	return qbft.CreateProposal(state, config, test.Value, test.RoundChangeJustifications, test.PrepareJustifications)
}

func (test *CreateMsgSpecTest) createRoundChange() (*qbft.SignedMessage, error) {
	ks := testingutils.Testing4SharesSet()
	msgId := types.NewBaseMsgID([]byte{1, 2, 3, 4}, types.BNRoleAttester)
	state := &qbft.State{
		Share: testingutils.TestingShare(ks),
		ID:    msgId,
	}
	config := testingutils.TestingConfig(ks)

	if len(test.PrepareJustifications) > 0 {
		state.LastPreparedRound = test.PrepareJustifications[0].Message.Round
		state.LastPreparedValue = &qbft.Data{
			Root:   test.ValueRoot,
			Source: test.Value,
		}

		for _, msg := range test.PrepareJustifications {
			_, err := state.PrepareContainer.AddFirstMsgForSignerAndRound(msg)
			if err != nil {
				return nil, errors.Wrap(err, "could not add first message for signer")
			}
		}
	}

	return qbft.CreateRoundChange(state, config, 1)
}

func (test *CreateMsgSpecTest) TestName() string {
	return "qbft create message " + test.Name
}
