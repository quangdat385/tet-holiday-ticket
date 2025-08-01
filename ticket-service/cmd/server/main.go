package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/quangdat385/holiday-ticket/ticket-service/cmd/swag/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/initialize"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service/impl"
	"github.com/quangdat385/holiday-ticket/ticket-service/pkg/utils/crypto"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Holiday Ticket Service
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/quangdat385/holiday-ticket/ticket-service

// @contact.name   TEAM DATNGUYEN
// @contact.url    github.com/quangdat385/holiday-ticket/ticket-service
// @contact.email  datnguyen03011985@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /ticket-service/api/v1/ticket
// @schemes   http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @type http
// @scheme bearer
// @bearerFormat JWT

// @security BearerAuth

func main() {
	r := initialize.Run()
	r.GET("/ping", func(c *gin.Context) {
		hashKey := crypto.GenerateHash("1", "123456")
		c.JSON(200, gin.H{
			"message": hashKey,
		})
	})
	r.GET("/ticket-service/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	go func() {
		impl.InitServer(":50051")
	}()

	fmt.Println("Starting ticket service...", global.Config.Server.Port)
	port := fmt.Sprintf(":%d", global.Config.Server.Port)
	r.Run(port) // listen and serve on
}
