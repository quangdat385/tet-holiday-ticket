package vo

type SegmentIDRequest struct {
	SegmentID int `uri:"segment_id" binding:"required"`
}
type SegmentFromStationIDRequest struct {
	FromStationID int `uri:"from_station_id" binding:"required"`
}
type SegmentToStationIDRequest struct {
	ToStationID int `uri:"to_station_id" binding:"required"`
}
type CreateSegmentRequest struct {
	TrainID       int `json:"train_id" binding:"required"`
	FromStationID int `json:"from_station_id" binding:"required"`
	ToStationID   int `json:"to_station_id" binding:"required"`
	SegmentOrder  int `json:"segment_order" binding:"required"`
	DistanceKm    int `json:"distance_km" binding:"required"`
}
type UpdateSegmentRequest struct {
	SegmentID     int `json:"segment_id" binding:"required"`
	TrainID       int `json:"train_id" binding:"optional"`
	FromStationID int `json:"from_station_id" binding:"optional"`
	ToStationID   int `json:"to_station_id" binding:"optional"`
	SegmentOrder  int `json:"segment_order" binding:"optional"`
	DistanceKm    int `json:"distance_km" binding:"optional"`
}
type SegmentListRequest struct {
	TrainID int `json:"train_id" binding:"-"`
	Page    int `json:"page" binding:"required"`
	Limit   int `json:"limit" binding:"required"`
}
