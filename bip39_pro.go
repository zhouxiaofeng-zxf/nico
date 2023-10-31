package nico

import (
	"crypto/sha1"
	"errors"
	"github.com/tyler-smith/go-bip39"
	"math/big"
	"strconv"
	"strings"
	"unicode/utf8"
)

func getHashCode(str string) []byte {
	data := []byte(str)
	hash := sha1.Sum(data)
	return hash[:]
}

func NewEntropyPro(words string, bitSize int) ([]byte, error) {
	if err := validateEntropyBitSize(bitSize); err != nil {
		return nil, err
	}
	byteLength := bitSize / 8 //byte length

	sentenceLength := byteLength / 2 // sentence Length
	wordList := strings.Split(strings.TrimSpace(words), " ")
	wordsStr := strings.ReplaceAll(words, " ", "")
	//validateWordsSize
	if len(wordList) < sentenceLength || len(wordList) > 3*sentenceLength || utf8.RuneCountInString(wordsStr) > 3*bitSize {
		return nil, errors.New("字符格式长度错误")
	}
	// words hash
	wordsHash := getHashCode(wordsStr)

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
		return bip39.ErrEntropyLengthInvalid
	}
	return nil
}
