package response

import "errors"

const (
	SuccessCodeStatus        = 20000
	CreateSuccessCodeStatus  = 20001
	CreateErrorCodeStatus    = 30000
	UpdateSuccessCodeStatus  = 30001
	DeleteSuccessCodeStatus  = 30002
	ParamInvalidCodeStatus   = 40000
	UnauthorizedCodeStatus   = 40001
	ForbiddenCodeStatus      = 40003
	FoundTicketErrCodeStatus = 40006
	TicketNotFoundCodeStatus = 40002
	UpdateErrorCodeStatus    = 40007
	DeleteErrorCodeStatus    = 40004
	NotFoundErrorCodeStatus  = 40005
	ErrorCodeStatus          = 50000
	ErrorCorsCodeStatus      = 40008
)

var msg = map[int]string{
	SuccessCodeStatus:        "success",
	CreateSuccessCodeStatus:  "create success",
	CreateErrorCodeStatus:    "create error",
	UpdateSuccessCodeStatus:  "update success",
	DeleteSuccessCodeStatus:  "delete success",
	ParamInvalidCodeStatus:   "param invalid",
	UnauthorizedCodeStatus:   "unauthorized",
	ForbiddenCodeStatus:      "forbidden",
	FoundTicketErrCodeStatus: "found ticket error",
	TicketNotFoundCodeStatus: "ticket not found",
	UpdateErrorCodeStatus:    "update error",
	DeleteErrorCodeStatus:    "delete error",
	NotFoundErrorCodeStatus:  "not found",
	ErrorCodeStatus:          "error",
	ErrorCorsCodeStatus:      "cors error",
}
var (
	ErrNotFoundDataErr = errors.New("not found data") // Type of internal error
)
