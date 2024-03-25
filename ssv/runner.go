package ssv

import (
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
	ssz "github.com/ferranbt/fastssz"
	"github.com/pkg/errors"
)

type Getters interface {
	GetBaseRunner() *BaseRunner
	GetBeaconNode() BeaconNode
	GetValCheckF() qbft.ProposedValueCheckF
	GetSigner() types.KeyManager
	GetNetwork() Network
}

type Runner interface {
	types.Encoder
	types.Root
	Getters

	// StartNewDuty starts a new duty for the runner, returns error if can't
	StartNewDuty(duty *types.Duty) error
	// HasRunningDuty returns true if it has a running duty
	HasRunningDuty() bool
	// ProcessPreConsensus processes all pre-consensus msgs, returns error if can't process
	ProcessPreConsensus(partialSignatureMessages *types.PartialSignatureMessages) error
	// ProcessConsensus processes all consensus msgs, returns error if can't process
	ProcessConsensus(msg *types.SignedSSVMessage) error
	// ProcessPostConsensus processes all post-consensus msgs, returns error if can't process
	ProcessPostConsensus(partialSignatureMessages *types.PartialSignatureMessages) error

	// expectedPreConsensusRootsAndDomain an INTERNAL function, returns the expected pre-consensus roots to sign
	expectedPreConsensusRootsAndDomain() ([]ssz.HashRoot, spec.DomainType, error)
	// expectedPostConsensusRootsAndDomain an INTERNAL function, returns the expected post-consensus roots to sign
	expectedPostConsensusRootsAndDomain() ([]ssz.HashRoot, spec.DomainType, error)
	// executeDuty an INTERNAL function, executes a duty.
	executeDuty(duty *types.Duty) error
}

type BaseRunner struct {
	State          *State
	Share          *types.Share
	QBFTController *qbft.Controller
	BeaconNetwork  types.BeaconNetwork
	BeaconRoleType types.BeaconRole

	// highestDecidedSlot holds the highest decided duty slot and gets updated after each decided is reached
	highestDecidedSlot spec.Slot
}

func NewBaseRunner(
	state *State,
	share *types.Share,
	controller *qbft.Controller,
	beaconNetwork types.BeaconNetwork,
	beaconRoleType types.BeaconRole,
	highestDecidedSlot spec.Slot,
) *BaseRunner {
	return &BaseRunner{
		State:              state,
		Share:              share,
		QBFTController:     controller,
		BeaconNetwork:      beaconNetwork,
		BeaconRoleType:     beaconRoleType,
		highestDecidedSlot: highestDecidedSlot,
	}
}

// SetHighestDecidedSlot set highestDecidedSlot for base runner
func (b *BaseRunner) SetHighestDecidedSlot(slot spec.Slot) {
	b.highestDecidedSlot = slot
}

// setupForNewDuty is sets the runner for a new duty
func (b *BaseRunner) baseSetupForNewDuty(duty *types.Duty) {
	// start new state
	b.State = NewRunnerState(b.Share.Quorum, duty)
}

// baseStartNewDuty is a base func that all runner implementation can call to start a duty
func (b *BaseRunner) baseStartNewDuty(runner Runner, duty *types.Duty) error {
	if err := b.ShouldProcessDuty(duty); err != nil {
		return errors.Wrap(err, "can't start duty")
	}

	b.baseSetupForNewDuty(duty)

	return runner.executeDuty(duty)
}

// baseStartNewBeaconDuty is a base func that all runner implementation can call to start a non-beacon duty
func (b *BaseRunner) baseStartNewNonBeaconDuty(runner Runner, duty *types.Duty) error {
	if err := b.ShouldProcessNonBeaconDuty(duty); err != nil {
		return errors.Wrap(err, "can't start non-beacon duty")
	}
	b.baseSetupForNewDuty(duty)
	return runner.executeDuty(duty)
}

// basePreConsensusMsgProcessing is a base func that all runner implementation can call for processing a pre-consensus msg
func (b *BaseRunner) basePreConsensusMsgProcessing(runner Runner, partialSignatureMessages *types.PartialSignatureMessages) (bool, [][32]byte, error) {
	if err := b.ValidatePreConsensusMsg(runner, partialSignatureMessages); err != nil {
		return false, nil, errors.Wrap(err, "invalid pre-consensus message")
	}

	hasQuorum, roots, err := b.basePartialSigMsgProcessing(partialSignatureMessages, b.State.PreConsensusContainer)
	return hasQuorum, roots, errors.Wrap(err, "could not process pre-consensus partial signature msg")
}

// baseConsensusMsgProcessing is a base func that all runner implementation can call for processing a consensus msg
func (b *BaseRunner) baseConsensusMsgProcessing(runner Runner, msg *types.SignedSSVMessage) (decided bool, decidedValue *types.ConsensusData, err error) {
	prevDecided := false
	if b.hasRunningDuty() && b.State != nil && b.State.RunningInstance != nil {
		prevDecided, _ = b.State.RunningInstance.IsDecided()
	}

	// TODO: revert `if false` after pre-consensus justification is fixed.
	if false {
		if err := b.processPreConsensusJustification(runner, b.highestDecidedSlot, msg); err != nil {
			return false, nil, errors.Wrap(err, "invalid pre-consensus justification")
		}
	}

	decidedSSVSignedMessage, err := b.QBFTController.ProcessMsg(msg)
	if err != nil {
		return false, nil, err
	}

	// we allow all consensus msgs to be processed, once the process finishes we check if there is an actual running duty
	// do not return error if no running duty
	if !b.hasRunningDuty() {
		return false, nil, nil
	}

	if decideCorrectly, err := b.didDecideCorrectly(prevDecided, decidedSSVSignedMessage); !decideCorrectly {
		return false, nil, err
	}

	// decode consensus data
	decidedMessage := &qbft.Message{}
	if err := decidedMessage.Decode(decidedSSVSignedMessage.SSVMessage.Data); err != nil {
		return true, nil, errors.Wrap(err, "could not decode decided message")
	}
	decidedValue = &types.ConsensusData{}
	if err := decidedValue.Decode(decidedMessage.FullData); err != nil {
		return true, nil, errors.Wrap(err, "failed to parse decided value to ConsensusData")
	}

	// update the highest decided slot
	b.highestDecidedSlot = decidedValue.Duty.Slot

	if err := b.validateDecidedConsensusData(runner, decidedValue); err != nil {
		return true, nil, errors.Wrap(err, "decided ConsensusData invalid")
	}

	runner.GetBaseRunner().State.DecidedValue = decidedValue

	return true, decidedValue, nil
}

// basePostConsensusMsgProcessing is a base func that all runner implementation can call for processing a post-consensus msg
func (b *BaseRunner) basePostConsensusMsgProcessing(runner Runner, partialSignatureMessages *types.PartialSignatureMessages) (bool, [][32]byte, error) {
	if err := b.ValidatePostConsensusMsg(runner, partialSignatureMessages); err != nil {
		return false, nil, errors.Wrap(err, "invalid post-consensus message")
	}

	hasQuorum, roots, err := b.basePartialSigMsgProcessing(partialSignatureMessages, b.State.PostConsensusContainer)
	return hasQuorum, roots, errors.Wrap(err, "could not process post-consensus partial signature msg")
}

// basePartialSigMsgProcessing adds an already validated partial msg to the container, checks for quorum and returns true (and roots) if quorum exists
func (b *BaseRunner) basePartialSigMsgProcessing(
	partialSignatureMessages *types.PartialSignatureMessages,
	container *PartialSigContainer,
) (bool, [][32]byte, error) {
	roots := make([][32]byte, 0)
	anyQuorum := false
	for _, msg := range partialSignatureMessages.Messages {
		prevQuorum := container.HasQuorum(msg.SigningRoot)

		container.AddSignature(msg)

		if prevQuorum {
			continue
		}

		quorum := container.HasQuorum(msg.SigningRoot)
		if quorum {
			roots = append(roots, msg.SigningRoot)
			anyQuorum = true
		}
	}

	return anyQuorum, roots, nil
}

// didDecideCorrectly returns true if the expected consensus instance decided correctly
func (b *BaseRunner) didDecideCorrectly(prevDecided bool, decidedMsg *types.SignedSSVMessage) (bool, error) {
	if decidedMsg == nil {
		return false, nil
	}

	if b.State.RunningInstance == nil {
		return false, errors.New("decided wrong instance")
	}

	// Decode
	message := &qbft.Message{}
	if err := message.Decode(decidedMsg.SSVMessage.Data); err != nil {
		return false, errors.Wrap(err, "could not decode decided Message to check if it decided correctly")
	}

	if message.Height != b.State.RunningInstance.GetHeight() {
		return false, errors.New("decided wrong instance")
	}

	// verify we decided running instance only, if not we do not proceed
	if prevDecided {
		return false, nil
	}

	return true, nil
}

func (b *BaseRunner) decide(runner Runner, input *types.ConsensusData) error {
	byts, err := input.Encode()
	if err != nil {
		return errors.Wrap(err, "could not encode ConsensusData")
	}

	if err := runner.GetValCheckF()(byts); err != nil {
		return errors.Wrap(err, "input data invalid")
	}

	if err := runner.GetBaseRunner().QBFTController.StartNewInstance(
		qbft.Height(input.Duty.Slot),
		byts,
	); err != nil {
		return errors.Wrap(err, "could not start new QBFT instance")
	}
	newInstance := runner.GetBaseRunner().QBFTController.InstanceForHeight(runner.GetBaseRunner().QBFTController.Height)
	if newInstance == nil {
		return errors.New("could not find newly created QBFT instance")
	}

	runner.GetBaseRunner().State.RunningInstance = newInstance
	return nil
}

// hasRunningDuty returns true if a new duty didn't start or an existing duty marked as finished
func (b *BaseRunner) hasRunningDuty() bool {
	if b.State == nil {
		return false
	}
	return !b.State.Finished
}

func (b *BaseRunner) ShouldProcessDuty(duty *types.Duty) error {
	if b.QBFTController.Height >= qbft.Height(duty.Slot) && b.QBFTController.Height != 0 {
		return errors.Errorf("duty for slot %d already passed. Current height is %d", duty.Slot,
			b.QBFTController.Height)
	}
	return nil
}

func (b *BaseRunner) ShouldProcessNonBeaconDuty(duty *types.Duty) error {
	// assume StartingDuty is not nil if state is not nil
	if b.State != nil && b.State.StartingDuty.Slot >= duty.Slot {
		return errors.Errorf("duty for slot %d already passed. Current slot is %d", duty.Slot,
			b.State.StartingDuty.Slot)
	}
	return nil
}

// Encapsulates a PartialSignatureMessages into a SignedSSVMessage and broadcast
func (b *BaseRunner) BroadcastPartialSignatureMessages(msg types.PartialSignatureMessages, signer types.KeyManager, network Network) error {
	// Create SSVMessage
	data, err := msg.Encode()
	if err != nil {
		return errors.Wrap(err, "failed to encode PartialSignatureMessages")
	}
	ssvMessage := &types.SSVMessage{
		MsgType: types.SSVPartialSignatureMsgType,
		MsgID:   types.NewMsgID(b.Share.DomainType, b.Share.ValidatorPubKey, b.BeaconRoleType),
		Data:    data,
	}

	// Sign SSVMessage
	signingData, err := ssvMessage.Encode()
	if err != nil {
		return errors.Wrap(err, "failed to encode SSVMessage for PartialSignatureMessages")
	}
	signature, err := signer.SignNetworkData(signingData, b.Share.NetworkPubkey)
	if err != nil {
		return errors.Wrap(err, "could not sign SSVMessage for PartialSignatureMessages")
	}

	// Create SignedSSVMessage
	msgToBroadcast := &types.SignedSSVMessage{
		OperatorID: []types.OperatorID{b.Share.OperatorID},
		Signature:  [][]byte{signature},
		SSVMessage: ssvMessage,
	}

	// Broadcast
	if err := network.Broadcast(msgToBroadcast); err != nil {
		return errors.Wrap(err, "can't broadcast SignedSSVMessage with PartialSignatureMessages")
	}
	return nil
}
