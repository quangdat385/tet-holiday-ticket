package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if global.AllowedOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		}

		if c.Request.Method == http.MethodOptions {
			response.ErrorResponse(c, response.ErrorCorsCodeStatus, "CORS preflight request")
			c.Abort()
			return
		}

		c.Next()
	}
}
