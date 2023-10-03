package types

import (
	"crypto/sha256"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// NetworkID are intended to separate different SSV networks. A network can have many forks in it.
type NetworkID [1]byte

var (
	MainnetNetworkID = NetworkID{0x0}
	PrimusNetworkID  = NetworkID{0x1}
	ShifuNetworkID   = NetworkID{0x2}
	JatoNetworkID    = NetworkID{0x3}
	JatoV2NetworkID  = NetworkID{0x4}
)

// DomainType is a unique identifier for signatures, 2 identical pieces of data signed with different domains will result in different sigs
type DomainType [4]byte

// DomainTypes represent specific forks for specific chains, messages are signed with the domain type making 2 messages from different domains incompatible
var (
	GenesisMainnet = DomainType{0x0, 0x0, MainnetNetworkID.Byte(), 0x0}
	PrimusTestnet  = DomainType{0x0, 0x0, PrimusNetworkID.Byte(), 0x0}
	ShifuTestnet   = DomainType{0x0, 0x0, ShifuNetworkID.Byte(), 0x0}
	ShifuV2Testnet = DomainType{0x0, 0x0, ShifuNetworkID.Byte(), 0x1}
	JatoTestnet    = DomainType{0x0, 0x0, JatoNetworkID.Byte(), 0x1}
	JatoV2Testnet  = DomainType{0x0, 0x0, JatoV2NetworkID.Byte(), 0x1}
)

// ForkData is a simple structure holding fork information for a specific chain (and its fork)
type ForkData struct {
	// Epoch in which the fork happened
	Epoch phase0.Epoch
	// Domain for the new fork
	Domain DomainType `ssz-size:"4"`
}

func (n NetworkID) Byte() byte {
	return n[0]
}

func (n NetworkID) GetForksData() []*ForkData {
	switch n {
	case MainnetNetworkID:
		return mainnetForks()
	case PrimusNetworkID:
		return []*ForkData{{Epoch: 0, Domain: PrimusTestnet}}
	case JatoNetworkID:
		return []*ForkData{{Epoch: 0, Domain: JatoTestnet}}
	case JatoV2NetworkID:
		return []*ForkData{{Epoch: 0, Domain: JatoV2Testnet}}
	default:
		return []*ForkData{}
	}
}

func (n NetworkID) DefaultFork() *ForkData {
	return n.GetForksData()[0]
}

// GetCurrentFork returns the ForkData with highest Epoch smaller or equal to "epoch"
func (n NetworkID) ForkAtEpoch(epoch phase0.Epoch) (*ForkData, error) {
	// Get list of forks
	forks := n.GetForksData()

	// If empty, raise error
	if len(forks) == 0 {
		return nil, errors.New("Fork list by GetForksData is empty. Unknown Network")
	}

	var current_fork *ForkData
	for _, fork := range forks {
		if fork.Epoch <= epoch {
			current_fork = fork
		}
	}
	return current_fork, nil
}

func (d DomainType) GetNetworkID() NetworkID {
	return NetworkID{d[2]}
}

func (f ForkData) GetRoot() ([]byte, error) {
	byts, err := f.MarshalSSZ()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal ForkData")
	}
	ret := sha256.Sum256(byts)
	return ret[:], nil
}

func (f *ForkData) Encode() ([]byte, error) {
	return f.MarshalSSZ()
}

func (f *ForkData) Decode(data []byte) error {
	return f.UnmarshalSSZ(data)
}

// mainnetForks returns all forks for the mainnet chain
func mainnetForks() []*ForkData {
	return []*ForkData{
		{
			Epoch:  0,
			Domain: GenesisMainnet,
		},
	}
}
