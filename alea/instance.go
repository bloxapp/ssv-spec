package alea

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pkg/errors"

	"github.com/MatheusFranco99/ssv-spec-AleaBFT/types"
)

type ProposedValueCheckF func(data []byte) error
type ProposerF func(state *State, round Round) types.OperatorID

// Instance is a single QBFT instance that starts with a Start call (including a value).
// Every new msg the ProcessMsg function needs to be called
type Instance struct {
	State  *State
	config IConfig

	processMsgF *types.ThreadSafeF
	startOnce   sync.Once
	StartValue  []byte
}

func NewInstance(
	config IConfig,
	share *types.Share,
	identifier []byte,
	height Height,
) *Instance {
	return &Instance{
		State: &State{
			Share:                share,
			ID:                   identifier,
			Round:                FirstRound,
			Height:               height,
			LastPreparedRound:    NoRound,
			ProposeContainer:     NewMsgContainer(),
			VCBCContainer:     	  NewMsgContainer(),
			ABAContainer:      	  NewMsgContainer(),
			FillGapContainer: 	  NewMsgContainer(),
			FillerContainer: 	  NewMsgContainer(),
			Priority: 			  Priority(0),
			AleaDefaultRound:	  FirstRound,
			queues:				  make(map[types.OperatorID]*VCBCQueue),
			S:					  NewVCBCQueue(),
			ACRound:			  FirstRound,
			StopAgreement:		  false,
			ABAState:			  NewABAState(FirstRound),
			FillerMsgReceived:	  make(chan bool),
			VCBCr:			   	  make(map[types.OperatorID]map[Priority]uint64),
			VCBCW:				  make(map[types.OperatorID]map[Priority][]*SignedMessage),
			VCBCm:				  make(map[types.OperatorID]map[Priority][]*ProposalData),
			VCBCu:				  make(map[types.OperatorID]map[Priority]types.Signature),

		},
		config:      config,
		processMsgF: types.NewThreadSafeF(),
	}
}

// Start is an interface implementation
func (i *Instance) Start(value []byte, height Height) {
	i.startOnce.Do(func() {
		i.StartValue = value
		i.State.Round = FirstRound
		i.State.Height = height

		fmt.Println("Starting instance")

		// -> Init
		// state variables are initiated on constructor NewInstance (namely, queues and S)

		// -> Broadcast
		// The broadcast part runs as an instance receives proposal or vcbc messages
		// 		proposal message: is the message that a client sends to the node
		// 		vcbc message: is the broadcast a node does after receiving a batch size number of proposals

		// The agreement component consists of an infinite loop and we shall call it with another Thread
		go i.StartAgreementComponent()
	})
}

func (i *Instance) Deliver(proposals []*ProposalData) int {
	// FIX ME : to be implemented
	return 1
}

func (i *Instance) Broadcast(msg *SignedMessage) error {
	byts, err := msg.Encode()
	if err != nil {
		return errors.Wrap(err, "could not encode message")
	}

	msgID := types.MessageID{}
	copy(msgID[:], msg.Message.Identifier)

	msgToBroadcast := &types.SSVMessage{
		MsgType: types.SSVConsensusMsgType,
		MsgID:   msgID,
		Data:    byts,
	}
	return i.config.GetNetwork().Broadcast(msgToBroadcast)
}

// ProcessMsg processes a new QBFT msg, returns non nil error on msg processing error
func (i *Instance) ProcessMsg(msg *SignedMessage) (decided bool, decidedValue []byte, aggregatedCommit *SignedMessage, err error) {
	if err := i.BaseMsgValidation(msg); err != nil {
		return false, nil, nil, errors.Wrap(err, "invalid signed message")
	}

	res := i.processMsgF.Run(func() interface{} {
		switch msg.Message.MsgType {
		case ProposalMsgType:
			return i.uponProposal(msg, i.State.ProposeContainer)
		case VCBCMsgType:
			return i.uponVCBC(msg, i.State.VCBCContainer)
		case FillGapMsgType:
			return nil
		case FillerMsgType:
			return nil
		case ABAInitMsgType:
			return nil
		case ABAAuxMsgType:
			return nil
		case ABAConfMsgType:
			return nil
		case ABAFinishMsgType:
			return nil
		// case PrepareMsgType:
		// 	return i.uponPrepare(msg, i.State.PrepareContainer, i.State.CommitContainer)
		// case CommitMsgType:
		// 	decided, decidedValue, aggregatedCommit, err = i.UponCommit(msg, i.State.CommitContainer)
		// 	if decided {
		// 		i.State.Decided = decided
		// 		i.State.DecidedValue = decidedValue
		// 	}
		// 	return err
		// case RoundChangeMsgType:
		// 	return i.uponRoundChange(i.StartValue, msg, i.State.RoundChangeContainer, i.config.GetValueCheckF())
		default:
			return errors.New("signed message type not supported")
		}
	})
	if res != nil {
		return false, nil, nil, res.(error)
	}
	return i.State.Decided, i.State.DecidedValue, aggregatedCommit, nil
}

func (i *Instance) BaseMsgValidation(msg *SignedMessage) error {
	if err := msg.Validate(); err != nil {
		return errors.Wrap(err, "invalid signed message")
	}

	if msg.Message.Round < i.State.Round {
		return errors.New("past round")
	}

	switch msg.Message.MsgType {
	case ProposalMsgType:
		return isValidProposal(
			i.State,
			i.config,
			msg,
			i.config.GetValueCheckF(),
			i.State.Share.Committee,
		)
	case VCBCMsgType:
		return isValidVCBC(
			i.State,
			i.config,
			msg,
			i.config.GetValueCheckF(),
			i.State.Share.Committee,
		)
	// case PrepareMsgType:
	// 	proposedMsg := i.State.ProposalAcceptedForCurrentRound
	// 	if proposedMsg == nil {
	// 		return errors.New("did not receive proposal for this round")
	// 	}
	// 	acceptedProposalData, err := proposedMsg.Message.GetCommitData()
	// 	if err != nil {
	// 		return errors.Wrap(err, "could not get accepted proposal data")
	// 	}
	// 	return validSignedPrepareForHeightRoundAndValue(
	// 		i.config,
	// 		msg,
	// 		i.State.Height,
	// 		i.State.Round,
	// 		acceptedProposalData.Data,
	// 		i.State.Share.Committee,
	// 	)
	// case CommitMsgType:
	// 	proposedMsg := i.State.ProposalAcceptedForCurrentRound
	// 	if proposedMsg == nil {
	// 		return errors.New("did not receive proposal for this round")
	// 	}
	// 	return validateCommit(
	// 		i.config,
	// 		msg,
	// 		i.State.Height,
	// 		i.State.Round,
	// 		i.State.ProposalAcceptedForCurrentRound,
	// 		i.State.Share.Committee,
	// 	)
	// case RoundChangeMsgType:
	// 	return validRoundChange(i.State, i.config, msg, i.State.Height, msg.Message.Round)
	default:
		return errors.New("signed message type not supported")
	}
}

// IsDecided interface implementation
func (i *Instance) IsDecided() (bool, []byte) {
	if state := i.State; state != nil {
		return state.Decided, state.DecidedValue
	}
	return false, nil
}

// GetConfig returns the instance config
func (i *Instance) GetConfig() IConfig {
	return i.config
}

// GetHeight interface implementation
func (i *Instance) GetHeight() Height {
	return i.State.Height
}

// GetRoot returns the state's deterministic root
func (i *Instance) GetRoot() ([]byte, error) {
	return i.State.GetRoot()
}

// Encode implementation
func (i *Instance) Encode() ([]byte, error) {
	return json.Marshal(i)
}

// Decode implementation
func (i *Instance) Decode(data []byte) error {
	return json.Unmarshal(data, &i)
}
