package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model/mapper"
)

type sNotification struct {
	r *database.Queries
}

func NewNotificationServiceImpl(r *database.Queries) *sNotification {
	return &sNotification{
		r: r,
	}
}
func (s *sNotification) GetNotificationById(context context.Context, id int64) (out model.NotificationOutput, err error) {
	notification, err := s.r.GetNotificationById(context, id)
	if err != nil {
		return out, err
	}
	if notification.ID == 0 {
		return out, errors.New("notification not found")
	}
	out = mapper.ToNotificationDTO(notification)
	return out, nil
}
func (s *sNotification) GetNotificationsByUserIDTo(context context.Context, userId int64) (out []model.NotificationOutput, err error) {
	notifications, err := s.r.GetNotificationsByUserIDTo(context, database.GetNotificationsByUserIDToParams{
		JSONQUOTE: fmt.Sprint(userId),
		Offset:    0,
		Limit:     50,
	})
	if err != nil {
		return out, err
	}
	for _, notification := range notifications {
		out = append(out, mapper.ToNotificationDTO(notification))
	}
	return out, nil
}
func (s *sNotification) GetNotificationsByUserIDFrom(context context.Context, userId int64) (out model.NotificationOutput, err error) {
	notification, err := s.r.GetNotificationsByUserIDFrom(context, userId)
	if err != nil {
		return out, err
	}
	out = mapper.ToNotificationDTO(notification)
	return out, nil
}
func (s *sNotification) GetNotificationsFromUserIDToIsNull(context context.Context) (out []model.NotificationOutput, err error) {
	notifications, err := s.r.GetNotificationWhenToIsNull(context, database.GetNotificationWhenToIsNullParams{
		Limit:  50,
		Offset: 0,
	})
	if err != nil {
		return out, err
	}
	for _, notification := range notifications {
		out = append(out, mapper.ToNotificationDTO(notification))
	}
	return out, nil
}
func (s *sNotification) CreateNotification(context context.Context, input model.NotificationInput) (out model.NotificationOutput, err error) {
	var toRawMessage json.RawMessage
	if input.To != nil {
		var ok bool
		toRawMessage, ok = input.To.(json.RawMessage)
		if !ok {
			return out, errors.New("input.To must be of type json.RawMessage")
		}
	}
	var contentRawMessage json.RawMessage
	if input.Content != nil {
		var ok bool
		contentRawMessage, ok = input.Content.(json.RawMessage)
		if !ok {
			return out, errors.New("input.Content must be of type json.RawMessage")
		}
	}
	notification, err := s.r.InsertNotification(context, database.InsertNotificationParams{
		From:    input.From,
		To:      toRawMessage,
		Content: contentRawMessage,
	})
	if err != nil {
		return out, err
	}
	ID, err := notification.LastInsertId()
	if err != nil {
		return out, err
	}
	notificationDb, err := s.r.GetNotificationById(context, ID)
	if err != nil {
		return out, nil
	}
	out = mapper.ToNotificationDTO(notificationDb)
	return out, nil
}
func (s *sNotification) DeleteNotificationById(context context.Context, id int64) (err error) {
	notification, err := s.r.GetNotificationById(context, id)
	if err != nil {
		return err
	}
	if notification.ID == 0 {
		return errors.New("notification not found")
	}
	_, err = s.r.DeleteNotificationById(context, id)
	if err != nil {
		return err
	}
	return nil
}
