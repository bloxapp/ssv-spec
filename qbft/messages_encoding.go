// Code generated by fastssz. DO NOT EDIT.
// Hash: af4f57350c4050394feaa6ace372808a25c0b8c410465c8bd051b691061aca1f
// Version: 0.1.3
package qbft

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
	offset := int(76)

	// Field (0) 'MsgType'
	dst = ssz.MarshalUint64(dst, uint64(m.MsgType))

	// Field (1) 'Height'
	dst = ssz.MarshalUint64(dst, uint64(m.Height))

	// Field (2) 'Round'
	dst = ssz.MarshalUint64(dst, uint64(m.Round))

	// Offset (3) 'Identifier'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.Identifier)

	// Field (4) 'Root'
	dst = append(dst, m.Root[:]...)

	// Field (5) 'DataRound'
	dst = ssz.MarshalUint64(dst, uint64(m.DataRound))

	// Offset (6) 'RoundChangeJustification'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(m.RoundChangeJustification); ii++ {
		offset += 4
		offset += len(m.RoundChangeJustification[ii])
	}

	// Offset (7) 'PrepareJustification'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(m.PrepareJustification); ii++ {
		offset += 4
		offset += len(m.PrepareJustification[ii])
	}

	// Field (3) 'Identifier'
	if size := len(m.Identifier); size > 56 {
		err = ssz.ErrBytesLengthFn("Message.Identifier", size, 56)
		return
	}
	dst = append(dst, m.Identifier...)

	// Field (6) 'RoundChangeJustification'
	if size := len(m.RoundChangeJustification); size > 13 {
		err = ssz.ErrListTooBigFn("Message.RoundChangeJustification", size, 13)
		return
	}
	{
		offset = 4 * len(m.RoundChangeJustification)
		for ii := 0; ii < len(m.RoundChangeJustification); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(m.RoundChangeJustification[ii])
		}
	}
	for ii := 0; ii < len(m.RoundChangeJustification); ii++ {
		if size := len(m.RoundChangeJustification[ii]); size > 51852 {
			err = ssz.ErrBytesLengthFn("Message.RoundChangeJustification[ii]", size, 51852)
			return
		}
		dst = append(dst, m.RoundChangeJustification[ii]...)
	}

	// Field (7) 'PrepareJustification'
	if size := len(m.PrepareJustification); size > 13 {
		err = ssz.ErrListTooBigFn("Message.PrepareJustification", size, 13)
		return
	}
	{
		offset = 4 * len(m.PrepareJustification)
		for ii := 0; ii < len(m.PrepareJustification); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(m.PrepareJustification[ii])
		}
	}
	for ii := 0; ii < len(m.PrepareJustification); ii++ {
		if size := len(m.PrepareJustification[ii]); size > 3700 {
			err = ssz.ErrBytesLengthFn("Message.PrepareJustification[ii]", size, 3700)
			return
		}
		dst = append(dst, m.PrepareJustification[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Message object
func (m *Message) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 76 {
		return ssz.ErrSize
	}

	tail := buf
	var o3, o6, o7 uint64

	// Field (0) 'MsgType'
	m.MsgType = MessageType(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'Height'
	m.Height = Height(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'Round'
	m.Round = Round(ssz.UnmarshallUint64(buf[16:24]))

	// Offset (3) 'Identifier'
	if o3 = ssz.ReadOffset(buf[24:28]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 76 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (4) 'Root'
	copy(m.Root[:], buf[28:60])

	// Field (5) 'DataRound'
	m.DataRound = Round(ssz.UnmarshallUint64(buf[60:68]))

	// Offset (6) 'RoundChangeJustification'
	if o6 = ssz.ReadOffset(buf[68:72]); o6 > size || o3 > o6 {
		return ssz.ErrOffset
	}

	// Offset (7) 'PrepareJustification'
	if o7 = ssz.ReadOffset(buf[72:76]); o7 > size || o6 > o7 {
		return ssz.ErrOffset
	}

	// Field (3) 'Identifier'
	{
		buf = tail[o3:o6]
		if len(buf) > 56 {
			return ssz.ErrBytesLength
		}
		if cap(m.Identifier) == 0 {
			m.Identifier = make([]byte, 0, len(buf))
		}
		m.Identifier = append(m.Identifier, buf...)
	}

	// Field (6) 'RoundChangeJustification'
	{
		buf = tail[o6:o7]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		m.RoundChangeJustification = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 51852 {
				return ssz.ErrBytesLength
			}
			if cap(m.RoundChangeJustification[indx]) == 0 {
				m.RoundChangeJustification[indx] = make([]byte, 0, len(buf))
			}
			m.RoundChangeJustification[indx] = append(m.RoundChangeJustification[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (7) 'PrepareJustification'
	{
		buf = tail[o7:]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		m.PrepareJustification = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 3700 {
				return ssz.ErrBytesLength
			}
			if cap(m.PrepareJustification[indx]) == 0 {
				m.PrepareJustification[indx] = make([]byte, 0, len(buf))
			}
			m.PrepareJustification[indx] = append(m.PrepareJustification[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Message object
func (m *Message) SizeSSZ() (size int) {
	size = 76

	// Field (3) 'Identifier'
	size += len(m.Identifier)

	// Field (6) 'RoundChangeJustification'
	for ii := 0; ii < len(m.RoundChangeJustification); ii++ {
		size += 4
		size += len(m.RoundChangeJustification[ii])
	}

	// Field (7) 'PrepareJustification'
	for ii := 0; ii < len(m.PrepareJustification); ii++ {
		size += 4
		size += len(m.PrepareJustification[ii])
	}

	return
}

// HashTreeRoot ssz hashes the Message object
func (m *Message) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the Message object with a hasher
func (m *Message) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'MsgType'
	hh.PutUint64(uint64(m.MsgType))

	// Field (1) 'Height'
	hh.PutUint64(uint64(m.Height))

	// Field (2) 'Round'
	hh.PutUint64(uint64(m.Round))

	// Field (3) 'Identifier'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(m.Identifier))
		if byteLen > 56 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(m.Identifier)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (56+31)/32)
	}

	// Field (4) 'Root'
	hh.PutBytes(m.Root[:])

	// Field (5) 'DataRound'
	hh.PutUint64(uint64(m.DataRound))

	// Field (6) 'RoundChangeJustification'
	{
		subIndx := hh.Index()
		num := uint64(len(m.RoundChangeJustification))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range m.RoundChangeJustification {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 51852 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (51852+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (7) 'PrepareJustification'
	{
		subIndx := hh.Index()
		num := uint64(len(m.PrepareJustification))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range m.PrepareJustification {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 3700 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (3700+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Message object
func (m *Message) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}
