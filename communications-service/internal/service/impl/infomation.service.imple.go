package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/utils/crypto"
)

type sInformation struct {
	r                *database.Queries
	distributedCache service.IRedisCache
}

func NewInformationImpl(r *database.Queries, redisCahe service.IRedisCache) *sInformation {
	return &sInformation{
		r:                r,
		distributedCache: redisCahe,
	}
}
func (s *sInformation) GetInformationByUserID(context context.Context, user_id int64) (out model.InformationOutput, err error) {
	information, err := s.r.GetCommunicationInfoByUserID(context, user_id)
	if err != nil {
		return out, err
	}
	out = mapper.ToInformationDTO(information)
	return out, nil
}
func (s *sInformation) UpdateInformationByUserID(context context.Context, input model.InfomationInput) (out string, err error) {
	information, err := s.r.GetCommunicationInfoByUserID(context, input.UserID)
	if err != nil {
		return out, err
	}
	if information.ID == 0 {
		return out, errors.New("information not found")
	}
	if input.Status != information.Status.Bool {
		information.Status.Bool = input.Status
	}
	if input.Type != "" {
		information.Type.String = input.Type
	}
	if input.Value != "" {
		information.Value = input.Value
	}

	_, err = s.r.UpdateCommunicationInfoByUserId(context, database.UpdateCommunicationInfoByUserIdParams{
		UserID: input.UserID,
		Status: sql.NullBool{Bool: information.Status.Bool, Valid: true},
		Type:   sql.NullString{String: information.Type.String, Valid: true},
		Value:  information.Value,
	})
	if err != nil {
		return out, err
	}
	out = "Update information successfully"
	return out, nil
}
func (s *sInformation) InsertInformationByUserID(context context.Context, input model.InfomationInput) (out string, err error) {
	_, err = s.r.InsertCommunicationInfo(context, database.InsertCommunicationInfoParams{
		UserID: input.UserID,
		Status: sql.NullBool{Bool: input.Status, Valid: true},
		Type:   sql.NullString{String: input.Type, Valid: true},
		Value:  input.Value,
	})
	if err != nil {
		return out, err
	}
	out = "Insert information successfully"
	return out, nil
}
func (s *sInformation) DeleteInformationByID(context context.Context, id int64) (out string, err error) {
	information, err := s.r.GetCommunicationInfoByID(context, id)
	if err != nil {
		return out, err
	}
	if information.ID == 0 {
		return out, errors.New("information not found")
	}
	_, err = s.r.DeleteCommunicationInfo(context, id)
	if err != nil {
		return out, err
	}
	out = "Delete information successfully"
	return out, nil
}
func (s *sInformation) SetUserConnected(context context.Context, user_id int64) (out string, err error) {
	hashKey := crypto.GenerateHash(string(rune(user_id)), "123456")
	userConnected := "connected:" + hashKey
	err = s.distributedCache.Set(context, userConnected, true, 0)
	if err != nil {
		return out, err
	}
	out = "User connected successfully"
	return out, nil
}
func (s *sInformation) DeleteUserConnected(context context.Context, user_id int64) (out string, err error) {
	hashKey := crypto.GenerateHash(string(rune(user_id)), "123456")
	userConnected := "connected:" + hashKey
	err = s.distributedCache.Del(context, userConnected)
	if err != nil {
		return out, err
	}
	out = "User disconnected successfully"
	return out, nil
}
func (s *sInformation) GetUserConnectedExists(context context.Context, user_id int64) (out bool, err error) {
	hashKey := crypto.GenerateHash(string(rune(user_id)), "123456")
	userConnected := "connected:" + hashKey
	exists, err := s.distributedCache.Exists(context, userConnected)
	if err != nil {
		return out, err
	}
	out = exists
	return out, nil
}
