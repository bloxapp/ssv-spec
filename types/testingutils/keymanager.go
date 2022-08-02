package testingutils

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/hex"

	"github.com/attestantio/go-eth2-client/spec/altair"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/go-bitfield"

	"github.com/bloxapp/ssv-spec/types"
)

type testingKeyManager struct {
	keys           map[string]*bls.SecretKey
	ecdsaKeys      map[string]*ecdsa.PrivateKey
	encryptionKeys map[string]*rsa.PrivateKey
	domain         types.DomainType
}

func NewTestingKeyManager() *testingKeyManager {
	ret := &testingKeyManager{
		keys:   map[string]*bls.SecretKey{},
		domain: types.GetDefaultDomain(),
	}

	ret.AddShare(Testing4SharesSet().ValidatorSK)
	for _, s := range Testing4SharesSet().Shares {
		ret.AddShare(s)
	}

	ret.AddShare(Testing7SharesSet().ValidatorSK)
	for _, s := range Testing7SharesSet().Shares {
		ret.AddShare(s)
	}

	ret.AddShare(Testing10SharesSet().ValidatorSK)
	for _, s := range Testing10SharesSet().Shares {
		ret.AddShare(s)
	}

	ret.AddShare(Testing13SharesSet().ValidatorSK)
	for _, s := range Testing13SharesSet().Shares {
		ret.AddShare(s)
	}
	return ret
}

// SignAttestation signs the given attestation
func (km *testingKeyManager) SignAttestation(data *spec.AttestationData, duty *types.Duty, pk []byte) (*spec.Attestation, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingAttestationRoot)
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		aggregationBitfield := bitfield.NewBitlist(duty.CommitteeLength)
		aggregationBitfield.SetBitAt(duty.ValidatorCommitteeIndex, true)

		return &spec.Attestation{
			AggregationBits: aggregationBitfield,
			Data:            data,
			Signature:       blsSig,
		}, TestingAttestationRoot, nil
	}
	return nil, nil, errors.New("pk not found")
}

// IsAttestationSlashable returns error if attestation is slashable
func (km *testingKeyManager) IsAttestationSlashable(data *spec.AttestationData) error {
	return nil
}

func (km *testingKeyManager) SignRoot(data types.Root, sigType types.SignatureType, pk []byte) (types.Signature, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		computedRoot, err := types.ComputeSigningRoot(data, types.ComputeSignatureDomain(km.domain, sigType))
		if err != nil {
			return nil, errors.Wrap(err, "could not sign root")
		}

		return k.SignByte(computedRoot).Serialize(), nil
	}
	return nil, errors.New("pk not found")
}

// SignRandaoReveal signs randao
func (km *testingKeyManager) SignRandaoReveal(epoch spec.Epoch, pk []byte) (types.Signature, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingRandaoRoot)
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		return sig.Serialize(), TestingRandaoRoot, nil
	}
	return nil, nil, errors.New("pk not found")
}

// IsBeaconBlockSlashable returns true if the given block is slashable
func (km *testingKeyManager) IsBeaconBlockSlashable(block *altair.BeaconBlock) error {
	return nil
}

// SignBeaconBlock signs the given beacon block
func (km *testingKeyManager) SignBeaconBlock(data *altair.BeaconBlock, duty *types.Duty, pk []byte) (*altair.SignedBeaconBlock, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingBeaconBlockRoot)
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		return &altair.SignedBeaconBlock{
			Message:   data,
			Signature: blsSig,
		}, TestingBeaconBlockRoot, nil
	}
	return nil, nil, errors.New("pk not found")
}

// SignSlotWithSelectionProof signs slot for aggregator selection proof
func (km *testingKeyManager) SignSlotWithSelectionProof(slot spec.Slot, pk []byte) (types.Signature, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingSelectionProofRoot)
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		return sig.Serialize(), TestingSelectionProofRoot, nil
	}
	return nil, nil, errors.New("pk not found")
}

// SignAggregateAndProof returns a signed aggregate and proof msg
func (km *testingKeyManager) SignAggregateAndProof(msg *spec.AggregateAndProof, duty *types.Duty, pk []byte) (*spec.SignedAggregateAndProof, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingSignedAggregateAndProofRoot)
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		return &spec.SignedAggregateAndProof{
			Message:   msg,
			Signature: blsSig,
		}, TestingSignedAggregateAndProofRoot, nil
	}
	return nil, nil, errors.New("pk not found")
}

// SignSyncCommitteeBlockRoot returns a signed sync committee msg
func (km *testingKeyManager) SignSyncCommitteeBlockRoot(slot spec.Slot, root spec.Root, validatorIndex spec.ValidatorIndex, pk []byte) (*altair.SyncCommitteeMessage, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingSyncCommitteeBlockRoot[:])
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		return &altair.SyncCommitteeMessage{
			Slot:            slot,
			BeaconBlockRoot: TestingSyncCommitteeBlockRoot,
			Signature:       blsSig,
		}, TestingSyncCommitteeBlockRoot[:], nil
	}
	return nil, nil, errors.New("pk not found")
}

func (km *testingKeyManager) SignContributionProof(slot spec.Slot, index uint64, pk []byte) (types.Signature, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingContributionProofRoots[index][:])
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		return sig.Serialize(), TestingContributionProofRoots[index][:], nil
	}
	return nil, nil, errors.New("pk not found")
}

func (km *testingKeyManager) SignContribution(contribution *altair.ContributionAndProof, pk []byte) (*altair.SignedContributionAndProof, []byte, error) {
	if k, found := km.keys[hex.EncodeToString(pk)]; found {
		sig := k.SignByte(TestingContributionRoots[contribution.Contribution.SubcommitteeIndex])
		blsSig := spec.BLSSignature{}
		copy(blsSig[:], sig.Serialize())

		return &altair.SignedContributionAndProof{
			Message:   contribution,
			Signature: blsSig,
		}, TestingContributionRoots[contribution.Contribution.SubcommitteeIndex], nil
	}
	return nil, nil, errors.New("pk not found")
}

// Decrypt given a rsa pubkey and a PKCS1v15 cipher text byte array, returns the decrypted data
func (km *testingKeyManager) Decrypt(pk *rsa.PublicKey, cipher []byte) ([]byte, error) {
	panic("implement")
}

// Encrypt given a rsa pubkey and data returns an PKCS1v15 e
func (km *testingKeyManager) Encrypt(pk *rsa.PublicKey, data []byte) ([]byte, error) {
	panic("implement")
}

// SignDKGOutput signs output according to the SIP https://docs.google.com/document/d/1TRVUHjFyxINWW2H9FYLNL2pQoLy6gmvaI62KL_4cREQ/edit
func (km *testingKeyManager) SignDKGOutput(output types.Root, address common.Address) (types.Signature, error) {
	panic("implemet")
}

func (km *testingKeyManager) SignETHDepositRoot(root []byte, address common.Address) (types.Signature, error) {
	panic("implemet")
}

func (km *testingKeyManager) AddShare(shareKey *bls.SecretKey) error {
	km.keys[hex.EncodeToString(shareKey.GetPublicKey().Serialize())] = shareKey
	return nil
}

func (km *testingKeyManager) RemoveShare(pubKey string) error {
	delete(km.keys, pubKey)
	return nil
}
