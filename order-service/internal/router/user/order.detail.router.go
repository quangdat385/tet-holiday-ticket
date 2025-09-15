package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/order-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/order-service/internal/middleware"
)

type OrderDetailRouter struct {
}

func (p *OrderDetailRouter) InitOrderDetailRoter(Router *gin.RouterGroup) {
	OrderDetailRouterPublicGroup := Router.Group("order/detail")
	OrderDetailRouterPublicGroup.Use()
	{

	}
	OrderDetailRouterPrivateGroup := Router.Group("order/detail")
	OrderDetailRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("User"))
	{
		OrderDetailRouterPrivateGroup.GET("/get-by-id/:id", controller.OrderDetailController.GetOrderDetailByID)
		OrderDetailRouterPrivateGroup.POST("/create", controller.OrderDetailController.CreateOrderDetail)
		OrderDetailRouterPrivateGroup.PUT("/update/:id", controller.OrderDetailController.UpdateOrderDetail)
	}
}
