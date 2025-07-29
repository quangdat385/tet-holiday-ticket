package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/middleware"
)

type TrainRouter struct {
}

func (r *TrainRouter) InitTrainRouter(Router *gin.RouterGroup) {
	TrainRouterPublicGroup := Router.Group("train")
	TrainRouterPublicGroup.Use()
	{
	}
	TrainRouterPrivateGroup := Router.Group("train")
	TrainRouterPrivateGroup.Use(middleware.AuthenMiddleWare())
	{
		TrainRouterPrivateGroup.POST("create", middleware.RoleMiddleware("Admin"), controller.TrainController.CreateTrain)
		TrainRouterPrivateGroup.PATCH("update/:train_id", middleware.RoleMiddleware("Admin"), controller.TrainController.UpdateTrain)
		TrainRouterPrivateGroup.PATCH("update-status/:train_id", middleware.RoleMiddleware("Admin"), controller.TrainController.UpdateTrainStatus)
		TrainRouterPrivateGroup.DELETE("delete/:train_id", middleware.RoleMiddleware("Admin"), controller.TrainController.DeleteTrain)
	}
}
