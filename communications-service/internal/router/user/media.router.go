package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
)

type MediaRouter struct{}

func (m *MediaRouter) InitMediaRouter(Router *gin.RouterGroup) {
	MediaRouterPrivateGroup := Router.Group("media")
	MediaRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("User"))
	{
		MediaRouterPrivateGroup.POST("upload", controller.MediaControllerRouter.UploadMedia)
		MediaRouterPrivateGroup.GET("get-by-id/:id", controller.MediaControllerRouter.GetMediaByID)
		MediaRouterPrivateGroup.GET("message/:message_id", controller.MediaControllerRouter.GetMediaByMessageID)
		MediaRouterPrivateGroup.GET("conversation", controller.MediaControllerRouter.GetMediaByConversationID)
		MediaRouterPrivateGroup.DELETE("delete/:id", controller.MediaControllerRouter.DeleteMedia)
	}
}
