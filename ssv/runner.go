package ssv

import (
	"crypto/sha256"
	"encoding/json"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/pkg/errors"
)

// Runner is manages the execution of a duty from start to finish, it can only execute 1 duty at a time.
// Prev duty must finish before the next one can start.
type Runner struct {
	BeaconRoleType types.BeaconRole
	BeaconNetwork  types.BeaconNetwork
	Share          *types.Share
	// State holds all relevant params for a full duty execution (consensus & post consensus)
	State          *State
	QBFTController *qbft.Controller
	storage        Storage
	beacon         BeaconNode
	valCheck       qbft.ProposedValueCheckF
}

func NewDutyRunner(
	beaconRoleType types.BeaconRole,
	beaconNetwork types.BeaconNetwork,
	share *types.Share,
	qbftController *qbft.Controller,
	storage Storage,
	beacon BeaconNode,
	valCheck qbft.ProposedValueCheckF,
) *Runner {
	return &Runner{
		BeaconRoleType: beaconRoleType,
		BeaconNetwork:  beaconNetwork,
		Share:          share,
		QBFTController: qbftController,
		storage:        storage,
		beacon:         beacon,
		valCheck:       valCheck,
	}
}

func (dr *Runner) StartNewDuty(duty *types.Duty) error {
	if err := dr.CanStartNewDuty(duty); err != nil {
		return err
	}
	dr.State = NewDutyExecutionState(dr.Share.Quorum, duty)
	return nil
}

// CanStartNewDuty returns nil if no running instance exists or already decided. Pre- / Post-consensus signature collections do not block a new duty from starting
func (dr *Runner) CanStartNewDuty(duty *types.Duty) error {
	if dr.State == nil {
		return nil
	}

	// check if instance running first as we can't start new duty if it does
	if dr.State.RunningInstance != nil {
		// check consensus decided
		if decided, _ := dr.State.RunningInstance.IsDecided(); !decided {
			return errors.New("consensus on duty is running")
		}
	}
	return nil
}

// HasRunningDuty returns true if a duty is already running (StartNewDuty called and returned nil)
func (dr *Runner) HasRunningDuty() bool {
	if dr.State == nil {
		return false
	}
	return dr.State.Finished != true
}

// GetRoot returns the root used for signing and verification
func (dr *Runner) GetRoot() ([]byte, error) {
	marshaledRoot, err := dr.Encode()
	if err != nil {
		return nil, errors.Wrap(err, "could not encode DutyRunnerState")
	}
	ret := sha256.Sum256(marshaledRoot)
	return ret[:], nil
}

// Encode returns the encoded struct in bytes or error
func (dr *Runner) Encode() ([]byte, error) {
	return json.Marshal(dr)
}

// Decode returns error if decoding failed
func (dr *Runner) Decode(data []byte) error {
	return json.Unmarshal(data, &dr)
}

func (dr *Runner) validatePartialSigMsg(signedMsg *SignedPartialSignatureMessage, slot spec.Slot) error {
	if err := signedMsg.Validate(); err != nil {
		return errors.Wrap(err, "SignedPartialSignatureMessage invalid")
	}

	if err := signedMsg.GetSignature().VerifyByOperators(signedMsg, dr.Share.DomainType, types.PartialSignatureType, dr.Share.Committee); err != nil {
		return errors.Wrap(err, "failed to verify PartialSignature")
	}

	for _, msg := range signedMsg.Message.Messages {
		if slot != msg.Slot {
			return errors.New("wrong slot")
		}

		if err := dr.verifyBeaconPartialSignature(msg); err != nil {
			return errors.Wrap(err, "could not verify Beacon partial Signature")
		}
	}

	return nil
}
