package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/controller"
)

type TrainRouter struct {
}

func (r *TrainRouter) InitTrainRouter(Router *gin.RouterGroup) {
	TrainRouterPublicGroup := Router.Group("train")
	TrainRouterPublicGroup.Use()
	{
		TrainRouterPublicGroup.GET("get-by-id/:train_id", controller.TrainController.GetTrainByID)
	}
	TrainRouterPrivateGroup := Router.Group("train")
	TrainRouterPrivateGroup.Use()
	{
	}
}
