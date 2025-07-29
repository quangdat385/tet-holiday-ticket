package model

type StationOutput struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int32  `json:"status"`
}
type StationInput struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int32  `json:"status"`
}
type StationListInput struct {
	Page   int   `json:"page"`
	Limit  int   `json:"limit"`
	Status int32 `json:"status"`
}
