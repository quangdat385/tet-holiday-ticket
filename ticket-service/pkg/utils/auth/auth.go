package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quangdat385/holiday-ticket/ticket-service/pkg/utils/random"
)

type Payload struct {
	jwt.StandardClaims
}

func ExtractBearerToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	fmt.Println("authHeader", authHeader)
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}
func ExtractXClientID(c *gin.Context) (int64, error) {
	clientId := c.GetHeader("x-client-id")
	if clientId == "" {
		return 0, errors.New("x-client-id header not found")
	}
	user_id, err := strconv.ParseInt(clientId, 10, 64)
	if err != nil {
		return 0, errors.New("x-client-id header not valid")
	}
	return user_id, nil
}
func ExtractXDeviceID(c *gin.Context) string {
	deviceId := c.GetHeader("x-device-id")
	if deviceId == "" {
		return strconv.Itoa(random.GenerateSixDigits())
	}
	return deviceId
}

func ParseToken(token string, secret string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Payload); ok && tokenClaims.Valid {
			return &claims.StandardClaims, nil
		}
	}
	return nil, err
}

func VerifyTokenSubject(token string, secret string) (*jwt.StandardClaims, error) {
	claims, err := ParseToken(token, secret)
	if err != nil {
		return nil, err
	}
	if err := claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}
