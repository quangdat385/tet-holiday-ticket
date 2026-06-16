package mapper

import (
	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

func ToMediaDTO(media database.PreGoCommunicationMedia99999) model.MediaOutput {
	out := model.MediaOutput{
		ID:           media.ID,
		UserID:       media.UserID,
		FileName:     media.FileName,
		OriginalName: media.OriginalName,
		MimeType:     media.MimeType,
		Size:         media.Size,
		URL:          media.URL,
		CreatedAt:    media.CreatedAt.Time,
		UpdatedAt:    media.UpdatedAt.Time,
	}
	if media.MessageID.Valid {
		out.MessageID = &media.MessageID.Int64
	}
	if media.ConversationID.Valid {
		out.ConversationID = &media.ConversationID.Int64
	}
	return out
}
