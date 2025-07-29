package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
)

type NotificationRouter struct {
}

func (n *NotificationRouter) InitNotificationRouter(Router *gin.RouterGroup) {
	NotificationRouterPublicGroup := Router.Group("notification")
	NotificationRouterPublicGroup.Use()
	{
		NotificationRouterPublicGroup.GET("/get-from-user-id-to-is-null", controller.NotificationControllerRouter.GetNotificationsByUserIDToNull)
	}

	NotificationRouterPrivateGroup := Router.Group("notification")
	NotificationRouterPrivateGroup.Use()
	{
		NotificationRouterPrivateGroup.GET("/get-by-id/:id", controller.NotificationControllerRouter.GetNotificationByID)
		NotificationRouterPrivateGroup.GET("/get-by-user-id-to/:userId", controller.NotificationControllerRouter.GetNotificationsByUserIDTo)
		NotificationRouterPrivateGroup.GET("/get-by-user-id-from/:userId", controller.NotificationControllerRouter.GetNotificationsByUserIDFrom)
		NotificationRouterPrivateGroup.POST("/create", controller.NotificationControllerRouter.CreateNotification)
		NotificationRouterPrivateGroup.DELETE("/delete-by-id/:id", controller.NotificationControllerRouter.DeleteNotification)

	}
}
