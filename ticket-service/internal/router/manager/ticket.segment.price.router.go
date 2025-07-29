package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type TicketSegmentPriceRouter struct {
}

func (r *TicketSegmentPriceRouter) InitTicketSegmentPriceRouter(Router *gin.RouterGroup) {
	TicketSegmentPriceRouterPublicGroup := Router.Group("ticket-segment-price")
	TicketSegmentPriceRouterPublicGroup.Use()
	{
	}
	TicketSegmentPriceRouterPrivateGroup := Router.Group("ticket-segment-price")
	TicketSegmentPriceRouterPrivateGroup.Use(middleware.AuthenMiddleWare())
	{
		TicketSegmentPriceRouterPrivateGroup.POST("create", controller.TicketSegmentPriceController.CreateTicketSegmentPrice)
		TicketSegmentPriceRouterPrivateGroup.PATCH("update/:ticket_segment_price_id", controller.TicketSegmentPriceController.UpdateTicketSegmentPrice)
		TicketSegmentPriceRouterPrivateGroup.DELETE("delete/:ticket_segment_price_id", controller.TicketSegmentPriceController.DeleteTicketSegmentPrice)
	}
}
