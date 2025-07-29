package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type StationRouter struct {
}

func (r *StationRouter) InitStationRouter(Router *gin.RouterGroup) {
	StationRouterPublicGroup := Router.Group("station")
	StationRouterPublicGroup.Use()
	{
	}
	StationRouterPrivateGroup := Router.Group("station")
	StationRouterPrivateGroup.Use(middleware.AuthenMiddleWare())
	{
		StationRouterPrivateGroup.POST("create", middleware.RoleMiddleware("Admin"), controller.StationController.CreateStation)
		StationRouterPrivateGroup.PATCH("update/:station_id", middleware.RoleMiddleware("Admin"), controller.StationController.UpdateStation)
		StationRouterPrivateGroup.PATCH("update-status/:station_id", middleware.RoleMiddleware("Admin"), controller.StationController.UpdateStationStatus)
		StationRouterPrivateGroup.DELETE("delete/:station_id", middleware.RoleMiddleware("Admin"), controller.StationController.DeleteStation)
	}
}
