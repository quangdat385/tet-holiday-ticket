package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
)

type SocketRouter struct{}

func (sr *SocketRouter) InitSocketRouter(r *gin.Engine) {
	r.GET("/ws", middleware.SocketAuthenMiddleWare(), middleware.SocketUserInfoMiddleware(), controller.SocketRooter.ConnectHandler)
}
