package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
)

type NotificationRouter struct {
}

func (n *NotificationRouter) InitNotificationRouter(Router *gin.RouterGroup) {
	NotificationRouterPublicGroup := Router.Group("notification")
	NotificationRouterPublicGroup.Use()
	{

	}

	NotificationRouterPrivateGroup := Router.Group("notification")
	NotificationRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("User"))
	{
		NotificationRouterPrivateGroup.GET("/get-by-id/:notification_id", controller.NotificationControllerRouter.GetNotificationByID)
		NotificationRouterPrivateGroup.GET("/get-by-user-id-to/:user_id", controller.NotificationControllerRouter.GetNotificationsByUserIDTo)
		NotificationRouterPrivateGroup.GET("/get-by-user-id-from/:user_id", controller.NotificationControllerRouter.GetNotificationsByUserIDFrom)
		NotificationRouterPrivateGroup.POST("/create", controller.NotificationControllerRouter.CreateNotification)
		NotificationRouterPrivateGroup.GET("/get-by-user-id-to-null", controller.NotificationControllerRouter.GetNotificationsByUserIDToNull)
	}
}
