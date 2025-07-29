package utils

import "fmt"

func GetUserKey(hashKey string, device string) string {
	return fmt.Sprintf("user:%s:%s", device, hashKey)
}
func GetUserBaseKey(hashKey string) string {
	return fmt.Sprintf("user_base:%s", hashKey)
}
