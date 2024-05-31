// Code generated by fastssz. DO NOT EDIT.
// Hash: 368858fe3c41187523cf56fcb1203180b9dc12e6167ef3085cfcb67e25e0a365
// Version: 0.1.3
package types

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Share object
func (s *Share) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the Share object to a target array
func (s *Share) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(510)

	// Field (0) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(s.ValidatorIndex))

	// Field (1) 'ValidatorPubKey'
	dst = append(dst, s.ValidatorPubKey[:]...)

	// Field (2) 'OwnValidatorShare'
	if dst, err = s.OwnValidatorShare.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'CommitteeID'
	dst = append(dst, s.CommitteeID[:]...)

	// Offset (4) 'ValidatorShares'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.ValidatorShares) * 350

	// Offset (5) 'Committee'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.Committee) * 56

	// Field (6) 'Quorum'
	dst = ssz.MarshalUint64(dst, s.Quorum)

	// Field (7) 'DomainType'
	dst = append(dst, s.DomainType[:]...)

	// Field (8) 'FeeRecipientAddress'
	dst = append(dst, s.FeeRecipientAddress[:]...)

	// Field (9) 'Graffiti'
	if size := len(s.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("Share.Graffiti", size, 32)
		return
	}
	dst = append(dst, s.Graffiti...)

	// Field (4) 'ValidatorShares'
	if size := len(s.ValidatorShares); size > 13 {
		err = ssz.ErrListTooBigFn("Share.ValidatorShares", size, 13)
		return
	}
	for ii := 0; ii < len(s.ValidatorShares); ii++ {
		if dst, err = s.ValidatorShares[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (5) 'Committee'
	if size := len(s.Committee); size > 13 {
		err = ssz.ErrListTooBigFn("Share.Committee", size, 13)
		return
	}
	for ii := 0; ii < len(s.Committee); ii++ {
		if dst, err = s.Committee[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Share object
func (s *Share) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 510 {
		return ssz.ErrSize
	}

	tail := buf
	var o4, o5 uint64

	// Field (0) 'ValidatorIndex'
	s.ValidatorIndex = phase0.ValidatorIndex(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'ValidatorPubKey'
	copy(s.ValidatorPubKey[:], buf[8:56])

	// Field (2) 'OwnValidatorShare'
	if err = s.OwnValidatorShare.UnmarshalSSZ(buf[56:406]); err != nil {
		return err
	}

	// Field (3) 'CommitteeID'
	copy(s.CommitteeID[:], buf[406:438])

	// Offset (4) 'ValidatorShares'
	if o4 = ssz.ReadOffset(buf[438:442]); o4 > size {
		return ssz.ErrOffset
	}

	if o4 < 510 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (5) 'Committee'
	if o5 = ssz.ReadOffset(buf[442:446]); o5 > size || o4 > o5 {
		return ssz.ErrOffset
	}

	// Field (6) 'Quorum'
	s.Quorum = ssz.UnmarshallUint64(buf[446:454])

	// Field (7) 'DomainType'
	copy(s.DomainType[:], buf[454:458])

	// Field (8) 'FeeRecipientAddress'
	copy(s.FeeRecipientAddress[:], buf[458:478])

	// Field (9) 'Graffiti'
	if cap(s.Graffiti) == 0 {
		s.Graffiti = make([]byte, 0, len(buf[478:510]))
	}
	s.Graffiti = append(s.Graffiti, buf[478:510]...)

	// Field (4) 'ValidatorShares'
	{
		buf = tail[o4:o5]
		num, err := ssz.DivideInt2(len(buf), 350, 13)
		if err != nil {
			return err
		}
		s.ValidatorShares = make([]*ValidatorShare, num)
		for ii := 0; ii < num; ii++ {
			if s.ValidatorShares[ii] == nil {
				s.ValidatorShares[ii] = new(ValidatorShare)
			}
			if err = s.ValidatorShares[ii].UnmarshalSSZ(buf[ii*350 : (ii+1)*350]); err != nil {
				return err
			}
		}
	}

	// Field (5) 'Committee'
	{
		buf = tail[o5:]
		num, err := ssz.DivideInt2(len(buf), 56, 13)
		if err != nil {
			return err
		}
		s.Committee = make([]*ShareMember, num)
		for ii := 0; ii < num; ii++ {
			if s.Committee[ii] == nil {
				s.Committee[ii] = new(ShareMember)
			}
			if err = s.Committee[ii].UnmarshalSSZ(buf[ii*56 : (ii+1)*56]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Share object
func (s *Share) SizeSSZ() (size int) {
	size = 510

	// Field (4) 'ValidatorShares'
	size += len(s.ValidatorShares) * 350

	// Field (5) 'Committee'
	size += len(s.Committee) * 56

	return
}

// HashTreeRoot ssz hashes the Share object
func (s *Share) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the Share object with a hasher
func (s *Share) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ValidatorIndex'
	hh.PutUint64(uint64(s.ValidatorIndex))

	// Field (1) 'ValidatorPubKey'
	hh.PutBytes(s.ValidatorPubKey[:])

	// Field (2) 'OwnValidatorShare'
	if err = s.OwnValidatorShare.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'CommitteeID'
	hh.PutBytes(s.CommitteeID[:])

	// Field (4) 'ValidatorShares'
	{
		subIndx := hh.Index()
		num := uint64(len(s.ValidatorShares))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.ValidatorShares {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (5) 'Committee'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Committee))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Committee {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (6) 'Quorum'
	hh.PutUint64(s.Quorum)

	// Field (7) 'DomainType'
	hh.PutBytes(s.DomainType[:])

	// Field (8) 'FeeRecipientAddress'
	hh.PutBytes(s.FeeRecipientAddress[:])

	// Field (9) 'Graffiti'
	if size := len(s.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("Share.Graffiti", size, 32)
		return
	}
	hh.PutBytes(s.Graffiti)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Share object
func (s *Share) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the ShareMember object
func (s *ShareMember) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the ShareMember object to a target array
func (s *ShareMember) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SharePubKey'
	if size := len(s.SharePubKey); size != 48 {
		err = ssz.ErrBytesLengthFn("ShareMember.SharePubKey", size, 48)
		return
	}
	dst = append(dst, s.SharePubKey...)

	// Field (1) 'Signer'
	dst = ssz.MarshalUint64(dst, uint64(s.Signer))

	return
}

// UnmarshalSSZ ssz unmarshals the ShareMember object
func (s *ShareMember) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 56 {
		return ssz.ErrSize
	}

	// Field (0) 'SharePubKey'
	if cap(s.SharePubKey) == 0 {
		s.SharePubKey = make([]byte, 0, len(buf[0:48]))
	}
	s.SharePubKey = append(s.SharePubKey, buf[0:48]...)

	// Field (1) 'Signer'
	s.Signer = OperatorID(ssz.UnmarshallUint64(buf[48:56]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ShareMember object
func (s *ShareMember) SizeSSZ() (size int) {
	size = 56
	return
}

// HashTreeRoot ssz hashes the ShareMember object
func (s *ShareMember) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the ShareMember object with a hasher
func (s *ShareMember) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SharePubKey'
	if size := len(s.SharePubKey); size != 48 {
		err = ssz.ErrBytesLengthFn("ShareMember.SharePubKey", size, 48)
		return
	}
	hh.PutBytes(s.SharePubKey)

	// Field (1) 'Signer'
	hh.PutUint64(uint64(s.Signer))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ShareMember object
func (s *ShareMember) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
