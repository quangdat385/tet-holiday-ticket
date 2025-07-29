package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type StationRouter struct {
}

func (r *StationRouter) InitStationRouter(Router *gin.RouterGroup) {
	StationRouterPublicGroup := Router.Group("station")
	StationRouterPublicGroup.Use()
	{
		StationRouterPublicGroup.GET("get-by-id/:station_id", controller.StationController.GetStationByID)
		StationRouterPublicGroup.GET("get-all", controller.StationController.GetAllStation)
	}
	StationRouterPrivateGroup := Router.Group("station")
	StationRouterPrivateGroup.Use()
	{
	}
}
