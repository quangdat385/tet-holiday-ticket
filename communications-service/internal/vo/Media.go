package vo

type MediaIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type MediaByMessageIDRequest struct {
	MessageID int64 `uri:"message_id" binding:"required"`
}

type MediaByConversationIDRequest struct {
	ConversationID int64 `form:"conversation_id" binding:"required"`
	Limit          int32 `form:"limit" binding:"omitempty,min=1,max=50"`
	Offset         int32 `form:"offset" binding:"omitempty,min=0"`
}

type UploadMediaRequest struct {
	MessageID      *int64 `form:"message_id"`
	ConversationID *int64 `form:"conversation_id"`
	UserID         int64  `form:"user_id" binding:"required"`
}
