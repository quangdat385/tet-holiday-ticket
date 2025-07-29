package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type TicketRouter struct {
}

func (p *TicketRouter) InitTicketRouter(Router *gin.RouterGroup) {
	TicketRouterPublicGroup := Router.Group("ticket")
	TicketRouterPublicGroup.Use()
	{
	}
	TicketRouterPrivateGroup := Router.Group("ticket")
	TicketRouterPrivateGroup.Use(middleware.AuthenMiddleWare()) // auth and permission middleware
	{
		TicketRouterPrivateGroup.POST("create", middleware.RoleMiddleware("Admin"), controller.TicketController.CreateTicket)              // create ticket
		TicketRouterPrivateGroup.PATCH("update/:ticket_id", middleware.RoleMiddleware("Admin"), controller.TicketController.UpdateTicket)  // update ticket
		TicketRouterPrivateGroup.DELETE("delete/:ticket_id", middleware.RoleMiddleware("Admin"), controller.TicketController.DeleteTicket) // get ticket
	}
}
