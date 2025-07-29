package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/payment-service/global"
	"github.com/quangdat385/holiday-ticket/payment-service/internal/routers"
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
	userRouter := routers.RouterGroupApp.User
	managerRouter := routers.RouterGroupApp.Manager

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.PaymentRouter.InitPaymentRouter(MainGroup)
	}
	{
		managerRouter.PaymentRouter.InitPaymentRouter(MainGroup)
	}
	return r
}
