package testingutils

import (
	"encoding/hex"
	"github.com/bloxapp/ssv-spec/dkg"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
)

type testingStorage struct {
	storage     map[string]*qbft.SignedMessage
	instances   map[string]*qbft.Instance
	operators   map[types.OperatorID]*dkg.Operator
	keygenoupts map[string]*dkg.KeyGenOutput
}

func NewTestingStorage() *testingStorage {
	ret := &testingStorage{
		storage:     make(map[string]*qbft.SignedMessage),
		instances:   make(map[string]*qbft.Instance),
		operators:   make(map[types.OperatorID]*dkg.Operator),
		keygenoupts: make(map[string]*dkg.KeyGenOutput),
	}

	for i, s := range Testing13SharesSet().DKGOperators {
		ret.operators[i] = &dkg.Operator{
			OperatorID:       i,
			ETHAddress:       s.ETHAddress,
			EncryptionPubKey: &s.EncryptionKey.PublicKey,
		}
	}

	return ret
}

// SaveHighestDecided saves the Decided value as highest for a validator PK and role
func (s *testingStorage) SaveHighestDecided(signedMsg *qbft.SignedMessage) error {
	s.storage[hex.EncodeToString(signedMsg.Message.Identifier)] = signedMsg
	return nil
}

// GetHighestDecided returns highest decided if found, nil if didn't
func (s *testingStorage) GetHighestDecided(identifier []byte) (*qbft.SignedMessage, error) {
	return s.storage[hex.EncodeToString(identifier)], nil
}

// SaveHighestInstance check if new instance is first or higher than last known the highest instance. if so, save to storage
func (s *testingStorage) SaveHighestInstance(instance *qbft.Instance) error {
	highestInstance, err := s.GetHighestInstance(instance.State.ID)
	if err != nil {
		return err
	}
	if highestInstance == nil || instance.GetHeight() > highestInstance.GetHeight() {
		s.instances[hex.EncodeToString(instance.State.ID)] = instance
	}
	return nil
}

func (s *testingStorage) GetHighestInstance(identifier []byte) (*qbft.Instance, error) {
	return s.instances[hex.EncodeToString(identifier)], nil
}

// GetDKGOperator returns true and operator object if found by operator ID
func (s *testingStorage) GetDKGOperator(operatorID types.OperatorID) (bool, *dkg.Operator, error) {
	if ret, found := s.operators[operatorID]; found {
		return true, ret, nil
	}
	return false, nil, nil
}

func (s *testingStorage) SaveKeyGenOutput(output *dkg.KeyGenOutput) error {
	s.keygenoupts[hex.EncodeToString(output.ValidatorPK)] = output
	return nil
}

func (s *testingStorage) GetKeyGenOutput(pk types.ValidatorPK) (*dkg.KeyGenOutput, error) {
	return s.keygenoupts[hex.EncodeToString(pk)], nil
}
