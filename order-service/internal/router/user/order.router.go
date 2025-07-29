package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/order-service/internal/controller"
)

type OrderRouter struct {
}

func (p *OrderRouter) InitOrderRoter(Router *gin.RouterGroup) {
	OrderRouterPublicGroup := Router.Group("order")
	OrderRouterPublicGroup.Use()
	{

	}
	OrderRouterPrivateGroup := Router.Group("order")
	OrderRouterPrivateGroup.Use()
	{
		OrderRouterPrivateGroup.POST("create", controller.OrderController.CreateOrder)
		OrderRouterPrivateGroup.GET(":id", controller.OrderController.GetOrderByID)
		OrderRouterPrivateGroup.PUT("update/:id", controller.OrderController.UpdateOrder)
		OrderRouterPrivateGroup.DELETE("delete/:id", controller.OrderController.DeleteOrder)

	}
}
