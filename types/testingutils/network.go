package testingutils

import (
	"sync"

	"github.com/bloxapp/ssv-spec/dkg"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
)

type TestingNetwork struct {
	broadcastMsgsLock *sync.Mutex
	BroadcastedMsgs   []*types.SSVMessage

	DKGOutputs                map[types.OperatorID]*dkg.SignedOutput
	BlameOutput               *dkg.BlameOutput
	SyncHighestDecidedCnt     int
	SyncHighestChangeRoundCnt int
	DecidedByRange            [2]qbft.Height
}

func NewTestingNetwork() *TestingNetwork {
	return &TestingNetwork{
		broadcastMsgsLock: new(sync.Mutex),
		BroadcastedMsgs:   make([]*types.SSVMessage, 0),
		DKGOutputs:        make(map[types.OperatorID]*dkg.SignedOutput, 0),
	}
}

func (net *TestingNetwork) Broadcast(message *types.SSVMessage) error {
	net.broadcastMsgsLock.Lock()
	defer net.broadcastMsgsLock.Unlock()

	net.BroadcastedMsgs = append(net.BroadcastedMsgs, message)
	return nil
}

// StreamDKGOutput will stream to any subscriber the result of the DKG
func (net *TestingNetwork) StreamDKGOutput(output map[types.OperatorID]*dkg.SignedOutput) error {
	for id, signedOutput := range output {
		net.DKGOutputs[id] = signedOutput
	}

	return nil
}

func (net *TestingNetwork) StreamDKGBlame(blame *dkg.BlameOutput) error {
	//TODO implement me
	net.BlameOutput = blame
	return nil
}

func (net *TestingNetwork) SyncHighestDecided(identifier types.MessageID) error {
	net.SyncHighestDecidedCnt++
	return nil
}

//func (net *TestingNetwork) SyncHighestDecided() error {
//	return nil
//}

// SyncDecidedByRange will sync decided messages from-to (including them)
func (net *TestingNetwork) SyncDecidedByRange(identifier types.MessageID, from, to qbft.Height) {
	net.DecidedByRange = [2]qbft.Height{from, to}
}

// BroadcastDKGMessage will broadcast a msg to the dkg network
func (net *TestingNetwork) BroadcastDKGMessage(msg *dkg.SignedMessage) error {
	data, err := msg.Encode()
	if err != nil {
		return err
	}

	net.broadcastMsgsLock.Lock()
	defer net.broadcastMsgsLock.Unlock()

	net.BroadcastedMsgs = append(net.BroadcastedMsgs, &types.SSVMessage{
		MsgType: types.DKGMsgType,
		MsgID:   types.MessageID{}, // TODO: what should we use for the MsgID?
		Data:    data,
	})
	return nil
}

func (net *TestingNetwork) GetBroadcastMessages() []*types.SSVMessage {
	net.broadcastMsgsLock.Lock()
	defer net.broadcastMsgsLock.Unlock()

	return net.BroadcastedMsgs
}
