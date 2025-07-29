package manager

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
		SeatReservationRouterPrivateGroup.POST(
			"create",
			middleware.RoleMiddleware("Amind", "Manager", "Staff", "User"),
			controller.SeatReservationController.ReserveSeat)
		SeatReservationRouterPrivateGroup.DELETE(
			"cancel",
			middleware.RoleMiddleware("Admin", "Manager"),
			controller.SeatReservationController.CancelSeatReservation,
		)
		SeatReservationRouterPrivateGroup.PATCH(
			"update/:seat_reservation_id",
			middleware.RoleMiddleware("Amind", "Manager", "Staff", "User"),
			controller.SeatReservationController.UpdateSeatReservation)
	}
}
