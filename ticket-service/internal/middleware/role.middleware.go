package middleware

import (
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/pkg/utils/context"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, err := context.GetUserBase(c.Request.Context())
		fmt.Println("userRoles", userRoles)
		if err != nil {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, err.Error())
			c.Abort()
			return
		}
		if len(userRoles.UserRoles) == 0 {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, "User has no roles")
			c.Abort()
			return
		}
		for _, role := range userRoles.UserRoles {
			if slices.Contains(roles, role.RoleName) {
				c.Next()
				return
			}
		}
		response.ErrorResponse(c, response.UnauthorizedCodeStatus, "User does not have permission")
		c.Abort()
	}
}
