package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/order-service/global"
	"github.com/quangdat385/holiday-ticket/order-service/internal/middleware"
	"github.com/quangdat385/holiday-ticket/order-service/internal/router"
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
	userRouter := router.RouterGroupApp.User
	managerRouter := router.RouterGroupApp.Manager

	MainGroup := r.Group("/ticket-order/api/v1")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.OrderRouter.InitOrderRoter(MainGroup)
		userRouter.OrderDetailRouter.InitOrderDetailRoter(MainGroup)
	}
	{
		managerRouter.OrderRouter.InitOrderRoter(MainGroup)
		managerRouter.OrderDetailRouter.InitOrderDetailRoter(MainGroup)
	}
	return r
}
