package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
)

func SocketUserInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.GetInt("UserID")
		if clientID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
			c.Abort()
			return
		}
		result, err := service.InformationService().GetInformationByUserID(c, int64(clientID))
		fmt.Println("User information:", result)
		if err != nil || result.ID == 0 {
			_, err = service.InformationService().InsertInformationByUserID(c, model.InfomationInput{
				UserID: int64(clientID),
				Status: true,
				Value:  "",
				Type:   "socket_id",
			})
			if err != nil {
				log.Println("Error inserting communication info:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert communication info"})
				return
			}
		}
		fmt.Println("User information:", result)
		c.Next()
	}
}
