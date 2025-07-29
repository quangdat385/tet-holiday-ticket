package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type SeatRouter struct {
}

func (r *SeatRouter) InitSeatRouter(Router *gin.RouterGroup) {
	SeatRouterPublicGroup := Router.Group("seat")
	SeatRouterPublicGroup.Use()
	{
	}
	SeatRouterPrivateGroup := Router.Group("seat")
	SeatRouterPrivateGroup.Use(middleware.AuthenMiddleWare())
	{
		SeatRouterPrivateGroup.POST("create", middleware.RoleMiddleware("Admin"), controller.SeatController.CreateSeat)
		SeatRouterPrivateGroup.PATCH("update/:seat_id", middleware.RoleMiddleware("Admin"), controller.SeatController.UpdateSeat)
		SeatRouterPrivateGroup.DELETE("delete/:seat_id", middleware.RoleMiddleware("Admin"), controller.SeatController.DeleteSeat)
	}
}
