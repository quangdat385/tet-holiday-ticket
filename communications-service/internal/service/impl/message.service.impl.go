package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model/mapper"
)

type sMessageSevice struct {
	r *database.Queries
}

func NewMessageServiceImpl(r *database.Queries) *sMessageSevice {
	return &sMessageSevice{
		r: r,
	}
}
func (s *sMessageSevice) GetMessageById(context context.Context, id int64) (out model.MessageOutput, err error) {
	message, err := s.r.GetCommunicationMessageById(context, id)
	if err != nil {
		return out, err
	}
	if message.ID == 0 {
		return out, errors.New("message not found")
	}
	out = mapper.ToMessageDTO(message)
	return out, nil
}
func (s *sMessageSevice) GetMessagesByConversationId(context context.Context, conversationId int64, limit, offset int32) (out []model.MessageOutput, err error) {
	messages, err := s.r.GetCommunicationMessagesByConversationId(context, database.GetCommunicationMessagesByConversationIdParams{
		ConversationID: conversationId,
		Limit:          limit,
		Offset:         offset,
	})
	if err != nil {
		return out, err
	}
	for _, message := range messages {
		out = append(out, mapper.ToMessageDTO(message))
	}
	return out, nil
}
func (s *sMessageSevice) GetMessageByUserId(context context.Context, userId int64, limit, offset int32) (out []model.MessageOutput, err error) {
	messages, err := s.r.GetCommunicationMessagesByUserId(context, database.GetCommunicationMessagesByUserIdParams{
		UserID: userId,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return out, err
	}
	for _, message := range messages {
		out = append(out, mapper.ToMessageDTO(message))
	}
	return out, nil
}
func (s *sMessageSevice) CreateMessage(context context.Context, in model.MessageInput) (out model.MessageOutput, err error) {
	result, err := s.r.InsertCommunicationMessage(context, database.InsertCommunicationMessageParams{
		ConversationID: in.ConversationID,
		UserID:         in.UserID,
		Type:           in.Type,
		Status:         sql.NullBool{Bool: in.Status, Valid: true},
		Message:        in.Message,
	})
	if err != nil {
		return out, err
	}
	ID, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	newMessage, err := s.r.GetCommunicationMessageById(context, ID)
	if err != nil {
		return out, err
	}
	out = mapper.ToMessageDTO(newMessage)
	return out, nil
}
func (s *sMessageSevice) UpdateMessageStatus(context context.Context, status bool, ID int64) (out model.MessageOutput, err error) {
	message, err := s.r.GetCommunicationMessageById(context, ID)
	if err != nil {
		return out, err
	}
	if message.ID == 0 {
		return out, errors.New("message not found")
	}
	_, err = s.r.UpdateCommunicationMessage(context, database.UpdateCommunicationMessageParams{
		ID:     ID,
		Status: sql.NullBool{Bool: status, Valid: true},
	})
	if err != nil {
		return out, err
	}
	message.Status = sql.NullBool{Bool: status, Valid: true}
	out = mapper.ToMessageDTO(message)
	return out, nil
}
func (s *sMessageSevice) DeleteMessage(context context.Context, ID int64) (err error) {
	message, err := s.r.GetCommunicationMessageById(context, ID)
	if err != nil {
		return err
	}
	if message.ID == 0 {
		return errors.New("message not found")
	}
	_, err = s.r.DeleteCommunicationMessage(context, ID)
	if err != nil {
		return err
	}
	return nil
}
