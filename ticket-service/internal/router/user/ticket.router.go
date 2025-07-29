package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type TicketRouter struct {
}

func (p *TicketRouter) InitTicketRouter(Router *gin.RouterGroup) {
	TicketRouterPublicGroup := Router.Group("ticket")
	TicketRouterPublicGroup.Use()
	{
		TicketRouterPublicGroup.GET("get-all-ticket", controller.TicketController.GetAllTicket)      // create ticket
		TicketRouterPublicGroup.GET("get-one/:ticket_id", controller.TicketController.GetTicketById) // get ticket by id
	}
	TicketRouterPrivateGroup := Router.Group("ticket")
	TicketRouterPrivateGroup.Use()
	{
	}
}
