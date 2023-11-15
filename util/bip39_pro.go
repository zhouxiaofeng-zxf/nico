package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
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

//对数据进行sha256处理
func getHashCode(str string) []byte {
	hash := sha256.Sum256([]byte(str))
	return hash[:]
}

func NewEntropyFromWords(words string, bitSize int) ([]byte, error) {
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
	wordsHashString := hex.EncodeToString(wordsHash)
	fmt.Println("wordsHash: ", wordsHashString)
	// Throw away big.Int for AND masking.
	sByte := make([]byte, 0)
	for i := byteLength - 1; i >= 0; i-- {
		word := wordList[i%len(wordList)]
		hash := getHashCode(strconv.Itoa(i) + word)
		hashString := hex.EncodeToString(hash)
		fmt.Println("hashString: ", strconv.Itoa(i)+":"+hashString)
		hashInt := new(big.Int).SetBytes(hash)

		endHashInt := new(big.Int).SetBytes(wordsHash)
		hashInt.Xor(hashInt, endHashInt)

		hashIntString := hex.EncodeToString(hashInt.Bytes())
		fmt.Println("hashIntString: ", strconv.Itoa(i)+":"+hashIntString)
		// Get the bytes representing the 11 bits as a 1 byte slice.
		wordBytes := padByteSliceSuper(hashInt.Bytes(), 1)
		fmt.Println("wordBytes: ", strconv.Itoa(i)+":", wordBytes)
		sByte = append(sByte, wordBytes...)
	}
	return sByte, nil
}
func NewEntropyFromWordsList(words []string, bitSize int) ([]byte, error) {
	if err := validateEntropyBitSize(bitSize); err != nil {
		return nil, err
	}
	byteLength := bitSize / 8        //byte length
	sentenceLength := byteLength / 2 // sentence Length
	if len(words) <= 0 || len(words) > sentenceLength {
		return nil, ErrInvalidWordsLength
	}
	holeWords := ""
	for _, w := range words {
		wordList := strings.Fields(w)
		//validateWordsSize
		if len(wordList) < sentenceLength || len(wordList) > 3*sentenceLength || utf8.RuneCountInString(w) > 3*bitSize {
			return nil, ErrInvalidWordsLength
		}
		holeWords = holeWords + w
	}

	wordList := strings.Fields(words[0])
	//validateWordsSize
	if len(wordList) < sentenceLength || len(wordList) > 3*sentenceLength || utf8.RuneCountInString(words[0]) > 3*bitSize {
		return nil, ErrInvalidWordsLength
	}
	// words hash
	wordsHash := getHashCode(holeWords)
	wordsHashString := hex.EncodeToString(wordsHash)
	fmt.Println("wordsHash: ", wordsHashString)
	// Throw away big.Int for AND masking.
	sByte := make([]byte, 0)
	for i := byteLength - 1; i >= 0; i-- {
		word := wordList[i%len(wordList)]
		hash := getHashCode(strconv.Itoa(i) + word)
		hashString := hex.EncodeToString(hash)
		fmt.Println("hashString: ", strconv.Itoa(i)+":"+hashString)
		hashInt := new(big.Int).SetBytes(hash)

		endHashInt := new(big.Int).SetBytes(wordsHash)
		hashInt.Xor(hashInt, endHashInt)

		hashIntString := hex.EncodeToString(hashInt.Bytes())
		fmt.Println("hashIntString: ", strconv.Itoa(i)+":"+hashIntString)
		// Get the bytes representing the 11 bits as a 1 byte slice.
		wordBytes := padByteSliceSuper(hashInt.Bytes(), 1)
		fmt.Println("wordBytes: ", strconv.Itoa(i)+":", wordBytes)
		sByte = append(sByte, wordBytes...)
	}
	return sByte, nil
}

//截取对应长度的字节数组
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
