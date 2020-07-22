package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(hexStr string) (string, error) {
	if len(hexStr)%2 != 0 {
		hexStr = "0" + hexStr
	}
	decoded, err := hex.DecodeString(hexStr)
	// Print the easter egg
	fmt.Printf("%s\n", decoded)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}

func FixedXOR(a string, b string) (string, error) {
	if len(a)%2 != 0 {
		a = "0" + a
		b = "0" + b
	}

	decodedA, err := hex.DecodeString(a)
	if err != nil {
		return "", err
	}

	decodedB, err := hex.DecodeString(b)
	if err != nil {
		return "", err
	}

	var result []byte
	for i, _ := range decodedA {
		result = append(result, decodedA[i]^decodedB[i])
	}

	// Print the easter egg
	fmt.Printf("%s\n", decodedA)
	fmt.Printf("%s\n", decodedB)
	fmt.Printf("%s\n", result)

	return hex.EncodeToString(result), nil
}
