package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	base58 "github.com/itchyny/base58-go"
)

func encrypt(input string) []byte {
	algo := sha256.New()
	_, err := algo.Write([]byte(input))
	if err != nil {
		fmt.Println("algo.Write()")
		return []byte{}
	}

	return algo.Sum(nil)
}

func base58Encode(b []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(b)
	if err != nil {
		fmt.Println("encoding.Encode()", err)
		os.Exit(1)
	}

	return string(encoded)
}

// GenerateShortLink generates a high entropy shortened URL given a link.
// For a given URL, we will encrypt the url using SHA256.
// And then we will encode the url using base58 (the same as bitcoin).
// Since we want to limit the url to only 8 characters, we will return the first
// 8 characters of that string.
func GenerateShortLink(initialLink, userId string) string {
	var i big.Int

	hash := encrypt(initialLink + userId)
	num := i.SetBytes(hash).Uint64()
	finalString := base58Encode(fmt.Appendf(nil, "%d", num))
	return finalString[:8]
}
