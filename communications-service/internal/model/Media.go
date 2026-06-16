package model

import "time"

type MediaInput struct {
	MessageID      *int64 `json:"message_id,omitempty"`
	ConversationID *int64 `json:"conversation_id,omitempty"`
	UserID         int64  `json:"user_id"`
	FileName       string `json:"file_name"`
	OriginalName   string `json:"original_name"`
	MimeType       string `json:"mime_type"`
	Size           int64  `json:"size"`
	URL            string `json:"url"`
}

type MediaOutput struct {
	ID             int64     `json:"id"`
	MessageID      *int64    `json:"message_id,omitempty"`
	ConversationID *int64    `json:"conversation_id,omitempty"`
	UserID         int64     `json:"user_id"`
	FileName       string    `json:"file_name"`
	OriginalName   string    `json:"original_name"`
	MimeType       string    `json:"mime_type"`
	Size           int64     `json:"size"`
	URL            string    `json:"url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
