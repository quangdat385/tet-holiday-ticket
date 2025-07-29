package middleware

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/identify_type"
	"github.com/quangdat385/holiday-ticket/communications-service/response"
	"github.com/quangdat385/holiday-ticket/communications-service/utils"
	"github.com/quangdat385/holiday-ticket/communications-service/utils/auth"
	"github.com/quangdat385/holiday-ticket/communications-service/utils/crypto"
)

func SocketAuthenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		client_id := c.Query("client_id")
		jwtToken := c.Query("token")
		device := c.Query("device")
		if client_id == "" || jwtToken == "" || device == "" {
			c.JSON(401, gin.H{
				"code":    response.UnauthorizedCodeStatus,
				"message": "client_id, token, and device are required",
				"data":    nil,
			},
			)
			c.Abort()
			return
		}
		userId, _ := strconv.Atoi(client_id)
		fmt.Println("jwtToken", jwtToken)
		hashKey := crypto.GenerateHash(strconv.Itoa(int(userId)), global.Config.Server.Salt)
		userKey := utils.GetUserKey(hashKey, device)
		userBaseKey := utils.GetUserBaseKey(hashKey)
		fmt.Println("userKey", userKey)
		fmt.Println("userBaseKey", userBaseKey)
		keyToken, err := global.Rdb.HGet(c.Request.Context(), "user_key_token", userKey).Result()
		fmt.Println("keyToken", keyToken)
		if err != nil {
			c.JSON(401, gin.H{
				"code":    response.UnauthorizedCodeStatus,
				"message": "Unauthorized: " + err.Error(),
			})
			c.Abort()
			return
		}
		userKeyToken := identify_type.UserKeyToken{}
		if err := json.Unmarshal([]byte(keyToken), &userKeyToken); err != nil {
			c.JSON(401, gin.H{
				"code":    response.UnauthorizedCodeStatus,
				"message": "Unauthorized: " + err.Error(),
			})
			c.Abort()
			return
		}
		fmt.Println("userKeyToken", userKeyToken)
		//check token in redis
		//validate token
		_, err = auth.VerifyTokenSubject(jwtToken, userKeyToken.PublicKey)
		if err != nil {
			c.JSON(401, gin.H{
				"code":    response.UnauthorizedCodeStatus,
				"message": "Unauthorized: " + err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("UserID", userId)
		c.Next()
	}
}
