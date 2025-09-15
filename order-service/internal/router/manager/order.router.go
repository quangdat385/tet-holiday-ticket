package manager

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
	OrderRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("Admin"))
	{
		OrderRouterPrivateGroup.POST("create", controller.OrderController.CreateOrder)
		OrderRouterPrivateGroup.DELETE("delete/:id", controller.OrderController.DeleteOrder)
	}
}
