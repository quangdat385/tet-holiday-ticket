package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

type (
	IMediaService interface {
		GetMediaById(ctx context.Context, id int64) (out model.MediaOutput, err error)
		GetMediaByMessageId(ctx context.Context, messageId int64) (out []model.MediaOutput, err error)
		GetMediaByConversationId(ctx context.Context, conversationId int64, limit, offset int32) (out []model.MediaOutput, err error)
		CreateMedia(ctx context.Context, in model.MediaInput) (out model.MediaOutput, err error)
		DeleteMedia(ctx context.Context, id int64) (err error)
	}
)

var (
	localMediaService IMediaService
)

func MediaService() IMediaService {
	if localMediaService == nil {
		panic("implement localMediaService not found for interface MediaService")
	}
	return localMediaService
}

func InitMediaService(i IMediaService) {
	localMediaService = i
}
