package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

type (
	IInformationService interface {
		GetInformationByUserID(contect context.Context, user_id int64) (out model.InformationOutput, err error)
		UpdateInformationByUserID(contect context.Context, input model.InfomationInput) (out string, err error)
		InsertInformationByUserID(contect context.Context, input model.InfomationInput) (out string, err error)
		DeleteInformationByID(contect context.Context, id int64) (out string, err error)
		SetUserConnected(contect context.Context, user_id int64) (out string, err error)
		DeleteUserConnected(context context.Context, user_id int64) (out string, err error)
		GetUserConnectedExists(context context.Context, user_id int64) (out bool, err error)
	}
)

var (
	localInformationService IInformationService
)

func InformationService() IInformationService {
	if localInformationService == nil {
		panic("implement localInformationService not found for interface InformationService")
	}
	return localInformationService
}
func InitInformationService(i IInformationService) {
	localInformationService = i
}
