package context

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/identify_type"
)

func GetUserBaseString(ctx context.Context) (string, error) {
	userBase, ok := ctx.Value("userBase").(string)
	if !ok {
		return "", nil
	}
	return userBase, nil
}
func GetUserBase(ctx context.Context) (out identify_type.UserBase, err error) {
	userBaseString, err := GetUserBaseString(ctx)
	if err != nil {
		return out, err
	}
	fmt.Println("userBaseString", userBaseString)
	var userBase identify_type.UserBase
	if err := json.Unmarshal([]byte(userBaseString), &userBase); err != nil {
		return out, err
	}
	return userBase, nil
}
