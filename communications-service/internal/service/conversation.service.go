package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

type (
	IConversationService interface {
		GetConversationById(context context.Context, id int64) (out model.ConversationOutput, err error)
		GetConversationByUserId(context context.Context, userId int64) (out []model.ConversationOutput, err error)
		CreateConversation(context context.Context, in model.ConversationInput) (out model.ConversationOutput, err error)
		UpdateConversation(context context.Context, in model.UpdateConversationInput) (out model.ConversationOutput, err error)
		AddUserToConversation(context context.Context, conversationId int64, userIds []int64) (out model.ConversationOutput, err error)
		RemoveUserFromConversation(context context.Context, conversationId int64, userId int64) (out model.ConversationOutput, err error)
		DeleteConversation(context context.Context, id int64) (err error)
	}
)

var (
	localConversationService IConversationService
)

func ConversationService() IConversationService {
	if localConversationService == nil {
		panic("implement localConversationService not found for interface ConversationService")
	}
	return localConversationService
}
func InitConversationService(i IConversationService) {
	localConversationService = i
}
