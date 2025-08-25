package mapper

import (
	"encoding/json"
	"log"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

func ToNotificationDTO(notification database.PreGoCommunicationNotification99999) model.NotificationOutput {
	var Contents any
	if err := json.Unmarshal(notification.Content, &Contents); err != nil {
		log.Printf("failed to unmarshal Content: %v", err)
		Contents = nil
	}
	var To int64
	if notification.To.Valid {
		To = notification.To.Int64
	}
	return model.NotificationOutput{
		ID:        notification.ID,
		From:      notification.From,
		To:        To,
		Content:   Contents,
		CreatedAt: notification.CreatedAt.Time,
		UpdatedAt: notification.UpdatedAt.Time,
	}
}
