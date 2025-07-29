package manager

import "github.com/gin-gonic/gin"

type PaymentRouter struct {
}

func (p *PaymentRouter) InitPaymentRouter(Router *gin.RouterGroup) {
	PaymentRouterPublicGroup := Router.Group("payment")
	{
		PaymentRouterPublicGroup.POST("create")
	}
	PaymentRouterPrivateGroup := Router.Group("payment")

	{
		PaymentRouterPrivateGroup.POST("create")
	}
}
