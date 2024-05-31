package types

import "github.com/attestantio/go-eth2-client/spec/phase0"

// SharedValidator holds all info about the validator share
type SharedValidator struct {
	ValidatorIndex      phase0.ValidatorIndex
	ValidatorPubKey     ValidatorPK `ssz-size:"48"`
	OwnValidatorShare   ValidatorShare
	CommitteeID         CommitteeID       `ssz-size:"32"`
	Committee           []*ValidatorShare `ssz-max:"13"`
	Quorum              uint64
	PartialQuorum       uint64
	DomainType          DomainType `ssz-size:"4"`
	FeeRecipientAddress [20]byte   `ssz-size:"20"`
	Graffiti            []byte     `ssz-size:"32"`
}

// HasQuorum returns true if at least 2f+1 items are present (cnt is the number of items). It assumes nothing about those items, not their type or structure
// https://github.com/ConsenSys/qbft-formal-spec-and-verification/blob/main/dafny/spec/L1/node_auxiliary_functions.dfy#L259
func (share *SharedValidator) HasQuorum(cnt int) bool {
	return uint64(cnt) >= share.Quorum
}
func (share *SharedValidator) HasPartialQuorum(cnt int) bool {
	return uint64(cnt) >= share.PartialQuorum
}

func (share *SharedValidator) Encode() ([]byte, error) {
	return share.MarshalSSZ()
}

func (share *SharedValidator) Decode(data []byte) error {
	return share.UnmarshalSSZ(data)
}
