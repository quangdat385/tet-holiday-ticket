package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model/mapper"
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
func (s *sConversation) GetConversationByUserId(context context.Context, userId int64) (out []model.ConversationOutput, err error) {
	userIdsJSON, err := json.Marshal([]int64{userId})
	if err != nil {
		return out, err
	}
	conversations, err := s.r.GetCommunicationConversationsByUserId(context, database.GetCommunicationConversationsByUserIdParams{
		JSONQUOTE: string(userIdsJSON),
		Limit:     50,
		Offset:    0,
	})
	if err != nil {
		return out, err
	}
	for _, conversation := range conversations {
		out = append(out, mapper.ToConversationDTO(conversation))
	}
	return out, nil
}
func (s *sConversation) CreateConversation(context context.Context, in model.ConversationInput) (out model.ConversationOutput, err error) {
	userIdsJSON, err := json.Marshal(in.UserIDS)
	if err != nil {
		return out, err
	}
	result, err := s.r.InsertCommunicationConversation(context, database.InsertCommunicationConversationParams{
		Title:       sql.NullString{String: in.Title, Valid: true},
		UserIds:     userIdsJSON,
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
	var userIdsJSON []byte
	if len(in.UserIDS) != 0 {
		userIdsJSON, err = json.Marshal(in.UserIDS)
		if err != nil {
			return out, err
		}
	}
	if err != nil {
		return out, err
	}
	if conversation.IsDeleted.Bool != in.IsDeleted {
		conversation.IsDeleted.Bool = in.IsDeleted
	}
	if in.Title != "" {
		conversation.Title = sql.NullString{String: in.Title, Valid: true}
	}
	_, err = s.r.UpdateCommunicationConversation(context, database.UpdateCommunicationConversationParams{
		ID:              in.ID,
		Title:           sql.NullString{String: conversation.Title.String, Valid: true},
		JSONARRAYAPPEND: string(userIdsJSON),
		Description:     sql.NullString{String: in.Description, Valid: in.Description != ""},
		Type:            sql.NullString{String: in.Type, Valid: in.Type != ""},
		Background:      sql.NullString{String: in.Background, Valid: in.Background != ""},
		Emoji:           sql.NullString{String: in.Emoji, Valid: in.Emoji != ""},
		IsDeleted:       sql.NullBool{Bool: conversation.IsDeleted.Bool, Valid: true},
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
func (s *sConversation) AddUserToConversation(context context.Context, conversationId int64, userIds []int64) (out model.ConversationOutput, err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, conversationId)
	if err != nil {
		return out, err
	}
	if conversation.ID == 0 {
		return out, errors.New("conversation not found")
	}
	userIdsJSON, err := json.Marshal(userIds)
	if err != nil {
		return out, err
	}
	_, err = s.r.AddUserToCommunicationConversation(context, database.AddUserToCommunicationConversationParams{
		ID:              conversationId,
		JSONARRAYAPPEND: string(userIdsJSON),
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
func (s *sConversation) RemoveUserFromConversation(context context.Context, conversationId int64, userId int64) (out model.ConversationOutput, err error) {
	conversation, err := s.r.GetCommunicationConversationById(context, conversationId)
	if err != nil {
		return out, err
	}
	if conversation.ID == 0 {
		return out, errors.New("conversation not found")
	}
	_, err = s.r.RemoveUserFromCommunicationConversation(context, database.RemoveUserFromCommunicationConversationParams{
		ID:         conversationId,
		JSONSEARCH: fmt.Sprint(userId),
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
