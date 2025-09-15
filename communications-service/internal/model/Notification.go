package model

import "time"

type NotificationInput struct {
	From    int64 `json:"from"`
	To      int64 `json:"to,omitempty"`
	Content any   `json:"content"`
}
type UpdateNotificationInput struct {
	UserID         int64 `json:"user_id"`
	NotificationID int64 `json:"notification_id"`
}
type NotificationOutput struct {
	ID        int64     `json:"id"`
	From      int64     `json:"from"`
	To        int64     `json:"to,omitempty"`
	Content   any       `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
