package vo

type NotificationIDRequest struct {
	NotificationID int `json:"notification_id"`
}
type CreateNotificationRequest struct {
	From    int64 `json:"from" binding:"required"`
	To      int64 `json:"to" binding:"-"`
	Content any   `json:"content" binding:"required"`
}
