package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type SeatRouter struct {
}

func (r *SeatRouter) InitSeatRouter(Router *gin.RouterGroup) {
	SeatRouterPublicGroup := Router.Group("seat")
	SeatRouterPublicGroup.Use()
	{
		SeatRouterPublicGroup.GET("get-one/:seat_id", controller.SeatController.GetSeat)
		SeatRouterPublicGroup.GET("get-by-train/:train_id", controller.SeatController.GetSeatsByTrain)
	}
	SeatRouterPrivateGroup := Router.Group("seat")
	SeatRouterPrivateGroup.Use()
	{
	}
}
