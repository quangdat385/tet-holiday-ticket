package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type TicketItemRouter struct {
}

func (r *TicketItemRouter) InitTicketItemRouter(Router *gin.RouterGroup) {
	TicketItemRouterPublicGroup := Router.Group("ticket-item")
	TicketItemRouterPublicGroup.Use()
	{
	}
	TicketItemRouterPrivateGroup := Router.Group("ticket-item")
	TicketItemRouterPrivateGroup.Use(middleware.AuthenMiddleWare())
	{
		TicketItemRouterPrivateGroup.POST("create", middleware.RoleMiddleware("Admin"), controller.TicketItemController.CreateTicketItem)
		TicketItemRouterPrivateGroup.PUT("update/:ticket_item_id", middleware.RoleMiddleware("Admin"), controller.TicketItemController.UpdateTicketItem)
		TicketItemRouterPrivateGroup.DELETE("delete/:ticket_item_id", middleware.RoleMiddleware("Admin"), controller.TicketItemController.DeleteTicketItem)
		TicketItemRouterPrivateGroup.PUT("set-stock-cache", middleware.RoleMiddleware("Admin"), controller.TicketItemController.SetStockCache)
		TicketItemRouterPrivateGroup.PUT("decrease-stock-cache", middleware.RoleMiddleware("Admin"), controller.TicketItemController.DecreaseStockCache)
	}
}
