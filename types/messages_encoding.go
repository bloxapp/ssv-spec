// Code generated by fastssz. DO NOT EDIT.
// Hash: 5bc4a86fe33d3dcdcd331912c5c1a2064bbd2c23879f4145c55653985cb5d4c0
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
	if size := len(s.Data); size > 705240 {
		err = ssz.ErrBytesLengthFn("SSVMessage.Data", size, 705240)
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
		if len(buf) > 705240 {
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
		if byteLen > 705240 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(s.Data)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (705240+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SSVMessage object
func (s *SSVMessage) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the SignedSSVMessage object
func (s *SignedSSVMessage) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedSSVMessage object to a target array
func (s *SignedSSVMessage) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(16)

	// Offset (0) 'Signatures'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Signatures); ii++ {
		offset += 4
		offset += len(s.Signatures[ii])
	}

	// Offset (1) 'OperatorIDs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.OperatorIDs) * 8

	// Offset (2) 'SSVMessage'
	dst = ssz.WriteOffset(dst, offset)
	if s.SSVMessage == nil {
		s.SSVMessage = new(SSVMessage)
	}
	offset += s.SSVMessage.SizeSSZ()

	// Offset (3) 'FullData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.FullData)

	// Field (0) 'Signatures'
	if size := len(s.Signatures); size > 13 {
		err = ssz.ErrListTooBigFn("SignedSSVMessage.Signatures", size, 13)
		return
	}
	{
		offset = 4 * len(s.Signatures)
		for ii := 0; ii < len(s.Signatures); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(s.Signatures[ii])
		}
	}
	for ii := 0; ii < len(s.Signatures); ii++ {
		if size := len(s.Signatures[ii]); size > 256 {
			err = ssz.ErrBytesLengthFn("SignedSSVMessage.Signatures[ii]", size, 256)
			return
		}
		dst = append(dst, s.Signatures[ii]...)
	}

	// Field (1) 'OperatorIDs'
	if size := len(s.OperatorIDs); size > 13 {
		err = ssz.ErrListTooBigFn("SignedSSVMessage.OperatorIDs", size, 13)
		return
	}
	for ii := 0; ii < len(s.OperatorIDs); ii++ {
		dst = ssz.MarshalUint64(dst, uint64(s.OperatorIDs[ii]))
	}

	// Field (2) 'SSVMessage'
	if dst, err = s.SSVMessage.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'FullData'
	if size := len(s.FullData); size > 4219065 {
		err = ssz.ErrBytesLengthFn("SignedSSVMessage.FullData", size, 4219065)
		return
	}
	dst = append(dst, s.FullData...)

	return
}

// UnmarshalSSZ ssz unmarshals the SignedSSVMessage object
func (s *SignedSSVMessage) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 16 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2, o3 uint64

	// Offset (0) 'Signatures'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 16 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'OperatorIDs'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'SSVMessage'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Offset (3) 'FullData'
	if o3 = ssz.ReadOffset(buf[12:16]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Field (0) 'Signatures'
	{
		buf = tail[o0:o1]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		s.Signatures = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 256 {
				return ssz.ErrBytesLength
			}
			if cap(s.Signatures[indx]) == 0 {
				s.Signatures[indx] = make([]byte, 0, len(buf))
			}
			s.Signatures[indx] = append(s.Signatures[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (1) 'OperatorIDs'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		s.OperatorIDs = ssz.ExtendUint64(s.OperatorIDs, num)
		for ii := 0; ii < num; ii++ {
			s.OperatorIDs[ii] = OperatorID(ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8]))
		}
	}

	// Field (2) 'SSVMessage'
	{
		buf = tail[o2:o3]
		if s.SSVMessage == nil {
			s.SSVMessage = new(SSVMessage)
		}
		if err = s.SSVMessage.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (3) 'FullData'
	{
		buf = tail[o3:]
		if len(buf) > 4219065 {
			return ssz.ErrBytesLength
		}
		if cap(s.FullData) == 0 {
			s.FullData = make([]byte, 0, len(buf))
		}
		s.FullData = append(s.FullData, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedSSVMessage object
func (s *SignedSSVMessage) SizeSSZ() (size int) {
	size = 16

	// Field (0) 'Signatures'
	for ii := 0; ii < len(s.Signatures); ii++ {
		size += 4
		size += len(s.Signatures[ii])
	}

	// Field (1) 'OperatorIDs'
	size += len(s.OperatorIDs) * 8

	// Field (2) 'SSVMessage'
	if s.SSVMessage == nil {
		s.SSVMessage = new(SSVMessage)
	}
	size += s.SSVMessage.SizeSSZ()

	// Field (3) 'FullData'
	size += len(s.FullData)

	return
}

// HashTreeRoot ssz hashes the SignedSSVMessage object
func (s *SignedSSVMessage) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedSSVMessage object with a hasher
func (s *SignedSSVMessage) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Signatures'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Signatures))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Signatures {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 256 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (256+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (1) 'OperatorIDs'
	{
		if size := len(s.OperatorIDs); size > 13 {
			err = ssz.ErrListTooBigFn("SignedSSVMessage.OperatorIDs", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range s.OperatorIDs {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(s.OperatorIDs))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	// Field (2) 'SSVMessage'
	if err = s.SSVMessage.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'FullData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(s.FullData))
		if byteLen > 4219065 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(s.FullData)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (4219065+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedSSVMessage object
func (s *SignedSSVMessage) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
