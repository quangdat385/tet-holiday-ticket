package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/middleware"
)

type InformationRouter struct{}

func (r *InformationRouter) InitInformationRouter(Router *gin.RouterGroup) {
	InformationRouterPublicGroup := Router.Group("information")
	InformationRouterPublicGroup.Use()
	{
		// Define public routes here
	}

	InformationRouterPrivateGroup := Router.Group("information")
	InformationRouterPrivateGroup.Use(middleware.AuthenMiddleWare(), middleware.RoleMiddleware("Admin", "Manager"))
	{
		InformationRouterPrivateGroup.DELETE("delete/:information_id", controller.InformationRouter.DeleteInformationByID)
	}
}
