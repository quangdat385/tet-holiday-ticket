package vo

type NotificationIDRequest struct {
	NotificationID int `uri:"notification_id" binding:"required"`
}
type CreateNotificationRequest struct {
	From    int64 `json:"from" binding:"required"`
	To      int64 `json:"to" binding:"-"`
	Content any   `json:"content" binding:"required"`
}
type UpdateNotificationRequest struct {
	UserID         int64 `json:"user_id" binding:"required"`
	NotificationID int64 `json:"notification_id" binding:"required"`
}
