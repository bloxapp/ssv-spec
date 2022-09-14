// Code generated by fastssz. DO NOT EDIT.
// Hash: fdc63158189bed20ffad91fe4d6bb4d5af53aeb98af12b67cd168fde7040d0fc
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Message object
func (m *Message) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the Message object to a target array
func (m *Message) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(64)

	// Field (0) 'ID'
	dst = append(dst, m.ID[:]...)

	// Offset (1) 'Data'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.Data)

	// Field (1) 'Data'
	if len(m.Data) > 394281 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, m.Data...)

	return
}

// UnmarshalSSZ ssz unmarshals the Message object
func (m *Message) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 64 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'ID'
	copy(m.ID[:], buf[0:60])

	// Offset (1) 'Data'
	if o1 = ssz.ReadOffset(buf[60:64]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 64 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Data'
	{
		buf = tail[o1:]
		if len(buf) > 394281 {
			return ssz.ErrBytesLength
		}
		if cap(m.Data) == 0 {
			m.Data = make([]byte, 0, len(buf))
		}
		m.Data = append(m.Data, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Message object
func (m *Message) SizeSSZ() (size int) {
	size = 64

	// Field (1) 'Data'
	size += len(m.Data)

	return
}

// HashTreeRoot ssz hashes the Message object
func (m *Message) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the Message object with a hasher
func (m *Message) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ID'
	hh.PutBytes(m.ID[:])

	// Field (1) 'Data'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(m.Data))
		if byteLen > 394281 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(m.Data)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (394281+31)/32)
	}

	hh.Merkleize(indx)
	return
}
