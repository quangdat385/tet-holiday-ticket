package mapper

import (
	"encoding/json"
	"log"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

func ToConversationDTO(conversation database.PreGoCommunicationConversation99999) model.ConversationOutput {
	var userIDs []int64
	if err := json.Unmarshal(conversation.UserIds, &userIDs); err != nil {
		log.Printf("failed to unmarshal UserIds: %v", err)
		userIDs = []int64{} // Fallback to empty slice if unmarshalling fails
	}
	return model.ConversationOutput{
		ID:          conversation.ID,
		Title:       conversation.Title.String,
		UserIDS:     userIDs,
		Description: conversation.Description.String,
		Type:        conversation.Type.String,
		Background:  conversation.Background.String,
		Emoji:       conversation.Emoji.String,
		IsDeleted:   conversation.IsDeleted.Bool,
		CreatedAt:   conversation.CreatedAt.Time,
		UpdatedAt:   conversation.UpdatedAt.Time,
	}
}
