package model

import "time"

type ConversationInput struct {
	Title       string  `json:"title"`
	UserIDS     []int64 `json:"user_ids"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Background  string  `json:"background"`
	Emoji       string  `json:"emoji"`
	IsDeleted   bool    `json:"is_deleted"`
}
type UpdateConversationInput struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	UserIDS     []int64 `json:"user_ids"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Background  string  `json:"background"`
	Emoji       string  `json:"emoji"`
	IsDeleted   bool    `json:"is_deleted"`
}
type ConversationOutput struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	UserIDS     []int64   `json:"user_ids"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Background  string    `json:"background"`
	Emoji       string    `json:"emoji"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
