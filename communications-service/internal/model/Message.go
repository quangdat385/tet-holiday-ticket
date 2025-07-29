package model

import "time"

type MessageInput struct {
	ConversationID int64  `json:"conversation_id"`
	UserID         int64  `json:"user_id"`
	Status         bool   `json:"status"`
	Message        string `json:"message"`
	Type           string `json:"type"`
}

type MessageOutput struct {
	ID             int64     `json:"id"`
	ConversationID int64     `json:"conversation_id"`
	UserID         int64     `json:"user_id"`
	Status         bool      `json:"status"`
	Message        string    `json:"message"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
