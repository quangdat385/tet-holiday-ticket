package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

type (
	INotificationService interface {
		GetNotificationById(context context.Context, id int64) (out model.NotificationOutput, err error)
		GetNotificationsByUserIDTo(context context.Context, userId int64) (out []model.NotificationOutput, err error)
		GetNotificationsByUserIDFrom(context context.Context, userId int64) (out model.NotificationOutput, err error)
		GetNotificationsFromUserIDToIsNull(context context.Context) (out []model.NotificationOutput, err error)
		CreateNotification(context context.Context, input model.NotificationInput) (out model.NotificationOutput, err error)
		DeleteNotificationById(context context.Context, id int64) (err error)
		UpdateNotification(context context.Context, input model.UpdateNotificationInput) (out bool, err error)
	}
)

var (
	localNotificationService INotificationService
)

func NotificationService() INotificationService {
	if localNotificationService == nil {
		panic("implement localNotificationService not found for interface NotificationService")
	}
	return localNotificationService
}
func InitNotificationService(i INotificationService) {
	localNotificationService = i
}
