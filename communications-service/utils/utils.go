package utils

import (
	"fmt"
	"math/rand"
)

func GetUserKey(hashKey string, device string) string {
	return fmt.Sprintf("user:%s:%s", device, hashKey)
}
func GetUserBaseKey(hashKey string) string {
	return fmt.Sprintf("user_base:%s", hashKey)
}
func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
