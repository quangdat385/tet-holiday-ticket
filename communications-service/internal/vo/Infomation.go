package vo

type InformationUserIDRequest struct {
	UserID int64 `uri:"user_id" binding:"required"`
}
type InformationIDRequest struct {
	InformationID int64 `uri:"information_id" binding:"required"`
}
type CreateInformationRequest struct {
	UserID int64  `json:"user_id" binding:"required"`
	Status bool   `json:"status" binding:"-"`
	Value  string `json:"value" binding:"-"`
	Type   string `json:"type" binding:"-"`
}
