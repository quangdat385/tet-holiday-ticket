package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type TicketItemRouter struct {
}

func (p *TicketItemRouter) InitTicketItemRouter(Router *gin.RouterGroup) {
	TicketItemRouterPublicGroup := Router.Group("ticket-item")
	TicketItemRouterPublicGroup.Use()
	{
		TicketItemRouterPublicGroup.GET("get-by-id/:ticket_id", controller.TicketItemController.GetTicketItem)
	}
	TicketItemRouterPrivateGroup := Router.Group("ticket-item")
	TicketItemRouterPrivateGroup.Use()
	{

	}
}
