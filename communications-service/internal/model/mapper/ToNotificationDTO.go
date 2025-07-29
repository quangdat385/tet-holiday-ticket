package mapper

import (
	"encoding/json"
	"log"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

func ToNotificationDTO(notification database.PreGoCommunicationNotification99999) model.NotificationOutput {
	var Tos any
	var Contents any
	if notification.To != nil {
		if err := json.Unmarshal(notification.To, &Tos); err != nil {
			log.Printf("failed to unmarshal To: %v", err)
			Tos = nil
		}
	} else {
		Tos = nil
	}
	if err := json.Unmarshal(notification.Content, &Contents); err != nil {
		log.Printf("failed to unmarshal Content: %v", err)
		Contents = nil
	}
	return model.NotificationOutput{
		ID:        notification.ID,
		From:      notification.From,
		To:        Tos,
		Content:   notification.Content,
		CreatedAt: notification.CreatedAt.Time,
		UpdatedAt: notification.UpdatedAt.Time,
	}
}
