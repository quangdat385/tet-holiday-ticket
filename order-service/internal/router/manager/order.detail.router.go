package manager

import "github.com/gin-gonic/gin"

type OrderDetailRouter struct {
}

func (p *OrderDetailRouter) InitOrderDetailRoter(Router *gin.RouterGroup) {
	OrderDetailRouterPublicGroup := Router.Group("order/detail")
	OrderDetailRouterPublicGroup.Use()
	{

	}
	OrderDetailRouterPrivateGroup := Router.Group("order/detail")
	OrderDetailRouterPrivateGroup.Use()
	{

	}
}
