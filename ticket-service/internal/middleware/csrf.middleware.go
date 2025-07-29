package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a http.HandlerFunc that calls the next Gin handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Writer.WriteHeaderNow() // Ensure headers are written
			c.Next()
		})

		csrfHandler := nosurf.New(handler)
		csrfHandler.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response.ErrorResponse(c, response.ForbiddenCodeStatus, "CSRF token mismatch")
			c.Abort()
		}))

		csrfHandler.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
