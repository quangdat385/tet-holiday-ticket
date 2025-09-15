package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/order-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/order-service/internal/middleware"
)

type OrderRouter struct {
}

func (p *OrderRouter) InitOrderRoter(Router *gin.RouterGroup) {
	OrderRouterPublicGroup := Router.Group("order")
	OrderRouterPublicGroup.Use()
	{

	}
	OrderRouterPrivateGroup := Router.Group("order")
	OrderRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("User"))
	{
		OrderRouterPrivateGroup.GET("get-by-id/:order_id", controller.OrderController.GetOrderByID)
		OrderRouterPrivateGroup.PUT("update/:order_id", controller.OrderController.UpdateOrder)
		OrderRouterPrivateGroup.GET("get-by-order-number/:order_number", controller.OrderController.GetOrderByOrderNumber)
		OrderRouterPrivateGroup.GET("get-by-user-id/:user_id", controller.OrderController.GetOrdersByUserID)
	}
}
