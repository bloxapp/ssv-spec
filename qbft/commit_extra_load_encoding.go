// Code generated by fastssz. DO NOT EDIT.
// Hash: a6c5f14e52d27353928e50bf7c810dbd51faa7f2ed36b8360997f7c96b0212fa
// Version: 0.1.3
package qbft

import (
	"github.com/bloxapp/ssv-spec/types"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the CommitExtraLoad object
func (c *CommitExtraLoad) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the CommitExtraLoad object to a target array
func (c *CommitExtraLoad) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Signatures'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(c.Signatures); ii++ {
		offset += 4
		offset += len(c.Signatures[ii])
	}

	// Field (0) 'Signatures'
	if size := len(c.Signatures); size > 4 {
		err = ssz.ErrListTooBigFn("CommitExtraLoad.Signatures", size, 4)
		return
	}
	{
		offset = 4 * len(c.Signatures)
		for ii := 0; ii < len(c.Signatures); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(c.Signatures[ii])
		}
	}
	for ii := 0; ii < len(c.Signatures); ii++ {
		if size := len(c.Signatures[ii]); size > 4 {
			err = ssz.ErrBytesLengthFn("CommitExtraLoad.Signatures[ii]", size, 4)
			return
		}
		dst = append(dst, c.Signatures[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the CommitExtraLoad object
func (c *CommitExtraLoad) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Signatures'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Signatures'
	{
		buf = tail[o0:]
		num, err := ssz.DecodeDynamicLength(buf, 4)
		if err != nil {
			return err
		}
		c.Signatures = make([]types.Signature, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 4 {
				return ssz.ErrBytesLength
			}
			if cap(c.Signatures[indx]) == 0 {
				c.Signatures[indx] = types.Signature(make([]byte, 0, len(buf)))
			}
			c.Signatures[indx] = append(c.Signatures[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the CommitExtraLoad object
func (c *CommitExtraLoad) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Signatures'
	for ii := 0; ii < len(c.Signatures); ii++ {
		size += 4
		size += len(c.Signatures[ii])
	}

	return
}

// HashTreeRoot ssz hashes the CommitExtraLoad object
func (c *CommitExtraLoad) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the CommitExtraLoad object with a hasher
func (c *CommitExtraLoad) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Signatures'
	{
		subIndx := hh.Index()
		num := uint64(len(c.Signatures))
		if num > 4 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range c.Signatures {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 4 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (4+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the CommitExtraLoad object
func (c *CommitExtraLoad) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
