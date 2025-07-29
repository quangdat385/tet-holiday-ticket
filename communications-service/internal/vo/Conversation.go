package vo

type ConversationIDRequest struct {
	ConversationID int64 `json:"conversation_id" binding:"required"`
}
type UserIDRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}
type UpdateConversationRequest struct {
	ConversationID int64   `json:"conversation_id" binding:"required"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	UserIDS        []int64 `json:"user_ids"`
	Type           string  `json:"type"`
	Background     string  `json:"background"`
	Emoji          string  `json:"emoji"`
	IsDeleted      bool    `json:"is_deleted"`
}
type AddUserToConversationRequest struct {
	ConversationID int64   `json:"conversation_id" binding:"required"`
	UserIDS        []int64 `json:"user_ids" binding:"required"`
}
type RemoveUserFromConversationRequest struct {
	ConversationID int64 `json:"conversation_id" binding:"required"`
	UserID         int64 `json:"user_id" binding:"required"`
}

type CreateConversationRequest struct {
	Title       string  `json:"title" binding:"required"`
	UserIDS     []int64 `json:"user_ids" binding:"required"`
	Description string  `json:"description"`
	Type        string  `json:"type" binding:"required"`
	Background  string  `json:"background"`
	Emoji       string  `json:"emoji"`
	IsDeleted   bool    `json:"is_deleted"`
}
