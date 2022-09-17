package testingutils

import (
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/ssv"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/herumi/bls-eth-go-binary/bls"
)

var AttesterMsgID = func() []byte {
	ret := types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleAttester)
	return ret[:]
}()

var ProposerMsgID = func() []byte {
	ret := types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleProposer)
	return ret[:]
}()
var AggregatorMsgID = func() []byte {
	ret := types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleAggregator)
	return ret[:]
}()
var SyncCommitteeMsgID = func() []byte {
	ret := types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleSyncCommittee)
	return ret[:]
}()
var SyncCommitteeContributionMsgID = func() []byte {
	ret := types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleSyncCommitteeContribution)
	return ret[:]
}()

var TestAttesterConsensusData = &types.ConsensusData{
	Duty:            TestingAttesterDuty,
	AttestationData: TestingAttestationData,
}
var TestAttesterConsensusDataByts, _ = TestAttesterConsensusData.MarshalSSZ()

var TestAggregatorConsensusData = &types.ConsensusData{
	Duty:              TestingAggregatorDuty,
	AggregateAndProof: TestingAggregateAndProof,
}
var TestAggregatorConsensusDataByts, _ = TestAggregatorConsensusData.MarshalSSZ()

var TestProposerConsensusData = &types.ConsensusData{
	Duty:      TestingProposerDuty,
	BlockData: TestingBeaconBlock,
}
var TestProposerConsensusDataByts, _ = TestProposerConsensusData.MarshalSSZ()

var TestSyncCommitteeConsensusData = &types.ConsensusData{
	Duty:                   TestingSyncCommitteeDuty,
	SyncCommitteeBlockRoot: TestingSyncCommitteeBlockRoot,
}
var TestSyncCommitteeConsensusDataByts, _ = TestSyncCommitteeConsensusData.MarshalSSZ()

var TestSyncCommitteeContributionConsensusData = &types.ConsensusData{
	Duty: TestingSyncCommitteeContributionDuty,
	SyncCommitteeContribution: map[phase0.BLSSignature]*altair.SyncCommitteeContribution{
		TestingContributionProofsSigned[0]: TestingSyncCommitteeContributions[0],
		TestingContributionProofsSigned[1]: TestingSyncCommitteeContributions[1],
		TestingContributionProofsSigned[2]: TestingSyncCommitteeContributions[2],
	},
}
var TestSyncCommitteeContributionConsensusDataByts, _ = TestSyncCommitteeContributionConsensusData.MarshalSSZ()

var TestConsensusUnkownDutyTypeData = &types.ConsensusData{
	Duty:            TestingUnknownDutyType,
	AttestationData: TestingAttestationData,
}
var TestConsensusUnkownDutyTypeDataByts, _ = TestConsensusUnkownDutyTypeData.MarshalSSZ()

var TestConsensusWrongDutyPKData = &types.ConsensusData{
	Duty:            TestingWrongDutyPK,
	AttestationData: TestingAttestationData,
}
var TestConsensusWrongDutyPKDataByts, _ = TestConsensusWrongDutyPKData.MarshalSSZ()

var SSVMsgAttester = func(qbftMsg *qbft.SignedMessage, partialSigMsg *ssv.SignedPartialSignature, msgType types.MsgType) *types.Message {
	return ssvMsg(qbftMsg, partialSigMsg, types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleAttester), msgType)
}

var SSVMsgWrongID = func(qbftMsg *qbft.SignedMessage, partialSigMsg *ssv.SignedPartialSignature, msgType types.MsgType) *types.Message {
	return ssvMsg(qbftMsg, partialSigMsg, types.NewBaseMsgID(TestingWrongValidatorPubKey[:], types.BNRoleAttester), msgType)
}

var SSVMsgProposer = func(qbftMsg *qbft.SignedMessage, partialSigMsg *ssv.SignedPartialSignature, msgType types.MsgType) *types.Message {
	return ssvMsg(qbftMsg, partialSigMsg, types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleProposer), msgType)
}

var SSVMsgAggregator = func(qbftMsg *qbft.SignedMessage, partialSigMsg *ssv.SignedPartialSignature, msgType types.MsgType) *types.Message {
	return ssvMsg(qbftMsg, partialSigMsg, types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleAggregator), msgType)
}

var SSVMsgSyncCommittee = func(qbftMsg *qbft.SignedMessage, partialSigMsg *ssv.SignedPartialSignature, msgType types.MsgType) *types.Message {
	return ssvMsg(qbftMsg, partialSigMsg, types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleSyncCommittee), msgType)
}

var SSVMsgSyncCommitteeContribution = func(qbftMsg *qbft.SignedMessage, partialSigMsg *ssv.SignedPartialSignature, msgType types.MsgType) *types.Message {
	return ssvMsg(qbftMsg, partialSigMsg, types.NewBaseMsgID(TestingValidatorPubKey[:], types.BNRoleSyncCommitteeContribution), msgType)
}

var ssvMsg = func(qbftMsg *qbft.SignedMessage, postMsg *ssv.SignedPartialSignature, msgID types.MessageID, msgType types.MsgType) *types.Message {
	var data []byte
	if qbftMsg != nil {
		data, _ = qbftMsg.Encode()
	} else if postMsg != nil {
		data, _ = postMsg.Encode()
	} else {
		panic("msg type undefined")
	}

	return &types.Message{
		ID:   types.PopulateMsgType(msgID, msgType),
		Data: data,
	}
}

var PostConsensusAttestationMsgWithWrongSig = func(sk *bls.SecretKey, id types.OperatorID, height qbft.Height) *ssv.SignedPartialSignature {
	return postConsensusAttestationMsg(sk, id, height, true, false)
}

var PostConsensusAttestationMsgWithWrongRoot = func(sk *bls.SecretKey, id types.OperatorID, height qbft.Height) *ssv.SignedPartialSignature {
	return postConsensusAttestationMsg(sk, id, height, true, false)
}

var PostConsensusAttestationMsg = func(sk *bls.SecretKey, id types.OperatorID, height qbft.Height) *ssv.SignedPartialSignature {
	return postConsensusAttestationMsg(sk, id, height, false, false)
}

var postConsensusAttestationMsg = func(
	sk *bls.SecretKey,
	id types.OperatorID,
	height qbft.Height,
	wrongRoot bool,
	wrongBeaconSig bool,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(TestingAttestationData.Target.Epoch, types.DomainAttester)
	signed, root, _ := signer.SignBeaconObject(TestingAttestationData, d, sk.GetPublicKey().Serialize())

	if wrongBeaconSig {
		signed, _, _ = signer.SignBeaconObject(TestingAttestationData, d, TestingWrongValidatorPubKey[:])
	}

	if wrongRoot {
		root = []byte{1, 2, 3, 4}
	}

	msgs := ssv.PartialSignatures{
		Messages: []*ssv.PartialSignature{
			{
				Slot:             TestingDutySlot,
				PartialSignature: signed,
				SigningRoot:      root,
				Signer:           id,
			},
		},
	}
	sig, _ := signer.SignRoot(msgs, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msgs,
		Signature: sig,
		Signer:    id,
	}
}

var PostConsensusProposerMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return postConsensusBeaconBlockMsg(sk, id, false, false)
}

var postConsensusBeaconBlockMsg = func(
	sk *bls.SecretKey,
	id types.OperatorID,
	wrongRoot bool,
	wrongBeaconSig bool,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()

	d, _ := beacon.DomainData(1, types.DomainProposer) // epoch doesn't matter here, hard coded
	sig, root, _ := signer.SignBeaconObject(TestingBeaconBlock, d, sk.GetPublicKey().Serialize())
	blsSig := phase0.BLSSignature{}
	copy(blsSig[:], sig)

	signed := bellatrix.SignedBeaconBlock{
		Message:   TestingBeaconBlock,
		Signature: blsSig,
	}

	if wrongBeaconSig {
		//signed, _, _ = signer.SignAttestation(TestingAttestationData, TestingAttesterDuty, TestingWrongSK.GetPublicKey().Serialize())
		panic("implement")
	}

	if wrongRoot {
		root = []byte{1, 2, 3, 4}
	}

	msgs := ssv.PartialSignatures{
		Messages: []*ssv.PartialSignature{
			{
				Slot:             TestingDutySlot,
				PartialSignature: signed.Signature[:],
				SigningRoot:      root,
				Signer:           id,
			},
		},
	}
	msgSig, _ := signer.SignRoot(msgs, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msgs,
		Signature: msgSig,
		Signer:    id,
	}
}

var PreConsensusFailedMsg = func(msgSigner *bls.SecretKey, msgSignerID types.OperatorID) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(TestingDutyEpoch, types.DomainRandao)
	signed, root, _ := signer.SignBeaconObject(types.SSZUint64(TestingDutyEpoch), d, msgSigner.GetPublicKey().Serialize())

	msg := ssv.PartialSignatures{
		Messages: []*ssv.PartialSignature{
			{
				Slot:             TestingDutySlot,
				PartialSignature: signed[:],
				SigningRoot:      root,
				Signer:           msgSignerID,
			},
		},
	}
	sig, _ := signer.SignRoot(msg, types.PartialSignatureType, msgSigner.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msg,
		Signature: sig,
		Signer:    msgSignerID,
	}
}

var PreConsensusRandaoMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return randaoMsg(sk, id, false, TestingDutyEpoch, TestingDutySlot, 1)
}

// PreConsensusRandaoNextEpochMsg testing for a second duty start
var PreConsensusRandaoNextEpochMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return randaoMsg(sk, id, false, TestingDutyEpoch2, TestingDutySlot2, 1)
}

var PreConsensusRandaoDifferentEpochMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return randaoMsg(sk, id, false, TestingDutyEpoch+1, TestingDutySlot, 1)
}

var PreConsensusRandaoWrongSlotMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return randaoMsg(sk, id, false, TestingDutyEpoch, TestingDutySlot+1, 1)
}

var PreConsensusRandaoMultiMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return randaoMsg(sk, id, false, TestingDutyEpoch, TestingDutySlot, 2)
}

var PreConsensusRandaoNoMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return randaoMsg(sk, id, false, TestingDutyEpoch, TestingDutySlot, 0)
}

var PreConsensusRandaoDifferentSignerMsg = func(msgSigner, randaoSigner *bls.SecretKey, msgSignerID, randaoSignerID types.OperatorID) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(TestingDutyEpoch, types.DomainRandao)
	signed, root, _ := signer.SignBeaconObject(types.SSZUint64(TestingDutyEpoch), d, randaoSigner.GetPublicKey().Serialize())

	msg := ssv.PartialSignatures{
		Messages: []*ssv.PartialSignature{
			{
				Slot:             TestingDutySlot,
				PartialSignature: signed[:],
				SigningRoot:      root,
				Signer:           randaoSignerID,
			},
		},
	}
	sig, _ := signer.SignRoot(msg, types.PartialSignatureType, msgSigner.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msg,
		Signature: sig,
		Signer:    msgSignerID,
	}
}

var randaoMsg = func(
	sk *bls.SecretKey,
	id types.OperatorID,
	wrongRoot bool,
	epoch phase0.Epoch,
	slot phase0.Slot,
	msgCnt int,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(epoch, types.DomainRandao)
	signed, root, _ := signer.SignBeaconObject(types.SSZUint64(epoch), d, sk.GetPublicKey().Serialize())

	msgs := ssv.PartialSignatures{
		Messages: []*ssv.PartialSignature{},
	}
	for i := 0; i < msgCnt; i++ {
		msg := &ssv.PartialSignature{
			Slot:             slot,
			PartialSignature: signed[:],
			SigningRoot:      root,
			Signer:           id,
		}
		if wrongRoot {
			msg.SigningRoot = make([]byte, 32)
		}
		msgs.Messages = append(msgs.Messages, msg)
	}

	sig, _ := signer.SignRoot(msgs, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msgs,
		Signature: sig,
		Signer:    id,
	}
}

var PreConsensusSelectionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return PreConsensusCustomSlotSelectionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot)
}

var PreConsensusSelectionProofNextEpochMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return selectionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot2, TestingDutySlot2, 1)
}

var PreConsensusMultiSelectionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return selectionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot, TestingDutySlot, 3)
}

var PreConsensusCustomSlotSelectionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID, slot phase0.Slot) *ssv.SignedPartialSignature {
	return selectionProofMsg(msgSK, beaconSK, msgID, beaconID, slot, TestingDutySlot, 1)
}

var PreConsensusWrongMsgSlotSelectionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return selectionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot, TestingDutySlot+1, 1)
}

var selectionProofMsg = func(
	sk *bls.SecretKey,
	beaconsk *bls.SecretKey,
	id types.OperatorID,
	beaconid types.OperatorID,
	slot phase0.Slot,
	msgSlot phase0.Slot,
	msgCnt int,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(1, types.DomainSelectionProof)
	signed, root, _ := signer.SignBeaconObject(types.SSZUint64(slot), d, beaconsk.GetPublicKey().Serialize())

	_msgs := make([]*ssv.PartialSignature, 0)
	for i := 0; i < msgCnt; i++ {
		_msgs = append(_msgs, &ssv.PartialSignature{
			Slot:             msgSlot,
			PartialSignature: signed[:],
			SigningRoot:      root,
			Signer:           beaconid,
		})
	}

	msgs := ssv.PartialSignatures{
		Messages: _msgs,
	}
	msgSig, _ := signer.SignRoot(msgs, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msgs,
		Signature: msgSig,
		Signer:    id,
	}
}

var PostConsensusAggregatorMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return postConsensusAggregatorMsg(sk, id, false, false)
}

var postConsensusAggregatorMsg = func(
	sk *bls.SecretKey,
	id types.OperatorID,
	wrongRoot bool,
	wrongBeaconSig bool,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(1, types.DomainAggregateAndProof)
	signed, root, _ := signer.SignBeaconObject(TestingAggregateAndProof, d, sk.GetPublicKey().Serialize())

	if wrongBeaconSig {
		//signed, _, _ = signer.SignAttestation(TestingAttestationData, TestingAttesterDuty, TestingWrongSK.GetPublicKey().Serialize())
		panic("implement")
	}

	if wrongRoot {
		root = []byte{1, 2, 3, 4}
	}

	msgs := ssv.PartialSignatures{
		Messages: []*ssv.PartialSignature{
			{
				Slot:             TestingDutySlot,
				PartialSignature: signed,
				SigningRoot:      root,
				Signer:           id,
			},
		},
	}
	sig, _ := signer.SignRoot(msgs, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msgs,
		Signature: sig,
		Signer:    id,
	}
}

var PostConsensusSyncCommitteeMsg = func(sk *bls.SecretKey, id types.OperatorID) *ssv.SignedPartialSignature {
	return postConsensusSyncCommitteeMsg(sk, id, false, false)
}

var postConsensusSyncCommitteeMsg = func(
	sk *bls.SecretKey,
	id types.OperatorID,
	wrongRoot bool,
	wrongBeaconSig bool,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(1, types.DomainSyncCommittee)
	signed, root, _ := signer.SignBeaconObject(types.SSZBytes(TestingSyncCommitteeBlockRoot[:]), d, sk.GetPublicKey().Serialize())

	if wrongBeaconSig {
		//signedAtt, _, _ = signer.SignAttestation(TestingAttestationData, TestingAttesterDuty, TestingWrongSK.GetPublicKey().Serialize())
		panic("implement")
	}

	if wrongRoot {
		root = []byte{1, 2, 3, 4}
	}

	msgs := ssv.PartialSignatures{
		Messages: []*ssv.PartialSignature{
			{
				Slot:             TestingDutySlot,
				PartialSignature: signed,
				SigningRoot:      root,
				Signer:           id,
			},
		},
	}
	sig, _ := signer.SignRoot(msgs, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   msgs,
		Signature: sig,
		Signer:    id,
	}
}

var PreConsensusContributionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return PreConsensusCustomSlotContributionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot)
}

var PreConsensusContributionProofNextEpochMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return contributionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot2, TestingDutySlot2, false, false)
}

var PreConsensusCustomSlotContributionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID, slot phase0.Slot) *ssv.SignedPartialSignature {
	return contributionProofMsg(msgSK, beaconSK, msgID, beaconID, slot, TestingDutySlot, false, false)
}

var PreConsensusWrongMsgSlotContributionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return contributionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot, TestingDutySlot+1, false, false)
}

var PreConsensusWrongOrderContributionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return contributionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot, TestingDutySlot, true, false)
}

var PreConsensusWrongCountContributionProofMsg = func(msgSK, beaconSK *bls.SecretKey, msgID, beaconID types.OperatorID) *ssv.SignedPartialSignature {
	return contributionProofMsg(msgSK, beaconSK, msgID, beaconID, TestingDutySlot, TestingDutySlot, false, true)
}

var contributionProofMsg = func(
	sk, beaconsk *bls.SecretKey,
	id, beaconid types.OperatorID,
	slot phase0.Slot,
	msgSlot phase0.Slot,
	wrongMsgOrder bool,
	dropLastMsg bool,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	d, _ := beacon.DomainData(1, types.DomainSyncCommitteeSelectionProof)

	msgs := make([]*ssv.PartialSignature, 0)
	for index := range TestingContributionProofIndexes {
		subnet, _ := beacon.SyncCommitteeSubnetID(uint64(index))
		data := &altair.SyncAggregatorSelectionData{
			Slot:              slot,
			SubcommitteeIndex: subnet,
		}
		sig, root, _ := signer.SignBeaconObject(data, d, beaconsk.GetPublicKey().Serialize())
		msg := &ssv.PartialSignature{
			Slot:             msgSlot,
			PartialSignature: sig[:],
			SigningRoot:      ensureRoot(root),
			Signer:           beaconid,
		}

		if dropLastMsg && index == len(TestingContributionProofIndexes)-1 {
			break
		}
		msgs = append(msgs, msg)
	}

	if wrongMsgOrder {
		m := msgs[0]
		msgs[0] = msgs[1]
		msgs[1] = m
	}

	msg := &ssv.PartialSignatures{
		Messages: msgs,
	}

	msgSig, _ := signer.SignRoot(msg, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   *msg,
		Signature: msgSig,
		Signer:    id,
	}
}

var PostConsensusSyncCommitteeContributionMsg = func(sk *bls.SecretKey, id types.OperatorID, keySet *TestKeySet) *ssv.SignedPartialSignature {
	return postConsensusSyncCommitteeContributionMsg(sk, id, TestingValidatorIndex, keySet, false, false)
}

var postConsensusSyncCommitteeContributionMsg = func(
	sk *bls.SecretKey,
	id types.OperatorID,
	validatorIndex phase0.ValidatorIndex,
	keySet *TestKeySet,
	wrongRoot bool,
	wrongBeaconSig bool,
) *ssv.SignedPartialSignature {
	signer := NewTestingKeyManager()
	beacon := NewTestingBeaconNode()
	dContribAndProof, _ := beacon.DomainData(1, types.DomainContributionAndProof)

	msgs := make([]*ssv.PartialSignature, 0)
	for index := range TestingSyncCommitteeContributions {
		// sign proof
		subnet, _ := beacon.SyncCommitteeSubnetID(uint64(index))
		data := &altair.SyncAggregatorSelectionData{
			Slot:              TestingDutySlot,
			SubcommitteeIndex: subnet,
		}
		dProof, _ := beacon.DomainData(1, types.DomainSyncCommitteeSelectionProof)

		proofSig, _, _ := signer.SignBeaconObject(data, dProof, keySet.ValidatorPK.Serialize())
		blsProofSig := phase0.BLSSignature{}
		copy(blsProofSig[:], proofSig)

		// get contribution
		contribution, _ := beacon.GetSyncCommitteeContribution(TestingDutySlot, subnet, TestingValidatorPubKey)

		// sign contrib and proof
		contribAndProof := &altair.ContributionAndProof{
			AggregatorIndex: validatorIndex,
			Contribution:    contribution,
			SelectionProof:  blsProofSig,
		}
		signed, root, _ := signer.SignBeaconObject(contribAndProof, dContribAndProof, sk.GetPublicKey().Serialize())

		if wrongRoot {
			root = []byte{1, 2, 3, 4}
		}

		msg := &ssv.PartialSignature{
			Slot:             TestingDutySlot,
			PartialSignature: signed,
			SigningRoot:      root,
			Signer:           id,
		}

		if wrongBeaconSig {
			//signedAtt, _, _ = signer.SignAttestation(TestingAttestationData, TestingAttesterDuty, TestingWrongSK.GetPublicKey().Serialize())
			panic("implement")
		}

		msgs = append(msgs, msg)
	}

	msg := &ssv.PartialSignatures{
		Messages: msgs,
	}

	sig, _ := signer.SignRoot(msg, types.PartialSignatureType, sk.GetPublicKey().Serialize())
	return &ssv.SignedPartialSignature{
		Message:   *msg,
		Signature: sig,
		Signer:    id,
	}
}

// ensureRoot ensures that SigningRoot will have sufficient allocated memory
// otherwise we get panic from bls:
// github.com/herumi/bls-eth-go-binary/bls.(*Sign).VerifyByte:738
func ensureRoot(root []byte) []byte {
	n := len(root)
	if n == 0 {
		n = 1
	}
	tmp := make([]byte, n)
	copy(tmp[:], root[:])
	return tmp[:]
}
