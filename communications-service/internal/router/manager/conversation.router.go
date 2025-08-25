package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
)

type ConversationRouter struct {
}

func (r *ConversationRouter) InitConversationRouter(Router *gin.RouterGroup) {
	ConversationRouterPublicGroup := Router.Group("conversation")
	ConversationRouterPublicGroup.Use()
	{

	}
	ConversationRouterPrivateGroup := Router.Group("conversation")
	ConversationRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("Admin", "Manager"))
	{
		ConversationRouterPrivateGroup.DELETE("delete/:conversation_id", controller.ConversationControllerRouter.DeleteConversation)
	}
}
