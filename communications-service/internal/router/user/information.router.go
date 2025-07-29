package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/controller"
)

type UserInformationRouter struct {
}

func (r *UserInformationRouter) InitUserInformationRouter(Router *gin.RouterGroup) {
	UserInformationRouterPublicGroup := Router.Group("information")
	UserInformationRouterPublicGroup.Use()
	{
		// Define public routes here
	}

	UserInformationRouterPrivateGroup := Router.Group("information")
	UserInformationRouterPrivateGroup.Use()
	{
		UserInformationRouterPrivateGroup.GET("get-by-user-id/:user_id", controller.InformationRouter.GetInformationByUserID)
		UserInformationRouterPrivateGroup.PUT("update-by-user-id/:user_id", controller.InformationRouter.UpdateInformationByUserID)
		UserInformationRouterPrivateGroup.POST("create", controller.InformationRouter.InsertInformationByUserID)
		UserInformationRouterPrivateGroup.DELETE("delete/:information_id", controller.InformationRouter.DeleteInformationByID)
	}
}
