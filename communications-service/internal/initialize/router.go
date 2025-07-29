package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/router"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	r.Use(middleware.Logger())
	r.Use(middleware.CORSMiddleware())
	socketRouter := router.RouterGroupApp.Socket
	{
		socketRouter.SocketRouter.InitSocketRouter(r)
	}
	userRouter := router.RouterGroupApp.User
	managerRouter := router.RouterGroupApp.Manager
	MainGroup := r.Group("/communicaition-service/api/v1")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.MessageRouter.InitMessageRoter(MainGroup)
	}
	{
		managerRouter.MessageRouter.InitMessageRoter(MainGroup)
	}
	return r
}
