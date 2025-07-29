package response

import "errors"

const (
	SuccessCodeStatus = 20000
)

var msg = map[int]string{
	SuccessCodeStatus: "success",
}

var (
	CouldNotGetTicketErr = errors.New("Could not get Ticket from MYSQL") //Type of Internal Error
)
