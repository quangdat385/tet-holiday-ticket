package manager

import "github.com/gin-gonic/gin"

type NotificationRouter struct {
}

func (r *NotificationRouter) InitNotificationRouter(Router *gin.RouterGroup) {
	NotificationRouterPublicGroup := Router.Group("notification")
	NotificationRouterPublicGroup.Use()
	{
		// Define public routes here
	}
	NotificationRouterPrivateGroup := Router.Group("notification")
	NotificationRouterPrivateGroup.Use()
	{
		// Define private routes here
	}
}
