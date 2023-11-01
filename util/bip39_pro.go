package util

import (
	"crypto/sha256"
	"errors"
	"math/big"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	// ErrInvalidWordsLength is returned when trying to use word set with an invalid size.
	ErrInvalidWordsLength = errors.New("Invalid words length")

	// ErrEntropyLengthInvalid is returned when trying to use an entropy set with an invalid size.
	ErrEntropyLengthInvalid = errors.New("Entropy length must be [128, 256] and a multiple of 32")
)

func getHashCode(str string) []byte {
	hash := sha256.Sum256([]byte(str))
	return hash[:]
}

func NewEntropyPro(words string, bitSize int) ([]byte, error) {
	if err := validateEntropyBitSize(bitSize); err != nil {
		return nil, err
	}
	byteLength := bitSize / 8 //byte length

	sentenceLength := byteLength / 2 // sentence Length
	wordList := strings.Fields(words)
	//validateWordsSize
	if len(wordList) < sentenceLength || len(wordList) > 3*sentenceLength || utf8.RuneCountInString(words) > 3*bitSize {
		return nil, ErrInvalidWordsLength
	}
	// words hash
	wordsHash := getHashCode(words)

	// Throw away big.Int for AND masking.
	sByte := make([]byte, 0)
	for i := byteLength - 1; i >= 0; i-- {
		word := wordList[i%len(wordList)]
		hash := getHashCode(strconv.Itoa(i) + word)
		hashInt := new(big.Int).SetBytes(hash)

		endHashInt := new(big.Int).SetBytes(wordsHash)
		hashInt.Xor(hashInt, endHashInt)

		// Get the bytes representing the 11 bits as a 1 byte slice.
		wordBytes := padByteSliceSuper(hashInt.Bytes(), 1)
		sByte = append(sByte, wordBytes...)
	}
	return sByte, nil
}

func padByteSliceSuper(slice []byte, length int) []byte {
	offset := len(slice) - length
	if offset <= 0 {
		return slice
	}

	newSlice := make([]byte, length)
	copy(newSlice, slice[offset:])

	return newSlice
}

func validateEntropyBitSize(bitSize int) error {
	if (bitSize%32) != 0 || bitSize < 128 || bitSize > 256 {
		return ErrEntropyLengthInvalid
	}
	return nil
}
