package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/payment-service/global"
	"github.com/quangdat385/holiday-ticket/payment-service/internal/initialize"
)

func main() {
	r := initialize.Run()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := fmt.Sprintln(":", global.Config.Server.Port)
	r.Run(port)
}
