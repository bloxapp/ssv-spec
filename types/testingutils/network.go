package testingutils

import (
	dkgtypes "github.com/bloxapp/ssv-spec/dkg/types"
	"github.com/bloxapp/ssv-spec/types"
)

type TestingNetwork struct {
	BroadcastedMsgs []*types.SSVMessage
}

func NewTestingNetwork() *TestingNetwork {
	return &TestingNetwork{
		BroadcastedMsgs: make([]*types.SSVMessage, 0),
	}
}

func (net *TestingNetwork) Broadcast(message types.Encoder) error {
	net.BroadcastedMsgs = append(net.BroadcastedMsgs, message.(*types.SSVMessage))
	return nil
}

func (net *TestingNetwork) BroadcastDecided(msg types.Encoder) error {
	return nil
}

// StreamDKGOutput will stream to any subscriber the result of the DKG
func (net *TestingNetwork) StreamDKGOutput(output map[types.OperatorID]*dkgtypes.SignedOutput) error {
	return nil
}

// BroadcastDKGMessage will broadcast a msg to the dkg network
func (net *TestingNetwork) BroadcastDKGMessage(msg *dkgtypes.SignedMessage) error {
	return nil
}
