package vo

type MessageIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type MessagesByConversationIDRequest struct {
	ConversationID int64 `form:"conversation_id" binding:"required"`
	Limit          int32 `form:"limit" binding:"omitempty,min=1,max=50"`
	Offset         int32 `form:"offset" binding:"omitempty,min=0"`
}
type MessagesByUserIDRequest struct {
	UserID int64 `form:"user_id" binding:"required"`
	Limit  int32 `form:"limit" binding:"omitempty,min=1,max=50"`
	Offset int32 `form:"offset" binding:"omitempty,min=0"`
}
type UpdateMessageStatusRequest struct {
	UserID    int64 `json:"user_id" binding:"required"`
	MessageID int64 `json:"message_id" binding:"required"`
}
type CreateMessageRequest struct {
	ConversationID int64  `json:"conversation_id" binding:"required"`
	UserID         int64  `json:"user_id" binding:"required"`
	Status         bool   `json:"status" binding:"-"`
	Message        string `json:"message" binding:"required"`
	Type           string `json:"type" binding:"required"`
}
