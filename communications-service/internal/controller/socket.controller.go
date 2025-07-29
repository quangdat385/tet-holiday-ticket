package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
)

var SocketRooter = new(cSocket)

type cSocket struct {
}

func (c *cSocket) ConnectHandler(ctx *gin.Context) {
	service.SocketService().Connect(ctx)
}
