package mapper

import (
	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

func ToMessageDTO(message database.PreGoCommunicationMessage99999) model.MessageOutput {
	return model.MessageOutput{
		ID:             message.ID,
		ConversationID: message.ConversationID,
		UserID:         message.UserID,
		Status:         message.Status.Bool,
		Message:        message.Message,
		Type:           message.Type,
		CreatedAt:      message.CreatedAt.Time,
		UpdatedAt:      message.UpdatedAt.Time,
	}
}
