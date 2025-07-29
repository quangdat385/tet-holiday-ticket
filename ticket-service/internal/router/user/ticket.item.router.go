package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type TicketItemRouter struct {
}

func (p *TicketItemRouter) InitTicketItemRouter(Router *gin.RouterGroup) {
	TicketItemRouterPublicGroup := Router.Group("ticket")
	TicketItemRouterPublicGroup.Use()
	{
		TicketItemRouterPublicGroup.GET("ticket-item/:ticket_id", controller.TicketItemController.GetTicketItem)
	}
	TicketItemRouterPrivateGroup := Router.Group("ticket")
	TicketItemRouterPrivateGroup.Use()
	{

	}
}
