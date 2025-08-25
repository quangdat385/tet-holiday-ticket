package mapper

import (
	"encoding/json"
	"log"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

func ToConversationDTO(conversation database.GetCommunicationConversationByIdRow) model.ConversationOutput {
	var userIDs []model.UserIDS
	if err := json.Unmarshal(conversation.UserIds, &userIDs); err != nil {
		log.Printf("failed to unmarshal UserIds: %v", err)
		userIDs = []model.UserIDS{} // Fallback to empty slice if unmarshalling fails
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
func ToConversationDTOBYUserID(conversation []database.GetListCommunicationConversationsRow) []model.ConversationOutput {
	var out []model.ConversationOutput
	for _, conv := range conversation {
		var userIDs []model.UserIDS
		if err := json.Unmarshal(conv.UserIds, &userIDs); err != nil {
			log.Printf("failed to unmarshal UserIds: %v", err)
			userIDs = []model.UserIDS{} // Fallback to empty slice if unmarshalling fails
		}
		out = append(out, model.ConversationOutput{
			ID:          conv.ID,
			Title:       conv.Title.String,
			UserIDS:     userIDs,
			Description: conv.Description.String,
			Type:        conv.Type.String,
			Background:  conv.Background.String,
			Emoji:       conv.Emoji.String,
			IsDeleted:   conv.IsDeleted.Bool,
			CreatedAt:   conv.CreatedAt.Time,
			UpdatedAt:   conv.UpdatedAt.Time,
		})
	}
	return out
}
