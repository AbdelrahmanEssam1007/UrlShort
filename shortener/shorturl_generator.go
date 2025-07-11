package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

// sha2560f computes the SHA-256 hash of the input string and returns the resulting byte slice.
func sha2560f(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

// EncodeBase58 encodes the given byte slice into a Base58 string using to avoid confusion with similar-looking characters.
func EncodeBase58(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// GenerateShortUrl generates a short URL based on the original URL and user ID.
func GenerateShortUrl(originalUrl string, userId string) string {
	urlHash := sha2560f(originalUrl + userId)
	generatedNum := new(big.Int).SetBytes(urlHash).Uint64()
	encodedUrl := EncodeBase58([]byte(fmt.Sprintf("%d", generatedNum)))
	return encodedUrl[:8] // Limit the length to 8 characters
}
