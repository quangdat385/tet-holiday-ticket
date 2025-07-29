package user

import "github.com/gin-gonic/gin"

type MessageRouter struct {
}

func (m *MessageRouter) InitMessageRoter(Router *gin.RouterGroup) {
	MessageRouterPublicGroup := Router.Group("message")
	MessageRouterPublicGroup.Use()
	{
	}
	MessageRouterPrivateGroup := Router.Group("message")
	MessageRouterPrivateGroup.Use()

	{
	}
}
