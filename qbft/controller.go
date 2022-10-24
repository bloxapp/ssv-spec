package qbft

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/bloxapp/ssv-spec/types"
)

// HistoricalInstanceCapacity represents the upper bound of InstanceContainer a controller can process messages for as messages are not
// guaranteed to arrive in a timely fashion, we physically limit how far back the controller will process messages for
const HistoricalInstanceCapacity int = 5

type InstanceContainer [HistoricalInstanceCapacity]*Instance

func (i InstanceContainer) FindInstance(height Height) *Instance {
	for _, inst := range i {
		if inst != nil {
			if inst.GetHeight() == height {
				return inst
			}
		}
	}
	return nil
}

// addNewInstance will add the new instance at index 0, pushing all other stored InstanceContainer one index up (ejecting last one if existing)
func (i *InstanceContainer) addNewInstance(instance *Instance) {
	for idx := HistoricalInstanceCapacity - 1; idx > 0; idx-- {
		i[idx] = i[idx-1]
	}
	i[0] = instance
}

// Controller is a QBFT coordinator responsible for starting and following the entire life cycle of multiple QBFT InstanceContainer
type Controller struct {
	Identifier types.MessageID
	Height     Height // incremental Height for InstanceContainer
	// StoredInstances stores the last HistoricalInstanceCapacity in an array for message processing purposes.
	StoredInstances InstanceContainer
	Domain          types.DomainType
	Share           *types.Share
	signer          types.SSVSigner
	valueCheck      ProposedValueCheckF
	storage         Storage
	network         Network
	proposerF       ProposerF
}

func NewController(
	identifier types.MessageID,
	share *types.Share,
	domain types.DomainType,
	signer types.SSVSigner,
	valueCheck ProposedValueCheckF,
	storage Storage,
	network Network,
	proposerF ProposerF,
) *Controller {
	return &Controller{
		Identifier:      identifier,
		Height:          -1, // as we bump the height when starting the first instance
		Domain:          domain,
		Share:           share,
		StoredInstances: InstanceContainer{},
		signer:          signer,
		valueCheck:      valueCheck,
		storage:         storage,
		network:         network,
		proposerF:       proposerF,
	}
}

// StartNewInstance will start a new QBFT instance, if can't will return error
func (c *Controller) StartNewInstance(value []byte) error {
	if err := c.canStartInstance(c.Height+1, value); err != nil {
		return errors.Wrap(err, "can't start new QBFT instance")
	}

	c.bumpHeight()
	newInstance := c.addAndStoreNewInstance()
	newInstance.Start(value, c.Height)

	return nil
}

// ProcessMsg processes a new msg, returns true if Decided, non nil byte slice if Decided (Decided value) and error
// Decided returns just once per instance as true, following messages (for example additional commit msgs) will not return Decided true
func (c *Controller) ProcessMsg(msg *types.Message) (bool, []byte, error) {
	msgID := msg.GetID()
	if !msgID.Compare(c.Identifier) {
		return false, nil, errors.New("message doesn't belong to Identifier")
	}

	var height Height
	signedMsg := &SignedMessage{}
	signedMsgH := &SignedMessageHeader{}

	switch msgID.GetMsgType() {
	case types.ConsensusProposeMsgType, types.ConsensusRoundChangeMsgType:
		if err := signedMsg.Decode(msg.GetData()); err != nil {
			return false, nil, errors.Wrap(err, "could not decode consensus msg from network msg")
		}
		height = signedMsg.Message.Height

	case types.ConsensusPrepareMsgType, types.ConsensusCommitMsgType:
		if err := signedMsgH.Decode(msg.GetData()); err != nil {
			return false, nil, errors.Wrap(err, "could not decode consensus msg header from network msg")
		}
		height = signedMsgH.Message.Height
	default:
		return false, nil, errors.New("message type not supported")
	}

	//signedMsg := &SignedMessage{}
	//if err := signedMsg.Decode(msg.GetData()); err != nil {
	//	return false, nil, errors.Wrap(err, "could not decode consensus Message from network Message")
	//}

	inst := c.InstanceForHeight(height)
	if inst == nil {
		return false, nil, errors.New("instance not found")
	}

	prevDecided, _ := inst.IsDecided()
	decided, decidedValue, aggregatedCommit, err := inst.ProcessMsg(msgID, signedMsg, signedMsgH)
	if err != nil {
		return false, nil, errors.Wrap(err, "could not process msg")
	}

	// if previously Decided we do not return Decided true again
	if prevDecided {
		return false, nil, err
	}

	// save the highest Decided
	if !decided {
		return false, nil, nil
	}

	// nolint
	if err := c.saveAndBroadcastDecided(aggregatedCommit); err != nil {
		// TODO - we do not return error, should log?
	}
	return decided, decidedValue, nil
}

func (c *Controller) InstanceForHeight(height Height) *Instance {
	return c.StoredInstances.FindInstance(height)
}

func (c *Controller) bumpHeight() {
	c.Height++
}

// GetIdentifier returns QBFT Identifier, used to identify messages
func (c *Controller) GetIdentifier() types.MessageID {
	return c.Identifier
}

// addAndStoreNewInstance returns creates a new QBFT instance, stores it in an array and returns it
func (c *Controller) addAndStoreNewInstance() *Instance {
	i := NewInstance(c.GenerateConfig(), c.Share, c.Identifier, c.Height)
	c.StoredInstances.addNewInstance(i)
	return i
}

func (c *Controller) canStartInstance(height Height, value []byte) error {
	if height > FirstHeight {
		// check prev instance if prev instance is not the first instance
		inst := c.StoredInstances.FindInstance(height - 1)
		if inst == nil {
			return errors.New("could not find previous instance")
		}
		if decided, _ := inst.IsDecided(); !decided {
			return errors.New("previous instance hasn't Decided")
		}
	}

	// check value
	if err := c.valueCheck(value); err != nil {
		return errors.Wrap(err, "value invalid")
	}

	return nil
}

// Encode implementation
func (c *Controller) Encode() ([]byte, error) {
	return json.Marshal(c)
}

// Decode implementation
func (c *Controller) Decode(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return errors.Wrap(err, "could not decode controller")
	}

	config := c.GenerateConfig()
	for _, i := range c.StoredInstances {
		if i != nil {
			i.config = config
		}
	}
	return nil
}

func (c *Controller) saveAndBroadcastDecided(aggregatedCommit *SignedMessage) error {
	if err := c.storage.SaveHighestDecided(c.GetIdentifier(), aggregatedCommit); err != nil {
		return errors.Wrap(err, "could not save decided")
	}

	// Broadcast Decided msg
	decidedEncoded, err := aggregatedCommit.Encode()
	if err != nil {
		return errors.Wrap(err, "could not encode decided message")
	}

	msgID := types.PopulateMsgType(c.Identifier, types.DecidedMsgType)

	broadcastMsg := &types.Message{
		ID:   msgID,
		Data: decidedEncoded,
	}

	if err := c.network.BroadcastDecided(broadcastMsg); err != nil {
		// We do not return error here, just Log broadcasting error.
		return errors.Wrap(err, "could not broadcast decided")
	}
	return nil
}

func (c *Controller) GenerateConfig() IConfig {
	return &Config{
		Signer:      c.signer,
		SigningPK:   c.Share.ValidatorPubKey,
		Domain:      c.Domain,
		ValueCheckF: c.valueCheck,
		Storage:     c.storage,
		Network:     c.network,
		ProposerF:   c.proposerF,
	}
}
