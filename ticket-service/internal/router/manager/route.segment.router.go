package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type RouteSegmentRouter struct {
}

func (p *RouteSegmentRouter) InitRouteSegmentRouter(Router *gin.RouterGroup) {
	RouteSegmentRouterPublicGroup := Router.Group("route-segment")
	RouteSegmentRouterPublicGroup.Use()
	{
	}
	RouteSegmentRouterPrivateGroup := Router.Group("route-segment")
	RouteSegmentRouterPrivateGroup.Use(middleware.AuthenMiddleWare())
	{
		RouteSegmentRouterPublicGroup.POST("create", middleware.RoleMiddleware("Admin"), controller.RouteSegmenController.CreateRouteSegment)
		RouteSegmentRouterPublicGroup.PATCH("update/:segment_id", middleware.RoleMiddleware("Admin"), controller.RouteSegmenController.UpdateRouteSegment)
		RouteSegmentRouterPublicGroup.DELETE("delete/:segment_id", middleware.RoleMiddleware("Admin"), controller.RouteSegmenController.DeleteRouteSegment)
	}
}
