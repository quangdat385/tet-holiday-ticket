package model

type RouteSegmentOutPut struct {
	ID            int64 `json:"id"`
	TrainID       int64 `json:"train_id"`
	FromStationID int64 `json:"from_station_id"`
	ToStationID   int64 `json:"to_station_id"`
	SegmentOrder  int32 `json:"segment_order"`
	DistanceKm    int32 `json:"distance_km"`
}

type RouteSegmentCreateInPut struct {
	TrainID       int64 `json:"train_id"`
	FromStationID int64 `json:"from_station_id"`
	ToStationID   int64 `json:"to_station_id"`
	SegmentOrder  int32 `json:"segment_order"`
	DistanceKm    int32 `json:"distance_km"`
}
type RouteSegmentUpdateInPut struct {
	ID            int64 `json:"id"`
	TrainID       int64 `json:"train_id"`
	FromStationID int64 `json:"from_station_id"`
	ToStationID   int64 `json:"to_station_id"`
	SegmentOrder  int32 `json:"segment_order"`
	DistanceKm    int32 `json:"distance_km"`
}
