package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
)

type ConversationRouter struct{}

func (c *ConversationRouter) InitConversationRouter(Router *gin.RouterGroup) {
	ConversationRouterPublicGroup := Router.Group("conversation")
	ConversationRouterPublicGroup.Use()
	{
		// Define public routes here
	}

	ConversationRouterPrivateGroup := Router.Group("conversation")
	ConversationRouterPrivateGroup.Use()
	{
		ConversationRouterPrivateGroup.POST("create", controller.ConversationControllerRouter.CreateConversation)
		ConversationRouterPrivateGroup.GET("get-by-id/{conversation_id}", controller.ConversationControllerRouter.GetConversationByID)
		ConversationRouterPrivateGroup.GET("get-by-user-id/{user_id}", controller.ConversationControllerRouter.GetConversationByUserID)
		ConversationRouterPrivateGroup.PUT("update/{conversation_id}", controller.ConversationControllerRouter.UpdateConversation)
		ConversationRouterPrivateGroup.PATCH("add-users/{conversation_id}", controller.ConversationControllerRouter.AddUserToConversation)
		ConversationRouterPrivateGroup.PATCH("remove-users/{conversation_id}/{user_id}", controller.ConversationControllerRouter.RemoveUserFromConversation)
		ConversationRouterPrivateGroup.DELETE("soft-delete/{conversation_id}", controller.ConversationControllerRouter.DeleteConversation)
	}
}
