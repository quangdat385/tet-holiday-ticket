package model

import "time"

type UserIDS struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
}

type ConversationInput struct {
	Title       string    `json:"title"`
	UserIDS     []UserIDS `json:"user_ids"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Background  string    `json:"background"`
	Emoji       string    `json:"emoji"`
	IsDeleted   bool      `json:"is_deleted"`
}
type UpdateConversationInput struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Background  string `json:"background"`
	Emoji       string `json:"emoji"`
	IsDeleted   bool   `json:"is_deleted"`
}
type SoftDeleteConversationInput struct {
	ID int64 `json:"id"`
}

type ConversationOutput struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	UserIDS     []UserIDS `json:"user_ids"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Background  string    `json:"background"`
	Emoji       string    `json:"emoji"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
