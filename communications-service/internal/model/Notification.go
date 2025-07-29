package model

import "time"

type NotificationInput struct {
	From    int64 `json:"from"`
	To      any   `json:"to"`
	Content any   `json:"content"`
}

type NotificationOutput struct {
	ID        int64     `json:"id"`
	From      int64     `json:"from"`
	To        any       `json:"to"`
	Content   any       `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
