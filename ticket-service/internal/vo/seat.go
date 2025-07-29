package vo

type SeatIDRequest struct {
	SeatID int `uri:"seat_id" binding:"required"`
}
type SeatListByTrainIDRequest struct {
	TrainID int `form:"train_id" binding:"required"`
	Limit   int `form:"limit" binding:"required"`
	Page    int `form:"page" binding:"required"`
}
type CreateSeatRequest struct {
	TrainID    int    `json:"train_id" binding:"required"`
	SeatNumber string `json:"seat_number" binding:"required"`
	SeatClass  string `json:"seat_class" binding:"required"`
	Status     int    `json:"status" binding:"required"`
}
type UpdateSeatRequest struct {
	SeatID     int    `json:"seat_id" binding:"required"`
	TrainID    int    `json:"train_id" binding:"-"`
	SeatNumber string `json:"seat_number" binding:"-"`
	SeatClass  string `json:"seat_class" binding:"-"`
	Status     int    `json:"status" binding:"-"`
}
type SeatListRequest struct {
	TrainID   int    `json:"train_id" binding:"-"`
	SeatClass string `json:"seat_class" binding:"-"`
	Status    int    `json:"status" binding:"-"`
	Limit     int    `json:"limit" binding:"required"`
	Page      int    `json:"page" binding:"required"`
}
