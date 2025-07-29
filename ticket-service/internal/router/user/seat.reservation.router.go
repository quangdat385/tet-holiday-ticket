package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type SeatReservationRouter struct {
}

func (r *SeatReservationRouter) InitSeatReservationRouter(Router *gin.RouterGroup) {
	SeatReservationRouterPublicGroup := Router.Group("seat-reservation")
	SeatReservationRouterPublicGroup.Use()
	{
	}
	SeatReservationRouterPrivateGroup := Router.Group("seat-reservation")
	SeatReservationRouterPrivateGroup.Use(middleware.AuthenMiddleWare())
	{
		SeatReservationRouterPrivateGroup.GET("get-by-id/:seat_reservation_id", controller.SeatReservationController.GetSeatReservationById)
		SeatReservationRouterPrivateGroup.GET("get-all-by-order-number", controller.SeatReservationController.GetAllSeatReservationsByOrderNumber)
		SeatReservationRouterPrivateGroup.GET("get-all-by-train-id", controller.SeatReservationController.GetAllSeatReservationsByTrainId)
	}
}
