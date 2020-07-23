package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"unicode/utf8"
)

// Solves Challenge 1
func HexToBase64(hs string) (string, error) {
	decoded, err := hexToBytes(hs)
	if err != nil {
		return "", err
	}
	// Print the easter egg
	fmt.Printf("%s\n", decoded)
	return base64.StdEncoding.EncodeToString(decoded), nil
}

// Solves Challenge 2
func FixedXOR(a string, b string) (string, error) {
	if len(a) != len(b) {
		return "", errors.New("FixedXOR: a and b must have the same length")
	}

	decodedA, err := hexToBytes(a)
	if err != nil {
		return "", err
	}

	decodedB, err := hexToBytes(b)
	if err != nil {
		return "", err
	}

	var result []byte
	for i := range decodedA {
		result = append(result, decodedA[i]^decodedB[i])
	}

	// Print the easter egg
	fmt.Printf("%s\n", decodedA)
	fmt.Printf("%s\n", decodedB)
	fmt.Printf("%s\n", result)

	return hex.EncodeToString(result), nil
}

// Solved Challenge 3
func SingleByteXORCipher(hs string, freq map[rune]float64) decryptionResult {
	ciphertext, err := hexToBytes(hs)
	if err != nil {
		log.Fatal(err)
	}

	bestDecryption := decryptionResult{}
	for i := 1; i < 256; i++ {
		newDecryption := decrypt(ciphertext, byte(i), freq)
		if newDecryption.Score > bestDecryption.Score {
			bestDecryption = newDecryption
		}
	}
	return bestDecryption
}

func FreqFromFile(name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal("failed to read file:", err)
	}
	return analyzeLetterFrequency(string(text))
}

func analyzeLetterFrequency(text string) map[rune]float64 {
	c := make(map[rune]float64)
	for _, char := range text {
		c[char]++
	}
	total := utf8.RuneCountInString(text)
	for char := range c {
		c[char] = c[char] / float64(total)
	}
	return c
}

type decryptionResult struct {
	Key   byte
	Score float64
	Text  string
}

func xorEachByte(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i, c := range input {
		result[i] = c ^ key
	}
	return result
}

func decrypt(ciphertext []byte, key byte, freq map[rune]float64) decryptionResult {
	plaintext := string(xorEachByte(ciphertext, key))
	score := 0.0

	for _, c := range plaintext {
		score += freq[c]
	}

	result := decryptionResult{
		Key:   key,
		Score: score / float64(len(plaintext)),
		Text:  plaintext,
	}
	return result
}

func hexToBytes(hs string) ([]byte, error) {
	if len(hs)%2 != 0 {
		hs = "0" + hs
	}
	return hex.DecodeString(hs)
}
