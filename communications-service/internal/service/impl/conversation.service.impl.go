package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/vo"
)

type sConversation struct {
	r *database.Queries
}

func NewConversationImpl(r *database.Queries) *sConversation {
	return &sConversation{
		r: r,
	}
}
func (s *sConversation) GetConversationById(context context.Context, id int64) (out model.ConversationOutput, err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, id)
	if err != nil {
		return out, err
	}
	if conversation.ID == 0 {
		return out, errors.New("conversation not found")
	}
	out = mapper.ToConversationDTO(conversation)
	return out, nil
}
func (s *sConversation) GetConversationByUserId(context context.Context, user_id int64, query vo.PageRequest) (out []model.ConversationOutput, err error) {
	conversations, err := s.r.GetListCommunicationConversations(context, database.GetListCommunicationConversationsParams{
		UserID: user_id,
		Limit:  int32(query.Limit),
		Offset: int32((query.Page - 1) * query.Limit),
	})
	if err != nil {
		return out, err
	}
	out = mapper.ToConversationDTOBYUserID(conversations)
	return out, nil
}
func (s *sConversation) CreateConversation(context context.Context, in model.ConversationInput) (out model.ConversationOutput, err error) {
	tx, err := global.Mdb.BeginTx(context, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return out, err
	}
	defer func() {
		if p := recover(); p != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	s.r = s.r.WithTx(tx)
	result, err := s.r.InsertCommunicationConversation(context, database.InsertCommunicationConversationParams{
		Title:       sql.NullString{String: in.Title, Valid: true},
		Description: sql.NullString{String: in.Description, Valid: true},
		Type:        sql.NullString{String: in.Type, Valid: true},
		Background:  sql.NullString{String: in.Background, Valid: true},
		Emoji:       sql.NullString{String: in.Emoji, Valid: true},
		IsDeleted:   sql.NullBool{Bool: in.IsDeleted, Valid: true},
	})
	if err != nil {
		return out, err
	}
	ID, err := result.LastInsertId()
	if err != nil {
		return out, err
	}

	for _, user := range in.UserIDS {
		_, err := s.r.AddUserToCommunicationConversation(context, database.AddUserToCommunicationConversationParams{
			ConversationID: ID,
			UserID:         user.UserID,
			NickName:       sql.NullString{String: user.NickName, Valid: true},
		})
		if err != nil {
			return out, err
		}
	}
	newConversation, err := s.r.GetCommunicationConversationById(context, ID)
	if err != nil {
		return out, err
	}
	out = mapper.ToConversationDTO(newConversation)
	return out, nil
}
func (s *sConversation) UpdateConversation(context context.Context, in model.UpdateConversationInput) (out model.ConversationOutput, err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, in.ID)
	if err != nil {
		return out, err
	}
	if conversation.ID == 0 {
		return out, errors.New("conversation not found")
	}
	if conversation.IsDeleted.Bool != in.IsDeleted {
		conversation.IsDeleted.Bool = in.IsDeleted
	}
	if in.Title != "" {
		conversation.Title = sql.NullString{String: in.Title, Valid: true}
	}
	_, err = s.r.UpdateCommunicationConversation(context, database.UpdateCommunicationConversationParams{
		ID:          in.ID,
		Title:       sql.NullString{String: conversation.Title.String, Valid: true},
		Description: sql.NullString{String: in.Description, Valid: in.Description != ""},
		Type:        sql.NullString{String: in.Type, Valid: in.Type != ""},
		Background:  sql.NullString{String: in.Background, Valid: in.Background != ""},
		Emoji:       sql.NullString{String: in.Emoji, Valid: in.Emoji != ""},
		IsDeleted:   sql.NullBool{Bool: conversation.IsDeleted.Bool, Valid: true},
	})
	if err != nil {
		return out, err
	}
	newConversation, err := s.r.GetCommunicationConversationById(context, in.ID)
	if err != nil {
		return out, err
	}
	out = mapper.ToConversationDTO(newConversation)
	return out, nil
}
func (s *sConversation) SoftDeleteConversation(context context.Context, in model.SoftDeleteConversationInput) (out bool, err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, in.ID)
	if err != nil {
		return out, err
	}
	if conversation.ID == 0 {
		return out, errors.New("conversation not found")
	}
	_, err = s.r.SoftDeleteCommunicationConversation(context, in.ID)
	if err != nil {
		return out, err
	}
	out = true
	return out, nil
}
func (s *sConversation) AddUserToConversation(context context.Context, conversationId int64, userIds []model.UserIDS) (out model.ConversationOutput, err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, conversationId)
	if err != nil {
		return out, err
	}
	if conversation.ID == 0 {
		return out, errors.New("conversation not found")
	}
	if conversation.Type.String == "personal" {
		return out, errors.New("cannot add user to personal conversation")
	}
	for _, userId := range userIds {
		if err != nil {
			return out, err
		}
		_, err = s.r.AddUserToCommunicationConversation(context, database.AddUserToCommunicationConversationParams{
			ConversationID: conversationId,
			UserID:         userId.UserID,
			NickName:       sql.NullString{String: userId.NickName, Valid: true},
		})
		if err != nil {
			return out, err
		}
	}
	newConversation, err := s.r.GetCommunicationConversationById(context, conversationId)
	if err != nil {
		return out, err
	}
	out = mapper.ToConversationDTO(newConversation)
	return out, nil
}
func (s *sConversation) RemoveUserFromConversation(context context.Context, conversationId int64, userId int64) (out model.ConversationOutput, err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, conversationId)
	if err != nil {
		return out, err
	}
	if conversation.ID == 0 {
		return out, errors.New("conversation not found")
	}
	if conversation.Type.String == "personal" {
		return out, errors.New("cannot remove user from personal conversation")
	}
	_, err = s.r.RemoveUserFromCommunicationConversation(context, database.RemoveUserFromCommunicationConversationParams{
		ConversationID: conversationId,
		UserID:         userId,
	})
	if err != nil {
		return out, err
	}
	newConversation, err := s.r.GetCommunicationConversationById(context, conversationId)
	if err != nil {
		return out, err
	}
	out = mapper.ToConversationDTO(newConversation)
	return out, nil
}
func (s *sConversation) DeleteConversation(context context.Context, id int64) (err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, id)
	if err != nil {
		return err
	}
	if conversation.ID == 0 {
		return errors.New("conversation not found")
	}
	_, err = s.r.DeleteCommunicationConversation(context, id)
	if err != nil {
		return err
	}
	return nil
}
