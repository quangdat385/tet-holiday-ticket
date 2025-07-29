package manager

import "github.com/gin-gonic/gin"

type OrderRouter struct {
}

func (p *OrderRouter) InitOrderRoter(Router *gin.RouterGroup) {
	OrderRouterPublicGroup := Router.Group("message")
	OrderRouterPublicGroup.Use()
	{
	}
	OrderRouterPrivateGroup := Router.Group("message")
	OrderRouterPrivateGroup.Use()
	{
	}
}
