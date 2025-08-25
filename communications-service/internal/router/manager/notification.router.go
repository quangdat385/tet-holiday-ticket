package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
)

type NotificationRouter struct {
}

func (r *NotificationRouter) InitNotificationRouter(Router *gin.RouterGroup) {
	NotificationRouterPublicGroup := Router.Group("notification")
	NotificationRouterPublicGroup.Use()
	{
		// Define public routes here
	}
	NotificationRouterPrivateGroup := Router.Group("notification")
	NotificationRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("Admin", "Manager"))
	{
		NotificationRouterPrivateGroup.DELETE("/delete-by-id/:id", controller.NotificationControllerRouter.DeleteNotification)

	}
}
