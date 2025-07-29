package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
)

type ConversationRouter struct {
}

func (r *ConversationRouter) InitConversationRouter(Router *gin.RouterGroup) {
	ConversationRouterPublicGroup := Router.Group("conversation")
	ConversationRouterPublicGroup.Use()
	{

	}
	ConversationRouterPrivateGroup := Router.Group("conversation")
	ConversationRouterPrivateGroup.Use()
	{
		ConversationRouterPrivateGroup.DELETE("delete/{conversation_id}", controller.ConversationControllerRouter.DeleteConversation)
	}
}
