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
	ParamInvalidCodeStatus:   "param invalid",
	FoundTicketErrCodeStatus: "found ticket error",
	TicketNotFoundCodeStatus: "ticket not found",
	CreateSuccessCodeStatus:  "create success",
	CreateErrorCodeStatus:    "create error",
	UpdateErrorCodeStatus:    "update error",
	UpdateSuccessCodeStatus:  "update success",
	DeleteErrorCodeStatus:    "delete error",
	DeleteSuccessCodeStatus:  "delete success",
	NotFoundErrorCodeStatus:  "not found",
	ErrorCodeStatus:          "error",
	UnauthorizedCodeStatus:   "unauthorized",
	ForbiddenCodeStatus:      "forbidden",
	ErrorCorsCodeStatus:      "cors error",
}

var (
	ErrCouldNotGetTicketErr          = errors.New("could not get ticket from MySQL")   // Type of internal error
	ErrTicketNotFoundErr             = errors.New("ticket not found")                  // Type of internal error
	ErrSeatNotFoundErr               = errors.New("seat not found")                    // Type of internal error
	ErrCreateSeatReservedErr         = errors.New("create seat reserved error")        // Type of internal error
	ErrSeatReserverNotFoundErr       = errors.New("seat reserved not found")           // Type of internal error
	ErrRouteSegmentNotFoundErr       = errors.New("route segment not found")           // Type of internal error
	ErrCreateRouteSegmentErr         = errors.New("create route segment error")        // Type of internal error
	ErrTicketSegmentPriceNotFoundErr = errors.New("ticket segment price not found")    // Type of internal error
	ErrCreateTicketSegmentPriceErr   = errors.New("create ticket segment price error") // Type of internal error
	ErrUpdateTicketSegmentPriceErr   = errors.New("update ticket segment price error") // Type of internal error
	ErrNotFoundDataErr               = errors.New("not found data")                    // Type of internal error
)
