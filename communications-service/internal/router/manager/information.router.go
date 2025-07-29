package manager

import "github.com/gin-gonic/gin"

type InformationRouter struct{}

func (r *InformationRouter) InitInformationRouter(Router *gin.RouterGroup) {
	InformationRouterPublicGroup := Router.Group("information")
	InformationRouterPublicGroup.Use()
	{
		// Define public routes here
	}

	InformationRouterPrivateGroup := Router.Group("information")
	InformationRouterPrivateGroup.Use()
	{
		// Define private routes here
	}
}
