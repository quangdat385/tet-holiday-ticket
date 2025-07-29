package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(input, salt string) string {
	key := []byte(salt)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
