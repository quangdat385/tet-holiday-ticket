package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type RouteSegmentRouter struct {
}

func (r *RouteSegmentRouter) InitRouteSegmentRouter(Router *gin.RouterGroup) {
	RouteSegmentRouterPublicGroup := Router.Group("route-segment")
	RouteSegmentRouterPublicGroup.Use()
	{
		RouteSegmentRouterPublicGroup.GET("get-one/:segment_id", controller.RouteSegmenController.GetRouteSegment)
		RouteSegmentRouterPublicGroup.GET("get-by-train-id/:train_id", controller.RouteSegmenController.GetRouteSegmentsByTrainID)
		RouteSegmentRouterPublicGroup.GET("get-by-from-station-id/:from_station_id", controller.RouteSegmenController.GetRouteSegmentsByFromStationID)
		RouteSegmentRouterPublicGroup.GET("get-by-to-station-id/:to_station_id", controller.RouteSegmenController.GetRouteSegmentsByToStationID)

	}
	RouteSegmentRouterPrivateGroup := Router.Group("route-segment")
	RouteSegmentRouterPrivateGroup.Use()
	{
	}
}
