package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type TicketSegmentPriceRouter struct {
}

func (r *TicketSegmentPriceRouter) InitTicketSegmentPriceRouter(Router *gin.RouterGroup) {
	TicketSegmentPriceRouterPublicGroup := Router.Group("ticket-segment-price")
	TicketSegmentPriceRouterPublicGroup.Use()
	{
	}
	TicketSegmentPriceRouterPrivateGroup := Router.Group("ticket-segment-price")
	TicketSegmentPriceRouterPrivateGroup.Use()
	{
		TicketSegmentPriceRouterPrivateGroup.GET("get-one/:ticket_segment_price_id", controller.TicketSegmentPriceController.GetTicketSegmentPriceById)
		TicketSegmentPriceRouterPrivateGroup.GET("get-by-route-segment-id/:route_segment_id", controller.TicketSegmentPriceController.GetTicketSegmentPriceByRouteSegmentId)
		TicketSegmentPriceRouterPrivateGroup.GET("get-by-from-to-route-segment-id", controller.TicketSegmentPriceController.GetAllTicketSegmentPriceByFromToRouteSegmentId)
		TicketSegmentPriceRouterPrivateGroup.GET("get-all", controller.TicketSegmentPriceController.GetAllTicketSegmentPrice)
	}
}
