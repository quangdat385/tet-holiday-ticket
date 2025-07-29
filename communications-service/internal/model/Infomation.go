package model

import "time"

type InfomationInput struct {
	UserID int64  `json:"user_id"`
	Status bool   `json:"status,omitempty"`
	Value  string `json:"value,omitempty"`
	Type   string `json:"type,omitempty"`
}
type InformationOutput struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Status    bool      `json:"status"`
	Value     string    `json:"value"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
