package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
)

type MessageRouter struct {
}

func (m *MessageRouter) InitMessageRoter(Router *gin.RouterGroup) {
	MessageRouterPublicGroup := Router.Group("message")
	MessageRouterPublicGroup.Use()
	{
	}
	MessageRouterPrivateGroup := Router.Group("message")
	MessageRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("User"))
	{
		MessageRouterPrivateGroup.POST("create", controller.MessageControllerRouter.CreateMessage)
		MessageRouterPrivateGroup.GET("conversation/:conversation_id", controller.MessageControllerRouter.GetMessagesByConversationID)
		MessageRouterPrivateGroup.PATCH("update/:id", controller.MessageControllerRouter.UpdateMessageStatus)
		MessageRouterPrivateGroup.DELETE("delete/:id", controller.MessageControllerRouter.DeleteMessage)
		MessageRouterPrivateGroup.GET("get-by-id/:id", controller.MessageControllerRouter.GetMessageByID)
		MessageRouterPrivateGroup.GET("user", controller.MessageControllerRouter.GetMessagesByUserID)
	}
}
