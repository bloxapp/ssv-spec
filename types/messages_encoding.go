// Code generated by fastssz. DO NOT EDIT.
// Hash: 5531456b0862c42236f8f810ebb855a85abe8316f714607db6ea03de509f5986
// Version: 0.1.3
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SSVMessage object
func (s *SSVMessage) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SSVMessage object to a target array
func (s *SSVMessage) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(68)

	// Field (0) 'MsgType'
	dst = ssz.MarshalUint64(dst, uint64(s.MsgType))

	// Field (1) 'MsgID'
	dst = append(dst, s.MsgID[:]...)

	// Offset (2) 'Data'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.Data)

	// Field (2) 'Data'
	if size := len(s.Data); size > 6291829 {
		err = ssz.ErrBytesLengthFn("SSVMessage.Data", size, 6291829)
		return
	}
	dst = append(dst, s.Data...)

	return
}

// UnmarshalSSZ ssz unmarshals the SSVMessage object
func (s *SSVMessage) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 68 {
		return ssz.ErrSize
	}

	tail := buf
	var o2 uint64

	// Field (0) 'MsgType'
	s.MsgType = MsgType(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'MsgID'
	copy(s.MsgID[:], buf[8:64])

	// Offset (2) 'Data'
	if o2 = ssz.ReadOffset(buf[64:68]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 68 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'Data'
	{
		buf = tail[o2:]
		if len(buf) > 6291829 {
			return ssz.ErrBytesLength
		}
		if cap(s.Data) == 0 {
			s.Data = make([]byte, 0, len(buf))
		}
		s.Data = append(s.Data, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SSVMessage object
func (s *SSVMessage) SizeSSZ() (size int) {
	size = 68

	// Field (2) 'Data'
	size += len(s.Data)

	return
}

// HashTreeRoot ssz hashes the SSVMessage object
func (s *SSVMessage) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SSVMessage object with a hasher
func (s *SSVMessage) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'MsgType'
	hh.PutUint64(uint64(s.MsgType))

	// Field (1) 'MsgID'
	hh.PutBytes(s.MsgID[:])

	// Field (2) 'Data'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(s.Data))
		if byteLen > 6291829 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(s.Data)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (6291829+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SSVMessage object
func (s *SSVMessage) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
