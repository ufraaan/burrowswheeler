package bwt

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrInvalidBlock = errors.New("invalid BW block")

var blockMagic = [4]byte{'B', 'W', 'T', 1}

const blockHeaderSize = 12

func EncodeBlocks(input []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("block size shoould be positive")
	}
	if uint64(blockSize) > uint64(^uint32(0)) {
		return nil, errors.New("block is too big")
	}
	var encoded []byte
	for offset := 0; offset < len(input) || offset == 0 && len(input) == 0; {
		end := offset + blockSize
		if end < offset || end > len(input) {
			end = len(input)
		}
		block := input[offset:end]
		transformed, primary := Transform(block)
		header := make([]byte, blockHeaderSize)
		copy(header[:4], blockMagic[:])
		binary.BigEndian.PutUint32(header[4:8], uint32(len(block)))
		binary.BigEndian.PutUint32(header[8:12], uint32(primary))
		encoded = append(encoded, header...)
		encoded = append(encoded, transformed...)
		offset = end
		if len(input) == 0 {
			break
		}
	}
	return encoded, nil
}

func DecodeBlocks(input []byte) ([]byte, error) {
	var decoded []byte
	for offset := 0; offset < len(input); {
		if len(input)-offset < blockHeaderSize {
			return nil, ErrInvalidBlock
		}
		header := input[offset : offset+blockHeaderSize]
		if string(header[:4]) != string(blockMagic[:]) {
			return nil, ErrInvalidBlock
		}
		length := binary.BigEndian.Uint32(header[4:8])
		primary := binary.BigEndian.Uint32(header[8:12])
		offset += blockHeaderSize


		// if uint()
		if uint64(length) > uint64(len(input)-offset) {
			return nil, io.ErrUnexpectedEOF
		}
		block, err := InverseTransform(input[offset:offset+int(length)], int(primary))
		if err != nil {
			return nil, ErrInvalidBlock
		}
		decoded = append(decoded, block...)
		offset += int(length)
	}
	return decoded, nil
}
