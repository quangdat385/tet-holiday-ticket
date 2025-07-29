package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

type (
	IMessageService interface {
		GetMessageById(context context.Context, id int64) (out model.MessageOutput, err error)
		GetMessagesByConversationId(context context.Context, conversationId int64, limit, offset int32) (out []model.MessageOutput, err error)
		GetMessageByUserId(context context.Context, userId int64, limit, offset int32) (out []model.MessageOutput, err error)
		CreateMessage(context context.Context, in model.MessageInput) (out model.MessageOutput, err error)
		UpdateMessageStatus(context context.Context, status bool, ID int64) (out model.MessageOutput, err error)
		DeleteMessage(context context.Context, ID int64) (err error)
	}
)

var (
	localMessageService IMessageService
)

func MessageService() IMessageService {
	if localMessageService == nil {
		panic("implement localMessageService not found for interface MessageService")
	}
	return localMessageService
}
func InitMessageService(i IMessageService) {
	localMessageService = i
}
