package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

func sha2560f(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func EncodeBase58(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortUrl(originalUrl string, userId string) string {
	urlHash := sha2560f(originalUrl + userId)
	generatedNum := new(big.Int).SetBytes(urlHash).Uint64()
	encodedUrl := EncodeBase58([]byte(fmt.Sprintf("%d", generatedNum)))
	return encodedUrl[:8]
}
