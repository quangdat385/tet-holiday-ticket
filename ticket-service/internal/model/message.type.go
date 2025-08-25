package model

type Message struct {
	Type             string           `json:"type"`
	MessageData      MessageData      `json:"message_data"`
	NotificationData NotificationData `json:"noti_data"`
}

type MessageData struct {
	ConversationID int64  `json:"conversation_id"`
	UserID         int64  `json:"user_id"`
	Status         bool   `json:"status"`
	Message        string `json:"message"`
	Type           string `json:"type"`
}
type NotificationData struct {
	From    int64 `json:"from"`
	To      any   `json:"to"`
	Content any   `json:"content"`
}
type ContentType struct {
	OrderNumber string `json:"order_number"`
	Message     string `json:"message"`
	Status      bool   `json:"status"`
}
