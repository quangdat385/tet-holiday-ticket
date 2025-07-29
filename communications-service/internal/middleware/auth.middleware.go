package middleware

import (
	"context"
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

func AuthenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		fmt.Println("uri", uri)
		jwtToken, ok := auth.ExtractBearerToken(c)
		userId, err := auth.ExtractXClientID(c)
		device := auth.ExtractXDeviceID(c)
		fmt.Println("jwtToken", jwtToken)
		if err != nil {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, err.Error())
			c.Abort()
			return
		}

		if !ok {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, "Unauthorized")
			c.Abort()
			return
		}
		hashKey := crypto.GenerateHash(strconv.Itoa(int(userId)), global.Config.Server.Salt)
		userKey := utils.GetUserKey(hashKey, device)
		userBaseKey := utils.GetUserBaseKey(hashKey)
		fmt.Println("userKey", userKey)
		fmt.Println("userBaseKey", userBaseKey)
		keyToken, err := global.Rdb.HGet(c.Request.Context(), "user_key_token", userKey).Result()
		fmt.Println("keyToken", keyToken)
		if err != nil {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, err.Error())
			c.Abort()
			return
		}
		userKeyToken := identify_type.UserKeyToken{}
		if err := json.Unmarshal([]byte(keyToken), &userKeyToken); err != nil {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, err.Error())
			c.Abort()
			return
		}
		fmt.Println("userKeyToken", userKeyToken)
		//check token in redis
		//validate token
		_, err = auth.VerifyTokenSubject(jwtToken, userKeyToken.PublicKey)
		if err != nil {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, err.Error())
			c.Abort()
			return
		}
		userBaseString, err := global.Rdb.HGet(c.Request.Context(), "user_base", userBaseKey).Result()
		if err != nil {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, err.Error())
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "userBase", userBaseString)
		c.Request = c.Request.WithContext(ctx)
		fmt.Println("userBaseString", userBaseString)
		c.Next()
	}
}
