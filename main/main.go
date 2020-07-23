package main

import (
	"fmt"
	"github.com/paulriddle/cryptopals/cryptopals"
)

func main() {
	result := cryptopals.SingleByteXORCipher(
		"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
		cryptopals.FreqFromFile("fixtures/aliceinwonderland.txt"),
	)
	fmt.Printf("%#v\n", result)
}
