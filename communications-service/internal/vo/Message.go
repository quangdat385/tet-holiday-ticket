package vo

type MessageIDResquest struct {
	MessageID int64 `json:"message_id"`
}

type MessagesByConversationIDRequest struct {
	ConversationID int64 `json:"conversation_id" binding:"required"`
	Limit          int32 `json:"limit" binding:"omitempty,min=1,max=50"`
	Offset         int32 `json:"offset" binding:"omitempty,min=0"`
}
type MessagesByUserIDRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
	Limit  int32 `json:"limit" binding:"omitempty,min=1,max=50"`
	Offset int32 `json:"offset" binding:"omitempty,min=0"`
}
type UpdateMessageStatusRequest struct {
	Status bool  `json:"status" binding:"required"`
	ID     int64 `json:"id" binding:"required"`
}
type CreateMessageRequest struct {
	ConversationID int64  `json:"conversation_id" binding:"required"`
	UserID         int64  `json:"user_id" binding:"required"`
	Status         bool   `json:"status" binding:"-"`
	Message        string `json:"message" binding:"required"`
	Type           string `json:"type" binding:"required"`
}
