package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
)

type MsgType [4]byte

var (
	// ConsensusProposeMsgType QBFT propose consensus message
	ConsensusProposeMsgType = MsgType{0x1, 0x0, 0x0, 0x0}
	// ConsensusPrepareMsgType QBFT prepare consensus message
	ConsensusPrepareMsgType = MsgType{0x1, 0x1, 0x0, 0x0}
	// ConsensusCommitMsgType QBFT commit consensus message
	ConsensusCommitMsgType = MsgType{0x1, 0x2, 0x0, 0x0}
	// ConsensusRoundChangeMsgType QBFT round change consensus message
	ConsensusRoundChangeMsgType = MsgType{0x1, 0x3, 0x0, 0x0}

	// DecidedMsgType are all QBFT decided messages
	DecidedMsgType = MsgType{0x2, 0x0, 0x0, 0x0}

	// PartialRandaoSignatureMsgType randao partial signature for proposer
	PartialRandaoSignatureMsgType = MsgType{0x3, 0x0, 0x0, 0x0}
	// PartialSelectionProofSignatureMsgType selection proof partial signature for aggregator
	PartialSelectionProofSignatureMsgType = MsgType{0x3, 0x1, 0x0, 0x0}
	// PartialContributionProofSignatureMsgType contribution proof for sync committee aggregator
	PartialContributionProofSignatureMsgType = MsgType{0x3, 0x2, 0x0, 0x0}
	// PartialPostConsensusSignatureMsgType post consensus partial signatures for all duties
	PartialPostConsensusSignatureMsgType = MsgType{0x3, 0x3, 0x0, 0x0}

	// DKGInitMsgType sent when DKG instance is started by requester
	DKGInitMsgType = MsgType{0x4, 0x0, 0x0, 0x0}
	// DKGProtocolMsgType contains all key generation protocol msgs
	DKGProtocolMsgType = MsgType{0x4, 0x1, 0x0, 0x0}
	// DKGDepositDataMsgType post DKG deposit data signatures
	DKGDepositDataMsgType = MsgType{0x4, 0x2, 0x0, 0x0}
	// DKGOutputMsgType final output msg used by requester to make deposits and register validator with SSV
	DKGOutputMsgType = MsgType{0x4, 0x3, 0x0, 0x0}

	// UnknownMsgType can't be identified
	UnknownMsgType = MsgType{0x0, 0x0, 0x0, 0x0}
)

// ValidatorPK is an eth2 validator public key
type ValidatorPK []byte

const (
	pubKeySize       = 48
	pubKeyStartPos   = 0
	roleTypeSize     = 4
	roleTypeStartPos = pubKeyStartPos + pubKeySize
	paddingSize      = 4
	paddingStartPos  = roleTypeStartPos + roleTypeSize
	msgTypeSize      = 4
	msgTypeStartPos  = paddingStartPos + paddingSize
)

type Validate interface {
	// Validate returns error if msg validation doesn't pass.
	// Msg validation checks the msg, it's variables for validity.
	Validate() error
}

// MessageIDBelongs returns true if message ID belongs to validator
func (vid ValidatorPK) MessageIDBelongs(msgID MessageID) bool {
	toMatch := msgID.GetPubKey()
	return bytes.Equal(vid, toMatch)
}

// MessageIDOld is used to identify and route messages to the right validator and Runner
type MessageIDOld [52]byte

type MessageID [60]byte

func (msgIDOld MessageIDOld) GetPubKey() []byte {
	return msgIDOld[pubKeyStartPos : pubKeyStartPos+pubKeySize]
}

func (msgIDOld MessageIDOld) GetRoleType() BeaconRole {
	roleByts := msgIDOld[roleTypeStartPos : roleTypeStartPos+roleTypeSize]
	return BeaconRole(binary.LittleEndian.Uint32(roleByts))
}

func NewMsgID(pk []byte, role BeaconRole) MessageIDOld {
	roleByts := make([]byte, 4)
	binary.LittleEndian.PutUint32(roleByts, uint32(role))

	ret := MessageIDOld{}
	copy(ret[pubKeyStartPos:pubKeyStartPos+pubKeySize], pk)
	copy(ret[roleTypeStartPos:roleTypeStartPos+roleTypeSize], roleByts)

	return ret
}

func (msgID MessageID) GetPubKey() []byte {
	return msgID[pubKeyStartPos:roleTypeStartPos]
}

func (msgID MessageID) GetRoleType() BeaconRole {
	roleByts := msgID[roleTypeStartPos:paddingStartPos]
	return BeaconRole(binary.LittleEndian.Uint32(roleByts))
}

func (msgID MessageID) GetMsgType() MsgType {
	var ret MsgType
	copy(ret[:], msgID[msgTypeStartPos:])
	return ret
}

func (msgID MessageID) Compare(identifier MessageID) bool {
	return bytes.Equal(msgID[pubKeyStartPos:paddingStartPos], identifier[pubKeyStartPos:paddingStartPos])
}

func NewBaseMsgID(pk []byte, role BeaconRole) MessageID {
	roleByts := make([]byte, roleTypeSize)
	binary.LittleEndian.PutUint32(roleByts, uint32(role))

	ret := MessageID{}
	copy(ret[pubKeyStartPos:roleTypeStartPos], pk)
	copy(ret[roleTypeStartPos:paddingStartPos], roleByts)
	return ret
}

func PopulateMsgType(msgID MessageID, msgType MsgType) MessageID {
	copy(msgID[msgTypeStartPos:], msgType[:])
	return msgID
}

func (msgIDOld MessageIDOld) String() string {
	return hex.EncodeToString(msgIDOld[:])
}

//type MsgType uint64
//
//const (
//	// SSVConsensusMsgType are all QBFT consensus related messages
//	SSVConsensusMsgType MsgType = iota
//	// SSVPartialSignatureMsgType are all partial signatures msgs over beacon chain specific signatures
//	SSVPartialSignatureMsgType
//	// DKGMsgType represent all DKG related messages
//	DKGMsgType
//)

type Root interface {
	// GetRoot returns the root used for signing and verification
	GetRoot() ([]byte, error)
}

// MessageSignature includes all functions relevant for a signed message (QBFT message, post consensus msg, etc)
type MessageSignature interface {
	Root
	GetSignature() Signature
	GetSigners() []OperatorID
	// MatchedSigners returns true if the provided signer ids are equal to GetSignerIds() without order significance
	MatchedSigners(ids []OperatorID) bool
	// Aggregate will aggregate the signed message if possible (unique signers, same digest, valid)
	Aggregate(signedMsg MessageSignature) error
}

// SSVMessage is the main message passed within the SSV network, it can contain different types of messages (QBTF, Sync, etc.)
type SSVMessage struct {
	MsgType MsgType
	MsgID   MessageIDOld
	Data    []byte
}

func (msg *SSVMessage) GetType() MsgType {
	return msg.MsgType
}

// GetID returns a unique msg ID that is used to identify to which validator should the message be sent for processing
func (msg *SSVMessage) GetID() MessageIDOld {
	return msg.MsgID
}

// GetData returns message Data as byte slice
func (msg *SSVMessage) GetData() []byte {
	return msg.Data
}

// Encode returns a msg encoded bytes or error
func (msg *SSVMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// Decode returns error if decoding failed
func (msg *SSVMessage) Decode(data []byte) error {
	return json.Unmarshal(data, &msg)
}

// Message is the main message passed within the SSV network, it can contain different types of messages (QBTF, Sync, etc.)
type Message struct {
	ID   MessageID `ssz-size:"60"`
	Data []byte    `ssz-max:"2048"`
}

// GetID returns a unique msg ID that is used to identify to which validator should the message be sent for processing
func (msg *Message) GetID() MessageID {
	return msg.ID
}

// GetData returns message Data as byte slice
func (msg *Message) GetData() []byte {
	return msg.Data
}

// Encode returns a msg encoded bytes or error
func (msg *Message) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// Decode returns error if decoding failed
func (msg *Message) Decode(data []byte) error {
	return json.Unmarshal(data, &msg)
}
