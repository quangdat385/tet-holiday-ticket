package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model/mapper"
)

type sMediaService struct {
	r *database.Queries
}

func NewMediaServiceImpl(r *database.Queries) *sMediaService {
	return &sMediaService{r: r}
}

func (s *sMediaService) GetMediaById(ctx context.Context, id int64) (out model.MediaOutput, err error) {
	media, err := s.r.GetCommunicationMediaById(ctx, id)
	if err != nil {
		return out, err
	}
	if media.ID == 0 {
		return out, errors.New("media not found")
	}
	out = mapper.ToMediaDTO(media)
	return out, nil
}

func (s *sMediaService) GetMediaByMessageId(ctx context.Context, messageId int64) (out []model.MediaOutput, err error) {
	items, err := s.r.GetCommunicationMediaByMessageId(ctx, sql.NullInt64{Int64: messageId, Valid: true})
	if err != nil {
		return out, err
	}
	for _, item := range items {
		out = append(out, mapper.ToMediaDTO(item))
	}
	return out, nil
}

func (s *sMediaService) GetMediaByConversationId(ctx context.Context, conversationId int64, limit, offset int32) (out []model.MediaOutput, err error) {
	items, err := s.r.GetCommunicationMediaByConversationId(ctx, database.GetCommunicationMediaByConversationIdParams{
		ConversationID: sql.NullInt64{Int64: conversationId, Valid: true},
		Limit:          limit,
		Offset:         offset,
	})
	if err != nil {
		return out, err
	}
	for _, item := range items {
		out = append(out, mapper.ToMediaDTO(item))
	}
	return out, nil
}

func (s *sMediaService) CreateMedia(ctx context.Context, in model.MediaInput) (out model.MediaOutput, err error) {
	params := database.InsertCommunicationMediaParams{
		UserID:       in.UserID,
		FileName:     in.FileName,
		OriginalName: in.OriginalName,
		MimeType:     in.MimeType,
		Size:         in.Size,
		URL:          in.URL,
	}
	if in.MessageID != nil {
		params.MessageID = sql.NullInt64{Int64: *in.MessageID, Valid: true}
	}
	if in.ConversationID != nil {
		params.ConversationID = sql.NullInt64{Int64: *in.ConversationID, Valid: true}
	}
	result, err := s.r.InsertCommunicationMedia(ctx, params)
	if err != nil {
		return out, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	newMedia, err := s.r.GetCommunicationMediaById(ctx, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToMediaDTO(newMedia)
	return out, nil
}

func (s *sMediaService) DeleteMedia(ctx context.Context, id int64) (err error) {
	media, err := s.r.GetCommunicationMediaById(ctx, id)
	if err != nil {
		return err
	}
	if media.ID == 0 {
		return errors.New("media not found")
	}
	_, err = s.r.DeleteCommunicationMedia(ctx, id)
	return err
}
