package vo

type StationIDRequest struct {
	StationID int `uri:"station_id" binding:"required"`
}
type CreateStationRequest struct {
	Name   string `json:"name" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Status int32  `json:"status" binding:"required"`
}
type UpdateStationRequest struct {
	StationID int    `json:"station_id" binding:"required"`
	Name      string `json:"name" binding:"-"`
	Code      string `json:"code" binding:"-"`
	Status    int32  `json:"status" binding:"-"`
}
type UpdateStationStatusRequest struct {
	StationID int   `json:"station_id" binding:"required"`
	Status    int32 `json:"status" binding:"required"`
}
type StationListRequest struct {
	Page   int `form:"page" binding:"required,min=1"`
	Limit  int `form:"limit" binding:"required,min=50,max=100"`
	Status int `form:"status" binding:"omitempty,oneof=0 1"`
}
