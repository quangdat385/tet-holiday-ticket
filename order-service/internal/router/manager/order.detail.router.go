package manager

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
	OrderDetailRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("Admin"))
	{
		OrderDetailRouterPrivateGroup.DELETE("/delete/:id", controller.OrderDetailController.DeleteOrderDetail)
	}
}
