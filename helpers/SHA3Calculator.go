package helpers

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA3Calculator generates hash of input string and returns hash as string
func SHA3Calculator(input string) (hashString string) {
	hash := sha512.New()
	hash.Write([]byte(input))
	hashString = hex.EncodeToString(hash.Sum(nil))
	return
}