package vo

type InformationUserIDRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
}
type InformationIDRequest struct {
	InformationID int64 `json:"information_id" validate:"required"`
}
type CreateInformationRequest struct {
	UserID int64  `json:"user_id" binding:"required"`
	Status bool   `json:"status" binding:"required"`
	Value  string `json:"value" binding:"required"`
	Type   string `json:"type" binding:"required"`
}
