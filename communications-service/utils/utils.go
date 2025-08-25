package utils

import (
	"fmt"
	"math/rand"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
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

// Loại phần tử có giá trị userId khỏi slice userIds
func RemoveUserId(userIds []model.UserIDS, userId int64) []model.UserIDS {
	result := make([]model.UserIDS, 0, len(userIds))
	for _, id := range userIds {
		if id.UserID != userId {
			result = append(result, id)
		}
	}
	return result
}
